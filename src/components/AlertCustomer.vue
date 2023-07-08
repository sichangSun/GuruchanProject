<template>
  <div>
    <div></div>
    <div>
      <b-button class="button-left"></b-button>
      <b-button class="button-right"></b-button>
    </div>
    <b-input-group> </b-input-group>
  </div>
</template>

<script>
export default {
  name: "AlertCustomer",
  props: {
    title: { type: String, require: true },
    msg: {
      type: String,
      require: true,
    },
    isshow: {
      type: Boolean,
    },
    buttonColor_left: {
      type: String,
    },
    buttonColor_right: {
      type: String,
    },
    button_left: {
      type: String,
    },
    button_right: {
      type: String,
    },
  },
};
</script>

<style></style>
