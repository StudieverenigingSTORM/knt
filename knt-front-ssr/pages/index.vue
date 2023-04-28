<template>
  <!-- d-flex align-items-center flex-column -->
  <div class="product-page">

    <!-- NAVIGATION TEST -->
    <!-- <router-link to="/products">Products</router-link> -->

    <!--Search Bar-->
    <input type="text" class="input user-search-bar" v-model="search" @input="onInputChange"
      placeholder="Search your name..." />

    <!-- align-self-start -->
    <div class="top-section">

      <!--Displaying users-->
      <NuxtLink
      :to="{ name: 'products', query:{user: user.id}}" 
        v-for="user in searchedUsers()" :key="user.id"
        class="d-flex user-card align-items-center border-bottom border-2 "
        >
        <div class="p-3">
          <h5>{{ user.firstName + " " + user.lastName }}</h5>
        </div>
      </NuxtLink >
    </div>

    <div class="bottom-section d-flex align-items-center flex-column">

      <div class="keyboard d-flex align-items-center">
        <SimpleKeyboard @onChange="onChange" @onKeyPress="onKeyPress" :input="input" />
      </div>

    </div>

  </div>
</template>

<script lang="ts">
  import { defineComponent } from "vue";
  import SimpleKeyboard from "../components/SimpleKeyboardUserPage.vue";

  interface User {
    id: number;
    firstName: string;
    lastName: string;
    balance: number;
  }

  export default defineComponent({

    components: {
      SimpleKeyboard
    },

    data() {
      return {
        search: "",
        input: "",
        users: useState<User[]>('userData'),
      };
    },

    methods: {
      onChange(input: string) {
        this.search = input;
      },
      onKeyPress(button: string) {
        console.log("button", button);
      },
      onInputChange(input: any) {
        this.search = input.target.value;
      },
      searchedUsers() {
        return this.users.filter(user => (user.firstName + " " + user.lastName).toLowerCase().includes(this.search.toLowerCase()))
      }
    },
  
  });
</script>
<script setup lang="ts">
  const runtimeConfig = useRuntimeConfig();

  if (process.server) {
    let userData: User[] = [];
    await useFetch('/users', {
      baseURL: runtimeConfig.public.apiBase,
      headers: {
        "X-API-Key": runtimeConfig.apiSecret
      },
      transform: function (data: string) {
        userData = JSON.parse(data)
      },
      server: true
    })
    console.log(userData)
    useState('userData', () => userData)
  }
</script>
<style>
.top-section,
.bottom-section {
  position: fixed;
  left: 0;
  right: 0;
}

.top-section {
  overflow: scroll;
  top: 55px;
  height: 45%;
}

.bottom-section {
  bottom: 0;
  height: 45%;
  display: flex;
  justify-content: center;
}

.header {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}

.navigation-button {
  display: inline-block;
  margin: 0.5rem 0.5rem 0;
}

.user-card{
  text-decoration: none; 
  color: inherit;
}

.user-card:hover {
  background-color: #e5e5e5;
  transition-timing-function: ease;
  transition: 0.5s;
  color: inherit;
}

.user-search-bar {
  display: block;
  margin: 0.5rem 0.5rem 0.5rem auto;
  padding: 5px;
  width: 20%;
}

.keyboard{
  width: 100%;
  height: 100%;
}
</style>