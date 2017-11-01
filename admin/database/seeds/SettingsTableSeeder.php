<?php

use Illuminate\Database\Seeder;
use TCG\Voyager\Models\Setting;

class SettingsTableSeeder extends Seeder
{
    /**
     * Auto generated seed file.
     */
    public function run()
    {
        $setting = $this->findSetting('title');

        if (!$setting->exists) {
            $setting->fill([
                'display_name' => 'Site Title',
                'value'        => 'Hermes',
                'details'      => '',
                'type'         => 'text',
                'order'        => 1,
                'group'        => 'Site'
            ])->save();
        }

        $setting = $this->findSetting('description');

        if (!$setting->exists) {
            $setting->fill([
                'display_name' => 'Site Description',
                'value'        => 'Gestión de feedback de las apps de la Ciudad',
                'details'      => '',
                'type'         => 'text',
                'order'        => 2,
                'group'        => 'Site'
            ])->save();
        }

        $setting = $this->findSetting('logo');

        if (!$setting->exists) {
            $setting->fill([
                'display_name' => 'Site Logo',
                'value'        => '',
                'details'      => '',
                'type'         => 'image',
                'order'        => 3,
                'group'        => 'Site'
            ])->save();
        }

        $setting = $this->findSetting('admin_bg_image');

        if (!$setting->exists) {
            $setting->fill([
                'display_name' => 'Admin Background Image',
                'value'        => '',
                'details'      => '',
                'type'         => 'image',
                'order'        => 9,
                'group'        => 'Admin'
            ])->save();
        }

        $setting = $this->findSetting('admin_title');

        if (!$setting->exists) {
            $setting->fill([
                'display_name' => 'Admin Title',
                'value'        => 'Hermes',
                'details'      => '',
                'type'         => 'text',
                'order'        => 4,
                'group'        => 'Admin'
            ])->save();
        }

        $setting = $this->findSetting('admin_description');

        if (!$setting->exists) {
            $setting->fill([
                'display_name' => 'Admin Description',
                'value'        => 'Gestión de feedback de las apps de la Ciudad',
                'details'      => '',
                'type'         => 'text',
                'order'        => 5,
                'group'        => 'Admin'
            ])->save();
        }

        $setting = $this->findSetting('admin_loader');

        if (!$setting->exists) {
            $setting->fill([
                'display_name' => 'Admin Loader',
                'value'        => '',
                'details'      => '',
                'type'         => 'image',
                'order'        => 6,
                'group'        => 'Admin'
            ])->save();
        }

        $setting = $this->findSetting('admin_icon_image');

        if (!$setting->exists) {
            $setting->fill([
                'display_name' => 'Admin Icon Image',
                'value'        => '',
                'details'      => '',
                'type'         => 'image',
                'order'        => 7,
                'group'        => 'Admin'
            ])->save();
        }

        $setting = $this->findSetting('google_analytics_client_id');

        if (!$setting->exists) {
            $setting->fill([
                'display_name' => 'Google Analytics Client ID',
                'value'        => '',
                'details'      => '',
                'type'         => 'text',
                'order'        => 9,
                'group'        => 'Admin'
            ])->save();
        }
    }

    /**
     * [setting description].
     *
     * @param [type] $key [description]
     *
     * @return [type] [description]
     */
    protected function findSetting($key)
    {
        return Setting::firstOrNew(['key' => $key]);
    }
}
