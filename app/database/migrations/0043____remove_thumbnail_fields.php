<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class RemoveThumbnailFields extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('images', function (Blueprint $table) {
            $table->dropColumn([
                'uploaded_filename',
                'thumbnail',
                'status',
                'image_id',
            ]);
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
            $table->string('uploaded_filename')->nullable();
            $table->string('thumbnail')->nullable();
            $table->integer('status')->default(0);
            $table->uuid('image_id')->nullable();
        });
    }
}
