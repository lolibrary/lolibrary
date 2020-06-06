{{
  collect($tags)->map(function ($name, $slug) {
    return compact('slug', 'name');
  })->values()->toJson()
}}
