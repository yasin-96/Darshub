import Vue from "vue";
import Vuex from "vuex";
import registrationModule from "@/store/modules/registrationModule";
Vue.use(Vuex);

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    registry: registrationModule,
  },
});
