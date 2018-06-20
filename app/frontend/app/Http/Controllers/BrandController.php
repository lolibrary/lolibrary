<?php

namespace App\Http\Controllers;

use App\Models\Brand;

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
        return view('brands.show', compact('brand'));
    }

    /**
     * Show a paginated list of brands.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        // todo: make this a static ::index() method.
        $brands = Brand::paginate(36);

        return view('brands.index', compact('brands'));
    }
}
