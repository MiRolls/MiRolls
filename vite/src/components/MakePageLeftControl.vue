<template>
  <div id="MakePageLeftControl">
    <div class="addControl" style="margin-top: 0">
      <span class="controlTitle">添加单选题</span><br>
      <span>选项数量</span>
      <input type="number" id="chooseNumber" class="controlInput" placeholder="3">
      <br>
      <button class="add" @click="$emit('add',{type:'choice',optionsNumber: radioNumber})">添加</button>
    </div>
    <div class="addControl">
      <span class="controlTitle">添加多选题</span><br>
      <span>选项数量</span>
      <input type="number" id="chooseNumber" class="controlInput" placeholder="3">
      <br>
      <button class="add" @click="$emit('add',{type:'choice',optionsNumber: choiceNumber})">添加</button>
    </div>
    <div class="addControl">
      <span class="controlTitle">添加填空题（单行）</span><br>
      <button class="add" @click="$emit('add',{type:'blank'})">添加</button>
    </div>
    <div class="addControl">
      <span class="controlTitle">添加填空题（多行）</span><br />
      <button class="add" @click="$emit('add',{type:'manyBlank'})">添加</button>
    </div>
    <div class="bottomControl">
      <span class="controlTitle">问卷标题：</span>
      <input type="text" id="RollsTitle" v-model="rollsTitle" @input="$emit('title-change',rollsTitle)">
      <button id="submitRolls" @click="submitRoll">提交问卷</button>
      <button id="submitRolls" @click="save()">保存为草稿</button>
    </div>
  </div>
</template>
<script>
export default {
  name: "MakePageLeftControl",
  data() {
    return {
      rollsTitle: '这是一张新的问卷',
      radioNumber: 3,
      choiceNumber: 3
    }
  },
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
  width: 25px;
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