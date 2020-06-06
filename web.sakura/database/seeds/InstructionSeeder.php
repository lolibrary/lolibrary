<?php

class InstructionSeeder extends Seeder
{
    /**
     * The model used for this seeder.
     *
     * @var string
     */
    protected static $model = App\Models\Instruction::class;

    /**
     * The name/value column.
     *
     * @var string
     */
    protected static $name = 'description';

    /**
     * A list of washing instructions, with slugs.
     *
     * @var array
     */
    protected static $content = [
        'dry-clean-any'            => 'Dry clean. Use any dry cleaning agent.',
        'dry-clean-petroleum'      => 'Dry clean. Use only a petroleum based agent.',
        'machine-wash-95'          => 'Machine wash in water temperature of 95°C or less. No other restrictions.',
        'machine-wash-60'          => 'Machine wash in water temperature of 60°C or less. No other restrictions.',
        'machine-wash-40'          => 'Machine wash in water temperature of 40°C or less. No other restrictions.',
        'machine-wash-delicate-40' => 'Machine wash at delicate cycle or hand wash in water temperature of 40°C or less.',
        'machine-wash-delicate-30' => 'Machine wash at delicate cycle or hand wash in water temperature of 30°C or less.',
        'hand-wash'                => 'Hand wash in water temperature of 30°C or less.',
        'can-bleach'               => 'Use chlorine bleach.',
        'do-not-wash'              => 'Do not wash (not washable).',
        'do-not-bleach'            => 'Do not use chlorine bleach.',
        'do-not-iron'              => 'Do not iron',
        'do-not-wring'             => 'Do not wring by hand.',
        'do-not-dry-clean'         => 'Do not dry clean.',
        'hang-dry'                 => 'Hang dry.',
        'hang-dry-shade'           => 'Hang dry in shade.',
        'dry-flat'                 => 'Lay flat to dry.',
        'dry-flat-shade'           => 'Lay flat to dry in shade.',
        'wring-or-spin-dry'        => 'Wring softly by hand or spin dry by machine quickly.',
        'iron-180-210'             => 'May be ironed directly at 180-210°C',
        'iron-140-160'             => 'May be ironed directly at 140-160°C.',
        'iron-80-120'              => 'May be ironed directly at 80-120°C',
        'iron-delicate-140-160'    => 'May be ironed at 140-160°C if a cloth is placed between iron and garment.',
        'iron-delicate-180-210'    => 'May be ironed directly at 180-210°C if a cloth is placed between iron and garment.',
    ];
}
