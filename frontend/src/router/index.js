import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Rework from "../views/Rework.vue";
import RegistrationView from "../views/RegistrationView.vue";
import LoginView from "../views/LoginView.vue";
import View_CourseOverview from "../views/View_CourseOverview.vue";
import View_CourseDashboard from "../views/Admin/Course/View_CourseDashboard.vue";
import store from "@/store/index.js";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
    authReq: 0,
  },
  {
    path: "/rework",
    name: "Rework",
    component: Rework,
    authReq: 0,
  },
  {
    path: "/registry",
    name: "Registry",
    component: RegistrationView,
    authReq: 0,
  },
  {
    path: "/logout",
    name: "Logout",
    authReq: 0,
    beforeRouteEnter: async (to, from, next) => {
      console.log("store:", store);
      store.dispatch("userStore/user/act_logUserOut");
      next("/");
    },
  },
  {
    path: "/login",
    name: "Login",
    component: LoginView,
    authReq: 0,
  },

  {
    path: "/course",
    name: "Courses",
    component: View_CourseOverview,
    beforeRouteEnter: async (to, from, next) => {
      console.log("store:", store);

      store.dispatch("courseStore/general/act_loadAllCourses");
      next("/");
    },
    authReq: 0,
    children: [
      {
        path: "List",
        component: View_CourseOverview,
      },
    ],
  },
  {
    path: "/admin/course",
    name: "AdminCourseSite",
    component: View_CourseDashboard,
    beforeRouteEnter: async (to, from, next) => {
      console.log("store:", store);

      store.dispatch("courseStore/general/act_loadAllCourses");
      next("");
    },
    authReq: 1,
  },
  {
    path: "/about",
    name: "About",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue"),
    authReq: 0,
  },
  {
    path: "/team",
    name: "Team",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue"),
    authReq: 0,
  },
  {
    path: "/team",
    name: "Team",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue"),
    authReq: 0,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

router.beforeEach(async (to, from, next) => {
  const isAuthenticated = store.getters["userStore/user/isAuthenticated"];
  console.log("to.name:", to.name);
  if(!to.authReq){
    next();
  }

  if(to.authReq && !isAuthenticated){
    next({ name: "Login" });
  }

  if(to.authReq && isAuthenticated){
    next();
  }
});

export default router;
