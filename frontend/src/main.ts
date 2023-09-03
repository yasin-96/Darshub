//Startpoint
import App from "./App.vue";
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

//I18n Languages
import { i18nLanguages } from "./plugins/i18nPlugin";

//SEO
import { createMetaManager, plugin as metaPlugin } from "vue-meta";
const metaManager = createMetaManager();

//AUTH0
import { authClient } from "./plugins/auth0";

//MD Editor
import { mdEditor } from "./plugins/md";

//Bind all option
app
  .use(router)
  .use(pinia)
  .use(i18nLanguages)
  .use(metaManager)
  .use(metaPlugin)
  .use(vuetify)
  .use(authClient)
  .use(mdEditor);

//Bind to containder #app -> index.html
app.mount("#app");
