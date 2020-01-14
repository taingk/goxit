<template>
  <div>
    <h1 class="title">Create Vote</h1>
    <Formik
      :initial-values="{
        title: '',
        description: ''
      }"
      @onSubmit="handleSubmit"
    >
      <form>
        <Field type="text" name="title" placeholder="Title" tag="input" />
        <Field
          class="textarea"
          type="textarea"
          name="description"
          placeholder="Description"
          tag="textarea"
        />
        <button type="submit">Post</button>
      </form>
    </Formik>
  </div>
</template>

<script>
import axios from '@/utils/axios';
import Formik from '@/components/formik/Formik.vue';
import Field from '@/components/formik/Field.vue';

export default {
  name: 'Create',
  components: {
    Formik,
    Field
  },
  methods: {
    handleSubmit({ event, values }) {
      event.preventDefault();
      axios
        .post('/vote', {
          title: values.title,
          description: values.description
        })
        .then(response => {
          if (response.status === 200) {
            this.$router.push('vote/' + response.data.uuid);
            console.log(response.data.uuid);
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
