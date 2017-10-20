<?php

namespace App;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

class App extends Model
{
    use SoftDeletes;

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'type', 'key', 'updated_by'
    ];

    /**
     * The attributes that should be mutated to dates.
     *
     * @var array
     */
    protected $dates = ['deleted_at'];

    protected $utils;

    /**
     * Create a new controller instance.
     *
     * @param  UtilsService $utils
     * @return void
     */
    public function __construct()
    {
        $this->utils = resolve('App\Services\UtilsService');
    }

    /**
     * Boot function for using with User Events
     *
     * @return void
     */
    protected static function boot() {
        parent::boot();

        static::creating(function ($model) {
            if (!$model->key) {
                $model->attributes['key'] = md5(date("Y-m-d H:i:s"));
            }

            $model->attributes['updated_at'] = null;
        });
    }

    /**
     * Get the ratings that belong to this app.
     */
    public function ratings() {
        return $this->hasMany('App\Rating', 'app_id', 'id');
    }

    /**
     * Get the users that belong to the app.
     */
    public function users() {
        return $this->belongsToMany('App\User');
    }

    /**
     * Get the app users of the app.
     */
    public function appusers() {
        return $this->belongsToMany('App\AppUser');
    }

    /**
     * Get the platforms the app is in.
     */
    public function platforms() {
        return $this->belongsToMany('App\Platform');
    }

    /**
     * Get the user that last updated the app.
     */
    public function updatedBy() {
        return $this->belongsTo('App\User', 'updated_by', 'id');
    }

    /**
     * For Voyager's CRUD.
     */
    public function updatedById() {
        return $this->belongsTo('App\User', 'updated_by', 'id');
    }

    public function setNameAttribute($value) {
        $this->attributes['name'] = ucfirst(filter_var(trim($value), FILTER_SANITIZE_SPECIAL_CHARS));
    }

    public function setTypeAttribute($value) {
        $type = strtolower($value);

        switch ($type) {
            case 'móvil': $this->attributes['type'] = 'M';
            case 'web': $this->attributes['type'] = 'W';
        }

        if (strlen($value) == 1) {
            $this->attributes['type'] = $value;
        }
    }

    public function getTypeAttribute() {
        if (array_key_exists('type', $this->attributes)) {
            switch ($this->attributes['type']) {
                case 'M': return 'Móvil';
                case 'W': return 'Web';
            }
        }
    }

    public function getCreatedAtAttribute() {
        return $this->utils->formatDate($this->attributes['created_at']);
    }

    public function getUpdatedAtAttribute() {
        return $this->attributes['updated_at'] ? $this->utils->formatDate($this->attributes['updated_at']) : '-';
    }
}