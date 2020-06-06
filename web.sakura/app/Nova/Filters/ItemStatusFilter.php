<?php

namespace App\Nova\Filters;

use App\Models\Item as BaseItem;
use Illuminate\Http\Request;
use Illuminate\Support\Str;
use Laravel\Nova\Filters\Filter;

class ItemStatusFilter extends Filter
{
    /**
     * The filter's component.
     *
     * @var string
     */
    public $component = 'select-filter';

    /**
     * Apply the filter to the given query.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Illuminate\Database\Eloquent\Builder  $query
     * @param  mixed  $value
     * @return \Illuminate\Database\Eloquent\Builder
     */
    public function apply(Request $request, $query, $value)
    {
        // lock away certain functions behind dev.
        if ($request->user()->developer()) {
            switch ($value) {
                case 'shoe-drafts':
                    return $query->where('status', BaseItem::SHOE_DRAFTS);
                case 'pending-items':
                    return $query->where('status', BaseItem::PENDING);
                case 'missing-images':
                    return $query->where('status', BaseItem::MISSING_IMAGES);
            }
        }

        switch ($value) {
            case 'published':
                return $query->where('status', BaseItem::PUBLISHED);
            case 'my-published':
                return $this->restrict($request, $query)
                    ->where('status', BaseItem::PUBLISHED);

            case 'drafts':
                return $query->where('status', BaseItem::DRAFT);
            case 'my-drafts':
                return $this->restrict($request, $query)
                    ->where('status', BaseItem::DRAFT);

            case 'my-items':
                return $this->restrict($request, $query);
            case 'my-pending-items':
                return $this->restrict($request, $query)
                    ->where('status', BaseItem::PENDING);

            case 'published-by-me':
                return $query->where('publisher_id', $request->user()->id)
                    ->where('status', BaseItem::PUBLISHED);

            case 'published-by-others':
                return $this->restrict($request, $query)
                    ->where('publisher_id', '!=', $request->user()->id)
                    ->where('status', BaseItem::PUBLISHED);
        }

        return $query;
    }

    /**
     * Get the filter's available options.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return array
     */
    public function options(Request $request)
    {
        if ($request->user()->developer()) {
            return [
                'My Items' => 'my-items',
                'My Drafts' => 'my-drafts',
                'Published by Me' => 'published-by-me',
                'My Items (Published by Others)' => 'published-by-others',

                'Show Drafts (status = 10)' => 'shoe-drafts',
                'Missing Images (status = 4)' => 'missing-images',

                'Pending Review (status = 2)' => 'pending-items',
                'All Published (status = 1)' => 'published',
                'All Drafts (status = 0)' => 'drafts',
            ];
        }

        // TODO: proper moderation queue resource, App\Nova\Queue.
        if ($request->user()->lolibrarian()) {
            return [
                'My Items' => 'my-items',
                'My Drafts' => 'my-drafts',
                'Published by Me' => 'published-by-me',
                'My Items (Published by Others)' => 'published-by-others',

                'All Drafts' => 'drafts',
                'All Published' => 'published',
            ];
        }

        // junior and under.

        return [
            'My Items' => 'my-items',
            'My Drafts' => 'my-drafts',
            'My Items (Published)' => 'my-published',
        ];
    }

    /**
     * Restrict a filter to the logged in user.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Illuminate\Database\Eloquent\Builder  $query
     * @return \Illuminate\Database\Eloquent\Builder
     */
    protected function restrict(Request $request, $query)
    {
        return $query->where('user_id', $request->user()->id);
    }
}
