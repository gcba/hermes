<?php

namespace App;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

class Rating extends Model
{
    use SoftDeletes;

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
     * The attributes that should be mutated to dates.
     *
     * @var array
     */
    protected $dates = ['deleted_at'];

    /**
     * Get the app the rating belongs to.
     */
    public function app()
    {
        return $this->belongsTo('App\App');
    }

    /**
     * Get the range the rating belongs to.
     */
    public function rating()
    {
        return $this->belongsTo('App\Rating');
    }

    /**
     * Get the app user the rating belongs to.
     */
    public function appUser()
    {
        return $this->belongsTo('App\AppUser', 'appuser_id');
    }

    /**
     * Get the platform the rating belongs to.
     */
    public function platform()
    {
        return $this->belongsTo('App\Platform');
    }

    /**
     * Get the device the rating belongs to.
     */
    public function device()
    {
        return $this->belongsTo('App\Device');
    }

    /**
     * Get the browser the rating belongs to.
     */
    public function browser()
    {
        return $this->belongsTo('App\Browser');
    }
}
