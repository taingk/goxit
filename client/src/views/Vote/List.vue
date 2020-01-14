<template>
  <div>
    <h1 class="title">Vote List</h1>
    <div class="centered">
      <table class="primary">
        <thead>
          <tr>
            <th>Title</th>
            <th>Description</th>
          </tr>
        </thead>
        <tbody id="v-for-object">
          <tr v-for="vote in votes" :key="vote.uuid">
            <td>
              <router-link
                :to="{ name: 'show-vote', params: { uuid: vote.uuid } }"
                >{{ vote.title }}
              </router-link>
            </td>
            <td>{{ vote.description }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import axios from '@/utils/axios';

export default {
  name: 'List',
  components: {},
  data: function() {
    return {
      votes: []
    };
  },
  created() {
    axios
      .get('vote')
      .then(response => {
        if (response.status === 200) {
          this.votes = response.data;
        }
      })
      .catch(response => {
        console.log(response);
      });
  }
};
</script>

<style>
.centered {
  display: flex;
  justify-content: center;
}
</style>
