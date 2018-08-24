<template>
  <div class="container">
    <div class="row">
      <div class="col-sm-12 col-md-4 col-lg-3 mb-2 mb-3">

        <div class="card">
            <div class="card-header">
                Filters
            </div>
            <div class="card-body">
                <div class="input-group pb-2">
                    <label class="control-label">Category</label>
                    <v-select style="width: 100%" v-model="state.categories" :options="categories" label="name" multiple></v-select>
                </div>
                <div class="input-group pb-2">
                    <label class="control-label">Brand</label>
                    <v-select style="width: 100%" v-model="state.brands" :options="brands" label="name" multiple></v-select>
                </div>
                <div class="input-group pb-2">
                    <label class="control-label">Features</label>
                    <v-select style="width: 100%" v-model="state.features" :options="features" label="name" multiple></v-select>
                </div>
                <div class="input-group pb-2">
                    <label class="control-label">Colorway</label>
                    <v-select style="width: 100%" v-model="state.colors" :options="colors" label="name" multiple></v-select>
                </div>
                <div class="input-group pb-2">
                    <label class="control-label">Tags</label>
                    <v-select style="width: 100%" v-model="state.tags" :options="tags" label="name" multiple></v-select>
                </div>
                <div class="input-group pb-2">
                    <label class="control-label">Year</label>
                    <v-select style="width: 100%" v-model="state.years" :options="years" multiple></v-select>
                </div>
            </div>
        </div>

      </div>

      <div class="col-sm-12 col-md-8 col-lg-9">
          <div class="row mb-3">
              <div class="col px-2">
                <div class="card">
                    <div class="card-body pb-0 pt-3">
                        <div class="row">
                            <div class="col-md-8 col-lg-9 col-xl-10 mb-3">
                                <p class="sr-only">Type to search or filter</p>
                                <input autocomplete="off" v-model="state.search" class="form-control input-lg" type="text" name="search" placeholder="Type to filter items by name" role="search">
                            </div>
                            <div class="col-md-4 col-lg-3 col-xl-2 mb-3">
                                <button class="btn btn-block btn-outline-primary">Search</button>
                            </div>
                        </div>
                    </div>
                </div>
              </div>
          </div>

         <div class="row">
            <div class="col">
            <h2>Debug</h2>

            <div class="row">
              <pre class="col">{{ state }}</pre>
              <pre class="col">{{ search }}</pre>
              <pre class="col">{{ query }}</pre>
            </div>

            </div>
        </div>
      </div>

    </div>
  </div>
</template>

<script>
    import qs from 'qs';

    export default {
        props: {
          categories: {
            type: Array,
            default() {
              return [
                { slug: 'jsk', name: 'JSK' },
                { slug: 'op', name: 'OP' },
              ];
            }
          },
          features: {
            type: Array,
            default() {
              return [
                { slug: 'shirring', name: 'Shirring' },
              ];
            }
          },
          brands: {
            type: Array,
            default() {
              return [
                { slug: 'angelic-pretty', name: 'Angelic Pretty' },
                { slug: 'innocent-world', name: 'Innocent World' },
              ];
            }
          },
          colors: {
            type: Array,
            default() {
              return [
                { slug: 'black', name: 'Black' },
                { slug: 'white', name: 'White' },
                { slug: 'rose-x-gold', name: 'Rose x Gold' },
              ];
            }
          },
          years: {
            type: Array,
            default() {
              return [
                '2001', '2002',
              ];
            }
          },
          tags: {
            type: Array,
            default() {
              return [
                { slug: 'stars', name: 'Stars' },
                { slug: 'stripes', name: 'Stripes' },
                { slug: 'bears', name: 'Bears' },
              ];
            }
          },
          url: {
            type: String,
            default: '',
          },
        },

        data() {
            return {
                results: {},

                // a function so we can read the initial state from the url.
                state: {
                  search: undefined,
                  categories: [],
                  brands: [],
                  features: [],
                  tags: [],
                  colors: [],
                  years: [],
                },
            }
        },

        computed: {
          search() {
            return {
                search: this.state.search,
                categories: this.state.categories.map(obj => obj.slug),
                brands: this.state.brands.map(obj => obj.slug),
                features: this.state.features.map(obj => obj.slug),
                tags: this.state.tags.map(obj => obj.slug),
                colors: this.state.colors.map(obj => obj.slug),
                years: this.state.years.map(year => parseInt(year, 10)),
            }
          },

          query() {
            const query = qs.stringify(this.search, { encodeValuesOnly: true, arrayFormat: 'brackets', indices: false });

            window.history.pushState(this.search, query, this.url + '?' + query);

            return query;
          }
        },

        mounted() {
            //
        },

        components: {}
    }
</script>
