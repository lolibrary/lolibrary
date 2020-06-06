<?php

use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;
use Illuminate\Support\Facades\Schema;

class CreateAttributesTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('attributes', function (Blueprint $table) {
            $table->uuid('id')->primary();
            $table->string('slug')->unique();
            $table->string('name');
            $table->timestampsTz();
        });

        Schema::create('attribute_item', function (Blueprint $table) {
            $table->increments('id');
            $table->uuid('attribute_id');
            $table->uuid('item_id');
            $table->text('value');
            $table->timestampsTz();

            $table->index(['attribute_id', 'item_id']);
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::drop('attribute_item');
        Schema::drop('attributes');
    }
}
