<template>
  <div>
    <div>
      <b-button class="backTop" @click="backTop">
        <svg xmlns="http://www.w3.org/2000/svg" width="50" height="50" fill="currentColor" class="bi bi-arrow-bar-up" viewBox="0 0 16 16">
          <path fill-rule="evenodd" d="M8 10a.5.5 0 0 0 .5-.5V3.707l2.146 2.147a.5.5 0 0 0 .708-.708l-3-3a.5.5 0 0 0-.708 0l-3 3a.5.5 0 1 0 .708.708L7.5 3.707V9.5a.5.5 0 0 0 .5.5zm-7 2.5a.5.5 0 0 1 .5-.5h13a.5.5 0 0 1 0 1h-13a.5.5 0 0 1-.5-.5z"/>
        </svg>
      </b-button>
    </div>
    <div class="container">
      <div class="homu-all">
        <div class="row">
        <!-- header title -->
        <h1>Guruchan</h1>
        <div></div>
        </div>
        <div class="row sousuo">
        <!-- sercharea -->
        <div class="col">
        </div>
        <div class="col-xl-6">
          <div>
            <SerchArea
            :tagList="tagList"/>
          </div>
        </div>
        <div class="col-xl-3">
          <like-button></like-button>
        </div>
        <div class="col">

        </div>
        </div>
        <div class="row">
          <b-button @click="addNewFood" class="outline-primary">追加</b-button>
        </div>
        <div class="row">
          <!-- <div class="col"></div> -->
        <!-- body list -->
          <!-- 默认最近.搜索以后变换,点击收藏,换成收藏里的 -->
              <div>
                <FoodList
              :FoodList="foodObj"/>
              </div>
          <!-- <div class="col"></div> -->
      </div>
      </div>
    </div>
    <div></div>
    <WriteNew ref="editModal"></WriteNew>
  </div>
</template>
<!-- 機能：行きたい店 -->
<script>
import LikeButton from './LikeButton.vue';
import SerchArea from './SerchArea.vue';
import WriteNew from './WriteNew.vue';
import FoodList from './FoodList.vue';

export default {
  components: {
    LikeButton,
    SerchArea,
    WriteNew,
    FoodList
  },
  name: "IndexPage",
  data(){
    return {
      foodObj:[
        {
          FoodId:1,
          MiseName:'味安',// 和tag连接
          FoodName:'背脂チャーハン',
          FoodPlace:'目黒',// TOdo 場所詳細
          FoodImg:'。。',
          FoodLikeFlag:false,
          FoodBunnRui:'cyuka',
          addTime:'2020/2/3',
          goToMeseTime:'2021/3/4',
          newUpdateTime:'2022/4/5',
          goOnFlag:'0'// 行ったことあるかどうか
        },
        {
          FoodId:2,
          MiseName:'ミラン',
          FoodName:'味噌ラーメン',
          FoodPlace:'五反田',
          FoodImg:'。。',
          FoodLikeFlag:true,
          FoodBunnRui:'ramenn',
          addTime:'2020/2/3',
          goToMeseTime:'2021/3/4',
          newUpdateTime:'2022/4/5',
          goOnFlag:'1'
        }],
      tagList: [
        {
          syuRui:'1',// タグの種類：食べ物の種類１、場所２、
          title:'目黒',
          tagId:1,
          addTime:'2021/3/4',
          isDel:0// 削除の制御  1:削除、０：あり

        },
        {
          syuRui:'1',// タグの種類：食べ物の種類１、場所２、
          title:'五反田',
          tagId:2,
          addTime:'2021/3/4',
          isDel:0// 削除の制御  1:削除、０：あり

        },
        {
          syuRui:'1',// タグの種類：食べ物の種類１、場所２、
          title:'中目黒',
          tagId:3,
          addTime:'2021/3/4',
          isDel:0// 削除の制御  1:削除、０：あり
        },
        {
          syuRui:'1',// タグの種類：食べ物の種類１、場所２、
          title:'新宿',
          tagId:4,
          addTime:'2021/3/4',
          isDel:0// 削除の制御  1:削除、０：あり
        },
        {
          syuRui:'1',// タグの種類：食べ物の種類１、場所２、
          title:'shibu',
          tagId:5,
          addTime:'2021/3/4',
          isDel:1// 削除の制御  1:削除、０：あり
        }
      ],
      editModal: false
    }
  },
  // created:{
  //   // this.$refs.eidtModal.$el.style.display = "block";
  // },
  methods:{
    backTop(){
      //Todo backtop
      console.log("hh")
    },
    addNewFood(){
      this.$refs.editModal.$refs['edit'].show();
    }
  }
};
</script>
<style scoped>
.sousuo {
  /* background-color: aquamarine; */
}
.backTop{
  float: left;
  margin-left: 5px;
  margin-top: 500px;

}
/* TODO right */
</style>
