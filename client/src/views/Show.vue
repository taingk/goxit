<template>
  <div>
    <h1 class="title">Vote</h1>
    <h3>{{ response.data.title }}</h3>
    <p>{{ response.data.description }}</p>
  </div>
</template>

<script>
import axios from '@/utils/axios';

export default {
  name: 'Show',
  components: {},
  methods: {
    handleSubmit({ event }) {
      event.preventDefault();
      axios
        .get('vote/' + this.$route.params.uuid)
        .then(response => {
          if (response.status === 200) {
            console.log(response);
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
