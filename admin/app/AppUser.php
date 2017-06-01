<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class AppUser extends Model
{
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'email',
    ];

    /**
     * Get the ratings that belong to this app user.
     */
    public function ratings()
    {
        return $this->hasMany('App\Rating');
    }

    /**
     * Get the messages that belong to this app user.
     */
    public function messages()
    {
        return $this->hasMany('App\Message');
    }

    /**
     * Get the apps of the app user.
     */
     public function apps() {
        return $this->belongsToMany('App\App');
     }
}
