<li class="nav-item dropdown">
    <a id="navbarDropdown" class="nav-link dropdown-toggle" href="#" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false" v-pre>
        {{ auth()->user()->getRoleAttribute() }} <span class="caret"></span>
    </a>
    <div class="dropdown-menu" aria-labelledby="navbarDropdown">
        <a class="dropdown-item" href="{{ url('/library') }}">
            <i class="fal fa-fw fa-columns"></i> {{ __('Dashboard') }}
        </a>
    </div>
</li>