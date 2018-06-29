<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

/**
 * A donation request.
 *
 * @property 
 */
class DonateRequest extends FormRequest
{
    /**
     * Check if this request is authorized.
     *
     * @return bool
     */
    public function authorize()
    {
        return true;
    }

    /**
     * Get a list of rules for this request.
     *
     * @return array
     */
    public function rules()
    {
        return [
            'stripeToken' => 'required|string',
            'amount' => 'required|integer|min:100',
        ];
    }
}
