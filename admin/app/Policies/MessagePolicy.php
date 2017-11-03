<?php

namespace App\Policies;

use App\Policies\BasePolicy;
use TCG\Voyager\Contracts\User as UserType;

class MessagePolicy extends BasePolicy
{
    public function before($user, $ability) {
        $role = $user->role()->pluck('name')->get('name');

        if (($role === 'admin' || $role === 'supervisor') &&
        ($ability === 'read' || $ability === 'create' || $ability === 'delete')) {
            return true;
        }

        return false;
    }

    protected function checkApp(UserType $user, $model) {
        $userApps = $user->apps()->pluck('id')->toArray();

        return in_array($model->app_id, $userApps);
    }
}