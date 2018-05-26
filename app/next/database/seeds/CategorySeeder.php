<?php

use App\Category;
use Illuminate\Database\Seeder;

class CategorySeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $categories = [
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

        foreach ($categories as $slug => $category) {
            if (is_numeric($slug)) {
                $slug = str_slug($category);
            }

            Category::create([
                'slug' => $slug,
                'name' => $category,
            ]);
        }
    }
}
