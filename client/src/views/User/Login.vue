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
      <form class="centered">
        <div class="half">
          <Field
            class="input stack"
            type="email"
            name="email"
            placeholder="Email"
            tag="input"
          />
          <Field
            class="input stack"
            type="password"
            name="password"
            placeholder="Password"
            tag="input"
          />
          <button class="button stack" type="submit">Se connecter</button>
        </div>
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
          }
        })
        .catch(response => {
          console.log(response);
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
