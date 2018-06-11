<?php

use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateUsersTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('users', function (Blueprint $table) {
            $table->uuid('id')->primary();
            $table->uuid('image_id')->nullable();

            $table->string('name')->nullable();
            $table->string('username')->unique();
            $table->string('slug')->unique();
            $table->string('email')->unique();
            $table->string('password');
            $table->boolean('admin')->default('false');
            $table->boolean('active')->default('false');
            $table->boolean('developer')->default('false');

            // required questions
            $table->string('style')->nullable(); // favourite style
            $table->string('brand')->nullable(); // favourite brand
            $table->string('inspiration')->nullable(); // free-form text

            // extras for the profile (inspiration is technically required)
            $table->string('favourite_books')->nullable();
            $table->string('occupation')->nullable();
            $table->string('location')->nullable();
            $table->string('website')->nullable();
            $table->string('since')->nullable();
            $table->string('casual_style')->nullable();
            $table->string('hobbies')->nullable();
            $table->string('livejournal')->nullable();

            $table->rememberToken();
            $table->timestampsTz();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::drop('users');
    }
}
