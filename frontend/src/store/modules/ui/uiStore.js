import notificationModule from "./modules/notificationModule";

const uiStore = {
  namespaced: true,
  state: {},
  mutations: {},
  actions: {},
  modules: {
    notification: notificationModule,
  },
};

export default uiStore;
