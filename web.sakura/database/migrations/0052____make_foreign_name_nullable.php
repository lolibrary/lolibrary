<?php

use App\Models\Item;
use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class MakeForeignNameNullable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('items', function (Blueprint $table) {
            $table->string('foreign_name', 300)->nullable()->change();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        if (Item::whereNull('foreign_name')->count() > 0) {
            throw new RuntimeException('Cannot make items.foreign_name non-null if we have any null values in the database');
        }

        Schema::table('items', function (Blueprint $table) {
            $table->string('foreign_name', 300)->change();
        });
    }
}
