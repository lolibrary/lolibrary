<?php

namespace App;

use App\Models\HasUuid;
use Illuminate\Notifications\Notifiable;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Laravel\Passport\HasApiTokens;
use NumberFormatter;

/**
 * An Item of Apparel.
 *
 * @property string $slug           The URL slug of an item.
 * @property string $english_name   The English Title of an Item.
 * @property string $foreign_name   The 'Japanese Title' of an Item.
 * @property int|null $year           The year an Item was released.
 * @property string|null $product_number An Item's product number.
 * @property int $status         The status of an item (stored internally as an int).
 * @property string $price          The price of this item, in a given currency.
 * @property string $currency       The currency of this item, as an ISO code.
 *
 * @property \App\Image $image     The primary {@link \App\Image} for this Item.
 * @property \App\Category $category  The {@link \App\Category} of this Item (e.g. JSK).
 * @property \App\User $submitter The {@link \App\User} who originally submitted this Item.
 * @property \App\User $publisher The {@link \App\User} who published this Item.
 * @property \App\Brand $brand     The {@link \App\Brand} of this Item (e.g. Angelic Pretty).
 *
 * @property \App\Image[]|\Illuminate\Database\Eloquent\Collection $images     The {@link \App\Image images} for this Item.
 * @property \App\Tag[]|\Illuminate\Database\Eloquent\Collection $tags       The {@link \App\Tag search tags} for this Item.
 * @property \App\Color[]|\Illuminate\Database\Eloquent\Collection $colors     The {@link \App\Color colorways} this Item comes in (e.g. Black).
 * @property \App\Feature[]|\Illuminate\Database\Eloquent\Collection $features   The {@link \App\Feature features} of this item (e.g. Back Shirring).
 * @property \App\Attribute[]|\Illuminate\Database\Eloquent\Collection $attributes The {@link \App\Attribute custom attributes} on this Item.
 * @property \App\User[]|\Illuminate\Database\Eloquent\Collection $stargazers The {@link \App\User users} who want to own this Item.
 * @property \App\User[]|\Illuminate\Database\Eloquent\Collection $owners     The {@link \App\Attribute users} who own this Item.
 * @property \App\Comment[]|\Illuminate\Database\Eloquent\Collection $comments   The {@link \App\Comment comments} on this Item.
 *
 * @property string $image_id The ID of this Item's {@link \App\Image image}.
 * @property string $type_id  The ID of this Item's {@link \App\Type type}.
 * @property string $user_id  The ID of the {@link \App\User user} who submitted this Item.
 * @property string $brand_id The ID of this Item's {@link \App\Brand brand}.
 * @property string $submitter_id The ID of this Item's {@link \App\User submitter}.
 * @property string $publisher_id The ID of this Item's {@link \App\User publisher}.
 *
 * @property \Carbon\Carbon $published_at The date this item was published.
 */
class User extends Authenticatable
{
    use Notifiable, HasApiTokens, HasUuid;

    public const DEVELOPER = 1000;
    public const ADMIN = 500;
    public const SENIOR_LOLIBRARIAN = 100;
    public const LOLIBRARIAN = 50;
    public const JUNIOR_LOLIBRARIAN = 10;
    public const REGULAR = 0;
    public const BANNED = -1;

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'name',
        'username',
        'email',
        'password',
    ];

    /**
     * The attributes excluded from the model's JSON form.
     *
     * @var array
     */
    protected $hidden = [
        'id',
        'password',
        'remember_token',
        'email',
        'auth',
        'url',
        'level',
        'banned',
        'updated_at',
        'created_at',
        'image_id',
        'profile',
    ];

    /**
     * Casts for attributes.
     *
     * @var array
     */
    protected $casts = [
        'banned' => 'boolean',
        'auth' => 'json',
        'level' => 'integer',
    ];

    /**
     * The items a user has submitted.
     *
     * @return \Illuminate\Database\Eloquent\Relations\HasMany
     */
    public function items()
    {
        return $this->hasMany(Item::class);
    }

    /**
     * The items a user has favourited/wishlisted.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsToMany|\App\Item[]
     */
    public function wishlist()
    {
        return $this->belongsToMany(Item::class, 'wishlist')->withTimestamps();
    }

    /**
     * The items a user owns.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsToMany|\App\Item[]
     */
    public function closet()
    {
        return $this->belongsToMany(Item::class, 'closet')->withTimestamps();
    }

    /**
     * The posts a user has.
     *
     * @return \Illuminate\Database\Eloquent\Relations\HasMany|\App\Post[]
     */
    public function posts()
    {
        return $this->hasMany(Post::class);
    }

    /**
     * The topics a user has created
     *
     * @return \Illuminate\Database\Eloquent\Relations\HasMany|\App\Topic[]
     */
    public function topics()
    {
        return $this->hasMany(Topic::class);
    }

    /**
     * The profile image for a user.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsTo|\App\Image
     */
    public function image()
    {
        return $this->belongsTo(Image::class);
    }

    /**
     * All messages for this user.
     *
     * @return \Illuminate\Database\Eloquent\Relations\HasMany|\App\Message[]
     */
    public function messages()
    {
        return $this->hasMany(Message::class, 'recipient_id');
    }

    /**
     * Sent messages for this user.
     *
     * @return \Illuminate\Database\Eloquent\Relations\HasMany|\App\Message[]
     */
    public function sent()
    {
        return $this->hasMany(Message::class, 'sender_id');
    }

    /**
     * The comments a user has left on items.
     *
     * @return \Illuminate\Database\Eloquent\Relations\HasMany|\App\Comment[]
     */
    public function comments()
    {
        return $this->hasMany(Comment::class);
    }

    /**
     * Get a user's profile.
     *
     * @return \Illuminate\Database\Eloquent\Relations\HasOne|\App\Profile
     */
    public function profile()
    {
        return $this->hasOne(Profile::class);
    }
}
