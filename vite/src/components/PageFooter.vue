<template>
  <div class="footer">
    <h1>{{name}}</h1>
    <p class="footerButton">一个免费且好用的国内问卷软件</p>
    <p class="footerButton" style="cursor: pointer;" @click="goIcp">{{ icp }}</p>
  </div>
</template>
<script>
  export default{
    name:"PageFooter",
    methods:{
      goIcp(){
        window.location.href="https://beian.miit.gov.cn/";
      }
    },
    data(){
      return {
        name:"米卷",
        icp: "鲁ICP备2022023454号-25"
      }
    },
    created(){
      fetch("/get/site",{
        method:"post",
      }).then(res=>res.json()).then(data=>{
        this.icp = data.icp;
        this.name = data.name;
        // noinspection JSUnresolvedVariable
      })
    },
  }
</script>
<style>
.footer{
  padding: 50px 0 50px 0;
  margin-top: 50px;
  background-color: rgb(23,27,33);
  text-align: center;
  color: white;
}
.footerButton{
  font-size: 17px;
}
</style>