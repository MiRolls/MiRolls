<template>
  <div id="savedRolls">
    <div class="line"></div>
    <div v-for="(item,index) in savedRolls" class="savedRoll" :key="item + index">
      <span style="margin-left: 10px;display: block;">{{ index + $t("savedRolls")}}</span>
    </div>
    <div class="bottomLine">{{ $t("savedNone") }}</div>
  </div>
</template>

<script>
export default {
  name: "SavedRolls",
  created(){
    // add hook
    this.savedRolls = JSON.parse(this.$cookies.get("draft"));
  },
  data(){
    return{
      savedRolls:[]
    }
  }
}
</script>

<style>
.savedRoll{
  padding-top:10px;
  padding-bottom: 10px;
  border-bottom:#bebebe 1px solid;
}

.bottomLine{
  margin-top: 20px;
  text-align: center;
  font-size: 17px;
}

#savedRolls{
  margin-top: 40px;
  border-radius: 10px;
  width: 75%;
  padding: 20px 25px 20px 25px;
  min-height: calc(100% - 100px - 40px - 260px);
  box-shadow: 0 0 6px rgba(0, 0, 0, 0.47);
  margin-left: auto;
  margin-right: auto;
}
</style>