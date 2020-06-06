<?php

class AttributeSeeder extends Seeder
{
    /**
     * The model to seed.
     *
     * @var string
     */
    protected static $model = App\Models\Attribute::class;

    /**
     * A list of attributes to seed.
     *
     * @var string[]
     */
    protected static $content = [
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
}
