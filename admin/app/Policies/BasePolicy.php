<?php

namespace App\Policies;

use App\User;
use TCG\Voyager\Policies\BasePolicy as VoyagerBasePolicy;
use TCG\Voyager\Contracts\User as UserType;

abstract class BasePolicy extends VoyagerBasePolicy
{
    /**
     * Determine if the given model can be viewed by the user.
     *
     * @param \App\User $user
     * @param  $model
     *
     * @return bool
     */
    public function read(User $user, $model)
    {
        return $this->checkPermission($user, $model, 'read');
    }

    /**
     * Determine if the given model can be edited by the user.
     *
     * @param \App\User $user
     * @param  $model
     *
     * @return bool
     */
    public function edit(User $user, $model)
    {
        return $this->checkPermission($user, $model, 'edit');
    }

    /**
     * Determine if the given model can be deleted by the user.
     *
     * @param \App\User $user
     * @param  $model
     *
     * @return bool
     */
    public function delete(User $user, $model)
    {
        return $this->checkPermission($user, $model, 'delete');
    }

    abstract protected function checkApp(UserType $user, $model);

    protected function checkPermission(UserType $user, $model, $action) {
        $userApps = $user->apps()->pluck('id')->toArray();

        return $this->checkApp($user, $model) || parent::checkPermission($user, $model, $action);
    }
}