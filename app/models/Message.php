<?php

namespace App\Models;

use Illuminate\Database\Eloquent\SoftDeletes;

/**
 * A private message between two people.
 *
 * In future, this could use a many-to-many relationship for group messaging.
 *
 * @property string               $sender_id
 * @property string               $recipient_id
 *
 * @property \App\User            $sender
 * @property \App\User            $recipient
 *
 * @property string               $message
 * @property bool                 $read
 *
 * @property \Carbon\Carbon|null  $deleted_at
 */
class Message extends Model
{
    use SoftDeletes;

    /**
     * Which attributes can be filled on this model via mass assignment.
     *
     * @var array
     */
    protected $fillable = ['message'];

    /**
     * An array of casts for attributes.
     *
     * @var array
     */
    protected $casts = ['read' => 'boolean'];

    /**
     * The user who sent the message.
     *
     * @return \Illuminate\Database\Eloquent\Relations\HasOne
     */
    public function sender()
    {
        return $this->hasOne(User::class, 'sender_id');
    }

    /**
     * The target of the message.
     *
     * @return \Illuminate\Database\Eloquent\Relations\HasOne
     */
    public function recipient()
    {
        return $this->hasOne(User::class, 'recipient_id');
    }

    /**
     * Mark this message as read.
     *
     * @return void
     */
    public function read()
    {
        $this->read = true;
        $this->save();
    }
}
