<?php

namespace App;

use Illuminate\Notifications\Notifiable;
use Illuminate\Foundation\Auth\User as Authenticatable;

class User extends Authenticatable
{
    use Notifiable;

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'email', 'password',
    ];

    /**
     * The attributes that should be hidden for arrays.
     *
     * @var array
     */
    protected $hidden = [
        'password', 'remember_token',
    ];

    /**
     * Get the messages that belong to this user.
     */
    public function messages()
    {
        return $this->hasMany('App\Message');
    }

     /**
     * Get the apps the user belongs to.
     */
     public function apps() {
        return $this->belongsToMany('App\App');
     }
}
