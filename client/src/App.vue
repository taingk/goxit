<template>
  <div id="app">
    <head>
      <meta name="viewport" content="width=device-width, initial-scale=1" />
    </head>
    <body>
      <nav id="nav" style="position: fixed;">
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
      </nav>
      <router-view style="padding-top: 45px" />
    </body>
    <link rel="stylesheet" href="https://unpkg.com/picnic" />
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
