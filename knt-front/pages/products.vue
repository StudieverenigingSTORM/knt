<template>
    <!--    Navigation bar-->
    <div class="product-page">
        <b-navbar class="header">
            <b-navbar-nav class="navigation-bar">
                <b-button v-b-modal.modal-cart class="navigation-button" href="#" variant="primary">
                    Cart
                </b-button>
                <b-button class="navigation-button">
                    Buy
                </b-button>
            </b-navbar-nav>
        </b-navbar>

        <!--Display cart-->
        <b-modal id="modal-cart" title="Products in Cart">

            <p v-for="line in cart.lines">
                {{ line.product.itemName }} : {{ line.qty }}
            </p>
            <p>Total: {{cart.total / 100}} €</p>

        </b-modal>

        <!--Display products-->
<!--        <b-card v-for="product in products" :key="product.id" :title="product.name + product.price" class="text-center">-->
<!--&lt;!&ndash;            Price: {{ product.price }} €&ndash;&gt;-->
<!--            <b-button v-on:click="addProductToCart(product.name, product.id, product.price)"-->
<!--                      href="#"-->
<!--                      variant="primary"-->
<!--                      class="btn btn-primary float-right align-top">-->
<!--                Add to cart-->
<!--            </b-button>-->
<!--        </b-card>-->
    <div v-for="product in products" :key="product.id" class="d-flex align-items-center border-bottom border-2">
        <div class="p-3"><h5>{{product.name}}</h5></div>
        <div class="p-3"><h5>{{product.price}} €</h5></div>
        <div class="ml-auto p-3">
            <b-button v-on:click="addProductToCart(product.name, product.id, product.price)"
                      href="#"
                      variant="primary"
                      class="btn btn-primary float-right align-top">
                Add to cart
            </b-button>
        </div>
    </div>

    </div>
</template>

<script lang="ts">
import Vue from 'vue';
interface Line {
    product: Product
    qty: number;
}
interface Cart {
    lines: Line[];
    total: number;
}
interface Product {
    itemName: string;
    sku: number;
    cost: number;
}
//cryptography library javascropt mdn
const ProductList: Product[] = []
JSON.stringify({ products: ProductList });
export default Vue.extend({
    name: 'IndexPage',
    data() {
        return {
            products: null,
            cart: {lines: [], total: 0} as Cart,
            viewCart: false,
        }
    },
    mounted() {
        this.getProducts()
    },
    methods: {
        getProducts() {
            this.$axios.get('/products').then((response) => {
                this.products = response.data
            })
        },
        addProductToCart(productName: string, productSku: number, productPrice: number) {
            let newProduct: Product = {itemName: productName, sku:productSku, cost:productPrice};
            const index = this.cart.lines.findIndex(line => line.product.itemName === productName);
            if (index > -1) {
                this.cart.lines[index].qty += 1;
                this.cart.total += this.cart.lines[index].product.cost * 100;
                return
            }
            let newLine = {product: newProduct, qty: 1};
            this.cart.total += productPrice * 100;
            this.cart.lines.push(newLine);
        }
    }
})
</script>

<style>
.header {
    display: flex;
    justify-content: end;
    align-items: center;
}
.navigation-button {
    display: inline-block;
    margin: 0.5rem 0.5rem 0;
}
</style>