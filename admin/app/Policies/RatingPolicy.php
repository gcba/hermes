<?php

namespace App\Policies;

use App\User;
use App\Rating;
use TCG\Voyager\Policies\BasePolicy;
use \TCG\Voyager\Contracts\User as UserType;

class RatingPolicy extends BasePolicy
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

    private function getUser(Rating $model) {
        $appId = $model->app()->select('id')->pluck('id');

        return User::select('id')->whereHas('apps', function ($query) use($appId) {
            $query->where('id', $appId);
        })->get()
        ->pluck('id')
        ->toArray();
    }

    protected function checkPermission(UserType $user, $model, $action) {
        $users = $this->getUser($model);

        return in_array($user->id, $users) || parent::checkPermission($user, $model, $action);
    }
}