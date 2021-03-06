<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateMessagesTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('messages', function (Blueprint $table) {
            $table->bigIncrements('id');
            $table->text('message');
            $table->enum('direction', ['in', 'out']);
            $table->index('direction');
            $table->smallInteger('status')->default(0);
            $table->index('status');
            $table->string('transport_id', 90)->nullable();
            $table->index('transport_id');
            $table->integer('rating_id')->unsigned();
            $table->foreign('rating_id')->references('id')->on('ratings')->onDelete('cascade');
            $table->index('rating_id');
            $table->integer('created_by')->unsigned()->nullable();
            $table->foreign('created_by')->references('id')->on('users')->onDelete('set null');
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
        Schema::dropIfExists('messages');
    }
}
