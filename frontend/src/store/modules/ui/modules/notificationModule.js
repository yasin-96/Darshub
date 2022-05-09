const notificationModule = {
  namespaced: true,
  state: () => ({
    timeout: 5000,
    active: false,
    message: null,
    color: null,
    defaultColor: "info",
  }),
  actions: {
    act_setMessage: function ({ commit }, payload) {
      if (payload.message) {
        commit("MUT_SET_NOTIFICATION_INFO");
      }
    },
    act_clearNotificationPanel: function ({ commit }) {
      commit("MUT_CLEAR_NOTIFICATION_INFO");
    },
    act_setSuccessNotification: function ({ commit }, message) {
      commit("MUT_SET_SUCCESS_NOTIFICATION", message);
    },
    act_setErrorNotification: function ({ commit }, message) {
      commit("MUT_SET_ERROR_NOTIFICATION", message);
    },
  },
  mutations: {
    MUT_SET_NOTIFICATION_INFO(state, { message, color }) {
      state.active = true;
      state.message = message;
      state.color = color;
    },
    MUT_CLEAR_NOTIFICATION_INFO(state) {
      state.active = false;
      state.message = null;
      state.color = null;
    },
    MUT_SET_ERROR_NOTIFICATION(state, message) {
      state.active = true;
      state.message = message;
      state.color = "error";
    },
    MUT_SET_SUCCESS_NOTIFICATION(state, message) {
      state.active = true;
      state.message = message;
      state.color = "success";
    },
  },
  getters: {},
};

export default notificationModule;
