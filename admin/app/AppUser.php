<?php

namespace App;

use DateTime;
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
     * Get the ratings that belong to this app user.
     */
    public function ratings() {
        return $this->hasMany('App\Rating', 'appuser_id', 'id');
    }

    /**
     * Get the apps of the app user.
     */
     public function apps() {
        return $this->belongsToMany('App\App', 'app_user_app');
     }

     /**
     * Get the platform that belong to the appuser.
     */
     public function platforms() {
        return $this->belongsToMany('App\Platform', 'app_user_platform');
     }

     /**
     * Get the devices that belong to the app user.
     */
     public function devices() {
        return $this->belongsToMany('App\Device', 'app_user_device');
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
