<?php

use App\Feature;
use Illuminate\Database\Seeder;

class FeatureSeeder extends Seeder
{

    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $features = [
            'Lining',
            'No shirring',
            'Corset lacing',
            'Long sleeves',
            'Detachable bow',
            'Partial shirring',
            'Side zip',
            'Detachable waist ties',
            'Tiered skirt',
            'Short sleeves',
            'High waist',
            'Pockets',
            'Back shirring',
            'Pintucks',
            'Peter pan collar',
            'Full shirring',
            'Adjustable straps',
            'Detachable trim',
            'Pleats',
            'High neck collar',
            'Neck ties',
            'Long sleeve',
            'Bustled',
            'Empire waist',
            'Removable waist ribbon',
            'Scalloped',
            'Detachable sleeves',
            'Dropped waist',
            'Removable collar',
            'Boning',
            'Tucks',
            'Removable belt',
            'Short sleeve',
            'Princess sleeves',
            'Halter neckline',
            'Jabot',
            'Capelet',
            'Removable sash',
            'Built-in petticoat',
            'Convertible straps',
            'Underbust',
            'Detachable apron',
            'Detachable yoke',
        ];

        foreach ($features as $feature) {
            Feature::create([
                'slug' => str_slug($feature),
                'name' => $feature,
            ]);
        }
    }
}
