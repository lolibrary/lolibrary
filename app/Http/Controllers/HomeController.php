<?php

namespace App\Http\Controllers\Api;

use App\Tag;
use Illuminate\Database\Eloquent\Builder;
use App\Http\Requests\Api\TagSearchRequest;

class HomeController extends Controller
{
    /**
     * Return a 200 at /
     *
     * @return \Illuminate\Contracts\Routing\ResponseFactory|\Symfony\Component\HttpFoundation\Response
     */
    public function welcome()
    {
        return response();
    }
}
