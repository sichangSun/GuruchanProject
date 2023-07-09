<template>
  <div>
    <div class="container">
      <div class="row">
        <!-- header title -->
        <h1>Guruchan</h1>
        <div></div>
      </div>
      <div class="row sousuo">
        <!-- sercharea -->
        <div class="col-xl-7">
          <div>
            <SerchArea
            :tagList="tagList"/>
          </div>
        </div>
        <div class="col-xl-4">
          2
          <like-button></like-button>
        </div>
      </div>
      <div class="row">
            <WriteNew />
      </div>
      <div class="row">
        <!-- body list -->
          <!-- 默认最近.搜索以后变换,点击收藏,换成收藏里的 -->
          <div v-for="food in foodObj" :key="food.FoodId">
            <FoodList
            :FoodList="foodObj"/>
          </div>
      </div>
    </div>
    <div></div>
  </div>
</template>

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
          FoodName:'背脂ちゃんはん',
          FoodPlace:'目黒',
          FoodImg:'。。',
          FoodLikeFlag:false,
          FoodBunnRui:'cyuka',
          addTime:'2020/2/3',
          goToMeseTime:'2021/3/4',
          newUpdateTime:'2022/4/5'
        },
        {
          FoodId:2,
          FoodName:'味噌ラーメン',
          FoodPlace:'五反田',
          FoodImg:'。。',
          FoodLikeFlag:true,
          FoodBunnRui:'ramenn',
          addTime:'2020/2/3',
          goToMeseTime:'2021/3/4',
          newUpdateTime:'2022/4/5'
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
          tagId:4,
          addTime:'2021/3/4',
          isDel:0// 削除の制御  1:削除、０：あり
        }
      ],
    }
  }
};
</script>
<style scoped>
.sousuo {
  background-color: aquamarine;
}
</style>
