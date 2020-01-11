<template>
  <div id="app">
    <div id="nav">
      <div v-if="isAuthenticated()">
        <span v-if="isAdmin()">
          <router-link to="/vote">Create vote</router-link> |
        </span>
        <router-link to="/votes">Vote List</router-link> |
        <router-link to="/user">Profile</router-link> |
        <router-link to="/logout">Logout</router-link>
      </div>
      <div v-else>
        <router-link to="/login">Login</router-link> |
        <router-link to="/register">Register</router-link>
      </div>
    </div>
    <router-view />
  </div>
</template>
<script>
import store from '@/store';

export default {
  methods: {
    isAuthenticated() {
      return store.state.token;
    },
    isAdmin() {
      return store.state.accessLevel;
    }
  }
};
</script>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
