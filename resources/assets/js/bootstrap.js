
window._ = require('lodash');
window.Popper = require('popper.js').default;

const Filter = require('./filter').default;
const Search = require('./search').default;

/**
 * We'll load jQuery and the Bootstrap jQuery plugin which provides support
 * for JavaScript based Bootstrap features such as modals and tabs. This
 * code may be modified to fit the specific needs of your application.
 */

try {
    window.$ = window.jQuery = require('jquery');

    require('bootstrap');
} catch (e) {}

/**
 * We'll load the axios HTTP library which allows us to easily issue requests
 * to our Laravel back-end. This library automatically handles sending the
 * CSRF token as a header based on the value of the "XSRF" token cookie.
 */

window.axios = require('axios');

window.axios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';

/**
 * Next we will register the CSRF Token as a common header with Axios so that
 * all outgoing HTTP requests automatically have it attached. This is just
 * a simple convenience so we don't have to attach every token manually.
 */

const token = document.head.querySelector('meta[name="csrf-token"]');

if (token) {
    window.axios.defaults.headers.common['X-CSRF-TOKEN'] = token.content;
} else {
    console.error('CSRF token not found: https://laravel.com/docs/csrf#csrf-x-csrf-token');
}

/**
 * Next we'll load chosen, for a nicer multi-select box with searching.
 * This will just attach itself to any .form-control-chosen element.
 */

window.chosen = require('chosen-js');

$(() => $('.form-control-chosen').chosen());

/**
 * Next we'll check if our search endpoint is set; if it's not set,
 * we'll emit an error in the console and refuse to load anything.
 */

const endpoint = document.head.querySelector('meta[name="search-endpoint"]');

if (! endpoint) {
    console.error('Search endpoint not set; set <meta name="search-endpoint" content="{{ route(...) }}">');
}

/**
 * Next we'll look for the "search bar" element; this will be the element
 * we'll bind our event listeners to, in order to actually search for
 * items. This also requires the endpoint from earlier to be set.
 */
else {
    const element = document.getElementById('search-bar');

    if (element) {
        const filter = new Filter(element);
        const search = new Search(element, endpoint.content);

        $(() => {
            $('.form-control-filter').on('change', () => filter.update());
            $('.form-control-search').on('input change', () => filter.update(true));
        });
    } else {
        console.log('Search Bar not found (looked for #search-bar)');
    }
}


