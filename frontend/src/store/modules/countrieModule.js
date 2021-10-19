import RestBackendCalls from "@/services/RestBackendCalls";

const API_HOST = process.env.VUE_APP_EXTERN_COUNTRIE_API;

const countrieModule = {
  namespaced: true,
  state: () => ({
    countries: [],
  }),
  actions: {
    act_loadAllCountries: function ({ commit }) {
      RestBackendCalls.doGetRequest(API_HOST)
        .then((res) => {
          console.log("res.data", res.data);
          if (res && res.data && res.status == 200) {
            commit("MUT_SET_ALL_COUNTRIES", res.data);
          }
        })
        .catch((err) => {
          console.error("act_loadAllCountries", err);
        });
    },
  },
  mutations: {
    MUT_SET_ALL_COUNTRIES: function (state, data) {
      state.countries = data;
    },
  },
  getters: {
    getCountriesWithFlags: function (state) {
      return state.countries.map((item) => {
        const name = item.name.common;
        const flag = item.flags.svg ? item.flags.svg : item.flags.png;
        const icon = item.flag;
        return { icon, flag, name };
      });
    },
  },
};

export default countrieModule;
