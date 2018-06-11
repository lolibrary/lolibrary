<li class="nav-item dropdown">
    <a id="navbarDropdown" class="nav-link dropdown-toggle" href="#" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false" v-pre>
        {{ Auth::user()->username }} <span class="caret"></span>
    </a>

    <div class="dropdown-menu" aria-labelledby="navbarDropdown">
        <a class="dropdown-item" href="{{ route('profile') }}">
            <i class="fal fa-fw fa-user"></i> {{ __('Profile') }}
        </a>

        <a class="dropdown-item" href="{{ route('wishlist') }}">
            <i class="fal fa-fw fa-star"></i> {{ __('Wishlist') }}
        </a>

        <a class="dropdown-item" href="{{ route('closet') }}">
            <i class="fal fa-fw fa-tags"></i> {{ __('Closet') }}
        </a>

        <a class="dropdown-item" href="{{ route('logout') }}"
            onclick="event.preventDefault(); document.getElementById('logout-form').submit();">
            <i class="fal fa-fw fa-sign-out-alt"></i> {{ __('Logout') }}
        </a>

        <form id="logout-form" action="{{ route('logout') }}" method="POST" style="display: none;">
            @csrf
        </form>
    </div>
</li>
