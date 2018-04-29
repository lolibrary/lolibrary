<?php

Route::get('/healthz', 'HealthCheckController@index')->name('healthz');
