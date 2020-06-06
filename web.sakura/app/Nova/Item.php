<?php

namespace App\Nova;

use App\Models\Item as BaseItem;
use App\Nova\Actions\{PublishItem, UnpublishItem};
use App\Nova\Filters\ItemStatusFilter;
use Illuminate\Support\Str;
use Laravel\Nova\Panel;
use Laravel\Nova\Fields\Text;
use Laravel\Nova\Fields\Badge;
use Laravel\Nova\Fields\Image;
use Laravel\Nova\Fields\Avatar;
use Laravel\Nova\Fields\Select;
use Laravel\Nova\Fields\BelongsTo;
use Laravel\Nova\Fields\BelongsToMany;
use Laravel\Nova\Fields\DateTime;
use Laravel\Nova\Fields\Trix;
use Illuminate\Http\Request;
use Whitecube\NovaFlexibleContent\Flexible;
use NovaAttachMany\AttachMany;

class Item extends Resource
{
    /**
     * The model the resource corresponds to.
     *
     * @var string
     */
    public static $model = BaseItem::class;

    /**
     * The single value that should be used to represent the resource when being displayed.
     *
     * @var string
     */
    public static $title = 'english_name';

    /**
     * The columns that should be searched.
     *
     * @var array
     */
    public static $search = [
        'english_name', 'foreign_name', 'slug',
    ];

    /**
     * The relations that should always be loaded.
     *
     * @var array
     */
    public static $with = [
        'submitter',
        'brand',
        'category',
        'tags',
        'publisher',
    ];

    /**
     * Get the fields displayed by the resource.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return array
     */
    public function fields(Request $request)
    {
        return [
            Text::make('ID')->onlyOnDetail(),
            Text::make('Slug')->onlyOnDetail(),

            Avatar::make('Image')
                ->disk('s3public')
                ->path('images')
                ->nullable()
                ->maxWidth(200),

            Text::make('English Name')
                ->sortable()
                ->rules('required', 'max:255')
                ->onlyOnForms(),

            Text::make('English Name')
                ->sortable()
                ->onlyOnIndex()
                ->displayUsing(function (string $text) {
                    return Str::limit($text, 40);
                }),

            Text::make('English Name')->onlyOnDetail(),

            Text::make('Foreign Name')
                ->sortable()
                ->rules('max:255')
                ->hideFromIndex(),

            Text::make('Product Number')
                ->sortable()
                ->rules('string', 'nullable', 'max:255')
                ->hideFromIndex(),

            Select::make('Year')
                ->options(
                    collect(range(1990, (int) date('Y') + 1))
                        ->reverse()
                        ->mapWithKeys(function ($value) {
                            return [$value => $value];
                        })
                )
                ->displayUsingLabels()
                ->rules('nullable', 'integer', 'min:1990', 'max:'.(date('Y') + 3))
                ->hideFromIndex(),

            BelongsTo::make('Brand')->sortable(),
            BelongsTo::make('Category')->sortable(),

            Trix::make('Notes')->alwaysShow(),

            new Panel('Submission Details', [
                BelongsTo::make('Submitter', 'submitter', User::class)->readonly()->sortable()->exceptOnForms(),
                BelongsTo::make('Publisher', 'publisher', User::class)->readonly()->nullable()->onlyOnDetail(),

                DateTime::make('Created', 'created_at')->onlyOnDetail(),
                DateTime::make('Updated', 'updated_at')->onlyOnDetail(),
                DateTime::make('Published', 'published_at')->onlyOnDetail(),
            ]),

            new Panel('Price Details', [
                Select::make('Currency')
                    ->options(BaseItem::CURRENCIES)
                    ->displayUsingLabels()
                    ->hideFromIndex(),
                Text::make('Price')
                    ->hideFromIndex(),
            ]),

            new Panel('Additional Images', [
                Flexible::make('Images')
                    ->addLayout('Image', 'image', [
                        Image::make('Image')
                            ->path('images')
                            ->disk('s3public')
                            ->maxWidth(100)
                            ->disableDownload(),
                    ])
                    ->button('Add images'),
            ]),

            // this panel is only shown on the creation page.
            new Panel('Tags and Features', [
                AttachMany::make('Features', 'features', Feature::class),
                AttachMany::make('Tags', 'tags', Tag::class),
                AttachMany::make('Colors', 'colors', Color::class),
            ]),

            // This panel is only shown on the view and edit page
            new Panel('Tags, Features and Colors', [
                BelongsToMany::make('Item Features', 'features', Feature::class)->display('name'),
                BelongsToMany::make('Item Tags', 'tags', Tag::class)->searchable(),
                BelongsToMany::make('Item Colors', 'colors', Color::class)->display('name'),
            ]),

            new Panel('Attributes', [
                // Need to make a custom attributes panel here to allow it on item creation.
                BelongsToMany::make('Attributes', 'attributes', Attribute::class)
                    ->fields(function () {
                        return [
                            Text::make('Value')->nullable(),
                        ];
                    }),
            ]),

            Badge::make('Status', function () {
                switch ($this->status) {
                    case BaseItem::PUBLISHED:
                        return 'published';
                    case BaseItem::DRAFT:
                        return 'draft';
                    case BaseItem::PENDING:
                        return 'pending';
                    case BaseItem::MISSING_IMAGES:
                    case BaseItem::SHOE_DRAFTS:
                        return 'dev-only';
                    default:
                        return 'unknown';
                }
            })->map([
                'published' => 'success',
                'draft' => 'danger',
                'pending' => 'info',
                'dev-only' => 'warning',
                'unknown' => 'warning',
            ])->exceptOnForms(),
        ];
    }

    /**
     * Get the cards available for the request.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return array
     */
    public function cards(Request $request)
    {
        return [];
    }

    /**
     * Get the filters available for the resource.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return array
     */
    public function filters(Request $request)
    {
        return [
            new ItemStatusFilter,
        ];
    }

    /**
     * Get the lenses available for the resource.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return array
     */
    public function lenses(Request $request)
    {
        return [];
    }

    /**
     * Get the actions available for the resource.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return array
     */
    public function actions(Request $request)
    {
        return [
            (new PublishItem)->canSee(function (Request $request) {
                /** @var \Laravel\Nova\Http\Requests\NovaRequest $request */
                $model = $request->findModelQuery()->first();

                /** @var \App\Models\Item $model */
                if ($model === null) {
                    return $request->user()->lolibrarian();
                }

                return $model->draft() && $request->user()->can('publish', $model);
            }),

            (new UnpublishItem)->canSee(function (Request $request) {
                /** @var \Laravel\Nova\Http\Requests\NovaRequest $request */
                $model = $request->findModelQuery()->first();

                /** @var \App\Models\Item $model */
                if ($model === null) {
                    return $request->user()->senior();
                }

                return $model->published() && $request->user()->can('publish', $model);
            }),
        ];
    }
}
