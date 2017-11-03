<?php

namespace App\Policies;

use App\Policies\BasePolicy;
use TCG\Voyager\Contracts\User as UserType;

class RatingPolicy extends BasePolicy
{
    public function before($user, $ability) {
        if ($user->role()->pluck('name')->get('name') === 'admin' && (
            $ability === 'read' || $ability === 'delete'
        )) {
            return true;
        }

        return false;
    }

    protected function checkApp(UserType $user, $model) {
        $userApps = $user->apps()->pluck('id')->toArray();

        return in_array($model->app_id, $userApps);
    }
}