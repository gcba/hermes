<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Brand extends Model
{
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name'
    ];

    /**
     * Get the devices that belong to this brand.
     */
    public function devices()
    {
        return $this->hasMany('App\Device');
    }
}
