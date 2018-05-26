<?php

use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateClosetTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('closet', function (Blueprint $table) {
            $table->increments('id');
            $table->uuid('user_id');
            $table->uuid('item_id');
            $table->timestampsTz();

            $table->index(['user_id', 'item_id']);
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::drop('closet');
    }
}
