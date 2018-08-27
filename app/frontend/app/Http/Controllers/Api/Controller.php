<?php

namespace App\Http\Controllers\Api;

use Illuminate\Http\JsonResponse;
use Illuminate\Contracts\View\View;
use Illuminate\Routing\Controller as BaseController;
use Illuminate\Foundation\Validation\ValidatesRequests;

class Controller extends BaseController
{
    use ValidatesRequests;

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
            return new JsonResponse([
                'status' => 500,
                'exception' => 'api.view.given',
                'message' => 'A view was supplied in an API route.',
            ], 500);
        }

        return new JsonResponse($result);
    }
}
