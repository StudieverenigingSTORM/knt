<template>
    <div class="root">

        <input type="text" class="user-search-bar" v-model="search" placeholder="Search your name..."/>

        <div v-for="user in searchedUsers" :key="user.id">
            <b-card v-b-modal.modal-1 :title="user.firstName + ' ' + user.lastName" class="user-card"
                    v-on:click="updateUser(user.firstName, user.lastName, user.password); userSelected = !userSelected"></b-card>
        </div>

        <!--     Section for th login pop-up-->
        <b-modal id="modal-1" title="Log in" @show="resetModal" @ok.prevent="handleOK">
            <b-container fluid>
                <b-row class="my-1">
                    <b-col sm="3">
                        <label for="username-input">Username: </label>
                    </b-col>
                    <b-col sm="9">
                        <!--                     TODO: Come up with something better-->
                        <p v-if="userSelected">{{ getSelectedUser() }}</p>
                        <p v-else>{{ getSelectedUser() }}</p>
                    </b-col>
                </b-row>
                <b-row class="my-1">
                    <b-col sm="3">
                        <label for="pincode-input">Pin-code: </label>
                    </b-col>
                    <b-col sm="9">
                        <b-form-input
                            type="password"
                            v-model="inputPassword"
                            id="pincode-input"
                            :state="passwordState"
                            placeholder="Enter your pin-code"
                            @keyup.enter="handleOK()"/>
                        <b-form-invalid-feedback id="input-live-feedback">Incorrect password!</b-form-invalid-feedback>
                    </b-col>
                </b-row>
            </b-container>
        </b-modal>

    </div>
</template>

<script lang="ts">
import Vue from 'vue'

interface User{
    username: string;
    password: string;
}

export default Vue.extend({
    name: 'IndexPage',
    data() {
        return {
            search: '',
            selectedUser: {} as User,
            userSelected: false,
            inputPassword: '',
            passwordState: (null as any),
            users: [],
        };
    },
    mounted() {
        this.getUsers()
    },
    methods: {
        getUsers() {
            this.$axios.get('/users').then((response) => {
                this.users = response.data
            })
        },

        getSelectedUser() {
            return this.selectedUser.username;
        },

        updateUser(firstName: string, lastName: string, password: string) {
            this.selectedUser.username = firstName + ' ' + lastName;
            this.selectedUser.password = password;
        },

        handleOK(){
            console.log("Actual password: ");
            console.log(this.selectedUser.password);
            console.log("Input: ");
            console.log(this.inputPassword);
            
    
            if(this.inputPassword == this.selectedUser.password){
                this.passwordState = true
                console.log("Logged in successfully")
                this.$router.push('/products')
            }
            else{
                this.passwordState = false
                console.log("Incorrect password")
            }
        },

        resetModal(){
            this.inputPassword = ''
            this.passwordState = null
        }
    },

    computed: {
        searchedUsers() {
            return (this as any).users.filter((user: any) => {
                let temp = (user.first_name + " " + user.last_name).toLowerCase()
                return temp.includes(this.search.toLowerCase())
            })
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
.user-card:hover {
    background-color: #e5e5e5;
    transition-timing-function: ease;
    transition: 0.5s;
}
.user-search-bar {
    display: block;
    margin: 0.5rem 0.5rem 0.5rem auto;
    padding: 5px;
    width: 20%;
}
</style>