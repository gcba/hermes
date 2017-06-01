<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class App extends Model
{
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'type', 'modified_by'
    ];

    /**
     * Get the user that last modified the App.
     */
     public function modifiedBy() {
        return $this->belongsTo('App\User', 'modified_by');
     }
}
