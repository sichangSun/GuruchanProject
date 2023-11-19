import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import PagedeNew from './components/PagedeNew.vue'
import indexList from './components/indexList.vue'

import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
// Install BootstrapVue
Vue.use(BootstrapVue)
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin)
Vue.use(VueRouter);

Vue.config.productionTip = false
const routes = [
  { path: '/PagedeNew', component: PagedeNew,name:'newpage'},
  { path: '/indexList', component: indexList,name:'indexList'},
  // { path: '/about', component: About },
];

const router = new VueRouter({
  routes,
});

new Vue({
  render: h => h(App),
  router
}).$mount('#app')
