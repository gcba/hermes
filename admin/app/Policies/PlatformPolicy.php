<?php

namespace App\Policies;

use App\Policies\BasePolicy;
use TCG\Voyager\Contracts\User as UserType;

class PlatformPolicy extends BasePolicy
{
    public function before($user, $ability) {
        $role = $user->role()->pluck('name')->get('name');

        if ($role === 'admin') {
            return true;
        }
    }

    protected function checkApp(UserType $user, $model) {
        $userApps = $user->apps()->pluck('id')->toArray();
        $platformApps = $model->apps()->pluck('id')->toArray();

        return count(array_intersect($userApps, $platformApps)) > 0;
    }
}