import userSettingsModule from "./modules/userSettingsModule";
import userModule from "./modules/userModule";

const userStore = {
  namespaced: true,
  modules: {
    settings: userSettingsModule,
    user: userModule,
  },
};

export default userStore;
