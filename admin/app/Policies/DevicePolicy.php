<?php

namespace App\Policies;

use App\Policies\BasePolicy;
use TCG\Voyager\Contracts\User as UserType;

class DevicePolicy extends BasePolicy
{
    public function before($user, $ability) {
        $role = $user->role()->pluck('name')[0];

        if (($role === 'admin' && ($ability === 'browse' || $ability === 'read' || $ability === 'delete')) ||
        ($role === 'supervisor' && ($ability === 'browse' || $ability === 'read'))) {
            return true;
        }

        return false;
    }

    protected function checkApp(UserType $user, $model) {
        return true;
    }
}