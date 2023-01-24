<template>
  <div class="MessageBody" />
  <div class="Message">{{}}</div>
</template>
<script lang="ts" setup>

const props = defineProps(['foo'])
</script>
<style>
.MessageBody{
  background-color: rgba(0, 0, 0, 0.24);
  width: 100%;
  height: 100%;
}
</style>