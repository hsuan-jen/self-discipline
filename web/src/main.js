import Vue from 'vue'
import App from './App.vue'
import router from './router'


new Vue({
    router,    // 注册到vue里
    render: h => h(App),
}).$mount('#app')
  
//createApp(App).mount('#app')
