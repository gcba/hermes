<?php

namespace App;

use App\Services\UtilsService;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

use Spatie\Activitylog\Traits\LogsActivity;

class Rating extends Model
{
    use SoftDeletes;
    use LogsActivity;

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'rating',
        'description',
        'app_version',
        'platform_version',
        'browser_version',
        'has_message',
        'app_id',
        'platform_id',
        'range_id',
        'appuser_id',
        'device_id',
        'browser_id'
    ];

    /**
     * Log all fillable attributes.
     *
     * @var bool
     */
    protected static $logFillable = true;

    /**
     * The attributes that should be mutated to dates.
     *
     * @var array
     */
    protected $dates = ['created_at', 'updated_at', 'deleted_at'];

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
     * Get the messages that belong to this rating.
     */
    public function messages() {
        return $this->hasMany('App\Message', 'rating_id', 'id');
    }

    /**
     * Get the app the rating belongs to.
     */
    public function app() {
        return $this->belongsTo('App\App', 'app_id', 'id');
    }

    /**
     * For Voyager's CRUD.
     */
    public function appId() {
        return $this->belongsTo('App\App', 'app_id', 'id');
    }

    /**
     * Get the range the rating belongs to.
     */
    public function range() {
        return $this->belongsTo('App\Range', 'range_id', 'id');
    }

     /**
     * For Voyager's CRUD.
     */
    public function rangeId() {
        return $this->belongsTo('App\Range', 'range_id', 'id');
    }

    /**
     * Get the app user the rating belongs to.
     */
    public function appuser() {
        return $this->belongsTo('App\AppUser', 'appuser_id', 'id');
    }

    /**
     * For Voyager's CRUD.
     */
    public function appuserId() {
        return $this->belongsTo('App\AppUser', 'appuser_id', 'id');
    }

    /**
     * Get the platform the rating belongs to.
     */
    public function platform() {
        return $this->belongsTo('App\Platform', 'platform_id', 'id');
    }

    /**
     * For Voyager's CRUD.
     */
    public function platformId() {
        return $this->belongsTo('App\Platform', 'platform_id', 'id');
    }

    /**
     * Get the device the rating belongs to.
     */
    public function device() {
        return $this->belongsTo('App\Device', 'device_id', 'id');
    }

    /**
     * For Voyager's CRUD.
     */
    public function deviceId() {
        return $this->belongsTo('App\Device', 'device_id', 'id');
    }

    /**
     * Get the browser the rating belongs to.
     */
    public function browser() {
        return $this->belongsTo('App\Browser', 'browser_id', 'id');
    }

    /**
     * For Voyager's CRUD.
     */
    public function browserId() {
        return $this->belongsTo('App\Browser', 'browser_id', 'id');
    }

    public function latestMessage() {
        return $this->hasOne('App\Message', 'rating_id', 'id')->latest();
    }

    public function getHasMessageAttribute(Bool $value) {
        return $value ? 'Sí' : 'No';
    }

    public function getCreatedAtAttribute(String $value) {
        return UtilsService::formatDate($value);
    }

    public function getUpdatedAtAttribute($value) {
        return $value ? UtilsService::formatDate($value) : '';
    }
}
