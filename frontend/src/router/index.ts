import { useCoreCourseStore } from "@/stores/course/coreCourseStore";
import { useLoginStore } from "@/stores/session/loginStore";
import {
  createRouter,
  createWebHistory,
  type NavigationGuardNext,
  type RouteLocationNormalized,
  type RouteMeta,
  type RouteParams,
  type RouteRecordRaw,
} from "vue-router";
const IndexView = () => import("@/views/index.vue");
const NotFoundView = () => import("@/views/404.vue");
const RegistrationView = () => import("@/views/registration.vue");
const LoginView = () => import("@/views/login.vue");
const UserDasboard = () => import("@/views/user/dashboard.vue");
const UserSettings = () => import("@/views/user/settings.vue");
const CourseOverview = () => import("@/views/course/overview.vue");

declare module "vue-router" {
  interface RouteMeta {
    // is optional
    isAdmin?: boolean;
    // must be declared by every route
    requiresAuth: boolean;
  }
}

const generalRoutes: Array<RouteRecordRaw> = [
  {
    path: "",
    name: "index",
    meta: { requiresAuth: false },
    component: IndexView,
  },
  {
    path: "/registry",
    name: "registry",
    meta: { requiresAuth: false },
    component: RegistrationView,
  },
  {
    path: "/logout",
    name: "logout",
    meta: { requiresAuth: false },
    component: {},
    beforeEnter: async (
      to: RouteLocationNormalized,
      from: RouteLocationNormalized,
      next: NavigationGuardNext
    ) => {
      useLoginStore().act_logUserOut();
      next({ name: "index" });
    },
  },
  {
    path: "/login",
    name: "login",
    meta: { requiresAuth: false },
    component: LoginView,
  },
  {
    path: "/about",
    name: "about",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    // component: () => import(/* webpackChunkName: "about" */ "../views/About.vue"),
    // authReq: 0,
    component: () => {},
  },
  {
    path: "/team",
    name: "team",
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
];

const userRoutes: Array<RouteRecordRaw> = [
  {
    path: "/user/:id",
    name: "user-dashboard",
    meta: { requiresAuth: true },
    component: UserDasboard,
    children: [
      {
        // UserProfile will be rendered inside User's <router-view>
        // when /user/:id/profile is matched
        path: "settings",
        name: "user-settings",
        meta: { requiresAuth: true },
        component: UserSettings,
        beforeEnter: async (
          to: RouteLocationNormalized,
          from: RouteLocationNormalized,
          next: NavigationGuardNext
        ) => {
          const params: RouteParams = to.params;
          const metaInfo: RouteMeta = to.meta;
          if (metaInfo.requiresAuth && useLoginStore().getUserId) {
            console.log(params.id);
            next();
          } else {
            alert("asdasd")
          }
        },
      },
    ],
    beforeEnter: async (
      to: RouteLocationNormalized,
      from: RouteLocationNormalized,
      next: NavigationGuardNext
    ) => {
      const params: RouteParams = to.params;
      const metaInfo: RouteMeta = to.meta;
      if (params.id && metaInfo.requiresAuth && useLoginStore().getUserId) {
        console.log(params.id);
        next();
      } else {
        next("/login");
      }
    },
  },
];

const courseRoutes: Array<RouteRecordRaw> = [
  {
    path: "/course/overview",
    name: "course-overview",
    component: CourseOverview,
    beforeEnter: async (
      to: RouteLocationNormalized,
      from: RouteLocationNormalized,
      next: NavigationGuardNext
    ) => {
      await useCoreCourseStore().act_loadAllCourseAsQuickInfo();
      next();
    },
  },
];

const adminRoutes: Array<RouteRecordRaw> = [
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
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [...generalRoutes, ...userRoutes, ...adminRoutes, ...courseRoutes],
});

router.beforeEach(async (to, from, next) => {
  // const isAuthenticated = store.getters["userStore/user/isAuthenticated"];
  // console.log("to.name:", to.name);
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
