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
        'name', 'type', 'key', 'modified_by'
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
    protected static function boot()
    {
        parent::boot();

        static::creating(function ($model)
        {
            if (!$model->key) {
                $model->attributes['key'] = md5(date("Y-m-d H:i:s"));
            }

            $model->attributes['name'] = sanitizeName($model->name);
        });

        static::updating(function ($model)
        {
            $model->attributes['name'] = sanitizeName($model->name);
        });
    }

    /**
     * Get the ratings that belong to this app.
     */
    public function ratings()
    {
        return $this->hasMany('App\Rating', 'rating_id', 'id');
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
     * Get the user that last modified the app.
     */
     public function modifiedBy() {
        return $this->belongsTo('App\User', 'modified_by', 'id');
     }

     private function sanitizeName($name) {
         return filter_var(trim($name), FILTER_SANITIZE_SPECIAL_CHARS);
    }
}
