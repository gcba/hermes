<?php

namespace App;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

class Platform extends Model
{
    use SoftDeletes;

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'key'
    ];

    /**
     * The attributes that should be mutated to dates.
     *
     * @var array
     */
    protected $dates = ['deleted_at'];

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
            if (!$model->key) {
                $model->attributes['key'] = md5(date("Y-m-d H:i:s"));
            }

            $model->attributes['updated_at'] = null;
        });
    }

    /**
     * Get the ratings that belong to the platform.
     */
    public function ratings() {
        return $this->hasMany('App\Rating', 'platform_id', 'id');
    }

     /**
     * Get the devices that belong to the platform.
     */
     public function devices() {
        return $this->hasMany('App\Device', 'platform_id', 'id');
     }

     /**
     * Get the apps that belong to the platform.
     */
     public function apps() {
        return $this->belongsToMany('App\App');
     }

     /**
     * Get the appusers that belong to the platform.
     */
     public function appusers() {
        return $this->belongsToMany('App\AppUser');
     }

    public function setNameAttribute($value) {
        $this->attributes['name'] = filter_var(trim($value), FILTER_SANITIZE_SPECIAL_CHARS);
    }

    public function getCreatedAtAttribute() {
        return $this->utils->formatDate($this->attributes['created_at']);
    }

    public function getUpdatedAtAttribute() {
        return $this->attributes['updated_at'] ? $this->utils->formatDate($this->attributes['updated_at']) : '-';
    }
}
