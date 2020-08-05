<template>
    <div class="jumbotron vertical-center">
        <div class="container">
            <div class="row center">
                <div class="col-sm-9 offset-sm-2">
                    <div class="alert" v-if="alert.message" :class="`alert ${alert.type}`">{{alert.message}}</div>
                    <router-view></router-view>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
    name: 'app',
    computed: {
        ...mapState({
            alert: state => state.alert
        })
    },
    methods: {
        ...mapActions({
            clearAlert: 'alert/clear' 
        })
    },
    watch: {
        $route (to, from){
            // clear alert on location change
            this.clearAlert();
        }
    } 
};
</script>

<style>

.vertical-center {
  min-height: 100%;  /* Fallback for browsers do NOT support vh unit */
  min-height: 100vh; /* These two lines are counted as one :-)       */
  display: flex;
  background-image: url("../../static/img/background.jpg");
  background-size: cover;
  background-repeat: no-repeat;
}
.center {
    margin-top: 200px;
}
.alert {
    font-size: 22px;
    margin-bottom: 20px;
}

</style>