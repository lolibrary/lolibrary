<template>
<div class="card">
    <div class="card-body text-center">
        <p class="mb-0 no-wrap" :title="item.english_name"
            style="white-space: nowrap; overflow-x: hidden; text-overflow: ellipsis;">
            <a :href="item.url">{{ item.english_name }}</a>
        </p>
        <p class="text-muted small"
            style="white-space: nowrap; overflow-x: hidden; text-overflow: ellipsis;"
            :title="item.foreign_name">
            {{ item.foreign_name ? item.foreign_name : '&nbsp;' }}
        </p>

        <div class="text-center item-image-container">
            <a :href="item.url">
                <img :src="item.image" class="mw-100 mh-100 rounded my-auto mx-auto" onerror="if (this.src !== window.defaultImage) this.src = window.defaultImage">
            </a>
        </div>
    </div>
    <ul class="list-group list-group-flush">
        <li class="list-group-item py-1 px-3">
            <div class="row small text-muted">
                <p class="col m-0 text-left">Brand</p>
                <p class="col m-0 text-right">Category</p>
            </div>
            <div class="d-flex small">
                <p class="p-0 m-0 text-left flex-fill" style="white-space: nowrap; overflow-x: ellipsis;">
                    <a :href="item.brand.url" :title="item.brand.name">
                        {{ item.brand.name.substring(0, 21) }}
                    </a>
                </p>
                <p class="p-0 m-0 text-right flex-fill" style="white-space: nowrap; overflow-x: hidden;">
                    <a :href="item.category.url" :title="item.category.name">
                        {{ item.category.name.substring(0, 9) }}
                    </a>
                </p>
            </div>
        </li>
    </ul>
    <a class="btn btn-outline-primary rounded-0" style="border: none;" :href="item.url">
        View Item
    </a>
</div>

</template>

<script>
export default {
  props: {
    item: {
      type: Object,
      default() {
        return {
          id: "",
          slug: "",
          url: "",
          english_name: "",
          foreign_name: "",
          product_number: "",
          brand: {
            image: "",
            name: "",
            short_name: "",
            slug: "",
          },
          category: {name: "", slug: "", image: "", url: ""},
          tags: [],
          features: [],
          colors: [],
          attributes: [],
          year: null,
          image: "",
          notes: "",
          price_details: {
            currency: "jpy",
            formatted: "",
            local_price: "",
            price: 0,
          },
          created_at: null,
          published_at: null,
          updated_at: null,
        };
      },
    },
  },


};
</script>
