<template>
  <div class="messageBackground"></div>
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
  position: absolute;
  z-index: 1145141;
  background-color: rgba(0, 0, 0, 0.18);
  width: 100%;
  height: 100%;
}
</style>