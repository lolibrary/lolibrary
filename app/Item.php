<?php

namespace App;

use App\Models\{ItemRelations, Publishable};
use Laravel\Scout\Searchable;
use NumberFormatter;

/**
 * An Item of Apparel.
 *
 * @property string      $slug            The URL slug of an item.
 * @property string      $english_name    The English Title of an Item.
 * @property string      $foreign_name    The 'Japanese Title' of an Item.
 * @property int|null    $year            The year an Item was released.
 * @property string|null $product_number  An Item's product number.
 * @property int         $status          The status of an item (stored internally as an int).
 * @property string      $price           The price of this item, in a given currency.
 * @property float       $price_formatted The price of this item, formatted to the rules of the given currency (e.g. /100 for gbp/usd)
 * @property string      $currency        The currency of this item, as an ISO code.
 *
 * @property \App\Image    $image     The primary {@link \App\Image} for this Item.
 * @property \App\Category $category  The {@link \App\Category} of this Item (e.g. JSK).
 * @property \App\User     $submitter The {@link \App\User} who originally submitted this Item.
 * @property \App\User     $publisher The {@link \App\User} who published this Item.
 * @property \App\Brand    $brand     The {@link \App\Brand} of this Item (e.g. Angelic Pretty).
 *
 * @property \App\Image[]|\Illuminate\Database\Eloquent\Collection     $images     The {@link \App\Image images} for this Item.
 * @property \App\Tag[]|\Illuminate\Database\Eloquent\Collection       $tags       The {@link \App\Tag search tags} for this Item.
 * @property \App\Color[]|\Illuminate\Database\Eloquent\Collection     $colors     The {@link \App\Color colorways} this Item comes in (e.g. Black).
 * @property \App\Feature[]|\Illuminate\Database\Eloquent\Collection   $features   The {@link \App\Feature features} of this item (e.g. Back Shirring).
 * @property \App\Attribute[]|\Illuminate\Database\Eloquent\Collection $attributes The {@link \App\Attribute custom attributes} on this Item.
 * @property \App\User[]|\Illuminate\Database\Eloquent\Collection      $stargazers The {@link \App\User users} who want to own this Item.
 * @property \App\User[]|\Illuminate\Database\Eloquent\Collection      $owners     The {@link \App\Attribute users} who own this Item.
 * @property \App\Comment[]|\Illuminate\Database\Eloquent\Collection   $comments   The {@link \App\Comment comments} on this Item.
 *
 * @property string $image_id The ID of this Item's {@link \App\Image image}.
 * @property string $category_id  The ID of this Item's {@link \App\Category category}.
 * @property string $user_id  The ID of the {@link \App\User user} who submitted this Item.
 * @property string $brand_id The ID of this Item's {@link \App\Brand brand}.
 * @property string $submitter_id The ID of this Item's {@link \App\User submitter}.
 * @property string $publisher_id The ID of this Item's {@link \App\User publisher}.
 *
 * @property \Carbon\Carbon $published_at The date this item was published.
 */
class Item extends Model
{
    use ItemRelations, Publishable;

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
     */
    const DRAFT = 0;

    /**
     * Indicates that an item should be visible to everyone.
     */
    const PUBLISHED = 1;

    /**
     * A shortcut for fully eager loading an item.
     *
     * Use: `Item::with(Item::FULLY_LOAD)`
     */
    const FULLY_LOAD = [
        'image',
        'images',
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
    const PARTIAL_LOAD = [
        'submitter',
        'brand',
        'category',
        'image',
        'tags'
    ];

    /**
     * The name of the index this model is searchable under.
     *
     * @var string
     */
    protected $index = 'items';

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
     * An array of keys to convert to dates.
     *
     * @var array
     */
    protected $dates = ['published_at'];

    /**
     * Hidden attributes.
     *
     * @var array
     */
    protected $hidden = [
        'id',
        'brand_id',
        'category_id',
        'image_id',
        'publisher_id',
        'user_id',
    ];

    /**
     * Get the indexable data array for the model.
     *
     * @return array
     */
    public function toSearchableArray()
    {
        return [
            'id' => $this->id,

            // regular attributes for searching.
            'searchable_brand_short' => $this->brand->short_name,
            'searchable_english_name' => $this->english_name,
            'searchable_foreign_name' => $this->foreign_name,

            // searchable facets
            'searchable_brand' => $this->brand->name,
            'searchable_category' => $this->category->name,
            'searchable_tags' => $this->tags->pluck('name')->all(),
            'searchable_colors' => $this->colors->pluck('name')->all(),
            'searchable_features' => $this->features->pluck('name')->all(),
            'searchable_product_number' => $this->product_number,

            // special: numerical filters
            'year' => $this->year,
            'price' => $this->price_formatted,

            // special: facets
            '_tags' => $this->tags->pluck('slug')->all(),
            'brand' => $this->brand->slug,
            'category' => $this->category->slug,
            'features' => $this->features->pluck('slug')->all(),
            'colors' => $this->colors->pluck('slug')->all(),
            'currency' => $this->currency,

            // special: image url to describe this item
            'thumbnail_url' => $this->image->thumbnail_url ?? null,
        ];
    }

    /**
     * Get a list of facet search values.
     *
     * @return array
     */
    public function getSearchFilters()
    {
        return [
            '_tags',
            'brand',
            'category',
            'colors',
            'currency',
            'features',
        ];
    }

    /**
     * Get a list of searchable fields to set, in ranking order.
     *
     * @return string[]
     */
    public function getSearchableFields()
    {
        return [
            'searchable_brand_short',
            'searchable_english_name',
            'searchable_foreign_name',
            'searchable_product_number',
            'searchable_brand',
            'searchable_category',
            'searchable_colors',
            'searchable_tags',
            'searchable_features',
            'year',
        ];
    }

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

        return (string) round($this->price / 100, 2);
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
     * Get a slug for an item.
     *
     * @param \App\Item $item
     * @return string
     */
    public static function slug(self $item)
    {
        $candidate = $item->brand->short_name . '-' . str_slug($item->english_name);

        if (! static::where('slug', $candidate)->exists()) {
            return $candidate;
        }

        $attempts = -1;

        do {
            if ($attempts > 255) {
                throw new \RuntimeException("Too many items have the slug prefix [{$candidate}]");
            }

            $try = $candidate . '-' . ++$attempts;
        } while (static::where('slug', $try)->exists());

        return $try;
    }
}
