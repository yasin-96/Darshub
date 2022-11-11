import Vue from "vue";
import Vuex from "vuex";
import registrationModule from "@/store/modules/registrationModule";
import coreModule from "@/store/modules/coreModule";
import typeModule from "@/store/modules/typeModule";
import countrieModule from "@/store/modules/countrieModule";
import uiStore from "./modules/ui/uiStore";
import courseStore from "./modules/course/courseStore";
import userStore from "./modules/user/userStore";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    registry: registrationModule,
    core: coreModule,
    types: typeModule,
    externApi: countrieModule,
    ui: uiStore,
    courseStore: courseStore,
    userStore: userStore
  },
});
