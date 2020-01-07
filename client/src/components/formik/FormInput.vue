<template>
  <div>
    <input v-bind="$props" :value="value" @input="handleInput" />
  </div>
</template>

<script>
export default {
  name: "FormInput",
  inject: ["setFormValue", "getFormInitialValue"],
  data() {
    return {
      value: ""
    };
  },
  props: {
    name: String,
    type: String,
    id: String,
    required: Boolean,
    disabled: Boolean,
    checked: Boolean
  },
  created() {
    this.value = this.getFormInitialValue(this.id);
    this.setFormValue(this.id, this.value);
  },
  methods: {
    handleInput(e) {
      if (e.target.type === "checkbox") {
        this.setFormValue(this.id, e.target.checked);
      } else {
        this.setFormValue(this.id, e.target.value);
      }
    }
  }
};
</script>
