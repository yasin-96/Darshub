import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Rework from "../views/Rework.vue";
import RegistrationView from "../views/RegistrationView.vue";
import View_CourseOverview from "../views/View_CourseOverview.vue";
import View_CourseDashboard from "../views/Admin/Course/View_CourseDashboard.vue";
import store from "@/store/index.js";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "home",
    component: Home,
  },
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
    path: "/course",
    name: "Courses",
    component: View_CourseOverview,
    beforeRouteEnter: async (to, from, next) => {
      console.log("store:", store);

      store.dispatch("courseStore/general/act_loadAllCourses");
      next("/");
    },
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

// router.beforeEach(async (to, from, next) => {
//   const allowedPages = ["login", "about", "notfound"];
//   const isAuthenticated = store.getters["user/isAuthenticated"];

//   if (allowedPages.includes(to.name)) {
//     next();
//   } else if (!isAuthenticated) {
//     next({ name: "login" });
//   } else {
//     next();
//   }
// });

export default router;
