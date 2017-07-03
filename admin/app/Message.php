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
     * Boot function for using with User Events
     *
     * @return void
     */
    protected static function boot()
    {
        parent::boot();

        static::creating(function ($model)
        {
            if (!$model->key) {
                $model->attributes['direction'] = 'out';
            }
        });
    }

    /**
     * Get the rating the message belongs to.
     */
    public function rating()
    {
        return $this->belongsTo('App\Rating', 'rating_id', 'id');
    }
}
