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
        'name',
    ];

    /**
     * Get the apps that belong to the platform.
     */
     public function apps() {
        return $this->belongsToMany('App\App');
     }
}
