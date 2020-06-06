<?php

namespace App\Nova\Actions;

use Illuminate\Http\Request;
use Illuminate\Bus\Queueable;
use Laravel\Nova\Actions\Action;
use Illuminate\Support\Collection;
use Laravel\Nova\Fields\ActionFields;
use Illuminate\Queue\SerializesModels;
use Illuminate\Queue\InteractsWithQueue;
use Illuminate\Contracts\Queue\ShouldQueue;

class UnpublishItem extends Action
{
    use InteractsWithQueue, Queueable, SerializesModels;

    /**
     * Perform the action on the given models.
     *
     * @param  \Laravel\Nova\Fields\ActionFields  $fields
     * @param  \Illuminate\Support\Collection  $models
     * @return mixed
     */
    public function handle(ActionFields $fields, Collection $models)
    {
        $models->each->unpublish();

        return Action::message('Unpublished!');
    }

    /**
     * Get the fields available on the action.
     *
     * @return array
     */
    public function fields()
    {
        return [];
    }

        /**
     * Check an item is authorized to run.
     * 
     * @param \Illuminate\Http\Request $request
     * @param \App\Models\Item $model
     * @return bool
     */
    public function authorizedToRun(Request $request, $model)
    {
        if ($model->draft()) {
            return false;
        }
        
        return $request->user()->can('publish', $model);
    }
}
