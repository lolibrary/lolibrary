<?php

use App\Topic;
use App\User;
use Illuminate\Database\Seeder;
use Illuminate\Support\Facades\Storage;

class UserSeeder extends Seeder
{

    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $user = User::forceCreate([
            'id' => '00000000-0000-0000-0000-000000000000',
            'email' => config('site.admin.email'),
            'username' => config('site.admin.username'),
            'slug' => str_slug(config('site.admin.username')),
            'name' => config('site.admin.name'),
            'inspiration' => "I'm just a seeded admin account!",
            'level' => User::DEVELOPER,
        ]);

        echo "Admin email: {$user->email}" . PHP_EOL;

        Topic::forceCreate([
            'id' => '8f3e5b62-45b2-4515-8f07-c24a748e068f',
            'user_id' => $user->id,
            'slug' => 'welcome-to-lolibrary',
            'title' => 'Welcome to Lolibrary!',
            'body' => file_get_contents(base_path('setup/seeds/topic.md')),
            'allow_comments' => false,
        ]);
    }
}
