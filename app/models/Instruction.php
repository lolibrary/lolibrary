<?php

namespace App\Models;

class Instruction extends Model
{
    /**
     * The values we're allowed to fill here.
     *
     * @var array
     */
    protected $fillable = [
        'slug',
        'description',
    ];

    /**
     * Get the items that have this care instruction.
     *
     * @return \Illuminate\Database\Eloquent\Relations\BelongsToMany
     */
    public function items()
    {
        return $this->belongsToMany(Item::class);
    }
}
