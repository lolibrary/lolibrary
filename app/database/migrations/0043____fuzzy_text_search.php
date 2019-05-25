<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class FuzzyTextSearch extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('items', function (Blueprint $table) {
            DB::statement('create extension pg_trgm;');
            DB::statement('ALTER DATABASE lolibrary SET pg_trgm.word_similarity_threshold = 0.8;');
            DB::statement('ALTER DATABASE lolibrary SET pg_trgm.similarity_threshold = 0.6;');
        });

        Schema::table('items', function (Blueprint $table) {
            DB::statement('create index english_name_trgm on items using gist (english_name gist_trgm_ops);');
            DB::statement('create index foreign_name_trgm on items using gist (foreign_name gist_trgm_ops);');
            DB::statement('create index product_number_trgm on items using gist (product_number gist_trgm_ops);');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::table('items', function (Blueprint $table) {
            $table->dropIndex('english_name_trgm');
            $table->dropIndex('foreign_name_trgm');
            $table->dropIndex('product_number_trgm');
        });

        Schema::table('items', function (Blueprint $table) {
            DB::statement('drop extension pg_trgm;');
        });
    }
}
