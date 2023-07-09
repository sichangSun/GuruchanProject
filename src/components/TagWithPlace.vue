<template>
  <div>
    <b-form-group label="Tagged input using dropdown" label-for="tags-with-dropdown">
      <b-form-tags id="tags-with-dropdown" v-model="value" no-outer-focus class="mb-2">
        <template v-slot="{ tags, disabled, addTag, removeTag }">
          <ul v-if="tags.length > 0" class="list-inline d-inline-block mb-2">
            <li v-for="tag in tags" :key="tag" class="list-inline-item">
              <b-form-tag
                @remove="removeTag(tag)"
                :title="tag"
                :disabled="disabled"
                variant="info"
                remove-label="Remove tag"
              >{{ tag.title }}</b-form-tag>
              <!-- variant 是bootstrap的变量.可能主要是控制css -->
            </li>
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
              :key="option.id"
              @click="onOptionClick({ option, addTag })"
            >
              {{ option.title}}
            </b-dropdown-item-button>
            <b-dropdown-text v-if="isEmptyOfList">
              保存されているタグはありません。
            </b-dropdown-text>
          </b-dropdown>
        </template>
      </b-form-tags>
    </b-form-group>
  </div>
</template>

<script>
export default {
  data() {
    return {
      options: [],// 源数组
      // filteroptions: [],// title的数组
      search: '',
      value: [],// tags用的数组
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
      const options = this.options.filter(opt => this.value.indexOf(opt) === -1 && opt.isDel !== 1)
      console.log(options);
      // titleのみ入ってる配列
      // const titleoptions = [];
      if (criteria) {
        // 検索アリの場合
        // const a = options.filter(opt => opt.title.indexOf(criteria) > -1);
        // console.log(a);
        // a.forEach(item => titleoptions.push(item.title));
        // return titleoptions;
        return options.filter(opt => opt.title.indexOf(criteria) > -1);

      }
      // 検索なしの場合
      // options.forEach(item => titleoptions.push(item.title));
      // return titleoptions;
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
    onOptionClick({ option, addTag }) {
      addTag(option.title)
      this.search = ''
    }
  }
}
</script>
<!-- <template>
  tagのcom,
    用插槽来实现在搜索组建里显示,v-for循环
  b-form-tags 可以实现tag
  <div>
    <b-form-group>
      <b-form-tags>来管理整个组建,包括选择的渲染表示 v-model="value" value是过滤以后出来的新数组
      <b-form-tags
      v-model="value"
      add-on-change
      no-outer-focus>

        <template v-slot="{ tags,inputAttrs, inputHandlers, disabled,removeTag}">

          <ul v-if="tags.length > 0" class="list-inline d-inline-block mb-2">
            <li v-for="tag in tags" :key="tag.tagId" class="list-inline-item">
              <b-form-tag
              @input="$emit('input', $event)"
              @remove="removeTag(tag)">
               <slot :tags="tagList"></slot>
                {{ tags.tagList.title }}
              </b-form-tag>
            </li>
          </ul>

          <b-dropdown>
            <template #button-content>
              Choose tags
            </template>
            <b-dropdown-item-button
              v-for="tag in availableOptions"
              :key="tag.tagId"
              v-bind="inputAttrs"
              @click="inputHandlers"
              :disabled="disabled"
              :options="availableOptions"
            >
              {{ tag.title }}
            </b-dropdown-item-button>
            <b-dropdown-text v-if="tagList.length === 0">
              There are no tags available to select
            </b-dropdown-text>
          </b-dropdown>
        </template>
      </b-form-tags>
    </b-form-group>
  </div>
</template> -->

<!-- <script>
export default {
  name: "TagWithPlace",
  data(){
    return {
      value:[],
      disabled:false
    }
  },
  prop:{
    tags:{}
  },
  watch:{
// Todo 监听value 有的话,就把下拉菜单里这一项非活性
  },
  computed:{
    // 原始数据过滤删掉的
    availableOptions(){
      return this.tagList.filter(tagOption => tagOption.isDel !== 1)
    }

  },
  methods:{
    inputHandlers(){
    }
  }
};
</script>
<style>

</style> -->
