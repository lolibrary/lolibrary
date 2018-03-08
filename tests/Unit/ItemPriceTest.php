<?php

namespace Tests\Unit;

use App\Item;
use Tests\TestCase;

class ItemPriceTest extends TestCase
{
    /**
     * A basic test example.
     *
     * @return void
     */
    public function testBasicAttributes()
    {
        $model = new Item(['price' => 12345, 'currency' => 'jpy']);

        $this->assertEquals(12345, $model->price);
        $this->assertEquals('jpy', $model->currency);
    }

    /**
     * Test the getFullPrice() method.
     *
     * @dataProvider getFullPriceData
     * @covers \App\Item::getFullPrice()
     * @param int $price
     * @param string $currency
     * @param string $expected
     * @return void
     */
    public function testFullPriceAccessor(int $price, string $currency, string $expected)
    {
        $model = new Item(compact('price', 'currency'));

        $this->assertEquals($expected, $model->getFullPrice());
    }

    public function getFullPriceData()
    {
        return [
            [12345, 'jpy', '12345'],
            [12345, 'krw', '12345'],
            [12345, 'cny', '12345'],
            [12345, 'usd', '123.45'],
            [12345, 'gbp', '123.45'],
            [12345, 'aud', '123.45'],
            [12345, 'eur', '123.45'],
            [12345, 'mxn', '123.45'],
            [12345, 'cad', '123.45'],
        ];
    }
}
