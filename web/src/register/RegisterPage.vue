<template>
    <div>
        <h2 class="title">Register</h2>
        <form @submit.prevent="handleSubmit">
            <div class="container">
                <div class="form-group">
                    <label for="email">Email</label>
                    <input type="text" v-model="user.email" v-validate="'required|email'" name="email" class="input-lg form-control" :class="{ 'is-invalid': submitted && errors.has('email') }" />
                    <div v-if="submitted && errors.has('email')" class="invalid-feedback">{{ errors.first('email') }}</div>
                </div>
                <div class="form-group">
                    <label htmlFor="password">Password</label>
                    <input type="password" v-model="user.password" v-validate="{ required: true }" name="password" class="input-lg form-control" :class="{ 'is-invalid': submitted && errors.has('password') }" />
                    <div v-if="submitted && errors.has('password')" class="invalid-feedback">{{ errors.first('password') }}</div>
                </div>
            </div>
            <div v-if="users.plan" class="container">
                <div class="card-deck mb-3 text-center">
                    <div class="card mb-4 box-shadow" v-for="plan in users.plan" :key="plan.id">
                        <div class="card-header">
                            <h4 class="card-title">{{plan.lookup_key}}</h4>
                        </div>
                        <div class="card-body">
                            <h5 class="card-title">${{(plan.unit_amount)/100}} <span style="color:gray">/ month</span></h5>
                            <p class="card-text">This is a wider card with supporting text below as a natural lead-in to additional content.</p>
                            <button v-on:click="user.plan=plan.lookup_key" class="btn btn-lg btn-block btn-outline-primary " :disabled="status.registering">Register</button>
                            <img v-show="status.registering" src="data:image/gif;base64,R0lGODlhEAAQAPIAAP///wAAAMLCwkJCQgAAAGJiYoKCgpKSkiH/C05FVFNDQVBFMi4wAwEAAAAh/hpDcmVhdGVkIHdpdGggYWpheGxvYWQuaW5mbwAh+QQJCgAAACwAAAAAEAAQAAADMwi63P4wyklrE2MIOggZnAdOmGYJRbExwroUmcG2LmDEwnHQLVsYOd2mBzkYDAdKa+dIAAAh+QQJCgAAACwAAAAAEAAQAAADNAi63P5OjCEgG4QMu7DmikRxQlFUYDEZIGBMRVsaqHwctXXf7WEYB4Ag1xjihkMZsiUkKhIAIfkECQoAAAAsAAAAABAAEAAAAzYIujIjK8pByJDMlFYvBoVjHA70GU7xSUJhmKtwHPAKzLO9HMaoKwJZ7Rf8AYPDDzKpZBqfvwQAIfkECQoAAAAsAAAAABAAEAAAAzMIumIlK8oyhpHsnFZfhYumCYUhDAQxRIdhHBGqRoKw0R8DYlJd8z0fMDgsGo/IpHI5TAAAIfkECQoAAAAsAAAAABAAEAAAAzIIunInK0rnZBTwGPNMgQwmdsNgXGJUlIWEuR5oWUIpz8pAEAMe6TwfwyYsGo/IpFKSAAAh+QQJCgAAACwAAAAAEAAQAAADMwi6IMKQORfjdOe82p4wGccc4CEuQradylesojEMBgsUc2G7sDX3lQGBMLAJibufbSlKAAAh+QQJCgAAACwAAAAAEAAQAAADMgi63P7wCRHZnFVdmgHu2nFwlWCI3WGc3TSWhUFGxTAUkGCbtgENBMJAEJsxgMLWzpEAACH5BAkKAAAALAAAAAAQABAAAAMyCLrc/jDKSatlQtScKdceCAjDII7HcQ4EMTCpyrCuUBjCYRgHVtqlAiB1YhiCnlsRkAAAOwAAAAAAAAAAAA==" />
                        </div>
                    </div>
                </div>
            </div>
        </form>
    </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
    data () {
        return {
            user: {
                email: '',
                password: '',
                plan: '',
            },
            submitted: false
        }
    },
    created () {
        this.getAllPlans();
    },
    computed: {
        ...mapState( 'account', ['status']),
        ...mapState({
            users: state => state.users.all,
        })
    },
    methods: {
                
        ...mapActions('users', {
            getAllPlans: 'getPlans',
        }),

        ...mapActions('account', ['register']),
        handleSubmit(e) {
            this.submitted = true;
            this.$validator.validate().then(valid => {
                if (valid) {
                    console.log(this.user)
                    this.register(this.user);
                }
            });
        }
    }
};
</script>
<style scoped>
    .margin {
        margin-bottom: 20px;
    }
</style>