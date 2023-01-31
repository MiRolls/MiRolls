<template>
  <div id="queryPage"></div>
  <app-bar></app-bar>
  <left-chooser mode="details"></left-chooser>
  <details-page></details-page>
  <page-footer style="margin-left: 20%"></page-footer>
</template>
<script>
import AppBar from "../components/AppBar.vue";
import LeftChooser from "../components/LeftChooser.vue";
import PageFooter from "../components/PageFooter.vue";
import DetailsPage from "../components/detailsPage.vue";

export default {
  name: "QueryPage",
  components: {DetailsPage, PageFooter, LeftChooser, AppBar},
}
</script>
<style scoped>
#queryPage{
  /*height: calc(100% - 73px);*/
  height: 100%;
}
</style>