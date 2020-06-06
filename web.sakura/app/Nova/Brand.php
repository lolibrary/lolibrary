<?php

namespace App\Nova;

use Laravel\Nova\Fields\ID;
use Laravel\Nova\Fields\Text;
use Laravel\Nova\Fields\Avatar;
use Laravel\Nova\Fields\DateTime;
use Illuminate\Http\Request;
use Laravel\Nova\Http\Requests\NovaRequest;

class Brand extends Resource
{
    /**
     * The model the resource corresponds to.
     *
     * @var string
     */
    public static $model = 'App\Models\Brand';

    /**
     * The single value that should be used to represent the resource when being displayed.
     *
     * @var string
     */
    public static $title = 'name';

    /**
     * The columns that should be searched.
     *
     * @var array
     */
    public static $search = [
        'slug', 'name', 'short_name',
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
            ID::make()->onlyOnDetail(),

            Avatar::make('Image')
                ->disk('s3public')
                ->path('brands')
                ->nullable(),

            Text::make('Name')
                ->sortable()
                ->rules('required', 'string', 'min:2', 'max:255'),

            Text::make('Slug')
                ->sortable()
                ->creationRules('required', 'string', 'regex:/[a-z0-9][a-z0-9\-]{1,100}/u', 'unique:brands,slug')
                ->updateRules('required', 'string', 'regex:/[a-z0-9][a-z0-9\-]{1,100}/u', 'unique:brands,slug,{{resourceId}}')
                ->hideFromIndex(),
            
            Text::make('Short Name')
                ->sortable()
                ->creationRules('required', 'string', 'regex:/[a-z0-9][a-z0-9\-]{1,50}/u', 'unique:brands,short_name')
                ->updateRules('required', 'string', 'regex:/[a-z0-9][a-z0-9\-]{1,50}/u', 'unique:brands,short_name,{{resourceId}}')
                ->hideFromIndex(),
            
            DateTime::make('Created', 'created_at')->onlyOnDetail(),
            DateTime::make('Updated', 'updated_at')->onlyOnDetail(),
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
        return [];
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
        return [];
    }
}
