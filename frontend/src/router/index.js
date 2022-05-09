import Vue from "vue";
import VueRouter from "vue-router";
// import Home from "../views/Home.vue";
import Rework from "../views/Rework.vue";
import RegistrationView from "../views/RegistrationView.vue";

Vue.use(VueRouter);

const routes = [
  
  // {
  //   path: "/",
  //   name: "home",
  //   component: Home,
  // },
  {
    path: "/",
    name: "Rework",
    component: Rework,
  },
  {
    path: "/registry",
    name: "Registry",
    component: RegistrationView,
  },
  {
    path: "/about",
    name: "About",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue"),
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
