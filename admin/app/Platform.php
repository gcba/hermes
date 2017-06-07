<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Platform extends Model
{
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'key'
    ];

    /**
     * Get the ratings that belong to the platform.
     */
    public function ratings()
    {
        return $this->hasMany('App\Rating');
    }

     /**
     * Get the devices that belong to the platform.
     */
     public function devices() {
        return $this->hasMany('App\Device');
     }

     /**
     * Get the apps that belong to the platform.
     */
     public function apps() {
        return $this->belongsToMany('App\App');
     }
}
