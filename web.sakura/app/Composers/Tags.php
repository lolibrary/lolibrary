<?php

namespace App\Composers;

use App\Models\Tag;

class Tags extends Composer
{
    /**
     * {@inheritdoc}
     */
    protected function load()
    {
        return Tag::select(['name', 'slug'])->get()->toSelectArray();
    }
}
