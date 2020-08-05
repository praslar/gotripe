<template>
  <stripe-checkout
    ref="checkoutRef"
    :pk="publishableKey"
    :items="items"
    :successUrl="successUrl"
    :cancelUrl="cancelUrl"
  >
    <template slot="checkout-button">
      <button @click="checkout">Shut up and take my money!</button>
    </template>
  </stripe-checkout>
</template>
 

<script>
import { StripeCheckout } from 'vue-stripe-checkout';

export default {
  components: {
    StripeCheckout
  },

  data: () => ({
    loading: false,
    publishableKey: process.env.PUBLISHABLE_KEY,
    successUrl: 'your-success-url',
    cancelUrl: 'your-cancel-url',
  }),

  methods: {
    checkout () {
      this.$refs.checkoutRef.redirectToCheckout();
    }
  }
}
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