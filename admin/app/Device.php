<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Device extends Model
{
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'screen_width', 'screen_height', 'ppi', 'brand_id', 'platform_id'
    ];

    /**
     * Get the ratings that belong to this device.
     */
    public function ratings()
    {
        return $this->hasMany('App\Rating');
    }

    /**
     * Get the device's brand.
     */
     public function brand() {
        return $this->belongsTo('App\Brand');
     }

     /**
     * Get the platform the device belongs to.
     */
     public function platform() {
        return $this->belongsTo('App\Platform');
     }
}