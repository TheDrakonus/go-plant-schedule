import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
    {
      path: "/",
      component: () => import("@/pages/Home.vue"),
    },
    {
      path: "/about",
      component: () => import("@/pages/About.vue"),
    },
    {
      path: "/boxes",
      component: () => import("@/pages/Boxes.vue"),
    },
    {
      path: "/plants",
      component: () => import("@/pages/Plants.vue"),
    },
    {
      path: "/gardens",
      component: () => import("@/pages/Gardens.vue"),
    },
    {
      path: "/settings",
      component: () => import("@/pages/Settings.vue"),
    }
];

export default routes;