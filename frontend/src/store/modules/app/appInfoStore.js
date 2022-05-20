import envInfoModule from "./modules/envInfoModule";
import gitModule from "./modules/gitModule";

const uiStore = {
  namespaced: true,
  state: {
      enviroment: process.env.VUE_APP_DEPLOYMENT_TYPE,

  },
  mutations: {},
  actions: {},
  modules: {
    version: gitModule,
    envInfo: envInfoModule,
  },
};

export default uiStore;
