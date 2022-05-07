const genderModule = {
  namespaced: true,
  state: () => ({
    genders: [
      { sex: "MALE", icon: "mdi-gender-male" },
      { sex: "FEMALE", icon: "mdi-gender-female" },
    ],
    role: {
      ADMIN: -41,
      USER: 0,
      AUTHOR: 1,
    },
  }),
  actions: {},
  mutations: {},
  getters: {},
};

export default genderModule;
