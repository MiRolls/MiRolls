<template>
  <div style="height: 100%;">
    <Message v-if="message" :message="message"></Message>
    <AppBar></AppBar>
    <MakePageLeftControl @add="addQuest" @submit="submitRoll" @title-change="changeTitle" @save-roll="saveRoll"/>
    <RollsPage ref="rollsPage" :title="rollsTitle"/>
    <PageFooter style="margin-left: 30%"/>
  </div>
</template>
<!--suppress JSUnresolvedFunction -->
<script>
import AppBar from "../components/AppBar.vue";
import MakePageLeftControl from "../components/MakePageLeftControl.vue";
import PageFooter from "../components/PageFooter.vue";
import RollsPage from "../components/RollsPage.vue";
import Message from "../components/Message.vue";

// noinspection JSUnresolvedVariable
export default {
  name:"MakePage",
  methods: {
    changeTitle(value) {
      this.rollsTitle = value;
    },
    addQuest(questSetting) {
      this.$refs.rollsPage.addQuestValue(questSetting.type, questSetting.optionsNumber)
    },
    saveRoll() {
      this.$refs.rollsPage.saveQuest();
    },
    submitRoll() {
      fetch("/roll/create",{
        method:"post",
        body:this.$refs.rollsPage.getRoll()
      }).then(res=>res.json()).then(date=>{
        if(date.message === "success"){
          // if success
          console.log("success",date)
          // noinspection JSUnresolvedVariable
          this.message = `问卷上传成功！查询码：${date.rollCode}，请务必务必务必务必牢记！访问${date.rollLink}即可开始答题。`
        }else {
          // if error
          // noinspection JSUnresolvedVariable
          console.log("error",date)
          this.message = `服务器出现问题，请截图发送给lm@lmfans.cn Error:${date.error}, ErrorType:${date.errorType}`
        }
      })
    }
  },
  data(){
    return{
      rollsTitle:this.$t("makeTitleNormal"),
      message: "",
    }
  },
  components: {RollsPage, PageFooter, MakePageLeftControl, AppBar,Message},

}
</script>