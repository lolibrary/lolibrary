<li class="nav-item dropdown">
    <a id="navbarDropdown" class="nav-link dropdown-toggle" href="#" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false" v-pre>
        {{ auth()->user()->getRoleAttribute() }} <span class="caret"></span>
    </a>
    <div class="dropdown-menu" aria-labelledby="navbarDropdown">
        <a class="dropdown-item" href="{{ route('admin.dashboard') }}">
            <i class="fal fa-fw fa-columns"></i> {{ __('Dashboard') }}
        </a>

    @junior
        <a class="dropdown-item" href="{{ route('items.create') }}">
            <i class="fal fa-fw fa-plus"></i> {{ __('Add Item') }}
        </a>
        <a class="dropdown-item" href="{{ route('items.index') }}">
            <i class="fal fa-fw fa-bookmark"></i> {{ __('All Items') }}
        </a>
    @endjunior

    @senior
        <a class="dropdown-item" href="{{ route('admin.items.queue') }}">
            <i class="fal fa-fw fa-plus"></i> {{ __('Drafts Queue') }}
        </a>
    @endsenior
    </div>
</li>