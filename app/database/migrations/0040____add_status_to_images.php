<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class AddStatusToImages extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('images', function (Blueprint $table) {
            $table->integer('status')->default(0);
            $table->string('uploaded_filename')->nullable();
        });

        DB::statement('alter table images alter column name drop not null');
        DB::statement('alter table images alter column filename drop not null');
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::table('images', function (Blueprint $table) {
            $table->dropColumn('status', 'uploaded_filename');
        });

        DB::statement('alter table images alter column name set not null');
        DB::statement('alter table images alter column filename set not null');
    }
}
