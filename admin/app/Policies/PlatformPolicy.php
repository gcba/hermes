<?php

namespace App\Policies;

use App\Policies\BasePolicy;
use TCG\Voyager\Contracts\User as UserType;

class PlatformPolicy extends BasePolicy
{
    public function before($user, $ability) {
        $role = $user->role()->pluck('name')[0];

        if ($role === 'admin' || $role === 'supervisor' &&
        ($ability === 'browse' || $ability === 'read' || $ability === 'create' || $ability === 'edit')) {
            return true;
        }

        return false;
    }

    protected function checkApp(UserType $user, $model) {
        $userApps = $user->apps()->pluck('id')->toArray();
        $platformApps = $model->apps()->pluck('id')->toArray();

        return count(array_intersect($userApps, $platformApps)) > 0;
    }
}