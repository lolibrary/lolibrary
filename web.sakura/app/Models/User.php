<?php

namespace App\Models;

use App\Models\Traits\{
    HasUuid, DateHandling, Wishlist, Closet, AccessLevels
};
use Illuminate\Support\Facades\DB;
use Laravel\Passport\HasApiTokens;
use Illuminate\Notifications\Notifiable;
use Illuminate\Database\Eloquent\Builder;
use Illuminate\Contracts\Auth\MustVerifyEmail;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Laravel\Nova\Actions\Actionable;

/**
 * A user of this application.
 *
 * @property string $email          The user's email.
 * @property string $name           The user's name.
 * @property string $username       The user's login username.
 * @property string $remember_token A strong random number that allows the user to use "remember me" sessions.
 *
 * @property int  $level    The user's level (permissions).
 * @property bool $banned   If the user is banned or not.
 * @property bool $verified Whether or not the user's email has been verified.
 *
 * @property \App\Models\Image $image The user's profile image.
 * @property string $image_id         The user's profile image ID.
 *
 * @property \App\Models\Item[]|\Illuminate\Database\Eloquent\Collection $items    The {@link \App\Item items} this user has submitted.
 * @property \App\Models\Item[]|\Illuminate\Database\Eloquent\Collection $wishlist The {@link \App\Item items} this user has favourited.
 * @property \App\Models\Item[]|\Illuminate\Database\Eloquent\Collection $closet   The {@link \App\Item items} this user owns.
 * @property \App\Models\Post[]|\Illuminate\Database\Eloquent\Collection $posts    The posts this user has created.
 */
class User extends Authenticatable implements MustVerifyEmail
{
    use Notifiable, HasApiTokens, HasUuid, DateHandling, Wishlist, Closet, AccessLevels, Actionable;

    public const DEVELOPER = 1000;
    public const ADMIN = 500;
    public const SENIOR_LOLIBRARIAN = 100;
    public const LOLIBRARIAN = 50;
    public const JUNIOR_LOLIBRARIAN = 10;
    public const REGULAR = 0;
    public const BANNED = -1;

    /**
     * Whether or not this model has an incrementing timestamp.
     *
     * @var bool
     */
    public $incrementing = false;

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
        'email_verified_at' => 'datetime',
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
     * The attributes that should be hidden for arrays.
     *
     * @var array
     */
    protected $hidden = [
        'password',
        'remember_token',
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
     * @return \Illuminate\Database\Eloquent\Relations\BelongsToMany|\App\Models\Item[]
     */
    public function wishlist()
    {
        return $this->belongsToMany(Item::class, 'wishlist')->withTimestamps();
    }

    /**
     * The items a user owns.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsToMany|\App\Models\Item[]
     */
    public function closet()
    {
        return $this->belongsToMany(Item::class, 'closet')->withTimestamps();
    }

    /**
     * The posts a user has.
     *
     * @return \Illuminate\Database\Eloquent\Relations\HasMany|\App\Models\Post[]
     */
    public function posts()
    {
        return $this->hasMany(Post::class);
    }

    /**
     * The profile image for a user.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsTo|\App\Models\Image
     */
    public function image()
    {
        return $this->belongsTo(Image::class);
    }

    /**
     * Get a user's profile.
     *
     * @return \Illuminate\Database\Eloquent\Relations\HasOne|\App\Models\Profile
     */
    public function profile()
    {
        return $this->hasOne(Profile::class);
    }

    /**
     * Scope a query to email address.
     *
     * @param \Illuminate\Database\Eloquent\Builder $query
     * @param string $email
     * @return \Illuminate\Database\Eloquent\Builder|\Illuminate\Database\Query\Builder
     */
    public function scopeEmail(Builder $query, string $email)
    {
        return $query->where(DB::raw('lower(email)'), mb_strtolower($email));
    }

}
