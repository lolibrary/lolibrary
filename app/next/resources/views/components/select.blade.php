<select multiple
        class="form-control form-control-chosen form-control-filter"
        data-placeholder="Tap or type to select"
        id="{{ $id }}"
>
    @foreach ($items as $key => $value)
        <option value="{{ $key }}">{{ $value }}</option>
    @endforeach
</select>
