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
        'name', 'screen_width', 'screen_height', 'ppi',
    ];

    /**
     * Get the device's brand.
     */
     public function brand() {
        return $this->belongsTo('App\Brand');
     }
}
