export const navbarModule = {
  namespaced: true,
  state: () => ({
    sidebarDrawer: false,
  }),
  actions: {
    act_toggleSidebarPanel({ commit }) {
      commit("MUT_TOGGLE_SIDEBAR_PANEL");
    },
  },
  mutations: {
    MUT_TOGGLE_SIDEBAR_PANEL(state) {
      state.sidebarDrawer = !state.sidebarDrawer;
    },
  },
  getters: {},
};
