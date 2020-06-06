@if (Auth::check())
    @if (! Auth::user()->owns($item))
        <a class="btn btn-outline-primary" href="{{ route('items.closet', $item) }}"
            onclick="event.preventDefault(); document.getElementById('closet-form').submit();">
            <i class="far fa-fw fa-shopping-bag"></i>  {{ __('Add to Closet') }}
        </a>
    @else
        <a class="btn btn-outline-danger" href="{{ route('items.closet', $item) }}"
            onclick="event.preventDefault(); document.getElementById('closet-form').submit();">
            <i class="far fa-fw fa-shopping-bag"></i>  {{ __('Remove from Closet') }}
        </a>
    @endif

    <form id="closet-form" action="{{ route('items.closet', $item) }}" method="POST" style="display: none;">
        @csrf
        <input type="hidden" name="_method" value="put">
    </form>
@endif
