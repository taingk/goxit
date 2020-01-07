<template>
  <div>
    <label v-if="label">{{ label }}</label>
    <component v-bind:is="component" v-bind="$props"></component>
  </div>
</template>

<script>
export default {
  name: "Field",
  props: {
    name: String,
    label: String,
    type: String,
    id: String,
    required: Boolean,
    checked: Boolean,
    component: String
  }
};
</script>
