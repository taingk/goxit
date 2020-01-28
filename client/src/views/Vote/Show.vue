<template>
  <div>
    <h1 class="title">Vote</h1>
    <h3>{{ vote.title }}</h3>
    <p>{{ vote.description }}</p>
    <p v-if="vote.uuid_votes">{{ voteCount }} votes</p>
    <p v-else>0</p>
    <button @click="handleSubmit">Vote</button>
  </div>
</template>

<script>
import axios from '@/utils/axios';

export default {
  name: 'Show',
  components: {},
  data: function() {
    return {
      vote: {},
      voteCount: 0
    };
  },
  created() {
    axios
      .get('vote/' + this.$route.params.uuid)
      .then(response => {
        if (response.status === 200) {
          this.vote = response.data;
          this.voteCount = response.data.uuid_votes.length;
          console.log(response);
        }
      })
      .catch(response => {
        console.log(response);
      });
  },
  methods: {
    handleSubmit() {
      axios
        .put('vote/' + this.$route.params.uuid)
        .then(response => {
          if (response.status === 200) {
            this.$toasted.success('You voted !');
            this.voteCount = this.voteCount + 1;
            console.log(this.voteCount);
          }
          if (response.status === 401) {
            this.$toasted.error('You can only vote once !');
          }
        })
        .catch(response => {
          if (response.status === 401) {
            this.$toasted.error('You can only vote once !');
          }
          console.log(response);
        });
    }
  }
};
</script>

<style>
.centered {
  display: flex;
  justify-content: center;
}
</style>
