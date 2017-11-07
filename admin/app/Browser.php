<?php

namespace App;

use App\Services\UtilsService;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

use Spatie\Activitylog\Traits\LogsActivity;

class Browser extends Model
{
    use SoftDeletes;
    use LogsActivity;

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name'
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
     * Get the ratings that belong to this browser.
     */
    public function ratings() {
        return $this->hasMany('App\Rating', 'browser_id', 'id');
    }

    public function getCreatedAtAttribute(String $value) {
        return UtilsService::formatDate($value);
    }

    public function getUpdatedAtAttribute($value) {
        return $value ? UtilsService::formatDate($value) : '';
    }
}
