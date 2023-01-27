<template>
    <div id="bar" :style="style">
        <span id="title">{{ site.name }}</span>
        <span id="secTitle">{{ site.link }}</span>
        <titleBtn :innerHtml="$t('appBarHome')" @click="goHome" />
        <titleBtn :innerHtml="$t('appBarSearch')" @click="resSearch" />
        <titleBtn :innerHtml="$t('appBarMake')" @click="makeQtn" />
        <!-- <LoginOrRegister loginStates="true"></LoginOrRegister> -->
        <LoginOrRegister loginStates :userHeadImg="site.logo"></LoginOrRegister>
    </div>
</template>
<script>
  import titleBtn from "./titleBtn.vue";
  import LoginOrRegister from "./LoginOrRegister.vue"
  export default {
    name: 'AppBar',
      methods: {
      goHome() {
        window.location.href = "/#/";
      },
      resSearch() {
        window.location.href = "/#/search";
      },
      makeQtn() {
        window.location.href = "/#/make";
      }
    },
    components: {titleBtn, LoginOrRegister},
    data(){
      return{
        site:{
          name:"米卷",
          link:"wj.lmfans.cn",
          logo:"https://img.lmfans.cn/i/2023/01/26/uqwc3n.png"
        },
        style:{
          backgroundColor: "rgb(21, 127, 248)"
        }
      }
    },
    created(){
      fetch("/get/site",{
        method:"post",
      }).then(res=>res.json()).then(data=>{
        console.log(data)
        this.site = data;
        // noinspection JSUnresolvedVariable
        this.style.backgroundColor = data.mainColor;
      })
    },
  }
</script>
<style>
    #bar {
        width: 100%;
        height: 33px;
        padding-top: 20px;
        padding-bottom: 20px;
        box-shadow: 0 0 12px rgba(0, 0, 0, 0.581);
        color: white;
        position: fixed;
        top: 0;
        margin-top: 0;
        z-index: 999;
    }

    #title {
        margin-left: 100px;
        margin-right: 5%;
        line-height: 28px;
        font-size: 28px;
        font-weight: 400;
    }

    #secTitle {
        position: absolute;
        top: 48px;
        left: 100px;
        font-size: 5px;
        margin-left: 3px;
    }
</style>