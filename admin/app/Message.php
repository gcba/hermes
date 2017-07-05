<?php

namespace App;

use DateTime;
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
    protected static function boot() {
        parent::boot();

        static::creating(function ($model) {
            $model->attributes['direction'] = 'out';
            $model->attributes['updated_at'] = null;
        });
    }

    /**
     * Get the rating the message belongs to.
     */
    public function rating() {
        return $this->belongsTo('App\Rating', 'rating_id', 'id');
    }

     /**
     * For Voyager's CRUD.
     */
    public function ratingId() {
        return $this->belongsTo('App\Rating', 'rating_id', 'id');
    }

    public function setMessageAttribute($value) {
        $this->attributes['message'] = ucfirst(filter_var(trim($value), FILTER_SANITIZE_SPECIAL_CHARS));
    }

    public function getDirectionAttribute(){
        return $this->attributes['direction'] == 'out' ? '➡️' : '⬅️';
    }

    public function getCreatedAtAttribute(){
        return $this->formatDate($this->attributes['created_at']);
    }

    public function getUpdatedAtAttribute(){
        return $this->attributes['updated_at'] ? $this->formatDate($this->attributes['updated_at']) : '-';
    }

    private function formatDate($dateString) {
        $date = new DateTime($dateString);

        return $date->format('d/m/Y H:i:s');
    }
}
