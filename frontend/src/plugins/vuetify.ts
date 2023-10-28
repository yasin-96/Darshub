import "@mdi/font/css/materialdesignicons.css";
import "vuetify/styles";
import { createVuetify } from "vuetify";
import { aliases, mdi } from "vuetify/iconsets/mdi";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";
import { VStepper } from "vuetify/labs/VStepper";
import { VDatePicker } from "vuetify/lib/labs/components.mjs";

export const vuetify = createVuetify({
  components: {
    components,
    VStepper,
    VDatePicker,
  },
  date: {
    locale: {
    },
  },
  directives,
  icons: {
    defaultSet: "mdi",
    aliases,
    sets: {
      mdi,
    },
  },
  ssr: false,
});
