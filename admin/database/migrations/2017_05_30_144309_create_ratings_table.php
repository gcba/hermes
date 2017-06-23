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
            $table->string('browser_version', 15)->nullable();
            $table->string('platform_version', 15);
            $table->boolean('has_message')->default(false);
            $table->index('has_message');
            $table->integer('app_id')->unsigned();
            $table->foreign('app_id')->references('id')->on('apps')->onDelete('cascade');
            $table->index('app_id');
            $table->integer('range_id')->unsigned();
            $table->foreign('range_id')->references('id')->on('ranges')->onDelete('cascade');
            $table->index('range_id');
            $table->integer('appuser_id')->unsigned()->nullable();
            $table->foreign('appuser_id')->references('id')->on('appusers')->onDelete('cascade');
            $table->index('appuser_id');
            $table->integer('platform_id')->unsigned();
            $table->foreign('platform_id')->references('id')->on('platforms')->onDelete('set null');
            $table->index('platform_id');
            $table->integer('device_id')->unsigned()->nullable();
            $table->foreign('device_id')->references('id')->on('devices')->onDelete('set null');
            $table->index('device_id');
            $table->integer('browser_id')->unsigned()->nullable();
            $table->foreign('browser_id')->references('id')->on('browsers')->onDelete('set null');
            $table->index('browser_id');
            $table->timestamps();
            $table->softDeletes();
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
