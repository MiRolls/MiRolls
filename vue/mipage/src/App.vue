<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
  export default {
    name: 'App',
  }
</script>

<style>
  #app {
    width: 100%;
    -webkit-font-smoothing: antialiased;
    -moz-font-smoothing: grayscale;
    height: 100%;
  }
  button{
    cursor: pointer;
  }
</style>