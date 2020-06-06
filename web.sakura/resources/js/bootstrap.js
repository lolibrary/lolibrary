
window._ = require('lodash');
window.Popper = require('popper.js').default;

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

/**
 * simple-lightbox is a lightweight lightbox interface, to make paging through
 * item photos a little nicer. It will register on elements with 'data-lightbox="show"'
 */

 window.SimpleLightbox = require('simple-lightbox');

/**
 * Both chosen and simple-lightbox need to be initialized *after* the rest of the
 * page has loaded - otherwise their elements may not be present yet.
 */

$(() => {
    $('.form-control-chosen').chosen()
    $('[data-toggle="tooltip"]').tooltip()
    let lightbox = new SimpleLightbox({elements: '[data-lightbox="show"]'});
});

const image = document.head.querySelector('meta[name="default-image"]');

window.defaultImage = image ? image.content : '/images/default.png';


