import notificationModule from "./modules/notificationModule";
import sidebarLeftModule from "./modules/sidebarLeftModule";

const uiStore = {
  namespaced: true,
  state: {},
  mutations: {},
  actions: {},
  modules: {
    notification: notificationModule,
    sidebarLeft: sidebarLeftModule
  },
};

export default uiStore;
