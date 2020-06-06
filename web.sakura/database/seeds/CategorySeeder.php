<?php

class CategorySeeder extends Seeder
{
    /**
     * The model to seed.
     *
     * @var string
     */
    protected static $model = App\Models\Category::class;

    /**
     * A listing of item categories.
     *
     * @var string[]
     */
    protected static $content = [
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
}
