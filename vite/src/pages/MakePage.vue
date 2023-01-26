<template>
  <div style="height: 100%;">
    <AppBar></AppBar>
    <MakePageLeftControl @add="addQuest" @submit="submitRoll" @title-change="changeTitle" @save-roll="saveRoll"/>
    <RollsPage ref="rollsPage" :title="rollsTitle"/>
    <PageFooter style="margin-left: 30%"/>
    <Message v-if="message"></Message>
  </div>
</template>
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
          // noinspection JSUnresolvedVariable
          this.message = `问卷上传成功！查询码：${date.rollCode}，请务必务必务必务必牢记！访问${date.rollLink}即可开始答题。`
        }else {
          // if error
        }
      })
    }
  },
  data(){
    return{
      rollsTitle:'这是一张新的问卷',
      message: String,
    }
  },
  components: {RollsPage, PageFooter, MakePageLeftControl, AppBar,Message},

}
</script>