<?php

namespace App\Composers;

use App\Brand;

class Brands extends Composer
{
    /**
     * {@inheritdoc}
     */
    protected function load()
    {
        return Brand::select(['name', 'slug', 'short_name'])->get();
    }
}
