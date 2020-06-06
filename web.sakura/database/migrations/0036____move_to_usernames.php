<?php

use App\Models\User;
use Illuminate\Database\PostgresConnection;
use Illuminate\Support\Facades\Cache;
use Illuminate\Support\Facades\DB;
use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class MoveToUsernames extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        // before we do this: back up all usernames to slugs.
        $users = User::query()->select(['slug', 'username'])->get();

        Cache::forever('migration.usernames', $users->toJson());

        User::where(function ($query) {
            return $query->whereNull('name')->orWhere('name', '');
        })->update([
            'name' => DB::raw('username'),
        ]);

        Schema::table('users', function (Blueprint $table) {
            $table->dropColumn('username', 'developer');
        });

        Schema::table('users', function (Blueprint $table) {
            $table->renameColumn('slug', 'username');
        });

        // now, make the name `not null`
        if (DB::connection() instanceof PostgresConnection) {
            DB::statement('alter table users alter column "name" set not null');
        }
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::table('users', function (Blueprint $table) {
            $table->renameColumn('username', 'slug');
        });

        Schema::table('users', function (Blueprint $table) {
            $table->string('username')->unique()->nullable();
            $table->boolean('developer')->default(false);
        });

        if (DB::connection() instanceof PostgresConnection) {
            DB::statement('alter table users alter column "name" set null');
        }
    }
}
