<template>
  <div class="">
    <h1 class="title">Create Vote</h1>
    <Formik
      :initial-values="{
        title: '',
        description: ''
      }"
      @onSubmit="handleSubmit"
    >
      <form class="centered">
        <div class="half">
          <Field
            type="text"
            name="title"
            placeholder="Title"
            tag="input"
            class="stack"
          />
          <Field
            class="textarea stack"
            type="textarea"
            name="description"
            placeholder="Description"
            tag="textarea"
          />
          <button class="button stack" type="submit">Post</button>
        </div>
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
            this.$toasted.success(
              `You have created the vote : ${values.title}`
            );
            this.$router.push('vote/' + response.data.uuid);
          }
        })
        .catch(({ response }) => {
          if (response.data.Code === '23505') {
            this.$toasted.error('Title already exists');
          }
          const { title, description } = values;

          if (title.length < 2) {
            this.$toasted.error('Title must be 2 characters length minimum');
          }
          if (description.length < 2) {
            this.$toasted.error(
              'Description must be 2 characters length minimum'
            );
          }
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
