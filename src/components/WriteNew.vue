<!-- Please remove this file from your project -->

import TagWithPlace from './TagWithPlace.vue';
<template>
  <div>
    <b-modal ref="edit" title="guruchan食べ物追加・編集"
    @ok="submitEvent" trigger="ok">
    <!-- @hidden="resetModal" Todo調べる -->
      <div class="d-block">
        <b-form @reset="onReset" ref="form">
          <div class="row">
            <b-form-group
              id="input-group-1"
              label="店名"
              label-for="input-1"
              description="必須">

                <b-form-input
                  id="input-1"
                  v-model="form.MiseName"
                  placeholder="店名入力ください"
                  required
                ></b-form-input>

            </b-form-group>
          </div>
          <b-form-group
              id="input-group-2"
              label="食べ物"
              label-for="input-2"
              description="必須">

                <b-form-input
                  id="input-2"
                  v-model="form.FoodName"
                  placeholder="食べ物入力ください"
                ></b-form-input>

            </b-form-group>
            <b-form-group
              id="input-group-3"
              label="場所"
              label-for="input-3"
              description="必須">

                <b-form-input
                  id="input-3"
                  v-model="form.FoodPlace"
                  placeholder="場所入力ください"
                ></b-form-input>
            </b-form-group>
            <span>tag</span>
            <TagWithPlace :tags="tagList"
            :foodTag="form.foodTag"
            ::tagChangeFlag="false"
            :tagSearchFlag="false"/>
            <!-- tag可能需要整体保存!!!! -->
            <span>
              <b-button>tag変更・追加</b-button>
              <b-form-checkbox>tagに追加</b-form-checkbox>
            </span>
            <b-form-group
              id="input-group-4"
              label="分類"
              label-for="input-4"
              description="">

              <b-form-input
                  id="input-4"
                  v-model="form.foodPlacesyousai"
                  placeholder="具体的に分類入力ください"
                ></b-form-input>
            <!-- 調べるネット、グーグルのAPIマップ開くとか
            groupとinputの関係調べる -->
            </b-form-group>
            <span>tag</span>
            <TagWithPlace :tags="tagList"
            :foodTag="form.foodTag"
            :tagChangeFlag="false"
            :tagSearchFlag="false"/>
            <span>
              <b-button>tag変更・追加</b-button>
              <b-form-checkbox>tagに追加</b-form-checkbox>
            </span>
            <span>行く予定？</span>
            <b-form-radio>行った？</b-form-radio>
            <b-form-radio>行ってない</b-form-radio>
            <template #modal-footer="{ ok, cancel}">
              <b-button type="submit" variant="primary" @click="ok()">Submit</b-button>
              <b-button variant="danger" @click="cancel()">cancel</b-button>
            </template>
            <b-button type="reset" variant="danger">Reset</b-button>
        </b-form>
      </div>
    </b-modal>
<!-- Todo alertcustomerでポップアップ作る、okPopupflagでvーif -->
    <!-- <div v-if="okPopupflag">
      <AlertCustomer
        :popOpMsg="AlertCustomerMsgOK">
        </AlertCustomer>
    </div>
    <div v-if="errPopupflag">
      <AlertCustomer
        :popOpMsg="AlertCustomerMsgErr">
        </AlertCustomer>
    </div> -->
  </div>
</template>

<script>
import TagWithPlace from "./TagWithPlace.vue";

export default {
  name: "WriteNew",
  data(){
    return {
      form:{
        FoodId:null,
        MiseName:'',// 和tag连接
        FoodName:'',
        FoodPlace:'',// TOdo 場所詳細
        FoodImg:'',
        FoodLikeFlag:false,
        FoodBunnRui:'',
        addTime:'',
        goToMeseTime:'',
        newUpdateTime:'',
        goOnFlag:0,// 行ったことあるかどうか
        foodTag:[],
      },
      okPopupflag:false,
      errPopupflag:false,
      AlertCustomerMsgOK:{
        title:'おけ',
        msg:'',
        isshow:'',
        button_left:'',
        button_right:'',
      },
      AlertCustomerMsgErr:{
        title:'失敗',
        msg:'',
        isshow:'',
        button_left:'',
        button_right:'',
      },
    }
  },
  props:['foodList','changeID','tagList'],
  components:{
    TagWithPlace,
  },
  mounted(){
  // Todo 編集の時にすでにあるデータを取得して画面表示でする
    // this.foodListTem =this.foodList;
    // this.changeId=this.changeID;
    // console.log(this.changeId)
    // if (this.changeID) {
    //   // 編集
    // console.log('changeID',this.changeID)
    //   const foodObj = this.foodList.find(foodObj => {
    //     foodObj.FoodID === this.changeID
    //   })
    //   if (foodObj) {
    //     const foodObjChange = Object.assign({},foodObj)
    //     this.form=foodObjChange;
    //   }
    // } else {
    //   //add new
    //   console.log('add');
    // }
  },
  watch:{
    // flush: 'post',
    changeID(newVal){
      // 編集
      if (typeof(newVal) === 'number') {
        const foodTemp = this.foodList.find(foodObj => foodObj.FoodId === newVal)
        if (foodTemp) {
          const foodObjChange = Object.assign({},foodTemp)
          this.form=foodObjChange;
        }
      }else{
        this.form.FoodId=null,
        this.form.MiseName='',// 和tag连接
        this.form.FoodName='',
        this.form.FoodPlace='',// TOdo 場所詳細
        this.form.FoodImg='',
        this.form.FoodLikeFlag=false,
        this.form.FoodBunnRui='',
        this.form.addTime='',
        this.form.goToMeseTime='',
        this.form.newUpdateTime='',
        this.form.goOnFlag=0// 行ったことあるかどうか
        this.form.foodTag=[]
      }
    }
  },
  methods:{
    onSubmit(event){
      event.preventDefault()
    },
    onReset(event){
      event.preventDefault()
    },
    submitEvent(event){
      //sumit のあとの処理
      event.preventDefault();
      this.handleSubmit();
      // this.resetModal();// クリア
      // this.okPopupflag=true;
      // this.$refs['edit'].hide();

    },
    handleSubmit(){
      // sumit処理
      //Todo await insert

      if(!this.checkFormValidity()){
        this.food
      }else{
        console.log(this.form)
        this.$emit("form-submitted"); // 触发自定义事件
        // 清空data组件里得值
      }
    },
    checkFormValidity() {
      //項目チエック
      const valid = this.$refs.form.checkValidity()
      return valid
    },
    // resetModal(){
    //   this.form.mise='',
    //   this.form.foddName='',
    //   this.form.foodPlace='',
    //   this.form.foodPlacesyousai='',//場所詳細
    //   this.form.FoodImg='',
    //   this.form. FoodLikeFlag='',
    //   this.form.FoodBunnRui='',
    //   this.form.goOnFlag=''
    // Todo 空にする 問題ある！！！！！！！
    //},
  }
}
</script>
<style scoped>
.btn-add {
  float: left;
}
</style>
