<?php

namespace App\Composers;

use App\Models\Feature;

class Features extends Composer
{
    /**
     * {@inheritdoc}
     */
    protected function load()
    {
        return Feature::select(['name', 'slug'])
            ->orderBy('slug', 'asc')
            ->get()
            ->toSelectArray();
    }
}
