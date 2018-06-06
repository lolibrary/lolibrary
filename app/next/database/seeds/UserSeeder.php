<?php

use App\User;
use Illuminate\Database\Seeder;

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
            'email' => config('site.admin.email') ?? 'admin@example.com',
            'username' => config('site.admin.username') ?? 'admin',
            'password' => bcrypt($password = str_random(64)),
            'name' => config('site.admin.name') ?? 'Admin',
            'level' => User::DEVELOPER,
        ]);

        echo "Admin email: {$user->email}" . PHP_EOL;
        echo "Admin password: {$password}" . PHP_EOL;
    }
}
