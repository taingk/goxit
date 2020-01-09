<template>
  <div></div
></template>
<script>
import store from '@/store';

export default {
  name: 'Logout',
  created() {
    const response = { data: {} };
    response.data.JWT = null;
    response.data.AccessLevel = 0;
    store.commit('authenticate', response.data);
    this.$router.push('/login');
  }
};
</script>

<style></style>
