<?php

use Illuminate\Database\Seeder;
use TCG\Voyager\Models\Menu;
use TCG\Voyager\Models\MenuItem;

class MenuItemsTableSeeder extends Seeder
{
    /**
     * Auto generated seed file.
     *
     * @return void
     */
    public function run()
    {
        if (file_exists(base_path('routes/web.php'))) {
            require base_path('routes/web.php');

            $adminMenu = Menu::where('name', 'admin')->firstOrFail();

            $dashboardMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Dashboard',
                'url'        => route('voyager.dashboard', [], false),
                'order'      => 1,
            ]);
            if (!$dashboardMenuItem->exists) {
                $dashboardMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-bar-chart',
                    'color'      => null,
                    'parent_id'  => null,
                ])->save();
            }

            /*
            $adminMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Media',
                'url'        => route('voyager.media.index', [], false),
            ]);
            if (!$adminMenuItem->exists) {
                $adminMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-images',
                    'color'      => null,
                    'parent_id'  => null,
                    'order'      => 5,
                ])->save();
            }
            */

            $ratingsMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Calificaciones',
                'route'      => 'voyager.ratings.index',
                'order'      => 2,
                'url'        => null,
                'parameters' => null,
            ]);
            if (!$ratingsMenuItem->exists) {
                $ratingsMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-star-two',
                    'color'      => null,
                    'parent_id'  => null,
                ])->save();
            }

            $messagesMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Mensajes',
                'route'      => 'voyager.messages.index',
                'order'      => 3,
                'url'        => null,
                'parameters' => null,
            ]);
            if (!$messagesMenuItem->exists) {
                $messagesMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-chat',
                    'color'      => null,
                    'parent_id'  => null,
                ])->save();
            }

            $appusersMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Usuarios',
                'route'      => 'voyager.appusers.index',
                'order'      => 4,
                'url'        => null,
                'parameters' => null,
            ]);
            if (!$appusersMenuItem->exists) {
                $appusersMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-people',
                    'color'      => null,
                    'parent_id'  => null,
                ])->save();
            }

            $appsMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Aplicaciones',
                'route'      => 'voyager.apps.index',
                'order'      => 5,
                'url'        => null,
                'parameters' => null,
            ]);
            if (!$appsMenuItem->exists) {
                $appsMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-categories',
                    'color'      => null,
                    'parent_id'  => null,
                ])->save();
            }

            $rangesMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Rangos',
                'route'      => 'voyager.ranges.index',
                'order'      => 6,
                'url'        => null,
                'parameters' => null,
            ]);
            if (!$rangesMenuItem->exists) {
                $rangesMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-star-half',
                    'color'      => null,
                    'parent_id'  => null,
                ])->save();
            }

            $contextMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Contexto',
                'order'      => 7,
                'url'        => null,
                'parameters' => null,
            ]);
            if (!$contextMenuItem->exists) {
                $contextMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-world',
                    'color'      => null,
                    'parent_id'  => null,
                ])->save();
            }

            $platformsMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Plataformas',
                'route'      => 'voyager.platforms.index',
                'parent_id'  => $contextMenuItem->id,
                'order'      => 1,
                'url'        => null,
                'parameters' => null,
            ]);
            if (!$platformsMenuItem->exists) {
                $platformsMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-laptop',
                    'color'      => null,
                ])->save();
            }

            $devicesMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Dispositivos',
                'route'      => 'voyager.devices.index',
                'parent_id'  => $contextMenuItem->id,
                'order'      => 2,
                'url'        => null,
                'parameters' => null,
            ]);
            if (!$devicesMenuItem->exists) {
                $devicesMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-phone',
                    'color'      => null,
                ])->save();
            }

            $brandsMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Marcas',
                'route'      => 'voyager.brands.index',
                'parent_id'  => $contextMenuItem->id,
                'order'      => 3,
                'url'        => null,
                'parameters' => null,
            ]);
            if (!$brandsMenuItem->exists) {
                $brandsMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-tag',
                    'color'      => null,
                ])->save();
            }

            $browsersMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Navegadores',
                'route'      => 'voyager.browsers.index',
                'parent_id'  => $contextMenuItem->id,
                'order'      => 4,
                'url'        => null,
                'parameters' => null,
            ]);
            if (!$browsersMenuItem->exists) {
                $browsersMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-browser',
                    'color'      => null,
                ])->save();
            }

            $administrationMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Administración',
                'order'      => 12,
                'url'        => null,
                'parameters' => null,
            ]);
            if (!$administrationMenuItem->exists) {
                $administrationMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-settings',
                    'color'      => null,
                    'parent_id'  => null,
                ])->save();
            }

            $usersMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Personal',
                'url'        => route('voyager.users.index', [], false),
                'parent_id'  => $administrationMenuItem->id,
                'order'      => 1,
                'parameters' => null,
            ]);
            if (!$usersMenuItem->exists) {
                $usersMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-person',
                    'color'      => null,
                ])->save();
            }

            $rolesMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Roles',
                'url'        => route('voyager.roles.index', [], false),
                'parent_id'  => $administrationMenuItem->id,
                'order'      => 2,
                'parameters' => null,
            ]);
            if (!$rolesMenuItem->exists) {
                $rolesMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-lock',
                    'color'      => null,
                ])->save();
            }

            $menusMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Menús',
                'url'        => route('voyager.menus.index', [], false),
                'parent_id'  => $administrationMenuItem->id,
                'order'      => 3,
                'parameters' => null,
            ]);
            if (!$menusMenuItem->exists) {
                $menusMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-list',
                    'color'      => null,
                ])->save();
            }

            $dbMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'DB',
                'url'        => route('voyager.database.index', [], false),
                'parent_id'  => $administrationMenuItem->id,
                'order'      => 4,
                'parameters' => null,
            ]);
            if (!$dbMenuItem->exists) {
                $dbMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-data',
                    'color'      => null,
                ])->save();
            }

            $configMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Configuración',
                'url'        => route('voyager.settings.index', [], false),
                'parent_id'  => $administrationMenuItem->id,
                'order'      => 5,
                'parameters' => null,
            ]);
            if (!$configMenuItem->exists) {
                $configMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-tools',
                    'color'      => null,
                ])->save();
            }
        }
    }
}
