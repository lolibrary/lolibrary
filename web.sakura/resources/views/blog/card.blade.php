<div class="card" style="margin-bottom: 10px">
    <img class="card-img-top"
         src="{{ $post->image ?? cdn_link('assets/backgrounds/pattern_dark_blog-cropped.png') }}"
         alt="" style="min-height: 120px; max-height: 150px; width: 100%; background-color: #ffffff;">
    <div class="card-body">
        <h5 class="card-title">{{ $post->title }}</h5>
        <h6 class="card-subtitle mb-2 text-muted text-right">
            {{ __('posted by') }} {{ $post->user->username ?? 'Anonymous' }},

            @if ($post->published_at)
                <time datetime="{{ $post->published_at->toRfc3339String() }}">{{ $post->published_at->format('jS M Y H:i') }}</time>
            @else
                <time datetime="{{ $post->created_at->toRfc3339String() }}">{{ $post->created_at->format('jS M Y H:i') }}</time>
            @endif
        </h6>

        <p class="card-text" style="min-height: 15rem; overflow-y: hidden">
            {{-- post preview is raw HTML. blog post permission is only given to admin or higher. --}}
            {!! $post->preview !!}
        </p>

        <div class="card-body text-right">
            <a href="{{ $post->url }}" class="card-link">
                <i class="far fa-book"></i> Read More
            </a>
        </div>
    </div>
</div>
