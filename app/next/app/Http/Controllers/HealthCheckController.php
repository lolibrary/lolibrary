<?php

namespace App\Http\Controllers;

class HealthCheckController extends Controller
{
    /**
     * Get the health check endpoint.
     * 
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        return response()->json(['alive' => true], 200);
    }
}