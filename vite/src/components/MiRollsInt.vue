<template>
  <div class="IntOfPage">
    <h1 class="IntOfTitle">{{name}}</h1>
    <img :src="logo" :alt="name" class="IntOfImage">
  </div>
</template>
<script>
export default {
  name:"MiRollsInt",
  setup(){
    fetch("/get/site",{
      method:"post",
    }).then(res=>res.json()).then(data=>{
      data.name
      // noinspection JSUnresolvedVariable
      this.logo = data.logo;
      this.name = data.name;
      //
    })
  },
  data(){
    return{
      logo:"https://img.lmfans.cn/i/2023/01/26/uqwc3n.png",
      name:"米卷"
    }
  }
}
</script>
<style>
.IntOfPage{
  text-align: center;
}
.IntOfTitle{
  margin-top: 3em;
}
.IntOfImage{
  max-width: 60%;
  height: 300px;
}
</style>