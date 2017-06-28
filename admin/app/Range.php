<?php

namespace App;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

class Range extends Model
{
    use SoftDeletes;

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'from', 'to', 'key', 'app_id'
    ];

    /**
     * The attributes that should be mutated to dates.
     *
     * @var array
     */
    protected $dates = ['deleted_at'];

    /**
     * Get the app this range belongs to.
     */
    public function app()
    {
        return $this->belongsTo('App\App');
    }

    /**
     * Get a readable range name.
     */
    public function name() {
        return $this->from . " a " . $this->to;
    }
}
