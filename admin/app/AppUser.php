<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class AppUser extends Model
{
    protected $table = 'appuser';

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'email'
    ];

    /**
     * Get the ratings that belong to this app user.
     */
    public function ratings()
    {
        return $this->hasMany('App\Rating');
    }

    /**
     * Get the apps of the app user.
     */
     public function apps() {
        return $this->belongsToMany('App\App');
     }
}
