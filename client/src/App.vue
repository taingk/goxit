<template>
  <div id="app">
    <head>
      <meta name="viewport" content="width=device-width, initial-scale=1" />
    </head>
    <body>
      <nav style="position: fixed;">
        <router-link to="/votes" class="brand">
          <img class="logo" src="./assets/logo2.png" />
          <span>GOXIT</span></router-link
        >
        <div class="menu">
          <div v-if="isAuthenticated()">
            <span v-if="isAdmin()">
              <router-link class="button" to="/vote">Create vote</router-link>
            </span>
            <router-link class=" pseudo button" to="/votes"
              >Vote List</router-link
            >
            <router-link class=" pseudo button" to="/user">Profile</router-link>
            <router-link class="pseudo button" to="/logout">Logout</router-link>
          </div>
          <div v-else>
            <router-link class=" pseudo button" to="/login">Login</router-link>
            <router-link class="  button" to="/register">Register</router-link>
          </div>
        </div>
      </nav>
      <div class="flex demo wrapper">
        <router-view style="padding-top: 45px" />
      </div>
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

  .wrapper {
    display: flex;
    align-items: center;
    justify-content: center;
  }
}
</style>
