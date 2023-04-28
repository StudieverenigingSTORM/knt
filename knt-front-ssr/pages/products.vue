<template>
<h1>Welcome to Product Page</h1>
<!-- Test for going back -->
<router-link to="/">Back</router-link>
<p>User id: {{ $route.query.user }}</p>

</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
    mounted(){
        console.log(this.$route.query.token);
    }
})
</script>

<style>
</style>