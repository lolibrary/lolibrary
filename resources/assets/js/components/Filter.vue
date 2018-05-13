<template>
    <div class="input-group" style="padding-bottom: 10px">
        <label class="control-label">{{ __('Category') }}</label>
        <select multiple
                class="form-control form-control-chosen form-control-filter"
                data-placeholder="Tap or type to select"
        >
            @foreach ($items as $key => $value)
            <option value="{{ $key }}">{{ $value }}</option>
            @endforeach
        </select>
    </div>


</template>