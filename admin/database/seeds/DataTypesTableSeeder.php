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
                'slug'                  => 'users',
                'icon'                  => 'voyager-person',
                'model_name'            => 'TCG\\Voyager\\Models\\User',
                'controller'            => 'Controller',
                'generate_permissions'  => 0,
                'description'           => 'Personal',
                'server_side'           => 1
            ])->save();
        }

        $dataType = $this->dataType('slug', 'apps');
        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'apps',
                'display_name_singular' => 'App',
                'display_name_plural'   => 'Apps',
                'slug'                  => 'apps',
                'icon'                  => 'voyager-categories',
                'model_name'            => 'App\\App',
                'controller'            => 'Controller',
                'generate_permissions'  => 0,
                'description'           => 'Aplicaciones',
                'server_side'           => 1
            ])->save();
        }

        $dataType = $this->dataType('slug', 'appusers');
        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'appusers',
                'display_name_singular' => 'App User',
                'display_name_plural'   => 'App Users',
                'slug'                  => 'appusers',
                'icon'                  => 'voyager-people',
                'model_name'            => 'App\\AppUser',
                'controller'            => 'Controller',
                'generate_permissions'  => 0,
                'description'           => 'Usuarios de las aplicaciones',
                'server_side'           => 1
            ])->save();
        }

        $dataType = $this->dataType('slug', 'brands');
        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'brands',
                'display_name_singular' => 'Brand',
                'display_name_plural'   => 'Brands',
                'slug'                  => 'brands',
                'icon'                  => 'voyager-tag',
                'model_name'            => 'App\\Brand',
                'controller'            => 'Controller',
                'generate_permissions'  => 0,
                'description'           => 'Marcas de los dispositivos',
                'server_side'           => 1
            ])->save();
        }

        $dataType = $this->dataType('slug', 'browsers');
        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'browsers',
                'display_name_singular' => 'Browser',
                'display_name_plural'   => 'Browsers',
                'slug'                  => 'browsers',
                'icon'                  => 'voyager-browser',
                'model_name'            => 'App\\Browser',
                'controller'            => 'Controller',
                'generate_permissions'  => 0,
                'description'           => 'Navegadores',
                'server_side'           => 1
            ])->save();
        }

        $dataType = $this->dataType('slug', 'devices');
        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'devices',
                'display_name_singular' => 'Device',
                'display_name_plural'   => 'Devices',
                'slug'                  => 'devices',
                'icon'                  => 'voyager-phone',
                'model_name'            => 'App\\Device',
                'controller'            => 'Controller',
                'generate_permissions'  => 0,
                'description'           => 'Dispositivos móviles',
                'server_side'           => 1
            ])->save();
        }

        $dataType = $this->dataType('slug', 'messages');
        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'messages',
                'display_name_singular' => 'Message',
                'display_name_plural'   => 'Messages',
                'slug'                  => 'messages',
                'icon'                  => 'voyager-chat',
                'model_name'            => 'App\\Message',
                'controller'            => 'Controller',
                'generate_permissions'  => 0,
                'description'           => 'Mensajes',
                'server_side'           => 1
            ])->save();
        }

        $dataType = $this->dataType('slug', 'platforms');
        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'platforms',
                'display_name_singular' => 'Platform',
                'display_name_plural'   => 'Platforms',
                'slug'                  => 'platforms',
                'icon'                  => 'voyager-laptop',
                'model_name'            => 'App\\Platform',
                'controller'            => 'Controller',
                'generate_permissions'  => 0,
                'description'           => 'Plataformas donde andan las aplicaciones',
                'server_side'           => 1
            ])->save();
        }

        $dataType = $this->dataType('slug', 'ranges');
        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'ranges',
                'display_name_singular' => 'Range',
                'display_name_plural'   => 'Ranges',
                'slug'                  => 'ranges',
                'icon'                  => 'voyager-star-half',
                'model_name'            => 'App\\Range',
                'controller'            => 'Controller',
                'generate_permissions'  => 0,
                'description'           => 'Rangos de calificaciones',
                'server_side'           => 1
            ])->save();
        }

        $dataType = $this->dataType('slug', 'ratings');
        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'ratings',
                'display_name_singular' => 'Rating',
                'display_name_plural'   => 'Ratings',
                'slug'                  => 'ratings',
                'icon'                  => 'voyager-star-two',
                'model_name'            => 'App\\Rating',
                'controller'            => 'Controller',
                'generate_permissions'  => 0,
                'description'           => 'Calificaciones',
                'server_side'           => 1
            ])->save();
        }

        $dataType = $this->dataType('slug', 'menus');
        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'menus',
                'display_name_singular' => 'Menu',
                'display_name_plural'   => 'Menus',
                'slug'                  => 'menus',
                'icon'                  => 'voyager-list',
                'model_name'            => 'TCG\\Voyager\\Models\\Menu',
                'controller'            => 'Controller',
                'generate_permissions'  => 0,
                'description'           => 'Menús',
                'server_side'           => 1
            ])->save();
        }

        $dataType = $this->dataType('slug', 'roles');
        if (!$dataType->exists) {
            $dataType->fill([
                'name'                  => 'roles',
                'display_name_singular' => 'Role',
                'display_name_plural'   => 'Roles',
                'slug'                  => 'roles',
                'icon'                  => 'voyager-lock',
                'model_name'            => 'TCG\\Voyager\\Models\\Role',
                'controller'            => 'Controller',
                'generate_permissions'  => 0,
                'description'           => 'Roles',
                'server_side'           => 1
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
