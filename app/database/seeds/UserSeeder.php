<?php

use App\Models\User;
use Illuminate\Database\Seeder;

class UserSeeder extends Seeder
{
    /**
     * A UUID for the admin user.
     *
     * @var string
     */
    protected const UUID = '00000000-0000-0000-0000-000000000000';

    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        if (User::where('id', static::UUID)->exists()) {
            return;
        }

        $dispatcher = User::getEventDispatcher();

        User::unsetEventDispatcher();

        $user = User::forceCreate([
            'id' => static::UUID,
            'email' => config('site.admin.email') ?? 'admin@example.com',
            'username' => config('site.admin.username') ?? 'admin',
            'password' => bcrypt($password = str_random(64)),
            'name' => config('site.admin.name') ?? 'Admin',
            'level' => User::DEVELOPER,
        ]);

        User::setEventDispatcher($dispatcher);

        echo "Admin email: {$user->email}" . PHP_EOL;
        echo "Admin password: {$password}" . PHP_EOL;
    }
}
