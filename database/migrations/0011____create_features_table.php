<?php

use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateFeaturesTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('features', function (Blueprint $table) {
            $table->uuid('id')->primary();
            $table->string('slug')->unique();
            $table->string('name');
            $table->timestampsTz();
        });

        Schema::create('feature_item', function (Blueprint $table) {
            $table->increments('id');
            $table->uuid('feature_id');
            $table->uuid('item_id');
            $table->timestampsTz();

            $table->index(['feature_id', 'item_id']);
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::drop('feature_item');
        Schema::drop('features');
    }
}
