<template>
  <div class="register">
    <h1>Register</h1>
    <Formik
      :initial-values="{
        firstname: '',
        lastname: '',
        email: '',
        password: ''
      }"
      @onSubmit="handleSubmit"
    >
      <form>
        <Field type="text" name="firstname" placeholder="Firstname" />
        <Field type="text" name="lastname" placeholder="Lastname" />
        <Field type="email" name="email" placeholder="Email" />
        <Field type="password" name="password" placeholder="Password" />
        <button type="submit">Register</button>
      </form>
    </Formik>
  </div>
</template>

<script>
import Formik from '@/components/formik/Formik.vue';
import Field from '@/components/formik/Field.vue';
import axios from '@/utils/axios';

export default {
  name: 'Register',
  components: {
    Formik,
    Field
  },
  methods: {
    handleSubmit({ event, values }) {
      event.preventDefault();
      axios
        .post('/user', {
          firstname: values.firstname,
          lastname: values.lastname,
          email: values.email,
          password: values.password
        })
        .then(response => {
          if (response.status === 200) {
            this.$router.push('/');
            console.log('Registered !');
          }
        })
        .catch(response => {
          console.log(response);
        });
    }
  }
};
</script>
