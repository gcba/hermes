<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateAppUserPlatformTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('app_user_platform', function (Blueprint $table) {
            $table->integer('app_user_id')->unsigned();
            $table->foreign('app_user_id')->references('id')->on('appusers')->onDelete('cascade');
            $table->index('app_user_id');
            $table->integer('platform_id')->unsigned();
            $table->foreign('platform_id')->references('id')->on('platforms')->onDelete('cascade');
            $table->index('platform_id');
            $table->primary(['platform_id', 'app_user_id']);
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('app_user_platform');
    }
}
