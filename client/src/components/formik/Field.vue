<template>
  <component :is="tag" v-bind="$props" :value="value" @input="handleInput" />
</template>

<script>
export default {
  name: 'Field',
  inject: ['setValue', 'getValue'],
  props: {
    name: String,
    type: String,
    id: String,
    required: Boolean,
    checked: Boolean,
    tag: String
  },
  data: () => ({
    value: ''
  }),
  created() {
    this.value = this.getValue(this.name);
    this.setValue(this.name, this.value);
    if (this.type === 'checkbox' && this.checked) {
      this.setValue(this.name, this.checked);
    }
  },
  methods: {
    handleInput(event) {
      const { type, checked, value } = event.target;

      if (type === 'checkbox') {
        this.setValue(this.name, checked);
      } else {
        this.setValue(this.name, value);
      }
    }
  }
};
</script>
