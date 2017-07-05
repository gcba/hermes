<?php

namespace App;

use DateTime;
use Illuminate\Notifications\Notifiable;
use Illuminate\Foundation\Auth\User as Authenticatable;

class User extends Authenticatable
{
    use Notifiable;

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'email', 'password', 'modified_by'
    ];

    /**
     * The attributes that should be hidden for arrays.
     *
     * @var array
     */
    protected $hidden = [
        'password', 'remember_token'
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
     * Get the apps the user belongs to.
     */
     public function apps() {
        return $this->belongsToMany('App\App');
     }

     /**
     * Get the user that last modified the app.
     */
     public function modifiedBy() {
        return $this->belongsTo('App\User', 'modified_by', 'id');
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
