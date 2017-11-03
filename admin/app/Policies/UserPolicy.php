<?php

namespace App\Policies;

use App\Policies\BasePolicy;
use TCG\Voyager\Contracts\User as UserType;

class UserPolicy extends BasePolicy
{
    public function before($user, $ability) {
        $role = $user->role()->pluck('name')[0];

        if ($role === 'admin' || $role === 'supervisor') {
            return true;
        }

        return false;
    }

    protected function checkApp(UserType $user, $model) {
        return false;
    }
}