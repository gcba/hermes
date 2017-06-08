<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateAppUserAppTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('app_user_app', function (Blueprint $table) {
            $table->increments('id');
            $table->integer('app_user_id')->unsigned();
            $table->foreign('app_user_id')->references('id')->on('appusers')->onDelete('cascade');
            $table->index('app_user_id');
            $table->integer('app_id')->unsigned();
            $table->foreign('app_id')->references('id')->on('apps')->onDelete('cascade');
            $table->index('app_id');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('app_user_app');
    }
}
