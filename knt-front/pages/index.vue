<template>
    <h1>{{res.greeting}} {{res.greeting2}}</h1>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
    name: 'IndexPage',
    data() {
        return {
        res: null
        }
    },
    methods: {
        getGreetings() {
        this.$axios.get('/hello').then((response) => {
            this.res = response.data;
            console.log(this.res);
        });
        }
    },
    mounted() {
        this.getGreetings();
    }
})
</script>
