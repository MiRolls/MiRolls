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
          this.message = $t('messageUpdateSuccess') + `${date.rollCode}, ${$t("messageUpdateSuccess2")} ${date.rollLink} ${$t("messageUpdateSuccess3")}`
        }else {
          // if error
          // noinspection JSUnresolvedVariable
          this.message = $t('messageDatabaseError') + ` Error:${date.error}, ErrorType:${date.errorType}`
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