<?php

use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateItemsTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('items', function (Blueprint $table) {
            $table->uuid('id')->primary();
            $table->uuid('type_id')->index();
            $table->uuid('brand_id')->index();
            $table->uuid('image_id'); // primary image
            $table->uuid('user_id')->index();

            $table->string('slug')->unique(); // URL slug
            $table->string('english_name', 300);
            $table->string('foreign_name', 300);
            $table->integer('year')->nullable();
            $table->string('product_number')->nullable();
            $table->text('notes')->nullable();

            $table->integer('status')->default(0);

            $table->timestampsTz();

            $table->index('created_at'); // we sort on this a *lot*.
            $table->index('updated_at');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::drop('items');
    }
}
