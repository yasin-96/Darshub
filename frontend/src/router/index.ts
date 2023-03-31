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
import IndexView from "@/views/startPage.vue";
import NotFoundView from "@/views/404.vue";
import ContactPage from "@/views/contactPage.vue";
import LoginView from "@/views/login.vue";
import UserDasboard from "@/views/user/userDashboard.vue";
import UserSettings from "@/views/user/userSettings.vue";
import UserCourses from "@/views/user/userCourses.vue";
import CourseOverview from "@/views/course/courseOverview.vue";
import CoursePreview from "@/views/course/coursePreview.vue";
import CourseManagement from "@/views/admin/courseManagement.vue";
import UserManagement from "@/views/admin/userManagement.vue";

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
    meta: { requiresAuth: false },

    component: () => {},
  },
  {
    path: "/impress",
    name: "impress",
    meta: { requiresAuth: false },
    component: () => {},
    beforeEnter: async (
      to: RouteLocationNormalized,
      from: RouteLocationNormalized,
      next: NavigationGuardNext
    ) => {
      next({ name: "index" });
    },
  },
  {
    path: "/team",
    name: "team",
    meta: { requiresAuth: false },
    component: () => {},
    beforeEnter: async (
      to: RouteLocationNormalized,
      from: RouteLocationNormalized,
      next: NavigationGuardNext
    ) => {
      next({ name: "index" });
    },
  },
  {
    path: "/contact",
    name: "contact",
    meta: { requiresAuth: false },
    component: ContactPage,
  },
  {
    path: "/:pathMatch(.*)*",
    name: "404",
    meta: { requiresAuth: false },

    component: NotFoundView,
  },
];

const userRoutes: Array<RouteRecordRaw> = [
  {
    path: "/user",
    name: "user",
    beforeEnter: async (
      to: RouteLocationNormalized,
      from: RouteLocationNormalized,
      next: NavigationGuardNext
    ) => {
      const metaInfo: RouteMeta = to.meta;
      if (metaInfo.requiresAuth && useLoginStore().getUserId) {
        next();
      } else {
        next({ name: "login" });
      }
    },
    children: [
      {
        path: "overview",
        name: "user-dashboard",
        meta: { requiresAuth: true },
        component: UserDasboard,
        beforeEnter: async (
          to: RouteLocationNormalized,
          from: RouteLocationNormalized,
          next: NavigationGuardNext
        ) => {
          next();
        },
      },
      {
        path: "course",
        name: "user-course",
        meta: { requiresAuth: true },
        component: UserCourses,
        beforeEnter: async (
          to: RouteLocationNormalized,
          from: RouteLocationNormalized,
          next: NavigationGuardNext
        ) => {
          next();
        },
      },
      {
        path: "settings",
        name: "user-settings",
        meta: { requiresAuth: true },
        component: UserSettings,
        beforeEnter: async (
          to: RouteLocationNormalized,
          from: RouteLocationNormalized,
          next: NavigationGuardNext
        ) => {
          const metaInfo: RouteMeta = to.meta;
          if (metaInfo.requiresAuth && useLoginStore().getUserId) {
            next();
          } else {
            console.log(useLoginStore().getUserId);
            // next({ name: "login" });
          }
        },
      },
    ],
  },
];

const courseRoutes: Array<RouteRecordRaw> = [
  {
    path: "/course",
    name: "course",
    children: [
      {
        path: "overview",
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
      {
        path: ":cId/preview",
        name: "course-preview",
        component: CoursePreview,
        beforeEnter: async (
          to: RouteLocationNormalized,
          from: RouteLocationNormalized,
          next: NavigationGuardNext
        ) => {
          const params: RouteParams = to.params;
          console.log(params);
          if (
            params.cId &&
            params.cId !== "undefined" &&
            params.cId !== "null" &&
            params.cId.length
          ) {
            await useCoreCourseStore().act_loadCourseById(String(params.cId));
            next();
          } else {
            next({ name: "course-overview" });
          }
        },
      },
    ],
  },
];

const adminRoutes: Array<RouteRecordRaw> = [
  {
    path: "/admin",
    name: "admin",
    meta: { requiresAuth: true },
    beforeEnter: async (
      to: RouteLocationNormalized,
      from: RouteLocationNormalized,
      next: NavigationGuardNext
    ) => {
      const params: RouteParams = to.params;
      const metaInfo: RouteMeta = to.meta;

      console.log(params);
      if (
        metaInfo.requiresAuth &&
        useLoginStore().getUserId &&
        useLoginStore().userHasAdminRights
      ) {
        next();
      } else {
        next({ name: "index" });
      }
    },
    children: [
      {
        path: "user/management",
        name: "user-manager",
        component: UserManagement,
        beforeEnter: async (
          to: RouteLocationNormalized,
          from: RouteLocationNormalized,
          next: NavigationGuardNext
        ) => {
          next();
        },
      },
      {
        path: "course/management",
        name: "course-manager",
        component: CourseManagement,
        beforeEnter: async (
          to: RouteLocationNormalized,
          from: RouteLocationNormalized,
          next: NavigationGuardNext
        ) => {
          await useCoreCourseStore().act_loadAllCourseAsQuickInfo();
          next();
        },
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [...generalRoutes, ...userRoutes, ...adminRoutes, ...courseRoutes],
});

//General Navigation Guards here
// router.beforeEach(async (to, from, next) => {
//  const isAuthenticated = store.getters["userStore/user/isAuthenticated"];
//  console.log("to.name:", to.name);
//  if (!to.authReq) {
//  next();
//  if (to.authReq && !isAuthenticated) {
//  next({ name: "Login" });
//  }
//  if (to.authReq && isAuthenticated) {
//  next();
//  }
// });

export default router;
