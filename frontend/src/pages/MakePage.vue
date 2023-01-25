<template>
  <div style="height: 100%;">
    <AppBar></AppBar>
    <MakePageLeftControl @add="addQuest" @submit="submitRoll" @title-change="changeTitle" @save-roll="saveRoll" />
    <RollsPage ref="rollsPage" :title="rollsTitle" />
    <PageFooter style="margin-left: 30%" />
  </div>
</template>
<script lang="ts">
  import AppBar from "../components/AppBar.vue";
  import MakePageLeftControl from "../components/MakePageLeftControl.vue";
  import PageFooter from "../components/PageFooter.vue";
  import RollsPage from "../components/RollsPage.vue";
  import {ref} from "vue";
  export default {
    name: "MakePage",
    components: { RollsPage, PageFooter, MakePageLeftControl, AppBar },
    setup(){
      interface questSetting{
        type:string,
        optionsNumber?:number,
      }
      interface response{
        message: "success"|"error",
        error?:string,
        errorType?:string
        rollCode?:string
      }

      const rollsPage = ref()
      let rollsTitle = ref("这是一张新的问卷");
      let message = ref("message")

      function addQuest(questSetting:questSetting){
        rollsPage.value.addQuestValue(questSetting.type, questSetting.optionsNumber)
      }
      function changeTitle(value: string) {
        rollsTitle.value = value;
      }
      function saveRoll() {
        rollsPage.value.saveQuest();
      }
      function submitRoll() {
        let roll = rollsPage.value.getRoll();
        fetch("/roll/create", {
          method: "POST",
          body: roll
        }).then(res => res.json()).then((date:response) => {
          if(date.message === "success"){
            // if success
            message.value = `已经提交！这是你的问卷查询码 ${date.rollCode} ，请务必记住！务必记住！`
          }else{
            // if error
            message.value = `提交出现错误，可能服务器出现问题，请联系lm@lmfans.cn(Error:${date.error},ErrorType:${date.errorType})`
          }
        })
      }

      return {
        addQuest,
        changeTitle,
        saveRoll,
        submitRoll,
        rollsTitle,
        message
      }
    },
  }
</script>