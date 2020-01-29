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
      <form class="centered">
        <div class="half">
          <Field
            type="text"
            name="firstname"
            placeholder="Firstname"
            class="stack"
            tag="input"
          />
          <Field
            type="text"
            name="lastname"
            placeholder="Lastname"
            tag="input"
            class="stack"
          />
          <Field
            type="email"
            name="email"
            placeholder="Email"
            tag="input"
            class="stack"
          />
          <Field
            type="password"
            name="password"
            placeholder="Password"
            tag="input"
            class="stack"
          />
          <button type="submit" class="button stack">Register</button>
        </div>
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
            this.$toasted.success('You are registered');
            this.$router.push('/');
          }
        })
        .catch(({ response }) => {
          if (response.data.Code === '23505') {
            this.$toasted.error('Email already exists');
          }
          const { firstname, lastname, password } = values;

          if (firstname.length < 2) {
            this.$toasted.error(
              'Firstname must be 2 characters length minimum'
            );
          }
          if (lastname.length < 2) {
            this.$toasted.error('Lastname must be 2 characters length minimum');
          }
          if (password.length < 6) {
            this.$toasted.error('Password must be 6 characters length minimum');
          }
        });
    }
  }
};
</script>
<style lang="css">
.centered {
  display: flex;
  justify-content: center;
}
</style>
