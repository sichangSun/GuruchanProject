<template>
  <div>
    <!-- 主页应该只放路由,然后在入口进入以后,导航把第一页东西放进来, -->
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
              :FoodList="foodObj"
              @emitFood = "emitFood"/>
              </div>
          <!-- <div class="col"></div> -->
      </div>
      </div>
    </div>
    <div></div>
    <WriteNew ref="editModal"
    @form-submitted="handleFormSubmitted"
    :changeID="changeID"
    :tagList="tagList"
    :foodList="foodObj"
    ></WriteNew>
    <div>
      <b-button @click="turnde" class="outline-primary">追加</b-button>
    </div>
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
          goOnFlag:0,// 行ったことあるかどうか
          foodTag:[]
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
          goOnFlag:1,
          foodTag:[1,2],
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
      editModal: false,
      changeID:''// 编辑时foodID
    }
  },
  created(){
    // TOdo 初期化
    // getDateAPI
    // 会社copy!
  },
  methods:{
    backTop(){
      //Todo backtop
      //console.log("hh")
    },
    addNewFood(id){
      this.changeID = id;
      this.$refs.editModal.$refs['edit'].show();
      // this.changeID = id;
    },
    handleFormSubmitted() {
      // 在提交表单后隐藏弹窗
      this.$nextTick(() => {
        this.$refs.editModal.$refs['edit'].hide();
      })
    },
    // 编辑食物
    emitFood(id){

      this.addNewFood(id);
      // this.changeID= id;
      // console.log( 'indexList  changeID!!!!!!!!',this.changeID)
    },
    // getDateAPI

    turnde(){
      console.log('ok')
      this.$router.push({ name:'newpage'})
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
