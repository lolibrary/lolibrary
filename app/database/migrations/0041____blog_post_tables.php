<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class BlogPostTables extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        (new CreatePostsTable)->down();
        (new CreateTopicsTable)->down();

        Schema::create('posts', function (Blueprint $table) {
            $table->uuid('id')->primary();
            $table->uuid('user_id')->index();
            $table->string('slug')->unique();

            $table->string('title');

            $table->text('preview');
            $table->text('body');

            // a URL to an image to use for the header.
            $table->string('image', 400)->nullable();

            $table->timestampsTz();
            $table->timestampTz('published_at')->nullable();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('posts');

        (new CreateTopicsTable)->up();
        (new CreatePostsTable)->up();
    }
}
