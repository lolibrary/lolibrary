<div class="card">
    <div class="card-body text-center">
        <p class="mb-0"
            title="{{ $item->english_name }}"
            style="white-space: nowrap; overflow-x: hidden; text-overflow: ellipsis;">
            {{ $item->english_name }}
        </p>
        <p class="text-muted small"
            title="{{ $item->foreign_name }}"
            style="white-space: nowrap; overflow-x: hidden; text-overflow: ellipsis;">
            @if ($item->foreign_name)
                {{ $item->foreign_name }}
            @else
                &nbsp;
            @endif
        </p>

        <div style="height: 7rem" class="text-center">
            <img src="{{ $item->image->url ?? asset('images/default.png') }}" class="mw-100 mh-100 rounded">
        </div>
    </div>
    <ul class="list-group list-group-flush">
        <li class="list-group-item">
            <div class="row small text-muted">
                <p class="col m-0 text-left">
                    {{ __('Brand') }}
                </p>
                <p class="col m-0 text-right">
                    {{ __('Category') }}
                </p>
            </div>
            <div class="d-flex small">
                <p class="p-0 m-0 text-left flex-fill" style="white-space: nowrap; overflow-x: ellipsis;">
                    <a href="{{ $item->brand->url }}" title="{{ $item->brand->name }}">
                        {{ str_limit($item->brand->name, 21) }}
                        {{-- deliberately chose 21 as the cutoff since lots of brand names fit on word boundaries. --}}
                    </a>
                </p>
                <p class="p-0 m-0 text-right flex-fill" style="white-space: nowrap; overflow-x: hidden">
                    <a href="{{ $item->category->url }}">
                        {{ $item->category->name }}
                    </a>
                </p>
            </div>
        </li>
    </ul>
    <a class="card-body py-3 text-center" href="{{ $item->url }}">
        View Item
    </a>
</div>
