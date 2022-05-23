const sidebarLeftModule = {
  namespaced: true,
  state: {
    showSidebarLeftDrawer: false,
    sidebarLeftNavigationLinks: {
      all: [
        { title: "router.sidebar.all.home", icon: "mdi-home", route: "/" },
        {
          title: "router.sidebar.all.courseOverview",
          icon: "mdi-view-dashboard",
          route: "/course",
        },
      ],
     
      author: [
        {
          title: "router.sidebar.author.dashboard",
          icon: "mdi-book-cog-outline",
          route: "/author/create",
        },
        {
          title: "router.sidebar.author.create",
          icon: "mdi-book-plus-outline",
          route: "/author/dashboard",
        },
      ],
      userManagement: [
        {
          title: "router.sidebar.admin.userSettings",
          icon: "mdi-account-cog-outline",
          route: "/admin/users",
        },
      ],
      courseManagement: [
        {
          title: "router.sidebar.admin.courseSettings",
          icon: "mdi-book-cog-outline",
          route: "/admin/course",
        },
      ],
      admin: [
        {
          title: "router.sidebar.admin.userSettings",
          icon: "mdi-account-cog-outline",
          route: "/admin/users",
        },
        {
          title: "router.sidebar.admin.courseSettings",
          icon: "mdi-book-cog-outline",
          route: "/admin/course",
        },
      ],
      app: [
        {
          title: "Über uns",
          icon: "mdi-domain",
          route: "/about-us",
        },
      ],
    },

    // [
    //   { title: "router.dashboard", icon: "mdi-view-dashboard", route: "/" },
    //   { title: "router.contact", icon: "mdi-email", route: "/contact" },
    //   { title: "router.about", icon: "mdi-information", route: "/about" },
    //   { title: "router.logout", icon: "mdi-logout-variant", route: "/logout" },
    // ],
  },
  actions: {
    act_toggleSidebarLeftPanel({ commit }, toggleValue) {
      commit("MUT_TOGGLE_SIDEBAR_LEFT_PANEL", toggleValue);
    },
  },
  mutations: {
    MUT_TOGGLE_SIDEBAR_LEFT_PANEL(state, toggleValue) {
      state.showSidebarLeftDrawer = toggleValue;
    },
  },
};

export default sidebarLeftModule;
