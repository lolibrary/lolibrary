<?php

use App\Models\Brand;
use App\Models\Image;
use Illuminate\Database\Seeder;
use Illuminate\Database\QueryException;

class BrandSeeder extends Seeder
{
    /**
     * A list of brands to seed.
     *
     * @var string[]
     */
    protected const BRANDS = [
        'ap' => 'Angelic Pretty',
        'btssb' => 'Baby, the Stars Shine Bright',
        'iw' => 'Innocent World',
        'aatp' => 'Alice and the Pirates',
        'meta' => 'Metamorphose Temps de Fille',
        'jane-marple' => 'Jane Marple',
        'victorian-maiden' => 'Victorian Maiden',
        'indie' => 'Indie Brand',
        'atelier-boz' => 'Atelier Boz',
        'emily-temple-cute' => 'Emily & Shirley Temple Cute',
        'excentrique' => 'Excentrique',
        'bodyline' => 'Bodyline',
        'taobao' => 'TaoBao',
        'moitie' => 'Moi-même-Moitié',
        'j-et-j' => 'Juliette et Justine',
        'mary-magdalene' => 'Mary Magdalene',
        'putumayo' => 'Putumayo',
        'h-naoto' => 'h. Naoto',
        'antique-beast' => 'Antique Beast',
        'atelier-pierrot' => 'Atelier Pierrot',
        'bpn' => 'Black Peace Now',
        'beth' => 'Beth',
        'max' => 'MAXICIMAM',
        'chantilly' => 'Chantilly',
        'cornet' => 'Cornet',
        'grimoire' => 'Grimoire',
        'heart-e' => 'Heart E',
        'chocomint' => 'Chocomint',
        'dokidoki' => '6%DOKIDOKI',
        'offbrand' => 'Offbrand',
        'milk' => 'MILK',
        'physical-drop' => 'Physical Drop',
        'millefleurs' => 'Millefleurs',
        'pink-house' => 'Pink House',
        'vivienne-westwood' => 'Vivienne Westwood',

        // ex-indie brands
        'haenuli' => 'Haenuli',
        'lief' => 'Leif',
        'chess-story' => 'Chess Story',
        'infanta' => 'Infanta',
        'surface-spell' => 'Surface Spell',
    ];

    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        foreach (static::BRANDS as $name => $brand) {
            $slug = str_slug($brand);

            if (Brand::where('slug', $slug)->exists()) {
                continue;
            }

            $image = Image::firstOrCreate([
                'name' => $brand . ' icon picture',
                'filename' => "{$slug}.png",
            ]);

            Brand::create([
                'slug' => str_slug($brand),
                'name' => $brand,
                'short_name' => $name,
                'image_id' => $image->id,
            ]);
        }
    }
}
