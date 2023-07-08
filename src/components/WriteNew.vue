<!-- Please remove this file from your project -->
<template>
  <div>
    <b-button @click="addNewFood" class="btn-add">追加</b-button>
  </div>
</template>

<script>
// import "~/assets/style/Style.scss";
export default {
  name: "WriteNew",
  methods:{
    addNewFood(){
      console.log("zhuijia");
    }
  }
};
</script>
<style scoped>
.btn-add {
  float: left;
}
</style>
