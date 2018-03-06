<?php

use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateInstructionsTable extends Migration
{
    /**
     * Run the migration.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('instructions', function (Blueprint $table) {
            $table->uuid('id')->primary();

            $table->string('slug')->unique();
            $table->string('description');

            $table->timestampsTz();
        });

        Schema::create('instruction_item', function (Blueprint $table) {
            $table->uuid('instruction_id');
            $table->uuid('item_id');

            $table->primary(['instruction_id', 'item_id']);
        });
    }

    /**
     * Reverse the migration.
     *
     * @return void
     */
    public function down()
    {
        Schema::drop('instructions');
        Schema::drop('instruction_item');
    }
}
