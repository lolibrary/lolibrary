@extends('layouts.app')

@section('content')
<div class="container">
    <h1 class="text-center" style="margin-bottom: 25px">
        <img src="{{ asset('images/banners/banner01.png') }}" alt="" style="max-height: 200px; max-width: 90%">
    </h1>

    <h2 class="sr-only">{{ __('News') }}</h2>
    <div class="row">
        <div class="col-md">
            <article class="card" style="margin-bottom: 10px">
            <img class="card-img-top" src="{{ asset('images/backgrounds/pattern_dark_blog-cropped.png') }}" alt="" style="object-fit: none; height: 160px; object-position: top">
                <div class="card-body">
                    <h5 class="card-title">We (finally) have a new site!</h5>
                    <h6 class="card-subtitle mb-2 text-muted text-right">posted by Amelia, <time datetime="2018-06-11T15:00:00Z">2018-06-11</time></h6>

                    <p class="card-text">
                        This one has been a long time coming; lots of delays as I've been the only one working on Lolibrary for close to 2 years now, and I've been moving house and changing jobs recently.
                    </p>

                    <div class="card-body text-right">
                        <a href="" class="card-link">
                            <i class="far fa-book"></i> Read More
                        </a>
                    </div>
                </div>
            </article>
        </div>
        <div class="col-md">
            <div class="card" style="margin-bottom: 10px">
                <img class="card-img-top" src="https://c5.patreon.com/external/logo/downloads_wordmark_navy@2x.png" alt="" style="height: 160px;">
                <div class="card-body">
                    <h5 class="card-title">We have a Patreon!</h5>
                    <h6 class="card-subtitle mb-2 text-muted text-right">posted by Amelia, <time datetime="2018-06-11T15:00:00Z">2018-06-11</time></h6>

                    <p class="card-text">
                        We started a Patreon so that people could donate to Lolibrary much, much easier <i class="far fa-heart"></i>.

                        <br>

                        If you can support us on Patreon, we'd greatly appreciate it - we're a pretty small team and we run entirely on donations.
                    </p>

                    <div class="card-body text-right">
                        <a href="" class="card-link">
                            <i class="far fa-book"></i> Read More
                        </a>
                    </div>
                </div>
            </div>
        </div>
        <div class="col-md">
            <div class="card" style="margin-bottom: 10px">
            <img class="card-img-top" src="{{ asset('images/backgrounds/pattern_dark_blog-cropped.png') }}" alt="" style="object-fit: none; height: 160px; object-position: top">
                <div class="card-body">
                    <h5 class="card-title">We (finally) have a new site!</h5>
                    <h6 class="card-subtitle mb-2 text-muted text-right">posted by Amelia, <time datetime="2018-06-11T15:00:00Z">2018-06-11</time></h6>

                    <p class="card-text">
                        This one has been a long time coming; lots of delays as I've been the only one working on Lolibrary for close to 2 years now, and I've been moving house and changing jobs recently.
                    </p>

                    <div class="card-body text-right">
                        <a href="" class="card-link">
                            <i class="far fa-book"></i> Read More
                        </a>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
@endsection
