<?php

namespace App\Http\Controllers\Items;

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
        return view('colors.show', compact('color'));
    }

    /**
     * Show a paginated list of colors.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        // todo: make this a static ::index() method.
        $colors = Color::paginate(36);

        return view('colors.index', compact('colors'));
    }
}
