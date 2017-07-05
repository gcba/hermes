<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateAppUserDeviceTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('app_user_device', function (Blueprint $table) {
            $table->integer('app_user_id')->unsigned();
            $table->foreign('app_user_id')->references('id')->on('appusers')->onDelete('cascade');
            $table->index('app_user_id');
            $table->integer('device_id')->unsigned();
            $table->foreign('device_id')->references('id')->on('devices')->onDelete('cascade');
            $table->index('device_id');
            $table->primary(['device_id', 'app_user_id']);
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('app_user_device');
    }
}
