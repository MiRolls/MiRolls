<template>
    <div id="homeMain" :style="style">
        <div id="mainText">
            <h1 class="titleOfMain">不用</h1>
            <h1 class="titleOfMain">登录账户 | 花费费用</h1>
            <span id="spanTitle">就可以享受优秀的免费问卷系统</span><br />
            <button id="btnMakePage" @click="goMake">创建问卷</button>
        </div>
    </div>
</template>
<script>
    export default {
      name: "MainPage",
      data(){
        return {
          style:{
            backgroundColor: "rgb(21, 127, 248)"
          }
        }
      },
      setup(){
        fetch("/get/site",{
          method:"post",
        }).then(res=>res.json()).then(data=>{
          data.name
          // noinspection JSUnresolvedVariable
          this.style.backgroundColor = data.mainColor;
        })
      },
      methods:{
        goMake(){
          window.location.href = "/#/make"
        }
      }
    }
</script>
<style>
    #homeMain {
        height: 80%;
        margin: 0;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }

    .titleOfMain {
        font-size: 3em;
        line-height: 1em;
    }

    #mainText {
        text-align: center;
        color: white;
    }

    #btnMakePage {
        margin-top: 1em;
        display: block;
        height: 35px;
        width: 100%;
        font-size: 16px;
        border: none;
        background-color: rgb(247, 227, 52);
        border-radius: 5px;
        cursor: pointer;
        transition: 0.1s;
    }

    #btnMakePage:hover {
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.497);
    }
</style>