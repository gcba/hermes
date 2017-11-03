<?php

namespace App\Policies;

use App\Policies\BasePolicy;
use TCG\Voyager\Contracts\User as UserType;

class AppPolicy extends BasePolicy
{
    public function before($user, $ability) {
        $role = $user->role()->pluck('name')->get('name');

        if ($role === 'admin' || $role === 'supervisor') {
            return true;
        }
    }

    protected function checkApp(UserType $user, $model) {
        return false;
    }
}