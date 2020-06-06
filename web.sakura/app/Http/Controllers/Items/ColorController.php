<?php

namespace App\Http\Controllers\Items;

use App\Models\Item;
use App\Models\Color;
use App\Http\Controllers\Controller;

class ColorController extends Controller
{
    /**
     * Show a color.
     *
     * @param \App\Models\Color $color
     * @return \Illuminate\Http\Response
     */
    public function show(Color $color)
    {
        return redirect()->to(search_route(['colors' => [$color->slug]]));
    }

    /**
     * Redirect to the search page.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        return redirect()->route('search');
    }
}
