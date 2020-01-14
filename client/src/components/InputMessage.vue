<template>
  <div class="input">
    <input type="text" @change="setMessage">
  </div>
</template>

<script>
export default {
  name: 'InputMessage',
  props: {
    setMessage: Function
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  div.input {
    background-color: red;
  }
</style>