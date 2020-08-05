<template>
    <div >
        <h3 >Movie list:</h3>
        <em v-if="users.loading">Loading users...</em>
        <span v-if="users.error" class="text-danger">ERROR: {{users.error}}</span>
        <form @submit.prevent="handleSubmit" >
            <div class="card-group">
                <div class="card" v-for="user in users.items" :key="user.id">
                    <img v-bind:src="user.thumb" :class="card-img-top" /> 
                    <div class="card-body">
                        <h5 class="card-title">{{user.name}}</h5>
                        <p class="card-text">{{user.amount}}$</p>
                        <button class="btn btn-primary" >Watch</button>
                    </div>
                </div>
            </div>
        </form>
        <p>
            <router-link class="btn  btn-primary" to="/login">Logout</router-link>
        </p>
    </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
    computed: {
        ...mapState({
            account: state => state.account,
            users: state => state.users.all
        })
    },
    created () {
        this.getAllUsers();
    },
    methods: {
        ...mapActions('users', {
            getAllUsers: 'getAll',
        }),
         handleSubmit (e) {
            this.submitted = true;
            const { username, password } = this;
            if (username && password) {
                this.login({ username, password })
            }
        }
    }
};
</script>

<style scoped>
    .card-group {
        padding: 10px;
    }
    .title {
        font-size:2em;
    }
    li {
        font-size:1em;
    }
</style>