<?php

namespace App\Composers;

use App\Models\Category;

class Categories extends Composer
{
    /**
     * {@inheritdoc}
     */
    protected function load()
    {
        return Category::select(['name', 'slug'])
            ->orderBy('slug', 'asc')
            ->get()
            ->toSelectArray();
    }
}
