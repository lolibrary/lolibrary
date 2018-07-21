<?php

namespace App\Http\Controllers\Api;

use App\Models\Tag;
use App\Models\Item;
use App\Models\Brand;
use App\Models\Color;
use App\Models\Feature;
use App\Models\Category;
use Illuminate\Support\Str;
use Illuminate\Database\Eloquent\Builder;

class SearchController extends Controller
{
    /**
     * An array of models that we allow searching on.
     *
     * @var string[]
     */
    protected const FILTERS = [
        Brand::class => 'brand',
        Category::class => 'category',
        Color::class => 'colors',
        Feature::class => 'features',
        Tag::class => 'tags',
    ];

    /**
     * Search for items.
     *
     * @param \App\Http\Controllers\Api\SearchRequest $request
     * @return \Illuminate\Contracts\Pagination\LengthAwarePaginator|\App\Item[]
     */
    public function search(SearchRequest $request)
    {
        $query = Item::query();

        $this->filters($request, $query);
        $this->years($request, $query);

        if ($request->search !== null) {
            $search = '%' . $request->search . '%';

            $query->where(function (Builder $query) use ($search) {
                $query->where('english_name', 'ilike', $search);
                $query->orWhere('foreign_name', 'ilike', $search);
                $query->orWhere('product_number', 'ilike', $search);
            });
        }

        $query->orderByDesc('created_at');

        $query->where('status', Item::PUBLISHED);

        return $query->paginate(24);
    }

    /**
     * Filter relationships.
     *
     * @param \App\Http\Controllers\Api\SearchRequest $request
     * @param \Illuminate\Database\Eloquent\Builder $query
     * @return void
     */
    protected function filters(SearchRequest $request, Builder $query)
    {
        foreach (static::FILTERS as $class => $relation) {
            [$singular, $plural] = [Str::singular($relation), Str::plural($relation)];

            $models = $request->input($plural) ?? $request->input($singular);

            if ($models !== null) {
                $query->whereHas($relation, function (Builder $query) use ($models) {
                    is_array($models)
                        ? $query->whereIn('slug', $models)
                        : $query->where('slug', $models);
                });
            }
        }
    }

    /**
     * Filter on year.
     *
     * @param \App\Http\Controllers\Api\SearchRequest $request
     * @param \Illuminate\Database\Eloquent\Builder $query
     * @return void
     */
    protected function years(SearchRequest $request, Builder $query)
    {
        $years = $request->years ?? $request->year;

        if ($years !== null) {
            is_array($years)
                ? $query->whereIn('year', $years)
                : $query->where('year', $years);
        }
    }
}
