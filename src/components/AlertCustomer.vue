<template>
  <div>
    <div>
      <b-button class="button-left">{{ popOpMsgObj.button_left }}</b-button>
      <b-button class="button-right">{{ popOpMsgObj.button_right }}</b-button>
    </div>
    <b-input-group> </b-input-group>
  </div>
</template>

<script>
export default {
  name: "AlertCustomer",
  //Todo propsをobjを修正！
  props: {
    // title: { type: String, require: true },
    // msg: {
    //   type: String,
    //   require: true,
    // },
    // isshow: {
    //   type: Boolean,
    // },
    // buttonColor_left: {
    //   type: String,
    // },
    // buttonColor_right: {
    //   type: String,
    // },
    // button_left: {
    //   type: String,
    // },
    // button_right: {
    //   type: String,
    // },
    popOpMsg:{
      type:Object,
      required:true
    },

  },
  data(){
    return{
      popOpMsgObj:{
        title:'',
        msg:'',
        isshow:'',
        button_left:'',
        button_right:'',
      }
    }
  },
  methods(){

  },
  mounted(){
    this.popOpMsgObj=this.popOpMsg;

  }
};
</script>

<style></style>
