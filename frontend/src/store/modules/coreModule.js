const genderModule = {
  namespaced: true,
  state: () => ({
    genders: ["MALE", "FEMALE"],
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
