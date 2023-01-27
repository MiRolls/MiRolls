<template>
  <div id="searchBar">
    <input type="text" id="searchInput" placeholder="Enter Your Code"/>
  </div>
</template>

<script>
export default {
  name: "SearchBar"
}
</script>

<style scoped>
#searchBar{
  margin-top: 120px;
  padding: 20px 25px 20px 25px;
  box-shadow: 0 0 6px rgba(0, 0, 0, 0.47);
  width: 75%;
  margin-left: auto;
  margin-right: auto;
  border-radius: 10px;
  text-align: center;
}
#searchInput{
  font-size: 16px;
  border: none;
  border-bottom: 1px black solid;
  width: calc(100% - 150px);
  outline: none;
  text-align: left;
}
</style>