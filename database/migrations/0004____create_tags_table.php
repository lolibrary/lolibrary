<?php

use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateTagsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('tags', function (Blueprint $table) {
            $table->uuid('id')->primary();
            $table->string('slug')->unique();
            $table->string('name');
            $table->timestampsTz();
        });

        Schema::create('item_tag', function (Blueprint $table) {
            $table->increments('id');
            $table->uuid('item_id');
            $table->uuid('tag_id');
            $table->timestampsTz();

            $table->index(['item_id', 'tag_id']);
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::drop('item_tag');
        Schema::drop('tags');
    }
}
