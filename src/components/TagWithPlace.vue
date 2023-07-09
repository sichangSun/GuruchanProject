<template>
  <div>
    <b-form-group label="" label-for="tags-with-dropdown" class="tag-waku">
      <b-form-tags id="tags-with-dropdown" v-model="value" no-outer-focus class="mb-2" >
        <ul v-if="filteroptions.length > 0" class="list-inline d-inline-block mb-2">
          <div v-for="tag in filteroptions" :key="tag.tagId" class="list-inline-item">
            <div class="tall">
              <b-form-tag
              :no-remove="true"
              :name="tag.tagId"
              ref="tagValue"
              >{{tag.title}}</b-form-tag>
              <b-button @click="removeTag(tag.tagId)">x</b-button>
            </div>
            <!-- variant 是bootstrap的变量.可能主要是控制css -->
          </div>
        </ul>

        <b-dropdown size="sm" variant="outline-secondary" block>
          <template #button-content>
            <b-icon icon="tag-fill"></b-icon> Choose tags
          </template>
          <b-dropdown-form @submit.stop.prevent="() => {}">
            <b-form-group
              label="タグ検索"
              label-for="tag-search-input"
              label-cols-md="auto"
              class="mb-0"
              label-size="sm"
              :description="searchDesc"
              :disabled="disabled"
            >
              <b-form-input
                v-model="search"
                id="tag-search-input"
                type="search"
                size="sm"
                autocomplete="off"
                ></b-form-input>
            </b-form-group>
          </b-dropdown-form>
          <b-dropdown-divider></b-dropdown-divider>
          <b-dropdown-item-button
            v-for="option in availableOptions"
            :key="option.tagId"
            :disabled="checkChoosed(option.tagId)"
            @click="optionClick(option.tagId)"
          >
            {{ option.title}}
          </b-dropdown-item-button>
          <b-dropdown-text v-if="isEmptyOfList">
            保存されているタグはありません。
          </b-dropdown-text>
        </b-dropdown>
      </b-form-tags>
    </b-form-group>
  </div>
</template>

<script>
export default {
  data() {
    return {
      options: [],// 源数组
      filteroptions: [],// title的数组
      search: '',
      value: [],
      isEmptyOfList: false,
      disabled: false
    }
  },
  props:['tags'],
  computed: {
    criteria() {
      // 検索のキーワードを処理
      return this.search.trim();
    },
    //计算属性 下拉菜单的数组
    availableOptions() {
      const criteria = this.criteria
      // valueに（上のtags）値あげる
      const options = this.options.filter(opt => opt.isDel !== 1)
      if (criteria) {
        // 検索アリの場合
        return options.filter(opt => opt.title.indexOf(criteria) > -1);

      }
      // 検索なしの場合

      console.log(options)
      return options
    },
    // 检索找不到时
    searchDesc() {
      if (!this.isEmptyOfList){
        if (this.criteria && this.availableOptions.length === 0) {
          return '一致するタグは存在していません。'
        }
      }
      return ''
    }
  },
  // watch:{
  //   criteria(newdata){
  //     if (!newdata && this.availableOptions.length === 0) {
  //       this.isEmptyOfList = true;
  //       this.disabled =true;
  //     }
  //   }
  // },
  created(){
    // 初期有无数据表示
    let isHavetagOption = false;
    if(!this.tags){
      this.isEmptyOfList = true;
    }else {
      this.tags.forEach(option =>{
        if (option.isDel === 0) {
          isHavetagOption = true;
        }})
      if(!isHavetagOption) {
        this.isEmptyOfList = true;
      }
    }
    this.options =this.tags;
    console.log(this.options);

  },
  methods: {
    optionClick(id) {
      console.log(id)
      const tagValue=this.options.find(opj => opj.tagId === id)
      if(tagValue){
        this.filteroptions.push(tagValue)
      }
      this.search = ''
      console.log("deeee",this.filteroptions)
    },
    removeTag(id){
      console.log("参数",id)
      const index = this.filteroptions.find(opj => opj.tagId===id)
      if (index) {
        this.filteroptions.splice(index, 1);
      }
    },
    checkChoosed(id){
      if (this.filteroptions.find(opj => opj.tagId===id)) {
        return true;
      }else {
        return false;
      }
    }
  }
}
</script>
<style>
.tall{
  weight: 100px;
  height: 40px;
  background-color: chocolate;
}
.tag-waku{
  background-color:transparent;
  border:0;
}
</style>

