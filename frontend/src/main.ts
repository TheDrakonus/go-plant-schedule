import { createApp } from "vue";
import { RouteRecordRaw, createRouter, createWebHistory } from "vue-router";
import App from "@/App.vue";
import Home from "@/pages/Home.vue";
import { createVuestic } from "vuestic-ui";
import "vuestic-ui/css";
import "material-design-icons-iconfont/dist/material-design-icons.min.css";
// Import our custom CSS
import "./scss/styles.scss";

import router from "./router";

createApp(App)
  .use(router)
  .use(createVuestic())
  .mount("#app");
