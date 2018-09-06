<div class="row">
    <h1 class="col p-2">{{ title_case($name) }} Info</h1>
</div>

<div class="row text-center mb-4">
    <div class="p-1 list-group text-center col-lg-4 col-sm-6 small p-2">
        <a class="list-group-item" href="{{ $model->url }}">
            @if (is_string($image ?? false))
                <img class="d-block mx-auto mw-100 py-2 px-4" src="{{ $image }}" height="200">
            @elseif (($image ?? false) === true)
                <img class="d-block mx-auto mw-100 py-2 px-4" src="{{ $model->image->url ?? default_asset() }}" height="200">
            @endif

            @if ($tag ?? false)
                #{{ $model->slug }}
            @else
                <h1 class="h5 text-center m-0">{{ $model->name }}</h1>
            @endif
        </a>
    </div>
    <div class="col-lg-8 col-sm-6 p-2">
        @if ($description ?? false)
            <div class="card">
                <div class="card-body text-left">
                    {{ $model->description ?? 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.' }}
                </div>
            </div>
        @endif
        <p class="@if ($description ?? false) my-4 @endif">
            <span class="text-regular">{{ $items->total() }}</span> items found with this {{ $name }}.
        </p>

        <a href="{{ search_route([str_plural($name) => [$model->slug]]) }}"
            class="btn text-white btn-secondary py-2">Search using this {{ $name }}</a>
    </div>
</div>
