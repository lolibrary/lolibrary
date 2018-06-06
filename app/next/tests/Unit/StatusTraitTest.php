<?php

namespace Tests\Unit;

use Tests\TestCase;
use App\Models\HasStatus;
use Illuminate\Support\Collection;

class StatusTraitTest extends TestCase
{
    public function testFetchingRawStatusCode()
    {
        $mock = new class {
            use HasStatus;

            protected static $statuses = [
                'foo' => 2 ** 0,
                'bar' => 2 ** 1,
            ];
            public $status = 1;
        };

        $this->assertEquals(2 ** 0, $mock->getRawStatus('foo'));
        $this->assertEquals(2 ** 1, $mock->getRawStatus('bar'));
    }

    /**
     * @expectedException \InvalidArgumentException
     * @expectedExceptionMessage unknown status 'baz'
     */
    public function testFetchingInvalidRawStatusCode()
    {
        $mock = new class {
            use HasStatus;

            protected static $statuses = [
                'foo' => 2 ** 0,
                'bar' => 2 ** 1,
            ];
            public $status = 1;
        };

        $mock->getRawStatus('baz');
    }

    public function testHasStatusChecks()
    {
        $mock = new class {
            use HasStatus;

            protected static $statuses = [
                'foo' => 2 ** 0,
                'bar' => 2 ** 1,
                'baz' => 2 ** 2,
            ];
            public $status = 5;
        };

        $this->assertTrue($mock->hasStatus('foo'));
        $this->assertFalse($mock->hasStatus('bar'));
        $this->assertTrue($mock->hasStatus('baz'));
    }

    /**
     * @expectedException \InvalidArgumentException
     * @expectedExceptionMessage unknown status 'invalid'
     */
    public function testHasInvalidStatus()
    {
        $mock = new class {
            use HasStatus;

            protected static $statuses = [
                'foo' => 2 ** 0,
                'bar' => 2 ** 1,
                'baz' => 2 ** 2,
            ];
            public $status = 5;
        };

        $mock->hasStatus('invalid');
    }

    public function testStatusMap()
    {
        $mock = new class {
            use HasStatus;

            protected static $statuses = [
                'foo' => 2 ** 0,
                'bar' => 2 ** 1,
                'baz' => 2 ** 2,
            ];
        };

        $this->assertEquals([
            'foo' => 2 ** 0,
            'bar' => 2 ** 1,
            'baz' => 2 ** 2,
        ], $mock->getStatusMap());
    }

    /**
     * @expectedException \InvalidArgumentException
     * @expectedExceptionMessage static::$statuses not found
     */
    public function testDefaultStatusMap()
    {
        $mock = new class {
            use HasStatus;
        };

        $mock->getStatusMap();
    }

    public function testStatuses()
    {
        $mock = new class {
            use HasStatus;

            protected static $statuses = [
                'foo' => 2 ** 0,
                'bar' => 2 ** 1,
                'baz' => 2 ** 2,
                'abc' => 2 ** 3,
            ];

            public $status = 5;
        };

        $collection = $mock->getStatusesAttribute();

        $this->assertInstanceOf(Collection::class, $collection);
        $this->assertEquals([
            'foo' => true,
            'bar' => false,
            'baz' => true,
            'abc' => false,
        ], $collection->all());
    }

    public function testGetRawStatusCode()
    {
        $mock = new class {
            use HasStatus;

            public $status = 5;
        };

        $this->assertEquals(5, $mock->getRawStatusCode());
        $this->assertEquals(5, $mock->status);
    }

    public function testGetRawStatusCodeAsString()
    {
        $mock = new class {
            use HasStatus;

            public $status = '5';
        };

        $this->assertEquals(5, $mock->getRawStatusCode());
        $this->assertEquals('5', $mock->status);
    }

    public function testSettingRawStatusCode()
    {
        $mock = new class {
            use HasStatus;

            public $status = '5';
        };

        $mock->setRawStatusCode(19);

        $this->assertEquals(19, $mock->getRawStatusCode());
        $this->assertEquals(19, $mock->status);
    }

    /**
     * @expectedException \TypeError
     */
    public function testSettingNonIntRawStatusCode()
    {
        $mock = new class {
            use HasStatus;

            public $status = '5';
        };

        $mock->setRawStatusCode('abc');
    }

    public function testSettingStringRawStatusCode()
    {
        $mock = new class {
            use HasStatus;

            public $status = '5';
        };

        $mock->setRawStatusCode('19');

        $this->assertEquals(19, $mock->getRawStatusCode());
        $this->assertEquals(19, $mock->status);
    }

    public function testStatusExists()
    {
        $mock = new class {
            use HasStatus;

            protected static $statuses = [
                'foo' => 2 ** 0,
                'bar' => 2 ** 1,
                'baz' => 2 ** 2,
                'abc' => 2 ** 3,
            ];
        };

        $this->assertTrue($mock->statusExists('foo'));
        $this->assertTrue($mock->statusExists('bar'));
        $this->assertTrue($mock->statusExists('baz'));
        $this->assertTrue($mock->statusExists('abc'));
        $this->assertFalse($mock->statusExists('invalid'));
    }

