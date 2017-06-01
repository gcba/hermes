<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Config extends Model
{
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'key', 'value', 'modified_by'
    ];

    /**
     * Get the user that last modified the config.
     */
     public function modifiedBy() {
        return $this->belongsTo('App\User', 'modified_by');
     }
}
