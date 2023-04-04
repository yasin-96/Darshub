//Startpoint
import App from "./App.vue";

//CSS (Tailwind)
import "./assets/main.css";

import { vuetify } from "./plugins/vuetify";

//Vue App Creation
import { createApp } from "vue";
const app = createApp(App);

//Enable or Disable Devtools [performance flag]
const currentEnviroment = import.meta.env.VITE_APP_NODE_ENV;

if (currentEnviroment === "develop" || currentEnviroment === "local") {
  app.config.performance = true;
} else {
  app.config.performance = false;
}

//Router
import router from "./router";

//Store def.
import { createPinia } from "pinia";
import { createPersistedState } from "pinia-plugin-persistedstate";
const pinia = createPinia();
pinia.use(
  createPersistedState({
    key: (id) => `_persisted_darshub_${id}`,
  })
);

//I18n Languages
import { i18nLanguages } from "./plugins/i18nPlugin";

//SEO
import { createMetaManager, plugin as metaPlugin } from "vue-meta";
const metaManager = createMetaManager();

//Bind all option
app.use(router).use(pinia).use(i18nLanguages).use(metaManager).use(metaPlugin).use(vuetify);

//Bind to containder #app -> index.html
app.mount("#app");
