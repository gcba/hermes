<?php

namespace App;

use Spatie\Activitylog\Traits\LogsActivity;
use Illuminate\Support\Facades\Auth;
use TCG\Voyager\Models\User as VoyagerUser;

class User extends VoyagerUser
{
    use LogsActivity;

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'email', 'role_id', 'updated_by'
    ];

    /**
     * Log all fillable attributes.
     *
     * @var array
     */
    protected static $logFillable = true;

    /**
     * Boot function for using with User Events
     *
     * @return void
     */
    protected static function boot() {
        parent::boot();

        static::creating(function ($model) {
            $model->attributes['updated_at'] = null;
            $model->attributes['updated_by'] = null;
        });

        static::updating(function ($model) {
            \Auth::user() !== null ?
                $model->attributes['updated_by'] = \Auth::user()->id :
                $model->attributes['updated_by'] = null;

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
    public function role() {
        return $this->belongsTo('App\Role', 'role_id', 'id');
    }

    /**
     * For Voyager's CRUD.
     */
    public function roleId() {
        return $this->belongsTo('App\Role', 'role_id', 'id');
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