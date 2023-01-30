<template>
  <div class="messageBackground" @click="close"></div>
  <div class="message">
    <h3>
      <svg role="img" xmlns="http://www.w3.org/2000/svg" width="24px" height="24px" viewBox="0 0 24 24" aria-labelledby="infoIconTitle" stroke="rgb(72,126,244)" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" fill="none" color="#2329D6"> <title id="infoIconTitle">Information</title> <path d="M12,12 L12,15"/> <line x1="12" y1="9" x2="12" y2="9"/> <circle cx="12" cy="12" r="10"/> </svg>
    </h3>
    {{ message }}
  </div>
</template>

<!--suppress JSUnresolvedFunction -->
<script>
import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";

export default {
  name: "Message",
  components: {FontAwesomeIcon},
  props:{
    message:String
  },
  methods:{
    close(){
      this.$router.push('/');
    }
  }
}
</script>

<style scoped>
.messageBackground{
  position: fixed;
  z-index: 1145141;
  background-color: rgba(0, 0, 0, 0.18);
  margin-top: -100px;
  width: 100%;
  height: 100%;
}
.message{
  position: fixed;
  z-index: 1145142;
  width: 40%;
  max-height: 170px;
  border-radius: 10px;
  background-color: white;
  padding: 10px 40px 10px 40px;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  margin:auto;

}
</style>