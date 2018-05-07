<?php

namespace App\Composers;

use App\Color;

class Categories extends Composer
{
    /**
     * {@inheritdoc}
     */
    protected function load()
    {
        return Color::select(['name', 'slug'])->get()->toSelectArray();
    }
}
