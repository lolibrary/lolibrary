<?php

use App\Models\Item;
use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class AddPublisherToItems extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('items', function (Blueprint $table) {
            $table->uuid('publisher_id')->nullable();
            $table->timestampTz('published_at')->nullable();
        });

        DB::table('items')
            ->where('status', Item::PUBLISHED)
            ->update([
                'published_at' => DB::raw('created_at'),
                'publisher_id' => DB::raw('user_id'),
            ]);
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::table('items', function (Blueprint $table) {
            $table->dropColumn(['publisher_id', 'published_at']);
        });
    }
}
