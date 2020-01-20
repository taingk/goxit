<template>
  <div>
    <h1>Profile</h1>
    <div v-if="!edit" class="centered">
      <article class="card half">
        <p>{{ user.firstname }} {{ user.lastname }}</p>
        <p>{{ user.email }}</p>
      </article>
    </div>
    <div v-else>
      <Formik
        :initial-values="{
          firstname: user.firstname,
          lastname: user.lastname,
          email: user.email,
          password: ''
        }"
        @onSubmit="handleSubmit"
      >
        <div class="centered">
          <form>
            <p>
              <Field
                tag="input"
                type="text"
                name="firstname"
                placeholder="Firstname"
              />
            </p>
            <p>
              <Field
                tag="input"
                type="text"
                name="lastname"
                placeholder="Lastname"
              />
            </p>
            <p>
              <Field
                tag="input"
                type="email"
                name="email"
                placeholder="Email"
              />
            </p>
            <p>
              <Field
                tag="input"
                type="password"
                name="password"
                placeholder="Password"
              />
            </p>
            <button type="submit">Update</button>
          </form>
        </div>
      </Formik>
    </div>
    <button @click="isEditable()">{{ edit ? 'Back' : 'Edit' }}</button>
  </div>
</template>

<script>
import Formik from '@/components/formik/Formik.vue';
import Field from '@/components/formik/Field.vue';
import axios from '@/utils/axios';
import store from '@/store';

export default {
  name: 'Profile',
  components: {
    Formik,
    Field
  },
  data() {
    return {
      user: {},
      edit: false
    };
  },
  created() {
    axios
      .get(`/user/${store.state.token}`)
      .then(response => {
        if (response.status === 200) {
          this.user = response.data;
        }
      })
      .catch(response => {
        console.log(response);
      });
  },
  methods: {
    isEditable() {
      this.edit = !this.edit;
    },
    handleSubmit({ event, values }) {
      event.preventDefault();
      let body = {};
      for (const value in values) {
        if (values[value]) body[value] = values[value];
      }
      axios
        .put(`/user/${store.state.token}`, body)
        .then(response => {
          if (response.status === 200) {
            console.log('Updated', response.data);
            this.user = response.data;
            this.edit = !this.edit;
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
