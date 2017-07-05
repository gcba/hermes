<?php

namespace App;

use DateTime;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

class Device extends Model
{
    use SoftDeletes;

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'screen_width', 'screen_height', 'ppi', 'brand_id', 'platform_id'
    ];

    /**
     * The attributes that should be mutated to dates.
     *
     * @var array
     */
    protected $dates = ['deleted_at'];

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
        return $this->hasMany('App\Rating', 'rating_id', 'id');
    }

    /**
     * Get the device's brand.
     */
     public function brand() {
        return $this->belongsTo('App\Brand', 'brand_id', 'id');
     }

     /**
     * Get the platform the device belongs to.
     */
     public function platform() {
        return $this->belongsTo('App\Platform', 'id', 'platform_id');
     }

    /**
     * For Voyager's CRUD.
     */
     public function brandId() {
        return $this->belongsTo('App\Brand', 'brand_id', 'id');
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

     public function getCreatedAtAttribute(){
        return $this->formatDate($this->attributes['created_at']);
    }

    public function getUpdatedAtAttribute(){
        return $this->attributes['updated_at'] ? $this->formatDate($this->attributes['updated_at']) : '';
    }

    private function formatDate($dateString) {
        $date = new DateTime($dateString);

        return $date->format('d/m/Y H:i:s');
    }
}
