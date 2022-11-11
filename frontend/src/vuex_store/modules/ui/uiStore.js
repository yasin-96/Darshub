import dialogsModule from "./modules/dialogsModule";
import notificationModule from "./modules/notificationModule";
import sidebarLeftModule from "./modules/sidebarLeftModule";

const uiStore = {
  namespaced: true,
  state: {},
  mutations: {},
  actions: {},
  modules: {
    notification: notificationModule,
    sidebarLeft: sidebarLeftModule,
    dialogs: dialogsModule
  },
};

export default uiStore;
