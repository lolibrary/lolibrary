<?php

namespace App\Providers;

use Aws\S3\S3Client;
use League\Flysystem\Filesystem;
use Illuminate\Support\Facades\Storage;
use Illuminate\Support\ServiceProvider;
use League\Flysystem\AwsS3v3\AwsS3Adapter;

class StorageServiceProvider extends ServiceProvider
{
    /**
     * Bootstrap any application services.
     *
     * @return void
     */
    public function boot()
    {
        Storage::extend('minio', function ($app, $config) {
            $client = new S3Client([
                'credentials' => [
                    'key'    => $config['key'],
                    'secret' => $config['secret'],
                ],
                'region'      => $config['region'],
                'version'     => 'latest',
                'bucket_endpoint' => false,
                'use_path_style_endpoint' => $config['local'] ?? false,
                'endpoint'    => $config['endpoint'],
            ]);

            $options = [
                'override_visibility_on_copy' => true,
                'url' => add_s3_bucket($config['endpoint'], $config['bucket']),
            ];

            return new Filesystem(new AwsS3Adapter($client, $config['bucket'], '', $options));
        });
    }

    /**
     * Register any application services.
     *
     * @return void
     */
    public function register()
    {
        //
    }
}
