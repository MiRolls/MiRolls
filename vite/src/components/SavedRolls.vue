<template>
  <div id="savedRolls">
    <div class="line"></div>
    <div v-for="(item,index) in savedRolls" class="savedRoll" :key="item + index" @click="$router.push('/make?draft=' + JSON.stringify(item))">
      <span style="margin-left: 10px;display: block;">{{ index+1 + ". " + item.title}}</span>
    </div>
    <div class="bottomLine" v-if="savedRolls.length === 0">{{ $t("savedNone") }}</div>
  </div>
</template>

<script>
export default {
  name: "SavedRolls",
  created(){
    // add hook
    if(this.$cookies.isKey("draft")){
      this.savedRolls = this.$cookies.get("draft").data;
    }
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
  cursor: pointer;
}

.bottomLine{
  margin-top: 20px;
  text-align: center;
  font-size: 17px;
  color: #a1a1a1;
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