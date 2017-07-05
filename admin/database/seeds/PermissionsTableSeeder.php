<?php

use Illuminate\Database\Seeder;
use TCG\Voyager\Models\Permission;

class PermissionsTableSeeder extends Seeder
{
    /**
     * Auto generated seed file.
     */
    public function run()
    {
        $keys = [
            'browse_admin',
            'browse_database',
            // 'browse_media',
            'browse_settings',
            'browse_apps',
            'browse_appusers',
            'browse_apps',
            'browse_platforms',
            'browse_brands',
            'browse_browsers',
            'browse_devices',
            'browse_ranges',
            'browse_messages',
            'browse_ratings',
        ];

        foreach ($keys as $key) {
            Permission::firstOrCreate([
                'key'        => $key,
                'table_name' => null,
            ]);
        }

        Permission::generateFor('menus');

        Permission::generateFor('roles');

        Permission::generateFor('users');

        Permission::generateFor('appusers');

        Permission::generateFor('apps');

        Permission::generateFor('platforms');

        Permission::generateFor('brands');

        Permission::generateFor('browsers');

        Permission::generateFor('devices');

        Permission::generateFor('ranges');

        Permission::generateFor('messages');

        Permission::generateFor('ratings');
    }
}
