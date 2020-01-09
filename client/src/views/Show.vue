<template>
  <div>
    <h1 class="title">Vote</h1>
    <h3>{{ vote.title }}</h3>
    <p>{{ vote.description }}</p>
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
      vote: {}
    };
  },
  created() {
    axios
      .get('vote/' + this.$route.params.uuid)
      .then(response => {
        if (response.status === 200) {
          this.vote = response.data;
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
            console.log('voted');
            // this.$router.push('/vote');
          }
        })
        .catch(response => {
          console.log(response);
        });
    }
  }
};
</script>

<style></style>
