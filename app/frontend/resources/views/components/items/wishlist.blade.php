@if (Auth::check())
    @if (Auth::user()->wants($item))
        <a class="btn btn-outline-primary" href="{{ route('items.wishlist', $item) }}"
            onclick="event.preventDefault(); document.getElementById('wishlist-form').submit();">
            <i class="fal fa-fw fa-star"></i>  {{ __('Add to Wishlist') }}
        </a>
    @else
        <a class="btn btn-outline-danger" href="{{ route('items.wishlist', $item) }}"
            onclick="event.preventDefault(); document.getElementById('wishlist-form').submit();">
            <i class="fal fa-fw fa-star"></i>  {{ __('Remove from Wishlist') }}
        </a>
    @endif

    <form id="wishlist-form" action="{{ route('items.wishlist', $item) }}" method="POST" style="display: none;">
        @csrf
        <input type="hidden" name="_method" value="put">
    </form>
@endif
