{{
  collect(array_reverse(range(1990, date('Y') + 1)))->map(function ($year) {
    return (string) $year;
  })->toJson()
}}
