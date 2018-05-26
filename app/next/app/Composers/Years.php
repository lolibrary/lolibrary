<?php

namespace App\Composers;

use Illuminate\View\View;

class Years
{
    /**
     * Add data to the given view.
     *
     * @param \Illuminate\View\View
     * @return void
     */
    public function compose(View $view)
    {
        $view->with('years', array_reverse(range(1990, date('Y') + 1)));
    }
}
