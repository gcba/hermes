<?php

namespace App\Policies;

use App\Policies\BasePolicy;
use TCG\Voyager\Contracts\User as UserType;

class RangePolicy extends BasePolicy
{
    protected function checkApp(UserType $user, $model) {
        $userApps = $user->apps()->pluck('id')->toArray();
        $rangeApps = $model->ratings()->pluck('app_id')->toArray();

        return count(array_intersect($userApps, $rangeApps)) > 0;
    }
}