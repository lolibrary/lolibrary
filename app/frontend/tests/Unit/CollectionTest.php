<?php

namespace Tests\Unit;

use Tests\TestCase;
use App\Models\User;
use App\Models\Brand;
use App\Models\Traits\Collection;
use Illuminate\Database\Eloquent\Collection as BaseCollection;

class CollectionTest extends TestCase
{
    public function testCreation()
    {
        $collection = new Collection(['foo', 'bar']);

        $this->assertInstanceOf(BaseCollection::class, $collection);
        $this->assertInstanceOf(Collection::class, $collection);
    }

    public function testSelectArray()
    {
        $models = [
            new Brand(['name' => 'Foo 1', 'slug' => 'slug-1']),
            new Brand(['name' => 'Foo 2', 'slug' => 'slug-2']),
            new Brand(['name' => 'Foo 3', 'slug' => 'slug-3']),
            new Brand(['name' => 'Foo 4', 'slug' => 'slug-4']),
            new Brand(['name' => 'Foo 5', 'slug' => 'slug-5']),
        ];

        $collection = new Collection($models);
        $array = $collection->toSelectArray();

        $this->assertCount(5, $collection);
        $this->assertCount(5, $array);

        $expected = [
            'slug-1' => 'Foo 1',
            'slug-2' => 'Foo 2',
            'slug-3' => 'Foo 3',
            'slug-4' => 'Foo 4',
            'slug-5' => 'Foo 5',
        ];

        $this->assertEquals($expected, $array);
    }

    public function testSelectArrayWithAlternativeKeys()
    {
        $models = [
            new User(['email' => 'Foo 1', 'username' => 'slug-1']),
            new User(['email' => 'Foo 2', 'username' => 'slug-2']),
            new User(['email' => 'Foo 3', 'username' => 'slug-3']),
            new User(['email' => 'Foo 4', 'username' => 'slug-4']),
            new User(['email' => 'Foo 5', 'username' => 'slug-5']),
        ];

        $collection = new Collection($models);
        $array = $collection->toSelectArray('username', 'email');

        $this->assertCount(5, $collection);
        $this->assertCount(5, $array);

        $expected = [
            'slug-1' => 'Foo 1',
            'slug-2' => 'Foo 2',
            'slug-3' => 'Foo 3',
            'slug-4' => 'Foo 4',
            'slug-5' => 'Foo 5',
        ];

        $this->assertEquals($expected, $array);
    }

    public function testSelectArrayWithAlternativeAdditionalKeys()
    {
        $models = [
            new Brand(['name' => 'Foo 1', 'slug' => 'slug-1', 'short_name' => 's-1']),
            new Brand(['name' => 'Foo 2', 'slug' => 'slug-2', 'short_name' => 's-2']),
            new Brand(['name' => 'Foo 3', 'slug' => 'slug-3', 'short_name' => 's-3']),
            new Brand(['name' => 'Foo 4', 'slug' => 'slug-4', 'short_name' => 's-4']),
            new Brand(['name' => 'Foo 5', 'slug' => 'slug-5', 'short_name' => 's-5']),
        ];

        $collection = new Collection($models);
        $array = $collection->toSelectArray('short_name');

        $this->assertCount(5, $collection);
        $this->assertCount(5, $array);

        $expected = [
            's-1' => 'Foo 1',
            's-2' => 'Foo 2',
            's-3' => 'Foo 3',
            's-4' => 'Foo 4',
            's-5' => 'Foo 5',
        ];

        $this->assertEquals($expected, $array);
    }
}
