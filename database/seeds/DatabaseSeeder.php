<?php

use Illuminate\Database\Seeder;

class DatabaseSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $this->call(BrandSeeder::class);
        $this->call(ColorSeeder::class);
        $this->call(TypeSeeder::class);
        $this->call(FeatureSeeder::class);
        $this->call(AttributeSeeder::class);
        $this->call(UserSeeder::class);
        $this->call(StyleSeeder::class);
        $this->call(InstructionSeeder::class);
        $this->call(ImageSeeder::class);
    }
}
