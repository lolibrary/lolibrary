<template>
    <div class="input-group" style="padding-bottom: 10px">
        <label class="control-label">Category</label>
        <select multiple
                class="form-control form-control-chosen form-control-filter"
                data-placeholder="Tap or type to select"
        >
            <option v-for="option in options" :value="option.value">{{ option.text }}</option>
        </select>
    </div>
</template>
