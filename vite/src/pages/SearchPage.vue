<template>
  <div id="SearchPage">
    <AppBar/>
    <SearchBar></SearchBar>
    <SavedRolls></SavedRolls>
    <PageFooter/>
  </div>
</template>
<script>
import AppBar from "../components/AppBar.vue";
import PageFooter from "../components/PageFooter.vue";
import SearchBar from "../components/SearchBar.vue";
import SavedRolls from "../components/SavedRolls.vue";

export default {
  name:"SearchPage",
  components: {SavedRolls, SearchBar, PageFooter, AppBar},
}
</script>
<style>
#SearchPage{
  height: 100%;
}
</style>