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
                'model_name'            => 'TCG\\Voyager\\Models\\User',
                'controller'            => '',
                'generate_permissions'  => 1,
                'description'           => '',
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
                'controller'            => 'AppController',
                'generate_permissions'  => 1,
                'description'           => 'App',
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
                'controller'            => '',
                'generate_permissions'  => 1,
                'description'           => 'Usuarios de las aplicaciones',
            ])->save();
        }

        $dataType = $this->dataType('slug', 'brands');
        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'brand',
                'display_name_singular' => 'Brand',
                'display_name_plural'   => 'Brands',
                'icon'                  => 'voyager-tag',
                'model_name'            => 'App\\Brand',
                'controller'            => '',
                'generate_permissions'  => 1,
                'description'           => 'Marcas de los dispositivos',
            ])->save();
        }

        $dataType = $this->dataType('slug', 'browsers');
        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'browser',
                'display_name_singular' => 'Browser',
                'display_name_plural'   => 'Browsers',
                'icon'                  => 'voyager-browser',
                'model_name'            => 'App\\Browser',
                'controller'            => '',
                'generate_permissions'  => 1,
                'description'           => 'Navegadores',
            ])->save();
        }

        $dataType = $this->dataType('slug', 'devices');
        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'device',
                'display_name_singular' => 'Device',
                'display_name_plural'   => 'Devices',
                'icon'                  => 'voyager-phone',
                'model_name'            => 'App\\Device',
                'controller'            => '',
                'generate_permissions'  => 1,
                'description'           => 'Dispositivos móviles',
            ])->save();
        }

        $dataType = $this->dataType('slug', 'messages');
        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'message',
                'display_name_singular' => 'Message',
                'display_name_plural'   => 'Messages',
                'icon'                  => 'voyager-chat',
                'model_name'            => 'App\\Message',
                'controller'            => '',
                'generate_permissions'  => 1,
                'description'           => 'Mensajes',
            ])->save();
        }

        $dataType = $this->dataType('slug', 'platforms');
        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'platform',
                'display_name_singular' => 'Platform',
                'display_name_plural'   => 'Platforms',
                'icon'                  => 'voyager-laptop',
                'model_name'            => 'App\\Platform',
                'controller'            => 'PlatformController',
                'generate_permissions'  => 1,
                'description'           => 'Plataformas donde andan las aplicaciones',
            ])->save();
        }

        $dataType = $this->dataType('slug', 'ranges');
        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'range',
                'display_name_singular' => 'Range',
                'display_name_plural'   => 'Ranges',
                'icon'                  => 'voyager-star-half',
                'model_name'            => 'App\\Range',
                'controller'            => '',
                'generate_permissions'  => 1,
                'description'           => 'Rangos de calificaciones',
            ])->save();
        }

        $dataType = $this->dataType('slug', 'ratings');
        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'rating',
                'display_name_singular' => 'Rating',
                'display_name_plural'   => 'Ratings',
                'icon'                  => 'voyager-star-two',
                'model_name'            => 'App\\Rating',
                'controller'            => '',
                'generate_permissions'  => 1,
                'description'           => 'Calificaciones',
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
                'description'           => '',
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
