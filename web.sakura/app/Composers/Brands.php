<?php

namespace App\Composers;

use App\Models\Brand;

class Brands extends Composer
{
    /**
     * {@inheritdoc}
     */
    protected function load()
    {
        return Brand::select(['name', 'slug'])
            ->orderBy('slug', 'asc')
            ->get()
            ->toSelectArray();
    }
}
