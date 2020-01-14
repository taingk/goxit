<template>
  <div>
    <h1>Login</h1>
    <Formik
      :initial-values="{
        email: '',
        password: ''
      }"
      @onSubmit="handleSubmit"
    >
      <form>
        <Field
          class="input"
          type="email"
          name="email"
          placeholder="Email"
          tag="input"
        />
        <Field
          class="input"
          type="password"
          name="password"
          placeholder="Password"
          tag="input"
        />
        <button class="button" type="submit">Se connecter</button>
      </form>
    </Formik>
  </div>
</template>

<script>
import Formik from '@/components/formik/Formik.vue';
import Field from '@/components/formik/Field.vue';
import axios from '@/utils/axios';
import store from '@/store';

export default {
  name: 'Login',
  components: {
    Formik,
    Field
  },
  methods: {
    handleSubmit({ event, values }) {
      event.preventDefault();
      axios
        .post('/login', {
          email: values.email,
          password: values.password
        })
        .then(response => {
          if (response.status === 200) {
            store.commit('authenticate', response.data);
            this.$router.push('/votes');
            console.log('Logged !');
          }
        })
        .catch(response => {
          console.log(response);
        });
    }
  }
};
</script>
