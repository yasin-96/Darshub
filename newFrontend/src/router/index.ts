import { createRouter, createWebHistory } from "vue-router";
const IndexView = () => import("@/views/index.vue");
const NotFoundView = () => import("@/views/404.vue");
const RegistrationView = () => import("@/views/registration.vue");
const LoginView = () => import("@/views/login.vue");

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "index",
      component: IndexView,
    },
    {
      path: "/registry",
      name: "registry",
      component: RegistrationView,
    },
    // {
    //   path: "/logout",
    //   name: "logout"
    //   // authReq: 0,
    // //   beforeRouteEnter: async (to, from, next) => {
    // //     console.log("store:", store);
    // //     store.dispatch("userStore/user/act_logUserOut");
    // //     next("/");
    // //   },
    // },
    {
      path: "/login",
      name: "Login",
      component: LoginView,
      // authReq: 0,
    },
    // {
    //   path: "/course",
    //   name: "Courses",
    //   component: View_CourseOverview,
    //   // beforeRouteEnter: async (to, from, next) => {
    //   //   console.log("store:", store);

    //   //   store.dispatch("courseStore/general/act_loadAllCourses");
    //   //   next("/");
    //   // },
    //   // authReq: 0,
    //   children: [
    //     {
    //       path: "List",
    //       component: View_CourseOverview,
    //     },
    //   ],
    // },
    {
      path: "/admin/course",
      name: "AdminCourseSite",
      // component: View_CourseDashboard,
      component: () => {},

      // beforeRouteEnter: async (to, from, next) => {
      //   console.log("store:", store);

      //   store.dispatch("courseStore/general/act_loadAllCourses");
      //   next("");
      // },
      // authReq: 1,
    },
    {
      path: "/about",
      name: "About",
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      // component: () => import(/* webpackChunkName: "about" */ "../views/About.vue"),
      // authReq: 0,
      component: () => {},
    },
    {
      path: "/team",
      name: "Team",
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      // component: () => import(/* webpackChunkName: "about" */ "../views/About.vue"),
      // authReq: 0,
      component: () => {},
    },
    {
      path: "/:pathMatch(.*)*",
      name: "404",
      component: NotFoundView,
    },
  ],
});

router.beforeEach(async (to, from, next) => {
  // const isAuthenticated = store.getters["userStore/user/isAuthenticated"];
  console.log("to.name:", to.name);
  // if (!to.authReq) {
  next();

  // if (to.authReq && !isAuthenticated) {
  // next({ name: "Login" });
  // }

  // if (to.authReq && isAuthenticated) {
  // next();
  // }
});

export default router;
