<!-- Please remove this file from your project -->
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
                  v-model="form.mise"
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
                  v-model="form.foddName"
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
                  v-model="form.foodPlace"
                  placeholder="場所入力ください"
                ></b-form-input>
            </b-form-group>
            <b-form-group
              id="input-group-4"
              label="場所の詳細"
              label-for="input-4"
              description="">>

              <b-form-input
                  id="input-4"
                  v-model="form.foodPlacesyousai"
                  placeholder="具体的に場所入力ください"
                ></b-form-input>
            <!-- 調べるネット、グーグルのAPIマップ開くとか
            groupとinputの関係調べる -->
            </b-form-group>
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
// import "~/assets/style/Style.scss";
// import AlertCustomer from './AlertCustomer.vue';
export default {
  name: "WriteNew",
  data(){
    return {
      form:{
        mise:'',
        foddName:'',
        foodPlace:'',
        foodPlacesyousai:'',//場所詳細
        FoodImg:'',
        FoodLikeFlag:'',
        FoodBunnRui:'',
        goOnFlag:'',
        // addTime:'2020/2/3',
        // goToMeseTime:'2021/3/4',
        // newUpdateTime:'2022/4/5',  apiで作る
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
      }
    }
  },
  props:['foodList'],
  components:{
    // AlertCustomer
  },
  created(){
  // Todo 編集の時にすでにあるデータを取得して画面表示でする
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
      this.$refs['edit'].hide();
    },
    handleSubmit(){
      // sumit処理
      //Todo await insert
      // okの場合、OK popup閉じる、だめの場合、チエック、await失敗の場合、失敗エラーmsg
      if(!this.checkFormValidity()){
        console.log('false')
      }else{
        console.log(this.form)
        this.$refs['edit'].hide();//写自定义事件回调事件
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
