<?php

use App\Color;
use Illuminate\Database\Seeder;

class ColorSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $colors = [
            'Black',
            'Pink',
            'White',
            'Offwhite',
            'Ivory',
            'Navy',
            'Brown',
            'Red',
            'Wine/Bordeaux',
            'Sax',
            'Blue',
            'Beige',
            'Lavender',
            'Green',
            'Mint',
            'Black x White',
            'Gray',
            'Black x Offwhite',
            'Yellow',
            'Purple',
            'Rose',
            'Pink x Offwhite',
            'Cream',
            'Gold',
            'Pink x White',
            'White x Pink',
            'Black x Pink',
            'Black x Red',
            'Antique Gold',
            'Silver',
            'White x Black',
            'Beige x Brown',
            'Sax x White',
            'Black x Navy',
            'Dark Pink',
            'Antique Silver',
            'Milk tea',
            'Orange',
            'Red x White',
            'Black x Beige',
            'Brown x Beige',
            'Black x Gray',
            'Black x Silver',
            'Red x Offwhite',
            'Black x Gold',
            'Offwhite x Black',
            'Black x Blue',
            'Navy x Offwhite',
            'Sax x Offwhite',
            'Brown x Pink',
            'Black x Wine',
            'Black x Purple',
            'Pink x Sax',
            'Navy x White',
            'Lavender x White',
            'Olive',
            'Gray x Black',
            'Ivory x Brown',
            'Red x Pink',
            'White x Red',
            'Wine/Bordeaux x Offwhite',
            'Pink x Black',
            'White x Navy',
            'Blue x White',
            'Black x Ivory',
            'Offwhite x Navy',
            'Mint x White',
            'White x Gray',
            'Black x Green',
            'Green/Mint x Brown',
            'Navy x Black',
            'Pink Gold',
            'Ivory x Black',
            'Wine x Offwhite',
            'Black x Plum',
            'White x Green',
            'Lavender x Black',
            'Offwhite/Cream',
        ];

        foreach ($colors as $color) {
            Color::create([
                'slug' => str_slug($color),
                'name' => $color,
            ]);
        }
    }
}
