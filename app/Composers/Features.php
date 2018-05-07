<?php

namespace App\Composers;

use App\Feature;

class Features extends Composer
{
    /**
     * {@inheritdoc}
     */
    protected function load()
    {
        return Feature::select(['name', 'slug'])->get()->toSelectArray();
    }
}
