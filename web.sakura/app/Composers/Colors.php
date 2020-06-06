<?php

namespace App\Composers;

use App\Models\Color;

class Colors extends Composer
{
    /**
     * {@inheritdoc}
     */
    protected function load()
    {
        return Color::select(['name', 'slug'])
            ->orderBy('slug', 'asc')
            ->get()
            ->toSelectArray();
    }
}
