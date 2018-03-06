<?php

use App\Category;
use Illuminate\Database\Seeder;

class TypeSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $types = [
            'JSK',
            'Hair Accessories',
            'Skirt',
            'Sets',
            'OP',
            'Blouse',
            'Socks',
            'Jewelry',
            'Bags',
            'Coats',
            'Bolero',
            'Pants',
            'bloomers' => 'Bloomers / Undergarments',
            'Salopette',
            'Unmentionables',
            'Other',
            'Cardigan',
            'Accessories',
            'Corset/Bustier',
            'Cape',
            'Vest',
            'Petticoat',
            'Parasols',
            'Cutsew',
        ];

        foreach ($types as $slug => $type) {
            if (is_numeric($slug)) {
                $slug = str_slug($type);
            }

            Category::create([
                'slug' => $slug,
                'name' => $type,
            ]);
        }
    }
}
