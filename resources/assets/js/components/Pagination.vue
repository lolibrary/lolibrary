<template>
    <nav aria-label="Search results pages" v-if="data.total > data.per_page">
        <ul class="pagination justify-content-center">
            <li class="page-item" v-if="data.prev_page_url">
                <a class="page-link" href="#" @click.prevent="page(--data.current_page)">Previous</a>
            </li>
            <li class="page-item"
                :class="{ 'active': n == data.current_page }"
                v-for="n in pages()">
                <a class="page-link" href="#" @click.prevent="page(n)">
                    {{ n }}
                </a>
            </li>
            <li class="page-item" v-if="data.next_page_url">
                <a class="page-link" href="#" @click.prevent="page(++data.current_page)">Next</a>
            </li>
        </ul>
    </nav>
</template>

<script>
    export default {
        props: {
            data: {
                type: Object,
                default() {
                    return {
                        current_page: 1,
                        data: [],
                        from: 1,
                        last_page: 1,
                        next_page_url: null,
                        per_page: 10,
                        prev_page_url: null,
                        to: 1,
                        total: 0,
                    }
                }
            },
            limit: {
                type: Number,
                default() {
                    return 1
                },
                validator(value) {
                    return value >= 0
                }
            },
        },
        data() {
            //
        },

        methods: {
            page(page) {
                if (page === '...') {
                    return;
                }

                this.$emit('pagination-page-change', page);
            },
            pages() {
                if (this.limit === -1) {
                    return 0;
                }

                if (this.limit === 0) {
                    return this.data.last_page;
                }

                const current = this.data.current_page;
                const last = this.data.last_page;
                const delta = this.limit;
                const left = current - delta;
                const right = current + delta + 1;

                let range = [];

                for (let i = 1; i <= last; i++) {
                    if (i === 1 || i === last || (i >= left && i < right)) {
                        range.push(i);
                    }
                }

                let l;
                let pages = [];

                range.forEach(function (i) {
                    if (l) {
                        if (i - l === 2) {
                            pages.push(l + 1);
                        } else if (i - l !== 1) {
                            pages.push('...');
                        }
                    }

                    pages.push(i);

                    l = i;
                });

                return pages;
            }
        }
    }
</script>
