<div class="row">
    @forelse ($items as $item)
        <div class="col-lg-3 col-md-4 col-sm-6 p-2">
            @include('items.card', ['item' => $item])
        </div>
    @empty
        <p class="text-center">No items found!</p>
    @endif
</div>

<div class="col p-2">
    {{ $items->links('pagination::simple-bootstrap-4') }}
</div>
