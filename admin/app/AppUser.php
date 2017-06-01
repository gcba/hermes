<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class AppUser extends Model
{
    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name', 'email',
    ];
}
