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
            $model->attributes['direction'] = 'out';
            $model->attributes['message'] = sanitizeMessage($model->message);
        });

        static::updating(function ($model)
        {
            $model->attributes['message'] = sanitizeMessage($model->message);
        });
    }

    /**
     * Get the rating the message belongs to.
     */
    public function rating()
    {
        return $this->belongsTo('App\Rating', 'rating_id', 'id');
    }

     /**
     * For Voyager's CRUD.
     */
    public function ratingId()
    {
        return $this->belongsTo('App\Rating', 'rating_id', 'id');
    }

    private function sanitizeMessage($message) {
         return filter_var(trim($message), FILTER_SANITIZE_SPECIAL_CHARS);
    }
}
