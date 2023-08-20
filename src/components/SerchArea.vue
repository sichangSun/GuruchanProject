<template>
  <div id="S">
    <b-form @submit="onSubmit" @reset="onReset">
      {{ message }}
      <b-form-group label="検索" placeholder="“ここから入力してください”">
        <b-form-input v-model="serchFood">
        </b-form-input>
      </b-form-group>
      <TagWithPlace :tags="tagList"
      :tagChangeFlag="true"
      :tagSearchFlag="true"/>
      <b-button type="submit" variant="primary">提交</b-button>
      <b-button type="reset" variant="danger">重置</b-button>
    </b-form>
  </div>
</template>

<script>
// eslint-disable-next-line no-unused-vars
import TagWithPlace from "./TagWithPlace.vue";

export default{
  name: "SerchArea",
  components:{
    TagWithPlace
  },
  props:['tagList'],
  data() {
    return {
      // tagList: [
      //   {
      //     syuRui:'1',// タグの種類：食べ物の種類１、場所２、
      //     title:'目黒',
      //     tagId:1,
      //     addTime:'2021/3/4',
      //     isDel:0// 削除の制御  1:削除、０：あり

      //   },
      //   {
      //     syuRui:'1',// タグの種類：食べ物の種類１、場所２、
      //     title:'gotnd',
      //     tagId:2,
      //     addTime:'2021/3/4',
      //     isDel:0// 削除の制御  1:削除、０：あり

      //   },
      //   {
      //     syuRui:'1',// タグの種類：食べ物の種類１、場所２、
      //     title:'中目黒',
      //     tagId:3,
      //     addTime:'2021/3/4',
      //     isDel:1// 削除の制御  1:削除、０：あり

      //   }
      // ],
      message: "guruchan",
      serchFood: "",
      placeFlag: "", // 画面から選択する場所のflag
      syurui: "", // 画面から選択する食べ物の種類flag
      foodSyurui: [], // 和食・洋食・麺など種類
      placeSyurui: [], // 場所の配列
    };
  },
  methods: {
    onSubmit() {
      console.log("submitok");
    },
    onReset() {
      alert(1);
    },
  },
};
</script>
