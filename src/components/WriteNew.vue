<!-- Please remove this file from your project -->
<template>
  <div>
    <b-modal ref="edit" title="guruchan食べ物追加・編集">
      <div class="d-block">
        <b-form @submit="onSubmit" @reset="onReset">
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

          <b-button type="submit" variant="primary">Submit</b-button>
          <b-button type="reset" variant="danger">Reset</b-button>
        </b-form>
      </div>
    </b-modal>
  </div>
</template>

<script>
// import "~/assets/style/Style.scss";
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
      }
    }
  },
  props:['foodList'],
  created:{
  // Todo 編集の時にすでにあるデータを取得して画面表示でする
  },
  methods:{
    onSubmit(event){
      event.preventDefault()
    },
    onReset(event){
      event.preventDefault()
    }
  }
}
</script>
<style scoped>
.btn-add {
  float: left;
}
</style>
