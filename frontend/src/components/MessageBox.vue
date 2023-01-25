<template>
  <div class="MessageBody" />
  <div class="Message">{{message}}</div>
</template>
<script lang="ts">
import {watch} from "vue";

export default {
  name:"MessageBox",
  props:{
    message:String
  },
}
</script>
<style>
.MessageBody{
  background-color: rgba(0, 0, 0, 0.24);
  width: 100%;
  height: 100%;
}
</style>