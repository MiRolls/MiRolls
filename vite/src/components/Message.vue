<template>
  <div class="messageBackground"></div>
  <div class="message">
    <h3>提示！</h3>
    {{ message }}
  </div>
</template>

<script>
export default {
  name: "Message",
  props:{
    message:String
  }
}
</script>

<style scoped>
.messageBackground{
  position: fixed;
  z-index: 1145141;
  background-color: rgba(0, 0, 0, 0.18);
  width: 100%;
  height: 100%;
}
.message{
  position: fixed;
  width: 50%;
  border-radius: 10px;
  background-color: white;
  padding: 30px 30px 30px 30px;
}
</style>