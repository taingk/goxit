<template>
  <Formik
    :initial-values="{
      email: '',
      password: ''
    }"
    @onSubmit="handleSubmit"
  >
    <form>
      <Field type="email" name="email" placeholder="Email" />
      <Field type="password" name="password" placeholder="Password" />
      <button type="submit">Se connecter</button>
    </form>
  </Formik>
</template>

<script>
import Formik from '@/components/formik/Formik.vue';
import Field from '@/components/formik/Field.vue';
import axios from '@/utils/axios';

export default {
  name: 'Login',
  components: {
    Formik,
    Field
  },
  methods: {
    handleSubmit({ event, values }) {
      event.preventDefault();
      console.log(values);
      axios
        .post('/login', {
          email: values.email,
          password: values.password
        })
        .then(response => {
          console.log(response);
        });
    }
  }
};
</script>
