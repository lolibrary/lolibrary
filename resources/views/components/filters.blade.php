<div class="card">
    <div class="card-header">
        {{ __('Filters') }}
    </div>
    <div class="card-body">
        <div class="input-group">
            <label class="control-label">{{ __('Category') }}</label>
            @include('components.categories')
        </div>
        <div class="input-group">
            <label class="control-label">{{ __('Brand') }}</label>
            @include('components.brands')
        </div>
        <div class="input-group">
            <label class="control-label">{{ __('Features') }}</label>
            @include('components.features')
        </div>
        <div class="input-group">
            <label class="control-label">{{ __('Colorway') }}</label>
            @include('components.colors')
        </div>
        <div class="input-group">
            <label class="control-label">{{ __('Year') }}</label>
            @include('components.year')
        </div>
    </div>
</div>
