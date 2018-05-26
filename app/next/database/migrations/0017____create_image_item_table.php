<?php

use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateImageItemTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('image_item', function (Blueprint $table) {
            $table->increments('id');
            $table->uuid('image_id');
            $table->uuid('item_id');
            $table->timestampsTz();

            $table->index(['image_id', 'item_id']);
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::drop('image_item');
    }
}
