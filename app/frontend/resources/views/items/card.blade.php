<div class="card">
    <div class="card-body text-center">
        <p class="mb-0" style="white-space: nowrap; overflow-x: hidden;">
            {{ $item->english_name }}
        </p>
        <p class="text-muted small" style="white-space: nowrap; overflow-x: hidden">
            {{ $item->foreign_name }}
        </p>

        <div style="height: 7rem" class="text-center">
            <img src="{{ $item->image->url }}" class="mw-100 mh-100 rounded">
        </div>
    </div>
    <ul class="list-group list-group-flush">
        <li class="list-group-item">
            <div class="d-flex">
                <div class="text-left flex-fill" style="white-space: nowrap; overflow-x: ellipsis">
                    <p class="m-0 p-0 small text-muted">
                        {{ __('Brand') }}
                    </p>
                    <p class="p-0 m-0 small">
                        <a href="{{ $item->brand->url }}">
                            {{ $item->brand->name }}
                        </a>
                    </p>
                </div>
                <div class="text-right flex-fill" style="white-space: nowrap; overflow-x: hidden">
                    <p class="m-0 p-0 small text-muted">
                        {{ __('Category') }}
                    </p>
                    <p class="p-0 m-0 small">
                        <a href="{{ $item->category->url }}">
                            {{ $item->category->name }}
                        </a>
                    </p>
                </div>
            </div>
        </li>
    </ul>
    <a class="card-body py-3 text-center" href="{{ $item->url }}">
        View Item
    </a>
</div>
