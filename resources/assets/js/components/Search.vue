<template>
    <div class="container">
        <!-- Search Bar -->
        <div class="row">
            <div class="col-sm-5 col-md-4 col-lg-3">
                <filter :name="Category" :values="categories" v-model="search.categories"></filter>
                <filter :name="Brand"    :values="brands"     v-model="search.brands"></filter>
                <filter :name="Features" :values="features"   v-model="search.features"></filter>
                <filter :name="Colorway" :values="colors"     v-model="search.colors"></filter>
                <filter :name="Tags" :values="tags"       v-model="search.tags"></filter>
                <filter :name="Year"     :values="years"      v-model="search.years"></filter>
            </div>
            <div class="col">
                <search-bar v-model="search.search" @input="typing"></search-bar>

                <div class="row">
                    <!-- results -->
                </div>

                <pagination :limit="2" :data="results"></pagination>
            </div>
        </div>
    </div>
</template>

<script>
    import SearchBar from './SearchBar.vue';
    import Filter from './Filter.vue';
    import Pagination from './Pagination.vue';
    import { VueSelect } from 'vue-select';
    import query from '../search';
    import debounce from 'lodash/debounce';

    export default {
        props: ['categories', 'brands', 'features', 'tags', 'colors', 'years'],

        data() {
            return {
                results: {},
                search: {
                    search: '',
                    categories: [],
                    brands: [],
                    features: [],
                    tags: [],
                    colors: [],
                    years: [],
                },
            }
        },

        mounted() {
            this.update()
        },

        // when "typing", we want to debounce our search.
        typing() {
            debounce(this.search, 250)
        },

        // when anything else happens, we just want to search.
        update() {
            query(this.search).resolve();
        },

        components: {
            SearchBar,
            Filter,
            Pagination,
            VueSelect,
        }
    }
</script>
