<?php

use App\Attribute;
use Illuminate\Database\Seeder;

class AttributeSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $attributes = [
            'Bust',
            'Length',
            'Price',
            'Waist',
            'Owner Height',
            'Owner Length',
            'Owner Waist',
            'Cuff',
            'Shoulder Width',
            'Sleeve Length',
            'Owner Notes',
            'Owner Bust',
            'Owner Underbust',

            // shoes
            'Heel Height',
            'Material',
            'Soles',
            'Finishes',
        ];

        foreach ($attributes as $attribute) {
            Attribute::create([
                'slug' => str_slug($attribute),
                'name' => $attribute,
            ]);
        }
    }
}
