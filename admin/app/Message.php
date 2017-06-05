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
        'message', 'direction', 'rating_id'
    ];

    /**
     * Get the rating the message belongs to.
     */
    public function rating()
    {
        return $this->belongsTo('App\Rating');
    }
}
