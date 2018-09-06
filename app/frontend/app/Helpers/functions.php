<?php

use App\Models\Image;
use App\Models\Model;
use GuzzleHttp\Psr7\Uri;
use Illuminate\Support\Facades\Storage;
use Illuminate\Support\Facades\Notification;

if (! function_exists('uuid')) {
    /**
     * Return a UUID without giving away our mac address.
     * Completely resistant to collisions.
     *
     * @return string
     */
    function uuid(): string
    {
        $node = dechex(random_int(0, 1 << 48) | 0x010000000000);

        return (string) Ramsey\Uuid\Uuid::uuid1($node);
    }
}

if (! function_exists('uuid4')) {
    /**
     * Return a random UUID, version 4.
     *
     * @return string
     */
    function uuid4(): string
    {
        return (string) Ramsey\Uuid\Uuid::uuid4();
    }
}

if (! function_exists('uuid5')) {
    /**
     * Return a sha1-based uuid, version 5.
     *
     * @param string $name
     *
     * @return string
     */
    function uuid5(string $name): string
    {
        return (string) Ramsey\Uuid\Uuid::uuid5(Model::NAMESPACE_UUID, $name);
    }
}

if (! function_exists('userify')) {
    /**
     * Return a username suitable for storage or searching from a display name.
     *
     * @param string $username
     *
     * @return string
     */
    function userify(string $username): string
    {
        $username = mb_strtolower($username);
        $username = preg_replace('/(\s+|[-]+)/u', '', $username);

        return $username;
    }
}

if (! function_exists('user')) {
    /**
     * Get a user by email or slug.
     *
     * @param string $id
     * @return \App\User|null
     */
    function user($id)
    {
        if (validator(['id' => $id], ['id' => 'required|email'])->passes()) {
            return App\User::where(DB::raw('lower(email)'), mb_strtolower($id))->first();
        }

        return App\User::where('username', $id)->first();
    }
}

if (! function_exists('slack')) {
    /**
     * Send a slack message notification.
     *
     * @param string $type
     * @return \Illuminate\Notifications\AnonymousNotifiable
     */
    function slack(string $type = 'notifications')
    {
        return Notification::route('slack', config("services.slack.$type"));
    }
}

if (! function_exists('add_s3_bucket')) {
    /**
     * Add an S3 bucket to a URL.
     *
     * @param string|null $url
     * @param string|null $bucket
     * @return string|null
     */
    function add_s3_bucket(?string $url, ?string $bucket)
    {
        if ($bucket === null || $url === null) {
            return null;
        }

        $uri = (new Uri($url));

        return (string) $uri->withHost("${bucket}." . $uri->getHost());
    }
}

if (! function_exists('search_route')) {
    /**
     * Get a search route for use with direct-linking.
     *
     * @param array $params
     * @return string
     */
    function search_route(array $params)
    {
        $results = [];

        foreach ($params as $key => $values) {
            $values = (array) $values;

            foreach ($values as $value) {
              $results[] = rawurlencode($key) . "[]=" . rawurlencode($value);
            }
        }

        return route('search') . '?' . implode('&', $results);
    }
}

if (! function_exists('default_asset')) {
    /**
     * Get the default asset, preferrably from the CDN.
     *
     * @return string
     */
    function default_asset()
    {
        return Storage::url((new Image)->getThumbnailsFolder() . '/default.jpeg');
    }
}
