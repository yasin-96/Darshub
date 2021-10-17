import Vue from "vue";
import Vuex from "vuex";
import registrationModule from "@/store/modules/registrationModule";
import coreModule from "@/store/modules/coreModule";
import typeModule from "@/store/modules/typeModule";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    registry: registrationModule,
    core: coreModule,
    types: typeModule,
  },
});
