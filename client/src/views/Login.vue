<template>
  <div class="login">
    <Formik
      :initialValues="{
        email: 'Test',
        password: ''
      }"
      @onSubmit="handleSubmit"
    >
      <form>
        <!-- <div v-if="errors.length">
          <b>Please correct the following error(s):</b>
          <ul>
            <li v-for="(error, index) in errors" :key="index">
              {{ error }}
            </li>
          </ul>
        </div> -->
        <Field component="FormInput" type="email" name="email" />
        <Field component="FormInput" name="password" />
        <button type="submit">Se connecter</button>
      </form>
    </Formik>
  </div>
</template>

<script>
import Formik from "@/components/formik/Formik.vue";
import Field from "@/components/formik/Field.vue";

export default {
  name: "login",
  components: {
    Formik,
    Field
  },
  data: function() {
    return {
      errors: []
    };
  },
  methods: {
    handleSubmit: function({ event, values }) {
      event.preventDefault();
      console.log(values);
      this.errors = [];

      if (values.email && values.password) {
        return true;
      }

      if (!values.email) {
        this.errors.push("Email required.");
      }
      if (!values.password) {
        this.errors.push("Password required.");
      }
    }
  }
};
</script>
