<?php

namespace Tests\Unit;

use Tests\TestCase;

class SearchRouteTest extends TestCase
{
    public function testBasicArrays()
    {
        $route = route('search');

        $this->assertEquals(
            $route . '?categories[]=jsk&categories[]=op',
            search_route(['categories' => ['jsk', 'op']])
        );
    }

    public function testSingleValue()
    {
        $route = route('search');

        $this->assertEquals(
            $route . '?categories[]=jsk',
            search_route(['categories' => 'jsk'])
        );
    }

    public function testMultipleArrays()
    {
        $route = route('search');

        $this->assertEquals(
            $route . '?categories[]=jsk&categories[]=pants&brands[]=angelic-pretty&brands[]=dokidoki6',
            search_route(['categories' => ['jsk', 'pants'], 'brands' => ['angelic-pretty', 'dokidoki6']])
        );
    }
}
