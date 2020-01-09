<template>
  <div>
    <h1 class="title">Vote List</h1>
    <ul id="v-for-object" class="demo">
      <li :v-for="vote in votes">
        <h3>{{ vote.title }}</h3>
        <p>{{ vote.description }}</p>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from '@/utils/axios';

export default {
  name: 'List',
  components: {},
  data: function() {
    return {
      vote: {},
      votes: {}
    };
  },
  created() {
    axios
      .get('vote')
      .then(response => {
        if (response.status === 200) {
          this.votes = response.data;
          console.log(response.data);
        }
      })
      .catch(response => {
        console.log(response);
      });
  }
};
</script>

<style></style>
