const dialogsModule = {
  namespaced: true,
  state: {
    signInDialog: {
      showDialog: false,
    },
  },
  actions: {
    act_toggleSignInDialog: function({ commit }, toggleValue) {
      commit("MUT_TOGGLE_SIGNIN_DIALOG", toggleValue);
    },
  },
  mutations: {
    MUT_TOGGLE_SIGNIN_DIALOG(state, toggleValue) {
      state.signInDialog.showDialog = toggleValue;
    },
  },
};

export default dialogsModule;
