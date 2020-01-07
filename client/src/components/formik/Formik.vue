<template>
  <div>
    <slot></slot>
  </div>
</template>

<script>
export default {
  name: "formik",
  props: {
    initialValues: Object
  },
  data: function() {
    return {
      values: {}
    };
  },
  methods: {
    setFormValue: function(key, value) {
      this.form.values[key] = value;
    },
    getFormInitialValue: function(key) {
      return this.initialValues[key];
    },
    handleSubmit: function(event) {
      this.$emit("submit", {
        event,
        values: this.values
      });
    }
  },
  provide: function() {
    return {
      setFormValue: this.setFormValue,
      getFormInitialValue: this.getFormInitialValue
    };
  }
};
</script>
