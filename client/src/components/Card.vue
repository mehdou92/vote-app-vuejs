<template>
  <div class="card" draggable="true" @dragstart="onDrag">
    <h1 class="card-name">{{ name }}</h1>
    <p v-if="status == 'done'">Done</p>
    <p v-show="status == 'done'">Done</p>
  </div>
</template>

<script>
export default {
  name: 'Card',
  props: {
    name: String,
    onDrag: Function,
    status: {
      type: String,
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  div.card {
    width: 100%;
    background-color: violet; 
    margin: 0 5px;
  }
  div.card:not(:first-child) {
    margin-top: 10px;
  }
  div.card-name {
    background-color: white;
  }
</style>