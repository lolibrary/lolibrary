<?php

use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;
use Illuminate\Support\Facades\Schema;

class ChangeUserAuthToArray extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('users', function (Blueprint $table) {
            $table->dropColumn('auth_token');
        });

        Schema::table('users', function (Blueprint $table) {
            $table->jsonb('auth')->default('{}');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::table('users', function (Blueprint $table) {
            $table->string('auth_token')->nullable();
        });

        Schema::table('users', function (Blueprint $table) {
            $table->dropColumn('auth');
        });
    }
}
