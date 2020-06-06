<?php

namespace App\Providers;

use App\Models\User;
use Illuminate\Support\Facades\Gate;
use Illuminate\Foundation\Support\Providers\AuthServiceProvider as ServiceProvider;

class AuthServiceProvider extends ServiceProvider
{
    /**
     * The policy mappings for the application.
     *
     * @var array
     */
    protected $policies = [
        'App\Models\Attribute' => 'App\Policies\AttributePolicy',
        'App\Models\Brand' => 'App\Policies\BrandPolicy',
        'App\Models\Category' => 'App\Policies\CategoryPolicy',
        'App\Models\Color' => 'App\Policies\ColorPolicy',
        'App\Models\Feature' => 'App\Policies\FeaturePolicy',
        'App\Models\Item' => 'App\Policies\ItemPolicy',
        'App\Models\User' => 'App\Policies\UserPolicy',
        'App\Models\Tag' => 'App\Policies\TagPolicy',
    ];

    /**
     * Register any authentication / authorization services.
     *
     * @return void
     */
    public function boot()
    {
        $this->registerPolicies();

        //
    }

    protected function registerItemGates()
    {
        //
    }
}
