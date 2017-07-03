<?php

namespace App;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

class AppUser extends Model
{
    use SoftDeletes;

    protected $table = 'appusers';

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'email', 'miba_id'
    ];

    /**
     * The attributes that should be mutated to dates.
     *
     * @var array
     */
    protected $dates = ['deleted_at'];

    /**
     * Get the ratings that belong to this app user.
     */
    public function ratings()
    {
        return $this->hasMany('App\Rating', 'rating_id', 'id');
    }

    /**
     * Get the apps of the app user.
     */
     public function apps() {
        return $this->belongsToMany('App\App');
     }

     /**
     * Get the platform that belong to the appuser.
     */
     public function platforms() {
        return $this->belongsToMany('App\Platform');
     }

     /**
     * Get the devices that belong to the app user.
     */
     public function devices() {
        return $this->belongsToMany('App\Device');
     }
}
