<?php

namespace App;

use Spatie\Activitylog\Traits\LogsActivity;
use Illuminate\Support\Facades\Auth;
use TCG\Voyager\Models\Role as VoyagerRole;

class Role extends VoyagerRole
{
    use LogsActivity;

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'display_name'
    ];

    /**
     * Log all fillable attributes.
     *
     * @var array
     */
    protected static $logFillable = true;

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
     * Get the users that belogn to this role.
     */
    public function users() {
        return $this->hasMany('App\User', 'role_id', 'id');
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