<template>
  <div id="app">
    <router-view :prop></router-view>
  </div>
</template>

<script>
export default {
  name: 'App',
  setup(){
    // be created
    fetch("/get/site",{
      method:"post",
    }).then(res=>res.json()).then(data=>{
      // data.name
      this.site = data;
    })
  },
  data(){
    return{
      site:{}
    }
  }
}
</script>

<style>
#app {
  width: 100%;
  height: 100%;
}
button{
  cursor: pointer;
}
</style>