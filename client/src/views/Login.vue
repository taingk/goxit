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
    }
  }
};
</script>
