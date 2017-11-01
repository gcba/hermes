<?php

use Illuminate\Database\Seeder;

class DatabaseSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $this->call(DataTypesTableSeeder::class);
        $this->call(DataRowsTableSeeder::class);

        $this->call(MenusTableSeeder::class);
        $this->call(MenuItemsTableSeeder::class);
        $this->call(RolesTableSeeder::class);
        $this->call(PermissionsTableSeeder::class);
        $this->call(PermissionRoleTableSeeder::class);
        $this->call(TranslationsTableSeeder::class);
        $this->call(SettingsTableSeeder::class);
        $this->call(UsersTableSeeder::class);
        $this->call(AppUsersTableSeeder::class);
        $this->call(Apps_Platforms_UsersSeeder::class);
        $this->call(RangesTableSeeder::class);
        $this->call(BrandsTableSeeder::class);
        $this->call(BrowsersTableSeeder::class);
        $this->call(DevicesTableSeeder::class);
        $this->call(RatingsTableSeeder::class);
        $this->call(MessagesTableSeeder::class);

        $this->call(AppDataRowsTableSeeder::class);
        $this->call(AppUserDataRowsTableSeeder::class);
        $this->call(BrandDataRowsTableSeeder::class);
        $this->call(BrowserDataRowsTableSeeder::class);
        $this->call(PlatformDataRowsTableSeeder::class);
        $this->call(RangeDataRowsTableSeeder::class);
        $this->call(DeviceDataRowsTableSeeder::class);
        $this->call(MessageDataRowsTableSeeder::class);
        $this->call(RatingDataRowsTableSeeder::class);
        $this->call(UserDataRowsTableSeeder::class);
    }
}
