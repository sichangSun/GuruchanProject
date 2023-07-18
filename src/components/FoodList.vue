<template>
    <table class="col-xl-7 table">
            <thead>
              <tr>
                <th scope="col">写真</th>
                <th scope="col">食べ物</th>
                <th scope="col">店・場所</th>
                <th scope="col">分類</th>
                <th scope="col">記録日付</th>
                <th scope="col">更新日付</th>
                <th scope="col">tag</th>
                <th scope="col">お気に入り</th>
                <th scope="col">操作</th>

              </tr>
            </thead>
            <tbody  >
              <tr v-for="food in FoodList" :key="food.FoodId">
                <th scope="row" >
                        <td class="">
                            <img src="" alt="写真">
                        </td>
                        </th>
                        <th scope="row" >
                        <td class="">
                            <div>{{ food.FoodName }}</div>
                        </td>
                        </th>
                        <th scope="row" >
                        <td class="">
                            <div>{{ food.FoodPlace }}</div>
                            <div>
                                <b-button>詳細へ</b-button>
                            </div>
                        </td>
                        </th>
                        <th scope="row" >
                        <td class="">
                            <div>{{ food.FoodBunnRui}}</div>
                        </td>
                        </th>
                        <th scope="row" >
                        <td class="">
                            <div>{{ food.goToMeseTime}}</div>
                        </td>
                        </th>
                        <!-- ToDo 店行く日付まとめて、配列にしたほうがいい -->
                        <th scope="row" >
                        <td class="col-xl-2">
                            <div>{{ food.newUpdateTime}}</div>
                        </td>
                        </th>
                        <th scope="row" >
                        <td class="">
                            <div>{{ 'tag'}}</div>
                        </td>
                        </th>
                        <th scope="row" >
                        <td class="">
                            <div>{{ food.FoodLikeFlag?food.FoodLikeFlag:'なし'}}</div>
                            <!-- タグ最後は配列で -->
                        </td>
                        </th>
                        <th scope="row" >
                        <td class="col delBtn">
                        <b-button variant="danger">削除</b-button>
                        </td>
                    </th>
              </tr>
            </tbody>
          </table>
</template>

<script>
export default {
  name:'FoodList',
  comments:{
  },
  props:['FoodList']


}
</script>

<style>
.test{
    /* background-color: red; */
    margin-left: 5px;
}
.foodItemStyle{
    border: 2px 1;
    /* Todo 这里写死width,但我其实想让他跟主页面一样宽,所以想用100%,但是她好像找不到他的父标签,所以没有效果,改 */
    height: 100px;
    background-color: gray;
    margin: 7px;

}
.col-xl-4{
    background-color: aquamarine;
    margin: 3px;
}
.col-xl-1 {
    background-color: aquamarine;
    margin: 3px;
    weight: 30px;
}
.delBtn{
    margin: 4px;
}
</style>