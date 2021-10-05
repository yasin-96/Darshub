import { createStore } from "vuex";
import { navbarModule } from "./modules/navbarModule/navbarModule";

export const store = createStore({
  modules: {
    navbar: navbarModule,
  },
  state() {
    return {
      count: 1,
    };
  },
});
