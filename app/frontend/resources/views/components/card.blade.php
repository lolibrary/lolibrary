<div class="card shadow-sm m-2 flex-grid-card">
    <a href="{{ $url }}">
        <div class="card-body p-0 m-0" style="height: {{ $size }}px; width: {{ $size }}px; display: table-cell">
            <div class="d-flex justify-content-center align-items-center" style="height: {{ $size - 50}}px">
                <img src="{{ $image }}"
                    onerror="this.src = '{{ asset('images/icon.svg') }}'"
                    class="p-3 mh-100 mw-100" alt=""
                    data-original-url="{{ $image }}">
            </div>
            <div style="height: 50px" class="text-center">
                <p class="text-muted small p-2">
                    {{ $title }}
                </p>
            </div>
        </div>
    </a>
</div>
