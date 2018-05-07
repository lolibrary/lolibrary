<select multiple class="form-control form-control-chosen">
    @foreach ($items as $key => $value)
        <option value="{{ $key }}">{{ $value }}</option>
    @endforeach
</select>
