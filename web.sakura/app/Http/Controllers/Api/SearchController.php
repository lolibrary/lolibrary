<?php

namespace App\Http\Controllers\Api;

use App\Models\Tag;
use App\Models\Item;
use App\Models\Brand;
use App\Models\Color;
use App\Models\Feature;
use App\Models\Category;
use Illuminate\Support\Str;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Storage;
use App\Http\Requests\Api\SearchRequest;
use Illuminate\Database\Eloquent\Builder;

use App\Http\Controllers\Controller as Base;

class SearchController extends Base
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
     * @param \Illuminate\Http\Request $request
     * @return \Illuminate\Contracts\Pagination\LengthAwarePaginator|\App\Item[]
     */
    public function search(SearchRequest $request)
    {
        $query = Item::query();

        $this->filters($request, $query);
        $this->years($request, $query);

        if (is_string($request->search) && strlen($request->search) > 0) {
            $search = $request->search;
            
            $query->where(function (Builder $query) use ($search) {
                $query->whereRaw('english_name %> ?',[$search]);
                $query->orWhereRaw('foreign_name %> ?', [$search]);
                $query->orWhereRaw('product_number %> ?', [$search]);
            });
        }

        $query->orderByDesc('published_at');

        $query->where('status', Item::PUBLISHED);

        $paginator = $query->paginate(24);

        $paginator->each(function (Item $item) {
            $item->image = Storage::cloud()->url($item->image);
            $item->makeVisible('image');
            
            if ($item->brand !== null) {
                $item->brand->image = Storage::cloud()->url($item->brand->image);
                $item->brand->makeVisible('image');
            }

            if ($item->category !== null) {
                $item->category->image = Storage::cloud()->url($item->category->image);
                $item->category->makeVisible('image');
            }
        });

        return $paginator;
    }

    /**
     * Filter relationships.
     *
     * @param \App\Requests\SearchRequest|\Illuminate\Http\Request $request
     * @param \Illuminate\Database\Eloquent\Builder $query
     * @return void
     */
    protected function filters(Request $request, Builder $query)
    {
        foreach (static::FILTERS as $class => $relation) {
            [$singular, $plural] = [Str::singular($relation), Str::plural($relation)];

            $models = (array) ($request->input($plural) ?? $request->input($singular));

            if (count($models) > 0) {
                $query->whereHas($relation, function (Builder $query) use ($models) {
                    $query->whereIn('slug', $models);
                });
            }
        }
    }

    /**
     * Filter on year.
     *
     * @param \App\Requests\SearchRequest|\Illuminate\Http\Request $request
     * @param \Illuminate\Database\Eloquent\Builder $query
     * @return void
     */
    protected function years(Request $request, Builder $query)
    {
        $years = (array) ($request->input('years') ?? $request->input('year'));

        if (count($years) > 0) {
            $query->whereIn('year', $years);
        }
    }
}
