<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateRatingsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('ratings', function (Blueprint $table) {
            $table->bigIncrements('id');
            $table->smallInteger('rating');
            $table->string('description', 30)->nullable();
            $table->string('app_version', 15)->nullable();
            $table->string('platform_version', 15);
            $table->string('browser_version', 15)->nullable();
            $table->boolean('has_message');
            $table->integer('app_id')->unsigned();
            $table->foreign('app_id')->references('id')->on('apps')->onDelete('cascade');
            $table->integer('appuser_id')->unsigned();
            $table->foreign('appuser_id')->references('id')->on('appusers')->onDelete('cascade');
            $table->integer('platform_id')->unsigned()->nullable();
            $table->foreign('platform_id')->references('id')->on('platforms')->onDelete('set null');
            $table->integer('device_id')->unsigned()->nullable();
            $table->foreign('device_id')->references('id')->on('devices')->onDelete('set null');
            $table->integer('browser_id')->unsigned()->nullable();
            $table->foreign('browser_id')->references('id')->on('browsers')->onDelete('set null');
            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('ratings');
    }
}
