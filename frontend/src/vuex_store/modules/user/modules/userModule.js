import RestBackendCalls from "@/services/RestBackendCalls";
import api_endpoints from "../../../../services/ResBackendAPIDefintions";

const userModule = {
  namespaced: true,
  state: () => ({
    info: {},
  }),
  actions: {
    act_logUserIn: async function ({ commit, dispatch }, data) {
      if (data.password && data.email) {
        RestBackendCalls.doPostRequest(api_endpoints.user.login, null, data)
          .then((resp) => {
            console.log("resp:", resp);

            if (resp && resp.data && resp.status == 200) {
              commit("MUT_SET_USER", resp.data);
              dispatch("ui/dialogs/act_toggleSignInDialog", false, {
                root: true,
              });
            }
          })
          .catch((err) => {
            commit("MUT_SET_USER", {});

            console.log(err);
          });
      }
    },
    act_logUserOut: function ({ commit }) {
      commit("MUT_CLEAR_USER");
    },
  },
  mutations: {
    MUT_SET_USER(state, userData) {
      state.info = userData;
    },
    MUT_CLEAR_USER(state) {
      state.info = {};
    },
  },
  getters: {
    getFullName(state) {
      const { First_Name, Last_Name, Email } = state.info;
      if (First_Name && Last_Name) {
        return `${First_Name} ${Last_Name}`;
      }

      if (First_Name) {
        return `${First_Name}`;
      }

      if (Last_Name) {
        return `${Last_Name}`;
      }
      if (Email) {
        const raw = Email.split("@")[0];
        return `${raw}`;
      }
      return "";
    },
    isAuthenticated(state) {
      return state.info && state.info.ID ? true : false;
    },
  },
};

export default userModule;
