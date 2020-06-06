<?php

namespace App\Models;

use NumberFormatter;
use App\Models\Traits\Sluggable;
use App\Models\Traits\Publishable;
use App\Models\Traits\ItemRelations;

/**
 * An Item of Apparel.
 *
 * @property string      $slug            The URL slug of an item.
 * @property string      $english_name    The English Title of an Item.
 * @property string|null $foreign_name    The 'Japanese Title' of an Item.
 * @property int|null    $year            The year an Item was released.
 * @property string|null $product_number  An Item's product number.
 * @property int         $status          The status of an item (stored internally as an int).
 * @property string      $price           The price of this item, in a given currency.
 * @property float       $price_formatted The price of this item, formatted to the rules of the given currency (e.g. /100 for gbp/usd)
 * @property string      $currency        The currency of this item, as an ISO code.
 *
 * @property \App\Models\Image $image     The primary {@link \App\Models\Image} for this Item.
 * @property \App\Models\Category $category  The {@link \App\Models\Category} of this Item (e.g. JSK).
 * @property \App\Models\User $submitter The {@link \App\Models\User} who originally submitted this Item.
 * @property \App\Models\User $publisher The {@link \App\Models\User} who published this Item.
 * @property \App\Models\Brand $brand     The {@link \App\Models\Brand} of this Item (e.g. Angelic Pretty).
 *
 * @property \App\Models\Image[]|\Illuminate\Database\Eloquent\Collection $images     The {@link \App\Models\Image images} for this Item.
 * @property \App\Models\Tag[]|\Illuminate\Database\Eloquent\Collection $tags       The {@link \App\Models\Tag search tags} for this Item.
 * @property \App\Models\Color[]|\Illuminate\Database\Eloquent\Collection $colors     The {@link \App\Models\Color colorways} this Item comes in (e.g. Black).
 * @property \App\Models\Feature[]|\Illuminate\Database\Eloquent\Collection $features   The {@link \App\Models\Feature features} of this item (e.g. Back Shirring).
 * @property \App\Models\Attribute[]|\Illuminate\Database\Eloquent\Collection $attributes The {@link \App\Models\Attribute custom attributes} on this Item.
 * @property \App\Models\User[]|\Illuminate\Database\Eloquent\Collection $stargazers The {@link \App\Models\User users} who want to own this Item.
 * @property \App\Models\User[]|\Illuminate\Database\Eloquent\Collection $owners     The {@link \App\Models\Attribute users} who own this Item.
 *
 * @property string $image_id The ID of this Item's {@link \App\Models\Image image}.
 * @property string $category_id  The ID of this Item's {@link \App\Models\Category category}.
 * @property string $user_id  The ID of the {@link \App\Models\User user} who submitted this Item.
 * @property string $brand_id The ID of this Item's {@link \App\Models\Brand brand}.
 * @property string $submitter_id The ID of this Item's {@link \App\Models\User submitter}.
 * @property string $publisher_id The ID of this Item's {@link \App\Models\User publisher}.
 *
 * @property \Carbon\Carbon $published_at The date this item was published.
 */
class Item extends Model
{
    use ItemRelations, Publishable, Sluggable;

    /**
     * A list of supported currencies.
     *
     * @var int
     */
    public const CURRENCIES = [
        'jpy' => 'Japanese Yen (¥)',
        'cny' => 'Chinese Yuan (RMB/¥)',
        'hkd' => 'Hong Kong Dollar (HK$)',
        'krw' => 'South Korean Won (₩)',
        'eur' => 'Euro (€)',
        'usd' => 'US Dollars ($)',
        'gbp' => 'Pound Sterling (£)',
        'cad' => 'Canadian Dollar (CA$)',
        'aud' => 'Australian Dollar (AU$)',
        'mxn' => 'Mexican Pesos ($)',
    ];

    /**
     * Indicates that an item is a draft and shouldn't be visible.
     *
     * @var int
     */
    public const DRAFT = 0;

    /**
     * Indicates that an item should be visible to everyone.
     *
     * @var int
     */
    public const PUBLISHED = 1;

    /**
     * Indicates that an item is ready for review and publishing.
     *
     * @var int
     */
    public const PENDING = 2;

    /**
     * Test status for missing image imports.
     *
     * @var int
     */
    public const MISSING_IMAGES = 4;

    /**
     * Test status for missing shoe drafts.
     *
     * @var int
     */
    public const SHOE_DRAFTS = 10;

    /**
     * A shortcut for fully eager loading an item.
     *
     * Use: `Item::with(Item::FULLY_LOAD)`
     */
    public const FULLY_LOAD = [
        'tags',
        'colors',
        'features',
        'category',
        'brand',
        'submitter',
        'attributes',
        'publisher',
    ];

    /**
     * The attributes required to show a listing of items.
     *
     * Use: `Item::with(Item::PARTIAL_LOAD)`
     */
    public const PARTIAL_LOAD = [
        'submitter',
        'brand',
        'category',
        'tags',
    ];

    /**
     * Non-fillable attributes.
     *
     * @var array
     */
    protected $guarded = [
        'id',
        'status',
        'slug',
        'created_at',
        'updated_at',
        'published_at',
    ];

    /**
     * Eager loads.
     *
     * @var array
     */
    protected $with = self::PARTIAL_LOAD;

    /**
     * Attributes to append.
     *
     * @var array
     */
    protected $appends = ['price_details', 'url'];

    /**
     * An array of keys to convert to dates.
     *
     * @var array
     */
    protected $dates = ['published_at'];

    /**
     * Visible attributes.
     *
     * @var array
     */
    protected $visible = [
        'id',
        'slug',
        'url',
        'english_name',
        'foreign_name',
        'notes',
        'price_details',
        'product_number',

        'tags',
        'colors',
        'features',
        'category',
        'brand',
        'submitter',
        'attributes',
        'publisher',

        'created_at',
        'updated_at',
        'published_at',
    ];

    /**
     * An array of column cast values.
     *
     * @var array
     */
    public $casts = [
        'images' => 'json',
    ];

    /**
     * Get the formatted price for this item.
     *
     * @return string
     */
    public function getFullPrice()
    {
        if (in_array($this->currency, ['jpy', 'krw', 'cny'])) {
            return (string) round($this->price);
        }

        return (string) round($this->price, 2);
    }

    /**
     * Get formatted price for an item.
     *
     * @return string
     */
    public function getPriceFormattedAttribute()
    {
        $price = $this->getFullPrice();

        $formatter = new NumberFormatter('en_US', NumberFormatter::CURRENCY);

        return $formatter->formatCurrency($price, $this->currency);
    }

    /**
     * Get a list of pricing details.
     *
     * @return array
     */
    public function getPriceDetailsAttribute()
    {
        return [
            'currency' => $this->currency,
            'price' => (int) $this->price,
            'local_price' => $this->getFullPrice(),
            'formatted' => $this->price_formatted,
        ];
    }
}