    public function testAddingStatus()
    {
        $mock = new class {
            use HasStatus;

            protected static $statuses = [
                'foo' => 2 ** 0,
                'bar' => 2 ** 1,
                'baz' => 2 ** 2,
                'abc' => 2 ** 3,
            ];

            public $status = 0;
        };

        $this->assertFalse($mock->hasStatus('abc'));

        $mock->addStatus('abc');

        $this->assertTrue($mock->hasStatus('abc'));
    }

    /**
     * @expectedException \InvalidArgumentException
     * @expectedExceptionMessage unknown status 'invalid'
     */
    public function testAddingNonExistentStatus()
    {
        $mock = new class {
            use HasStatus;

            protected static $statuses = [
                'foo' => 2 ** 0,
                'bar' => 2 ** 1,
                'baz' => 2 ** 2,
                'abc' => 2 ** 3,
            ];

            public $status = 0;
        };

        $mock->addStatus('invalid');
    }

    public function testModelSavingOnAddStatus()
    {
        $mock = new class {
            use HasStatus;

            protected static $statuses = [
                'foo' => 2 ** 0,
                'bar' => 2 ** 1,
                'baz' => 2 ** 2,
                'abc' => 2 ** 3,
            ];

            public $status = 0;

            public $saved = false;

            public function save()
            {
                $this->saved = true;
            }
        };

        $this->assertFalse($mock->saved);

        $mock->addStatus('abc');

        $this->assertTrue($mock->saved);
    }

    public function testModelNotSavingOnAddStatus()
    {
        $mock = new class {
            use HasStatus;

            protected static $statuses = [
                'foo' => 2 ** 0,
                'bar' => 2 ** 1,
                'baz' => 2 ** 2,
                'abc' => 2 ** 3,
            ];

            public $status = 0;

            public $saved = false;

            public function save()
            {
                $this->saved = true;
            }
        };

        $this->assertFalse($mock->saved);

        $mock->addStatus('abc', false);

        $this->assertFalse($mock->saved);
    }

    public function testAddingMultipleStatuses()
    {
        $mock = new class {
            use HasStatus;

            protected static $statuses = [
                'foo' => 2 ** 0,
                'bar' => 2 ** 1,
                'baz' => 2 ** 2,
                'abc' => 2 ** 3,
            ];

            public $status = 0;
        };

        $this->assertFalse($mock->hasStatus('foo'));
        $this->assertFalse($mock->hasStatus('bar'));
        $this->assertFalse($mock->hasStatus('baz'));
        $this->assertFalse($mock->hasStatus('abc'));

        $mock->addStatuses(['abc', 'bar', 'baz']);

        $this->assertFalse($mock->hasStatus('foo'));
        $this->assertTrue($mock->hasStatus('bar'));
        $this->assertTrue($mock->hasStatus('baz'));
        $this->assertTrue($mock->hasStatus('abc'));
    }

    public function testAddingMultipleStatusesSingleSave()
    {
        $mock = new class {
            use HasStatus;

            protected static $statuses = [
                'foo' => 2 ** 0,
                'bar' => 2 ** 1,
                'baz' => 2 ** 2,
                'abc' => 2 ** 3,
            ];

            public $status = 0;

            public $saved = false;
            public $called = 0;

            public function save()
            {
                $this->saved = true;
                $this->called++;
            }
        };

        $this->assertFalse($mock->saved);
        $this->assertEquals(0, $mock->called);

        $mock->addStatuses(['abc', 'bar', 'baz']);

        $this->assertTrue($mock->saved);
        $this->assertEquals(1, $mock->called);
    }

    public function testAddingMultipleStatusesWithoutSave()
    {
        $mock = new class {
            use HasStatus;

            protected static $statuses = [
                'foo' => 2 ** 0,
                'bar' => 2 ** 1,
                'baz' => 2 ** 2,
                'abc' => 2 ** 3,
            ];

            public $status = 0;

            public $saved = false;
            public $called = 0;

            public function save()
            {
                $this->saved = true;
                $this->called++;
            }
        };

        $this->assertFalse($mock->saved);
        $this->assertEquals(0, $mock->called);

        $mock->addStatuses(['abc', 'bar', 'baz'], false);

        $this->assertFalse($mock->saved);
        $this->assertEquals(0, $mock->called);
    }

    /**
     * @expectedException \InvalidArgumentException
     * @expectedExceptionMessage unknown status 'invalid'
     */
    public function testAddingMultipleStatusesWithInvalid()
    {
        $mock = new class {
            use HasStatus;

            protected static $statuses = [
                'foo' => 2 ** 0,
                'bar' => 2 ** 1,
                'baz' => 2 ** 2,
                'abc' => 2 ** 3,
            ];

            public $status = 0;
        };

        $mock->addStatuses(['abc', 'bar', 'invalid', 'baz']);
    }
}
