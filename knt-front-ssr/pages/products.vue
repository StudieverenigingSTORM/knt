<template>
    <!-- <h1>Welcome to Product Page</h1> -->
    <!-- Test for going back -->
    <!-- <router-link to="/">Back</router-link> -->
    <!-- <p>User id: {{ $route.query.user }}</p> -->

    <!-- Top section -->
    <div class="product-page-section d-flex flex-row">
        <div class="left-section">
            <div class="header-section d-flex justify-content-between">
                <h5 class="align-self-center username-display">Test user | Balance: $$$</h5>
                <button type="button" class="btn btn-primary">Back</button>
            </div>

            <!-- POPULATED WITH DUMMY VALUES -->
            <div class="product-section border border-2">
                <div>
                    <div class="p-3 product-item border-bottom border-2" v-for="product in products">
                        <h5>{{ product.name + ": " + product.price / 100 + "€" }}</h5>
                    </div>
                </div>
            </div>
        </div>

        <div class="right-section">
            <div class="header-section d-flex justify-content-center">
                <h5 class="align-self-center">Total: $$$</h5>
            </div>
            <!-- POPULATED WITH DUMMY VALUES -->
            <div class="cart-section">
                <div class="d-flex flex-row justify-content-between">
                    <h5 class="p-3 w-50">Product1:</h5>
                    <h5 class="p-3">qty</h5>
                    <div class="p-3 w-50 d-flex justify-content-end">
                        <button type="button" class="btn btn-primary mr-auto">
                            <h5>+</h5>
                        </button>
                        <button type="button" class="btn btn-danger mr-auto">
                            <h5>-</h5>
                        </button>
                    </div>
                </div>

            </div>
            <!-- Purchase Button -->
            <div class="purchase-section d-flex align-items-center justify-content-center border-top border-2">
                <button type="button" class="purchase-button btn btn-secondary">Purchase</button>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
interface User {
    id: number;
    firstName: string;
    lastName: string;
    balance: number;
}

interface Product {
    id: number;
    price: number;
    name: string;
}

import { defineComponent } from "vue";

export default defineComponent({
    data() {
        return {
            products: useState<Product[]>('productData'),
        };
    },

    mounted() {
        console.log(this.$route.query.user);
    }
})
</script>


<script setup lang="ts">
//Get Products
const runtimeConfig = useRuntimeConfig();

if (process.server) {
    let productData: Product[] = [];
    await useFetch('/users/products', {
        baseURL: runtimeConfig.public.backendBase,
        headers: {
            "X-API-Key": runtimeConfig.apiSecret
        },
        transform: function (data: string) {
            productData = JSON.parse(data)
        },
        server: true
    })
    //console.log(productData)
    useState('productData', () => productData)
}
</script>

<style>
.left-section,
.right-section {
    width: 50vw;
}

.header-section {
    height: 50px;
}

.username-display {
    padding-left: 1rem;
}

.product-section {
    height: calc(100vh - 50px);
    overflow: scroll;
}

.purchase-section {
    height: calc(25vh);
}

.cart-section {
    height: calc(75vh - 50px);
    overflow: scroll;
}

.product-item:hover {
    background-color: #e5e5e5;
    transition-timing-function: ease;
    transition: 0.5s;
    color: inherit;
}
</style>