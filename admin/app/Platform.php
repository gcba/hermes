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

    /**
     * Boot function for using with User Events
     *
     * @return void
     */
    protected static function boot()
    {
        parent::boot();

        static::creating(function ($model)
        {
            if (!$model->key) {
                $model->attributes['key'] = md5(date("Y-m-d H:i:s"));
            }
        });
    }

    /**
     * Get the ratings that belong to the platform.
     */
    public function ratings()
    {
        return $this->hasMany('App\Rating', 'rating_id', 'id');
    }

     /**
     * Get the devices that belong to the platform.
     */
     public function devices() {
        return $this->hasMany('App\Device', 'device_id', 'id');
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
}
