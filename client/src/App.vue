<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
import Home from './components/Home.vue'
export default {
  name: 'app',
  components: {
    Home
  },
}
</script>

<style>
#app {
  display: flex;
  flex-direction: column;
  height: 100%;
  width: 100%;
  min-width: 100%;
  max-width: 100%;
  min-width: 100%;
  max-width: 100%;
}
</style>