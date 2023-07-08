<template>
    <div class="row test">
        <div>
            <b-row>
                <b-col>
                    <div>
        <li>
            <div class="row foodItemStyle">
                <div class="col-xl-1">
                    <img src="" alt="写真">
                </div>
                <div class="col-xl-2">名称
                    <div>{{ FoodList.FoodName }}</div>
                </div>

                <div class="col-xl-1">地址
                    <div>{{ FoodList.FoodPlace }}</div>
                    <div>
                        <b-button>詳細へ</b-button>
                    </div>
                </div>
                <div class="col-xl-1">分類
                    <div>{{ FoodList.FoodBunnRui}}</div>
                </div>
                <div class="col-xl-1"> 初店行く日付
                    <div>{{ FoodList.goToMeseTime}}</div>
                </div>
                <!-- ToDo 店行く日付まとめて、配列にしたほうがいい -->
                <div class="col-xl-2"> 最後店行く日付
                    <div>{{ FoodList.newUpdateTime}}</div>
                </div>
                <div class="col-xl-1"> Tag
                    <div>{{ 'tag'}}</div>
                </div>
                <div class="col-xl-1"> 気に入る
                    <div>{{ FoodList.FoodLikeFlag?FoodList.FoodLikeFlag:'なし'}}</div>
                </div>
                <div class="col-xl-1 delBtn">
                   <b-button variant="danger">削除</b-button>
                </div>
            </div>
        </li>
    </div>
                </b-col>
            </b-row>
        </div>
    </div>
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
    width: 1130px;
    height: 100px;
    background-color: gray;
    margin: 4px;

}
.col-xl-4{
    background-color: aquamarine;
    margin: 1px;
}
.col-xl-1 {
    background-color: aquamarine;
    margin: 1px;
}
.delBtn{
    margin: 4px;
}
</style>