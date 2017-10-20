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
        'message', 'direction', 'status', 'transport_id', 'rating_id', 'created_by'
    ];

    /**
     * Boot function for using with User Events
     *
     * @return void
     */
    protected static function boot() {
        parent::boot();

        static::creating(function ($model) {
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

    /**
     * Get the user that created the message.
     */
    public function createdBy() {
        return $this->belongsTo('App\User', 'created_by', 'id');
    }

    /**
     * For Voyager's CRUD.
     */
    public function createdById() {
        return $this->belongsTo('App\User', 'created_by', 'id');
    }


    public function setMessageAttribute($value) {
        $filteredValue = filter_var(
            trim(mb_strimwidth($value, 0, 1500, '')),
            FILTER_SANITIZE_STRING,
            FILTER_FLAG_STRIP_LOW | FILTER_FLAG_STRIP_BACKTICK | FILTER_FLAG_ENCODE_LOW
        );

        $this->attributes['message'] = ucfirst(htmlspecialchars($filteredValue));
    }

    public function setTransportIdAttribute($value) {
        $this->attributes['transport_id'] = filter_var(trim($value), FILTER_SANITIZE_EMAIL);
    }

    public function getMessageAttribute() {
        return html_entity_decode( $this->attributes['message']);
    }

    public function getCreatedAtAttribute() {
        $utils = resolve('App\Services\UtilsService');

        return $utils->formatDate($this->attributes['created_at']);
    }

    public function getUpdatedAtAttribute() {
        $utils = resolve('App\Services\UtilsService');

        return $this->attributes['updated_at'] ? $utils->formatDate($this->attributes['updated_at']) : '-';
    }
}
