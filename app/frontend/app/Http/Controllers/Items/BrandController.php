<?php

namespace App\Http\Controllers\Items;

use App\Models\Item;
use App\Models\Brand;
use App\Http\Controllers\Controller;

class BrandController extends Controller
{
    /**
     * Show a brand.
     *
     * @param \App\Models\Brand $brand
     * @return \Illuminate\Http\Response
     */
    public function show(Brand $brand)
    {
        return redirect()->to(search_route(['brands' => [$brand->slug]]));
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
