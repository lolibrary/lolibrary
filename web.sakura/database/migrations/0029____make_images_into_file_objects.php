<?php

use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;
use Illuminate\Support\Facades\Schema;

class MakeImagesIntoFileObjects extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('images', function (Blueprint $table) {
            $table->renameColumn('url', 'filename');
        });

        Schema::table('images', function (Blueprint $table) {
            $table->uuid('image_id')->index()->nullable();
            $table->string('thumbnail')->nullable();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::table('images', function (Blueprint $table) {
            $table->renameColumn('filename', 'url');
        });

        Schema::table('images', function (Blueprint $table) {
            $table->dropColumn(['image_id', 'thumbnail']);
        });
    }
}
