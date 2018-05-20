<?php

namespace App;

use App\Models\DateHandling;
use Illuminate\Database\Eloquent\Relations\Pivot as BasePivot;

/**
 * A custom pivot model for this application.
 *
 * @property \Carbon\Carbon $created_at
 * @property \Carbon\Carbon $updated_at
 */
class Pivot extends BasePivot
{
    use DateHandling;
}
