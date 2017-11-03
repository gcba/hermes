<?php

namespace App\Policies;

use App\Policies\BasePolicy;
use TCG\Voyager\Contracts\User as UserType;

class MessagePolicy extends BasePolicy
{
    public function before($user, $ability) {
        $role = $user->role()->pluck('name')[0];

        if (($role === 'admin' || $role === 'supervisor') &&
        ($ability === 'browse' || $ability === 'read' || $ability === 'create' || $ability === 'delete')) {
            return true;
        }

        return false;
    }

    protected function checkApp(UserType $user, $model) {
        $userApps = $user->apps()->pluck('id')->toArray();
        $appId = $model->rating()->pluck('app_id')->get('app_id');

        return in_array($appId, $userApps);
    }
}