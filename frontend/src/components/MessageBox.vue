<template>
  <div class="MessageBody" />
  <div class="Message">{{}}</div>
</template>
<script lang="ts">
export default {
  name:"MessageBox",
  setup(){
    defineProps<{
      message:
    }>()
  }
}

// const props = defineProps(['foo'])
</script>
<style>
.MessageBody{
  background-color: rgba(0, 0, 0, 0.24);
  width: 100%;
  height: 100%;
}
</style>