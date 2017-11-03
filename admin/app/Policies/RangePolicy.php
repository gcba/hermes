<?php

namespace App\Policies;

use App\Policies\BasePolicy;
use TCG\Voyager\Contracts\User as UserType;

class RangePolicy extends BasePolicy
{
    public function before($user, $ability) {
        $role = $user->role()->pluck('name')[0];

        if (($role === 'admin' && ($ability === 'read' || $ability === 'create' || $ability === 'delete')) ||
        ($role === 'supervisor' && ($ability === 'read' || $ability === 'create'))) {
            return true;
        }

        return false;
    }

    protected function checkApp(UserType $user, $model) {
        $userApps = $user->apps()->pluck('id')->toArray();
        $rangeApps = $model->ratings()->pluck('app_id')->toArray();

        return count(array_intersect($userApps, $rangeApps)) > 0;
    }
}