<?php

use App\Style;
use Illuminate\Database\Seeder;

class StyleSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $styles = [
            'Gothic',
            'Sweet',
            'Classic',
            'Casual',
            'Hime',
            'Shiro',
            'Kuro',
            'Country',
            'Sailor',
            'Guro',
            'Punk',
            'Ero',
            'Pirate',
            'Steampunk',
            'Fairy',
            'Deco',
            'Mori',
            'Kodona',
            'Aristocrat',
        ];

        foreach ($styles as $style) {
            Style::create([
                'slug' => str_slug($style),
                'name' => $style,
            ]);
        }
    }
}
