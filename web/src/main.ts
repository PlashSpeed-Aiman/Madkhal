import { createApp } from 'vue'
import './index.css'
import App from './App.vue'
import Wifi from "./components/Wifi.vue";
import IMaalum from "./components/iMaalum.vue";
import Settings from "./components/Settings.vue";
import {createRouter,createWebHashHistory} from "vue-router";
import About from "./components/About.vue";

const routes = [
    { path: '/', component: Wifi },
    { path: '/imaalum', component: IMaalum },
    {path: '/settings',component: Settings},
    {path: '/about',component: About}
]

// 3. Create the router instance and pass the `routes` option
// You can pass in additional options here, but let's
// keep it simple for now.
const router = createRouter({
    // 4. Provide the history implementation to use. We are using the hash history for simplicity here.
    history: createWebHashHistory(),
    routes, // short for `routes: routes`
})
createApp(App).use(router).mount('#app')
