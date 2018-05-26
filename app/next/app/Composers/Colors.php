<?php

namespace App\Composers;

use App\Color;

class Colors extends Composer
{
    /**
     * {@inheritdoc}
     */
    protected function load()
    {
        return Color::select(['name', 'slug'])->get()->toSelectArray();
    }
}
