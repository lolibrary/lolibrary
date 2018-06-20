<?php

use App\Models\Post;
use Illuminate\Support\Carbon;
use Illuminate\Database\Seeder;

class BlogSeeder extends Seeder
{
    /**
     * Run the database seeds.
     *
     * @return void
     */
    public function run()
    {
        $user = '00000000-0000-0000-0000-000000000000';

        if (! Post::where('slug', 'lolibrary-open-source')->exists()) {
            $body = <<<BODY
In order to make it much easier for people to contribute to Lolibrary from a development point of view,
we've <a href="https://github.com/lolibrary/lolibrary" target="_blank" rel="external nofollow">open-sourced</a> Lolibrary's code on our GitHub.
<br><br>
If you're a developer with PHP, Javascript or Go experience, feel free to have a look!
<br>
We also post tickets that need doing on our <a href="https://trello.com/b/5nJUFgVo/lolibrary-kanban" rel="external nofollow" target="_blank">public kanban board</a>.
BODY;

            Post::forceCreate([
                'slug' => 'lolibrary-open-source',
                'title' => 'Lolibrary is now open source',
                'user_id' => $user,
                'preview' => $body,
                'body' => $body,
                'created_at' => Carbon::parse('2018-06-11T16:00:00+00:00'),
                'published_at' => Carbon::parse('2018-06-11T16:30:00+00:00'),
            ]);
        }

        if (! Post::where('slug', 'patreon-is-here')->exists()) {
            $body = <<<BODY
We started a Patreon so that people could donate to Lolibrary much, much easier <i class="far fa-heart"></i>.
<br>
If you could <a href="https://patreon.com/lolibrary" rel="external" target="_blank">support us on Patreon</a>, we'd greatly appreciate it - we're a pretty small team and we run entirely on donations!
BODY;

            Post::forceCreate([
                'slug' => 'patreon-is-here',
                'title' => 'We have a Patreon!',
                'user_id' => $user,
                'preview' => $body,
                'body' => $body,
                'image' => 'https://c5.patreon.com/external/logo/downloads_wordmark_navy@2x.png',
                'created_at' => Carbon::parse('2018-06-11T16:00:00+00:00'),
                'published_at' => Carbon::parse('2018-06-11T16:00:00+00:00'),
            ]);
        }

        if (! Post::where('slug', 'new-site')->exists()) {
            $preview = <<<BODY
This one has been a long time coming! Lots of delays, as I've been the only one working on Lolibrary for close to 2 years now, and I've recently moved house and changed job.
<br><br>
The good news is that the site is now under active development and I can start getting out the features that people have been missing! And you can register for accounts again!
BODY;

            $body = <<<BODY
This one has been a long time coming! Lots of delays, as I've been the only one working on Lolibrary for close to 2 years now, and I've recently moved house and changed job.
<br><br>
The good news is that the site is now under active development and I can start getting out the features that people have been missing! And you can register for accounts again!
<br>
If you experience any site-breaking bugs and do not get an error message telling you your error has been reported, you can help by
sending bug reports to <a href="mailto:support@lolibrary.org">support@lolibrary.org</a>
or as a GitHub issue on <a href="https://github.com/lolibrary/lolibrary/issues" target="_blank" rel="external nofollow">our bug tracker</a>.
<br><br>
If you want to keep up with our development, you can follow us on <a href="https://patreon.com/lolibrary" target="_blank" rel="external">Patreon</a>.
<br><br>
If you'd like to join as a developer, email me at <a href="mailto:amelia@lolibrary.org">amelia@lolibrary.org</a>.
BODY;


            Post::forceCreate([
                'slug' => 'new-site',
                'title' => 'Welcome to the new site~',
                'user_id' => $user,
                'preview' => $preview,
                'body' => $body,
                'created_at' => Carbon::parse('2018-06-11T15:00:00+00:00'),
                'published_at' => Carbon::parse('2018-06-11T15:00:00+00:00'),
            ]);
        }
    }
}
