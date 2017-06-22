<?php

namespace App;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

class Range extends Model
{
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'from', 'to', 'key', 'app_id'
    ];

    /**
     * Get the app this range belongs to.
     */
    public function app()
    {
        return $this->belongsTo('App\App');
    }
}
