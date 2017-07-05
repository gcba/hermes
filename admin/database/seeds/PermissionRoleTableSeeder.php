<?php

use Illuminate\Database\Seeder;
use TCG\Voyager\Models\Permission;
use TCG\Voyager\Models\Role;

class PermissionRoleTableSeeder extends Seeder
{
    /**
     * Auto generated seed file.
     *
     * @return void
     */
    public function run()
    {
        $admin = Role::where('name', 'admin')->firstOrFail();
        $supervisor = Role::where('name', 'supervisor')->firstOrFail();
        $support = Role::where('name', 'support')->firstOrFail();
        $user = Role::where('name', 'user')->firstOrFail();

        $allPermissions = Permission::all();

        $adminCannot = [
            'add_appusers',
            'add_ratings',
            'add_devices',
            'add_browsers',
            'add_brands',
            'edit_appusers',
            'edit_ratings',
            'edit_messages',
            'edit_ranges',
            'edit_devices',
            'edit_browsers',
            'edit_brands'
        ];
        $supervisorCannot = [
            'delete_appusers',
            'delete_ratings',
            'delete_ranges',
            'delete_platforms',
            'delete_browsers',
            'delete_devices',
            'delete_brands',
            'browse_settings',
            'read_settings',
            'edit_settings',
            'add_settings',
            'delete_settings',
            'browse_menus',
            'read_menus',
            'edit_menus',
            'add_menus',
            'delete_menus',
            'browse_roles',
            'read_roles',
            'edit_roles',
            'add_roles',
            'delete_roles'
        ];
        $supportCannot = [
            'add_apps',
            'add_users',
            'add_ranges',
            'add_platforms',
            'edit_apps',
            'edit_users',
            'edit_platforms',
            'delete_apps',
            'browse_users',
            'read_users',
            'delete_users'
        ];
        $userCannot = [
            'add_messages',
            'delete_messages'
        ];

        // Set admin permissions
        $adminPermissions = $allPermissions->filter(function ($value) use ($adminCannot) {
            return !in_array($value->key, $adminCannot);
        });

        $admin->permissions()->sync(
            $adminPermissions->pluck('id')->all()
        );

        // Set supervisor permissions
        $supervisorPermissions = $adminPermissions->filter(function ($value) use ($supervisorCannot) {
            return !in_array($value->key, $supervisorCannot);
        });

        $supervisor->permissions()->sync(
            $supervisorPermissions->pluck('id')->all()
        );

        // Set support permissions
        $supportPermissions = $supervisorPermissions->filter(function ($value) use ($supportCannot) {
            return !in_array($value->key, $supportCannot);
        });

        $support->permissions()->sync(
            $supportPermissions->pluck('id')->all()
        );

        // Set user permissions
        $userPermissions = $supportPermissions->filter(function ($value) use ($userCannot) {
            return !in_array($value->key, $userCannot);
        });

        $user->permissions()->sync(
            $userPermissions->pluck('id')->all()
        );
    }
}
