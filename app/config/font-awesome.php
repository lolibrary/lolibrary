<?php

return [

    /*
    |--------------------------------------------------------------------------
    | Font Awesome CDN URL
    |--------------------------------------------------------------------------
    |
    | While we do have a build chain, it's far easier to just use the CDN
    | for Font Awesome, and it lets us select between pro and free easier.
    |
    */

    'url' => env('FONT_AWESOME_URL', 'https://use.fontawesome.com/releases/v5.7.2/js/all.js'),

    /*
    |--------------------------------------------------------------------------
    | Font Awesome CDN URL
    |--------------------------------------------------------------------------
    |
    | We need an integrity hash so that we can verify the javascript that
    | we're loading from Font Awesome's CDN to make sure it hasn't been
    | man-in-the-middle'd or tampered with in any way.
    |
    */

    'hash' => env('FONT_AWESOME_HASH', 'sha384-0pzryjIRos8mFBWMzSSZApWtPl/5++eIfzYmTgBBmXYdhvxPc+XcFEk+zJwDgWbP'),
];
