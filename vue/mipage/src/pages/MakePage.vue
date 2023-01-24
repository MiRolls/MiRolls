<template>
  <div style="height: 100%;">
    <AppBar></AppBar>
    <MakePageLeftControl @add="addQuest" @submit="submitRoll" @title-change="changeTitle" @save-roll="saveRoll"/>
    <RollsPage ref="rollsPage" :title="rollsTitle"/>
    <PageFooter style="margin-left: 30%"/>
  </div>
</template>
<script>
import AppBar from "@/components/AppBar.vue";
import MakePageLeftControl from "@/components/MakePageLeftControl.vue";
import PageFooter from "@/components/PageFooter.vue";
import RollsPage from "@/components/RollsPage.vue";

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
      this.$refs.rollsPage.getRoll();
    }
  },
  data(){
    return{
      rollsTitle:'这是一张新的问卷',
    }
  },
  components: {RollsPage, PageFooter, MakePageLeftControl, AppBar},

}
</script>