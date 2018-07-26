<?php

namespace App\Http\Controllers;

use App\Http\Requests\SearchRequest;

class SearchController extends Controller
{
    public function index(SearchRequest $request)
    {
        return view('search');
    }
}
