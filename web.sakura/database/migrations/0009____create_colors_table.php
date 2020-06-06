<?php

use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateColorsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('colors', function (Blueprint $table) {
            $table->uuid('id')->primary();
            $table->string('slug')->unique();
            $table->string('name');
            $table->timestampsTz();
        });

        Schema::create('color_item', function (Blueprint $table) {
            $table->increments('id');
            $table->uuid('color_id');
            $table->uuid('item_id');
            $table->timestampsTz();

            $table->index(['color_id', 'item_id']);
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::drop('color_item');
        Schema::drop('colors');
    }
}
