<?php

use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class AddLevelToUsersTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::table('users', function (Blueprint $table) {
            $table->dropColumn(['admin', 'active']);
        });

        Schema::table('users', function (Blueprint $table) {
            $table->integer('level')->default(App\Models\User::JUNIOR_LOLIBRARIAN);
            $table->boolean('banned')->default(false);
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
            $table->dropColumn('level');
            $table->dropColumn('banned');
        });

        Schema::table('users', function (Blueprint $table) {
            $table->boolean('admin')->default(false);
            $table->boolean('active')->default(true);
        });
    }
}
