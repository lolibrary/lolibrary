<?php

use App\Models\User;
use App\Models\Profile;
use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class CreateProfilesTable extends Migration
{
    /**
     * Profile fields to add.
     *
     * @var array
     */
    protected $fields = [
        'inspiration',
        'favourite_books',
        'occupation',
        'location',
        'website',
        'since',
        'casual_style',
        'hobbies',
        'livejournal',
        'style',
        'brand',
    ];

    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('profiles', function (Blueprint $table) {
            $table->uuid('id')->primary();
            $table->uuid('user_id')->unique();

            foreach ($this->fields as $field) {
                $table->string($field)->nullable();
            }

            $table->timestampsTz();
        });

        $cursor = User::query()->select(array_merge($this->fields, ['id']))->cursor();

        foreach ($cursor as $user) {
            $user->profile()->create($user->toArray());
        }

        Schema::table('users', function (Blueprint $table) {
            $table->dropColumn($this->fields);
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('profiles');

        Schema::table('users', function (Blueprint $table) {
            foreach ($this->fields as $field) {
                $table->string($field)->nullable();
            }
        });
    }
}
