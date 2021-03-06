<?php

namespace App;

use App\Services\UtilsService;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

use Spatie\Activitylog\Traits\LogsActivity;

class Device extends Model
{
    use SoftDeletes;
    use LogsActivity;

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'screen_width', 'screen_height', 'ppi', 'brand_id', 'platform_id'
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
     * Get the ratings that belong to this device.
     */
    public function ratings() {
        return $this->hasMany('App\Rating', 'device_id', 'id');
    }

    /**
     * Get the device's brand.
     */
     public function brand() {
        return $this->belongsTo('App\Brand', 'brand_id', 'id');
    }

    /**
     * For Voyager's CRUD.
     */
    public function brandId() {
        return $this->belongsTo('App\Brand', 'brand_id', 'id');
    }

    /**
     * Get the platform the device belongs to.
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
     * Get the app users that belong to the device.
     */
    public function appusers() {
        return $this->belongsToMany('App\AppUser');
    }

    public function getCreatedAtAttribute(String $value) {
        return UtilsService::formatDate($value);
    }

    public function getUpdatedAtAttribute($value) {
        return $value ? UtilsService::formatDate($value) : '';
    }
}
