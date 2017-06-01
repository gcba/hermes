<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateBrandsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('brands', function (Blueprint $table) {
            $table->increments('id');
            $table->text('message');
            $table->enum('direction', ['in', 'out']);
            $table->integer('rating_id')->unsigned();
            $table->foreign('rating_id')->references('id')->on('ratings')->onDelete('cascade');
            $table->integer('appuser_id')->unsigned();
            $table->foreign('appuser_id')->references('id')->on('appusers')->onDelete('cascade');
            $table->integer('app_id')->unsigned();
            $table->foreign('app_id')->references('id')->on('apps')->onDelete('cascade');
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
        Schema::dropIfExists('brands');
    }
}
