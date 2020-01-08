<template>
  <component :is="input" v-bind="$props" />
</template>

<script>
import Input from '@/components/formik/Input.vue';

export default {
  name: 'Field',
  components: {
    Input
  },
  props: {
    name: String,
    type: String,
    id: String,
    required: Boolean,
    checked: Boolean
  },
  data() {
    return {
      input: 'Input'
    };
  }
};
</script>
