<template>
  <div id="app">
    <router-link to="/">Board </router-link>
    <router-link to="/list/1">List 1 </router-link>
    <router-link :to="{name: 'HomeList', params: {listid: 2}}">List 2 </router-link>
    <router-view name="header"></router-view>
    <router-view></router-view>
  </div>
</template>

<script>
import Board from './components/Board.vue'
export default {
  name: 'app',
  components: {
    Board
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