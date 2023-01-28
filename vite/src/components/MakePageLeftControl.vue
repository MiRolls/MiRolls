<template>
  <div id="MakePageLeftControl">
    <div class="addControl" style="margin-top: 0">
      <span class="controlTitle">{{ $t("makeRadio") }}</span><br>
      <span>{{$t("makeOptionsNumber")}}</span>
      <input type="number" id="chooseNumber" class="controlInput" placeholder="3" v-model="radioNumber">
      <br>
      <!--suppress JSUnresolvedVariable -->
      <button class="add" @click="$emit('add',{type:'choice',optionsNumber: radioNumber})" :style="{backgroundColor:style}">
        {{ $t("makeLeftControlButton") }}</button>
    </div>
    <div class="addControl">
      <span class="controlTitle">{{$t("makeChoice")}}</span><br>
      <span>{{$t("makeOptionsNumber")}}</span>
      <input type="number" id="chooseNumber" class="controlInput" placeholder="3" v-model="choiceNumber">
      <br>
      <!--suppress JSUnresolvedVariable -->
      <button class="add" @click="$emit('add',{type:'choice',optionsNumber: choiceNumber})" :style="{backgroundColor:style}">{{ $t("makeLeftControlButton") }}</button>
    </div>
    <div class="addControl">
      <span class="controlTitle">{{$t("makeBlank")}}</span><br>
      <button class="add" @click="$emit('add',{type:'blank'})" :style="{backgroundColor:style}">{{ $t("makeLeftControlButton") }}</button>
    </div>
    <div class="addControl">
      <span class="controlTitle">{{$t("makeManyBlank")}}</span><br />
      <button class="add" @click="$emit('add',{type:'manyBlank'})" :style="{backgroundColor:style}">{{ $t("makeLeftControlButton") }}</button>
    </div>
    <div class="bottomControl">
      <span class="controlTitle">{{$t("makeRollTitle")}}</span>
      <input type="text" id="RollsTitle" :value="$t('makeTitleNormal')" @input="$emit('title-change',$event.target.value)">
      <button id="submitRolls" @click="submitRoll" :style="{backgroundColor:style}">{{$t("makeSubmit")}}</button>
      <button id="submitRolls" @click="save()" :style="{backgroundColor:style}">{{$t("makeSave")}}</button>
    </div>
  </div>
</template>
<!--suppress JSUnresolvedVariable -->
<script>
export default {
  name: "MakePageLeftControl",
  data() {
    return {
      radioNumber: 3,
      choiceNumber: 3,
      style:"rgb(21, 127, 248)"
    }
  },
  created() {
    fetch("/get/site",{method:"post"}).then(res=>res.json()).then(data=>{
      this.style = data.mainColor;
    })
  }
  ,
  methods: {
    submitRoll() {
      console.log("submit")
      this.$emit('submit')
    },
    save() {
      this.$emit('save-roll');
    }
  }
}
</script>
<style>
#submitRolls{
  margin-top: 10px;
  font-size: 16px;
  border: none;
  background-color: rgb(21, 127, 248);
  color: white;
  border-radius: 3px;
  height: 32px;
  width: 90%;
  transition: 0.3s;
}

#submitRolls:hover{
  box-shadow: 0 0 6px rgba(0, 0, 0, 0.55);
}

#RollsTitle{
  width: 90%;
  border: none;
  border-bottom: 1px solid black;
  background: none;
  height: 20px;
  font-size: 16px;
  outline: none;
  transition: 0.5s;
}

#RollsTitle:focus{
  border-bottom: 2px solid rgb(59,125,240);
}

.bottomControl {
  position: absolute;
  bottom: 0;
  width: 100%;
  height: 150px;
  text-align: center;
}

#MakePageLeftControl{
  position: fixed;
  height: calc(100% - 73px);
  margin-top: -23px;
  width: 30%;
  background-color: #f5f5f5;
}

.controlInput{
  border: none;
  outline: none;
  border-bottom: solid 1px black;
  width: 35px;
  font-size: 16px;
}

.addControl{
  width: calc(100% - 1px);
  background-color: white;
  padding: 18px 0 18px 0;
  text-align: center;
  margin-top: 2px;
}
.controlTitle{
  font-weight: bold;
}
.add{
  font-size: 16px;
  border: none;
  background-color: rgb(21, 127, 248);
  border-radius: 5px;
  width: 60%;
  color: white;
  margin-top: 3px;
  padding: 3px 0 3px 0;
}
.add:hover{
  box-shadow: 0 0 5px rgba(0, 0, 0, 0.36);
}
</style>