<?php

namespace App\Http\Controllers\Items;

use App\Models\Item;
use App\Models\Feature;
use App\Http\Controllers\Controller;

class FeatureController extends Controller
{
    /**
     * Show a feature.
     *
     * @param \App\Models\Feature $feature
     * @return \Illuminate\Http\Response
     */
    public function show(Feature $feature)
    {
        $items = $feature->items()->with(Item::PARTIAL_LOAD)->paginate(24);

        return view('features.show', compact('feature', 'items'));
    }

    /**
     * Show a paginated list of features.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        // todo: make this a static ::index() method.
        $features = Feature::paginate(36);

        return view('features.index', compact('features'));
    }
}
