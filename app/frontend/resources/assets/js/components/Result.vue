<template>

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
          english_name: "",
          foreign_name: "",
          product_number: "",
          brand: {
            image: {name: "", thumbnail_url: "", url: ""},
            name: "",
            short_name: "",
            slug: "",
          },
          category: {name: "", slug: ""},
          tags: [],
          features: [],
          colors: [],
          attributes: [],
          year: null,
          image: {name: "", thumbnail_url: "/images/default.png", url: ""},
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
