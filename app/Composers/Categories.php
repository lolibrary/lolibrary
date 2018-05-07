<?php

namespace App\Composers;

use App\Category;

class Categories extends Composer
{
    /**
     * {@inheritdoc}
     */
    protected function load()
    {
        return Category::select(['name', 'slug'])->get();
    }
}
