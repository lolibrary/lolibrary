<?php

namespace App\Http\Requests\Api;

use Illuminate\Foundation\Http\FormRequest;

/**
 * Class TagSearchRequest.
 *
 * @property string $search
 */
class TagSearchRequest extends FormRequest
{
    /**
     * Authorize this request.
     *
     * @return bool
     */
    public function authorize()
    {
        return true;
    }

    /**
     * Get the rules for this request.
     *
     * @return array
     */
    public function rules()
    {
        return [
            'search' => 'required|string|min:1,max:30',
        ];
    }
}
