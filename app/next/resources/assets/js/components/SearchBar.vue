<template>
    <div class="card">
        <div class="card-body">
            <input class="form-control form-control-lg"
                   placeholder="Start typing to search..."
                   role="search"
                   v-model="search"
            >
        </div>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                search: '',
            }
        }
    }
</script>