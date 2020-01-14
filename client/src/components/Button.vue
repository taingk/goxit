<template>
  <button>{{ text }}</button>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    text: {
      type: String,
      required: true
    }
  }
};
</script>

<style></style>
