<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Message extends Model
{
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'message', 'direction', 'rating_id', 'appuser_id', 'user_id',
    ];

    /**
     * Get the rating the message belongs to.
     */
    public function rating()
    {
        return $this->belongsTo('App\Rating');
    }

    /**
     * Get the app user the message belongs to.
     */
    public function appUser()
    {
        return $this->belongsTo('App\AppUser', 'appuser_id');
    }

    /**
     * Get the user the message belongs to.
     */
    public function user()
    {
        return $this->belongsTo('App\User');
    }
}
