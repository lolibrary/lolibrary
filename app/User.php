<?php

namespace App;

use App\Models\HasUuid;
use Illuminate\Notifications\Notifiable;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Laravel\Passport\HasApiTokens;
use NumberFormatter;

/**
 * A user of this application.
 *
 * @property string $email          The user's email.
 * @property string $name           The user's name.
 * @property string $username       The user's login username.
 * @property string $display_name   The user's displayed username.
 * @property string $remember_token A strong random number that allows the user to use "remember me" sessions.
 *
 * @property int  $level  The user's level (permissions).
 * @property bool $banned If the user is banned or not.
 *
 * @property \App\Image $image The user's profile image.
 * @property string $image_id The user's profile image ID.
 *
 * @property \App\Item[]|\Illuminate\Database\Eloquent\Collection $items    The {@link \App\Item items} this user has submitted.
 * @property \App\Item[]|\Illuminate\Database\Eloquent\Collection $wishlist The {@link \App\Item items} this user has favourited.
 * @property \App\Item[]|\Illuminate\Database\Eloquent\Collection $closet   The {@link \App\Item items} this user owns.
 * @property \App\Post[]|\Illuminate\Database\Eloquent\Collection $posts    The posts this user has created.
 * @property \App\Topic[]|\Illuminate\Database\Eloquent\Collection $topics   The topics this user has created.
 * @property \App\Message[]|\Illuminate\Database\Eloquent\Collection $messages The messages received by this user.
 * @property \App\Message[]|\Illuminate\Database\Eloquent\Collection $sent     The messages sent by this user.
 * @property \App\Comment[]|\Illuminate\Database\Eloquent\Collection $comments The comments this user has left.
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
     * Casts for attributes.
     *
     * @var array
     */
    protected $casts = [
        'banned' => 'boolean',
        'level' => 'integer',
    ];

    /**
     * Visible attributes.
     *
     * @var array
     */
    protected $visible = [
        'username',
        'profile',
        'created_at',
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
