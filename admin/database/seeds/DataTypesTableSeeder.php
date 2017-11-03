<?php

use Illuminate\Database\Seeder;
use TCG\Voyager\Models\DataType;

class DataTypesTableSeeder extends Seeder
{
    /**
     * Auto generated seed file.
     */
    public function run()
    {
        $dataType = $this->dataType('slug', 'users');

        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'users',
                'display_name_singular' => 'User',
                'display_name_plural'   => 'Users',
                'icon'                  => 'voyager-person',
                'model_name'            => 'App\\User',
                'controller'            => '\\App\\Http\\Controllers\\Controller',
                'policy_name'           => '\\App\\Policies\\UserPolicy',
                'generate_permissions'  => 1,
                'description'           => 'Personal',
                'server_side'           => 0
            ])->save();
        }

        $dataType = $this->dataType('slug', 'apps');

        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'apps',
                'display_name_singular' => 'App',
                'display_name_plural'   => 'Apps',
                'icon'                  => 'voyager-categories',
                'model_name'            => 'App\\App',
                'controller'            => '\\App\\Http\\Controllers\\Controller',
                'policy_name'           => '\\App\\Policies\\AppPolicy',
                'generate_permissions'  => 1,
                'description'           => 'Aplicaciones',
                'server_side'           => 0
            ])->save();
        }

        $dataType = $this->dataType('slug', 'appusers');

        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'appusers',
                'display_name_singular' => 'App User',
                'display_name_plural'   => 'App Users',
                'icon'                  => 'voyager-people',
                'model_name'            => 'App\\AppUser',
                'controller'            => '\\App\\Http\\Controllers\\Controller',
                'policy_name'           => '\\App\\Policies\\AppUserPolicy',
                'generate_permissions'  => 1,
                'description'           => 'Usuarios de las aplicaciones',
                'server_side'           => 0
            ])->save();
        }

        $dataType = $this->dataType('slug', 'brands');

        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'brands',
                'display_name_singular' => 'Brand',
                'display_name_plural'   => 'Brands',
                'icon'                  => 'voyager-tag',
                'model_name'            => 'App\\Brand',
                'controller'            => '\\App\\Http\\Controllers\\Controller',
                'policy_name'           => '\\App\\Policies\\BrandPolicy',
                'generate_permissions'  => 1,
                'description'           => 'Marcas de los dispositivos',
                'server_side'           => 0
            ])->save();
        }

        $dataType = $this->dataType('slug', 'browsers');

        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'browsers',
                'display_name_singular' => 'Browser',
                'display_name_plural'   => 'Browsers',
                'icon'                  => 'voyager-browser',
                'model_name'            => 'App\\Browser',
                'controller'            => '\\App\\Http\\Controllers\\Controller',
                'generate_permissions'  => 1,
                'description'           => 'Navegadores',
                'server_side'           => 0
            ])->save();
        }

        $dataType = $this->dataType('slug', 'devices');

        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'devices',
                'display_name_singular' => 'Device',
                'display_name_plural'   => 'Devices',
                'icon'                  => 'voyager-phone',
                'model_name'            => 'App\\Device',
                'controller'            => '\\App\\Http\\Controllers\\DataTablesController',
                'policy_name'           => '\\App\\Policies\\DevicePolicy',
                'generate_permissions'  => 1,
                'description'           => 'Dispositivos móviles',
                'server_side'           => 0
            ])->save();
        }

        $dataType = $this->dataType('slug', 'messages');

        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'messages',
                'display_name_singular' => 'Message',
                'display_name_plural'   => 'Messages',
                'icon'                  => 'voyager-chat',
                'model_name'            => 'App\\Message',
                'controller'            => '\\App\\Http\\Controllers\\MessagesController',
                'policy_name'           => '\\App\\Policies\\MessagePolicy',
                'generate_permissions'  => 1,
                'description'           => 'Mensajes',
                'server_side'           => 0
            ])->save();
        }

        $dataType = $this->dataType('slug', 'platforms');

        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'platforms',
                'display_name_singular' => 'Platform',
                'display_name_plural'   => 'Platforms',
                'icon'                  => 'voyager-laptop',
                'model_name'            => 'App\\Platform',
                'controller'            => '\\App\\Http\\Controllers\\Controller',
                'policy_name'           => '\\App\\Policies\\PlatformPolicy',
                'generate_permissions'  => 1,
                'description'           => 'Plataformas donde andan las aplicaciones',
                'server_side'           => 0
            ])->save();
        }

        $dataType = $this->dataType('slug', 'ranges');

        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'ranges',
                'display_name_singular' => 'Range',
                'display_name_plural'   => 'Ranges',
                'icon'                  => 'voyager-star-half',
                'model_name'            => 'App\\Range',
                'controller'            => '\\App\\Http\\Controllers\\Controller',
                'policy_name'           => '\\App\\Policies\\RangePolicy',
                'generate_permissions'  => 1,
                'description'           => 'Rangos de calificaciones',
                'server_side'           => 0
            ])->save();
        }

        $dataType = $this->dataType('slug', 'ratings');

        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'ratings',
                'display_name_singular' => 'Rating',
                'display_name_plural'   => 'Ratings',
                'icon'                  => 'voyager-star-two',
                'model_name'            => 'App\\Rating',
                'controller'            => '\\App\\Http\\Controllers\\DataTablesController',
                'policy_name'           => '\\App\\Policies\\RatingPolicy',
                'generate_permissions'  => 1,
                'description'           => 'Calificaciones',
                'server_side'           => 0
            ])->save();
        }

        $dataType = $this->dataType('slug', 'menus');

        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'menus',
                'display_name_singular' => 'Menu',
                'display_name_plural'   => 'Menus',
                'icon'                  => 'voyager-list',
                'model_name'            => 'TCG\\Voyager\\Models\\Menu',
                'controller'            => '',
                'generate_permissions'  => 1,
                'description'           => 'Menús',
                'server_side'           => 0
            ])->save();
        }

        $dataType = $this->dataType('slug', 'roles');

        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'roles',
                'display_name_singular' => 'Role',
                'display_name_plural'   => 'Roles',
                'icon'                  => 'voyager-lock',
                'model_name'            => 'TCG\\Voyager\\Models\\Role',
                'controller'            => '',
                'generate_permissions'  => 1,
                'description'           => '',
                'server_side'           => 0
            ])->save();
        }
    }

    /**
     * [dataType description].
     *
     * @param [type] $field [description]
     * @param [type] $for   [description]
     *
     * @return [type] [description]
     */
    protected function dataType($field, $for)
    {
        return DataType::firstOrNew([$field => $for]);
    }
}
