import envInfoModule from "./modules/envInfoModule";

const uiStore = {
  namespaced: true,
  state: {
      enviroment: process.env.VUE_APP_DEPLOYMENT_TYPE,

  },
  mutations: {},
  actions: {},
  modules: {
    envInfo: envInfoModule,
  },
};

export default uiStore;
