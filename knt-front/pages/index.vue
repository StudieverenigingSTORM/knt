<template>
    <div>
        <b-card v-for="product in products" :key="product.id" :title="product.name">
            Price: {{product.price}}
        </b-card>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
    name: 'IndexPage',
    data() {
        return {
            products: null
        }
    },
    methods: {
        getProducts() {
            this.$axios.get('/api/products').then((response) => {
                this.products = response.data
            });
        }
    },
    mounted() {
        this.getProducts()
    }
})
</script>
