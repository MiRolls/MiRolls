<template>
  <div id="savedRolls">
    <div class="line"></div>
    <div v-for="(item,index) in savedRolls" class="savedRoll" :key="item + index">
      <span style="margin-left: 10px;display: block;">{{ index + $t("savedRolls")}}</span>
    </div>
  </div>
</template>

<script>
export default {
  name: "SavedRolls",
  created(){
    // add hook

  },
  data(){
    return{
      savedRolls:[

      ]
    }
  }
}
</script>

<style>
.savedRoll{
  /*background: black;*/
  padding-top:10px;
  padding-bottom: 10px;
  border-bottom:#bebebe 1px solid;
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