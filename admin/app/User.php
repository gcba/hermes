<?php

namespace App;

use TCG\Voyager\Models\User as VoyagerUser;

class User extends VoyagerUser
{
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'email', 'password', 'modified_by'
    ];

    /**
     * The attributes that should be hidden for arrays.
     *
     * @var array
     */
    protected $hidden = [
        'password', 'remember_token'
    ];

    protected $utils;

    /**
     * Create a new controller instance.
     *
     * @param  UtilsService $utils
     * @return void
     */
    public function __construct()
    {
        $this->utils = resolve('App\Services\UtilsService');
    }

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
    }

     /**
     * Get the apps the user belongs to.
     */
     public function apps() {
        return $this->belongsToMany('App\App');
     }

     /**
     * Get the user that last modified the app.
     */
     public function modifiedBy() {
        return $this->belongsTo('App\User', 'modified_by', 'id');
     }

    public function getCreatedAtAttribute() {
        return $this->utils->formatDate($this->attributes['created_at']);
    }

    public function getUpdatedAtAttribute() {
        return $this->attributes['updated_at'] ? $this->utils->formatDate($this->attributes['updated_at']) : '-';
    }
}
