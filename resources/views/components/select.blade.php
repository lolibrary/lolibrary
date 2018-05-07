<select multiple class="form-control form-control-chosen" data-placeholder="Click or type to select">
    @foreach ($items as $key => $value)
        <option value="{{ $key }}">{{ $value }}</option>
    @endforeach
</select>
