Nova.booting((Vue, router, store) => {
    Vue.component('index-image-array', require('./components/IndexField'))
    Vue.component('detail-image-array', require('./components/DetailField'))
    Vue.component('form-image-array', require('./components/FormField'))
})
