<template>
  <div @submit="handleSubmit">
    <slot />
  </div>
</template>

<script>
export default {
  name: 'Formik',
  props: {
    initialValues: Object
  },
  data() {
    return {
      defaultValues: this.initialValues,
      values: {}
    };
  },
  methods: {
    setValue(key, value) {
      this.values[key] = value;
    },
    getValue(key) {
      return this.defaultValues[key];
    },
    handleSubmit(event) {
      this.$emit('onSubmit', {
        values: this.values,
        event
      });
    }
  },
  provide() {
    return {
      setValue: this.setValue,
      getValue: this.getValue
    };
  }
};
</script>
