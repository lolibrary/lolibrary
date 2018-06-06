<?php

namespace App\Http\Controllers\Api;

use Illuminate\Contracts\View\View;
use App\Http\Controllers\Controller as BaseController;

class Controller extends BaseController
{
    /**
     * Execute an action on the controller.
     *
     * @param  string  $method
     * @param  array   $parameters
     * @return \Symfony\Component\HttpFoundation\Response
     */
    public function callAction($method, $parameters)
    {
        $result = call_user_func_array([$this, $method], $parameters);

        if ($result instanceof View) {
            return response()->json([
                'status' => 500,
                'exception' => 'api.view.given',
                'message' => 'A view was supplied in an API route.',
            ], 500);
        }

        return response()->json($result);
    }
}
