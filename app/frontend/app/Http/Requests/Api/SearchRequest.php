<?php

namespace App\Http\Requests\Api;

use Illuminate\Foundation\Http\FormRequest;

/**
 * Class SearchRequest.
 *
 * @property string|null $search
 *
 * @property string|null $category
 * @property string[]|null $categories
 *
 * @property string|null $brand
 * @property string[]|null $brands
 *
 * @property string|null $color
 * @property string[]|null $colors
 *
 * @property string|null $feature
 * @property string[]|null $features
 *
 * @property string|null $tag
 * @property string[]|null $tags
 *
 * @property int|null $year
 * @property int[]|null $years
 */
class SearchRequest extends FormRequest
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
            'search' => 'sometimes|nullable|string|min:0,max:60',

            'category' => 'sometimes|required|exists:category,slug',
            'categories' => 'sometimes|array',
            'categories.*' => 'required|string|exists:categories,slug',

            'brand' => 'sometimes|required|exists:brands,slug',
            'brands' => 'sometimes|array',
            'brands.*' => 'required|string|exists:brands,slug',

            'color' => 'sometimes|required|exists:colors,slug',
            'colors' => 'sometimes|array',
            'colors.*' => 'required|string|exists:colors,slug',

            'feature' => 'sometimes|required|exists:features,slug',
            'features' => 'sometimes|array',
            'features.*' => 'required|string|exists:features,slug',

            'tag' => 'sometimes|required|exists:tags,slug',
            'tags' => 'sometimes|array',
            'tags.*' => 'required|string|exists:tags,slug',

            'year' => 'sometimes|required|integer|min:1990|max:' . (date('Y') + 3),
            'years' => 'sometimes|array',
            'years.*' => 'required|integer|min:1990|max:' . (date('Y') + 3),
        ];
    }
}
