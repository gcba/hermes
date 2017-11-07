<?php

use App\Role;
use Illuminate\Database\Seeder;
use TCG\Voyager\Models\Permission;

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
            // Add
            'add_appusers',
            'add_ratings',
            'add_devices',
            'add_browsers',
            'add_brands',
            // Edit
            'edit_appusers',
            'edit_ratings',
            'edit_messages',
            'edit_ranges',
            'edit_devices',
            'edit_browsers',
            'edit_brands'
        ];

        $supervisorCannot = [
            // Browse
            'browse_compass',
            'browse_database',
            'browse_hooks',
            'browse_settings',
            'browse_menus',
            'browse_roles',
            // Read
            'read_database',
            'read_hooks',
            'read_settings',
            'read_menus',
            'read_roles',
            // Add
            'add_database',
            'add_hooks',
            'add_settings',
            'add_menus',
            'add_roles',
            // Edit
            'edit_database',
            'edit_hooks',
            'edit_settings',
            'edit_menus',
            'edit_roles',
            // Delete
            'delete_appusers',
            'delete_ratings',
            'delete_ranges',
            'delete_platforms',
            'delete_browsers',
            'delete_devices',
            'delete_brands',
            'delete_database',
            'delete_hooks',
            'delete_settings',
            'delete_menus',
            'delete_roles'
        ];

        $supportCannot = [
            // Browse
            'browse_users',
            // Read
            'read_users',
            // Add
            'add_apps',
            'add_users',
            'add_ranges',
            'add_platforms',
            // Edit
            'edit_apps',
            'edit_users',
            'edit_platforms',
            // Delete
            'delete_apps',
            'delete_users'
        ];

        $userCannot = [
            // Add
            'add_messages',
            // Delete
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
