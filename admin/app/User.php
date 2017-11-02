<?php

namespace App;

use Illuminate\Support\Facades\Auth;
use TCG\Voyager\Models\User as VoyagerUser;
use TCG\Voyager\Models\Role;

class User extends VoyagerUser
{
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'email', 'password', 'updated_by'
    ];

    /**
     * The attributes that should be hidden for arrays.
     *
     * @var array
     */
    protected $hidden = [
        'password', 'remember_token'
    ];

    /**
     * Boot function for using with User Events
     *
     * @return void
     */
    protected static function boot() {
        parent::boot();

        static::creating(function ($model) {
            $model->attributes['updated_at'] = null;
        });

        static::updating(function ($model) {
            \Auth::user() !== null ?
                $model->attributes['updated_by'] = \Auth::user()->id :
                $model->attributes['updated_at'] = null;

            $adminRole = Role::where('name', 'admin')->firstOrFail();

            if ($model->role_id === $adminRole->id) {
                $apps = App::select('id')->pluck('id')->toArray();

                $model->apps()->attach($apps);
            }
        });
    }

     /**
      * Get the apps the user belongs to.
      */
     public function apps() {
        return $this->belongsToMany('App\App');
     }

     /**
     * Get the messages that were created by this user.
     */
    public function messages() {
        return $this->hasMany('App\Message', 'created_by', 'id');
    }

    /**
     * Get the user that last modified this user.
     */
    public function updatedBy() {
        return $this->belongsTo('App\User', 'updated_by', 'id');
    }

    /**
     * For Voyager's CRUD.
     */
    public function updatedById() {
        return $this->belongsTo('App\User', 'updated_by', 'id');
    }

    public function getCreatedAtAttribute() {
        $utils = resolve('App\Services\UtilsService');

        return $utils->formatDate($this->attributes['created_at']);
    }

    public function getUpdatedAtAttribute() {
        $utils = resolve('App\Services\UtilsService');

        return $this->attributes['updated_at'] ? $utils->formatDate($this->attributes['updated_at']) : '-';
    }
}