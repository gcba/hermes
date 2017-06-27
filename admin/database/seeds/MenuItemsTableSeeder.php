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

            $adminMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Dashboard',
                'url'        => route('voyager.dashboard', [], false),
            ]);
            if (!$adminMenuItem->exists) {
                $adminMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-boat',
                    'color'      => null,
                    'parent_id'  => null,
                    'order'      => 1,
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

            $adminMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Usuarios',
                'url'        => route('voyager.users.index', [], false),
            ]);
            if (!$adminMenuItem->exists) {
                $adminMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-person',
                    'color'      => null,
                    'parent_id'  => null,
                    'order'      => 3,
                ])->save();
            }

            $adminMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Aplicaciones',
                'route'      => 'voyager.apps.index',
                'url'        => null,
                'parameters' => null,
            ]);
            if (!$adminMenuItem->exists) {
                $adminMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-categories',
                    'color'      => null,
                    'parent_id'  => null,
                    'order'      => 4,
                ])->save();
            }

            $adminMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Plataformas',
                'route'      => 'voyager.platforms.index',
                'url'        => null,
                'parameters' => null,
            ]);
            if (!$adminMenuItem->exists) {
                $adminMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-laptop',
                    'color'      => null,
                    'parent_id'  => null,
                    'order'      => 5,
                ])->save();
            }

            $adminMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Rangos',
                'route'      => 'voyager.ranges.index',
                'url'        => null,
                'parameters' => null,
            ]);
            if (!$adminMenuItem->exists) {
                $adminMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-star-half',
                    'color'      => null,
                    'parent_id'  => null,
                    'order'      => 6,
                ])->save();
            }

            $adminMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Roles',
                'url'        => route('voyager.roles.index', [], false),
            ]);
            if (!$adminMenuItem->exists) {
                $adminMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-lock',
                    'color'      => null,
                    'parent_id'  => null,
                    'order'      => 7,
                ])->save();
            }

            $toolsMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Herramientas',
                'url'        => '',
            ]);
            if (!$toolsMenuItem->exists) {
                $toolsMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-tools',
                    'color'      => null,
                    'parent_id'  => null,
                    'order'      => 8,
                ])->save();
            }

            $adminMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Menús',
                'url'        => route('voyager.menus.index', [], false),
            ]);
            if (!$adminMenuItem->exists) {
                $adminMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-list',
                    'color'      => null,
                    'parent_id'  => $toolsMenuItem->id,
                    'order'      => 9,
                ])->save();
            }

            $adminMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'DB',
                'url'        => route('voyager.database.index', [], false),
            ]);
            if (!$adminMenuItem->exists) {
                $adminMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-data',
                    'color'      => null,
                    'parent_id'  => $toolsMenuItem->id,
                    'order'      => 10,
                ])->save();
            }

            $adminMenuItem = MenuItem::firstOrNew([
                'menu_id'    => $adminMenu->id,
                'title'      => 'Configuración',
                'url'        => route('voyager.settings.index', [], false),
            ]);
            if (!$adminMenuItem->exists) {
                $adminMenuItem->fill([
                    'target'     => '_self',
                    'icon_class' => 'voyager-settings',
                    'color'      => null,
                    'parent_id'  => null,
                    'order'      => 11,
                ])->save();
            }
        }
    }
}
