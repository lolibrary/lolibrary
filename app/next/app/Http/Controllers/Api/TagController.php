<?php

namespace App\Http\Controllers\Api;

use App\Tag;
use Illuminate\Database\Eloquent\Builder;
use App\Http\Requests\Api\TagSearchRequest;

class TagController extends Controller
{
    /**
     * Get all tags, paginated.
     *
     * @return \App\Tag[]|\Illuminate\Pagination\LengthAwarePaginator
     */
    public function index()
    {
        return Tag::orderBy('created_at')->paginate(100);
    }

    /**
     * Get a specific category.
     *
     * @param \App\Tag $tag
     * @return \App\Tag
     */
    public function show(Tag $tag)
    {
        return $tag;
    }

    /**
     * Search for a tag.
     *
     * @param \App\Http\Requests\Api\TagSearchRequest $request
     * @return \App\Tag[]|\Illuminate\Pagination\LengthAwarePaginator
     */
    public function search(TagSearchRequest $request)
    {
        return Tag::orderBy('created_at')->where(function (Builder $query) use ($request) {
            $query->where('slug', 'ilike', "%{$request->search}%")
                ->orWhere('slug', 'ilike', "%{$request->search}%");
        })->paginate(100);
    }
}
