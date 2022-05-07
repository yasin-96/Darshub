import { shallowMount } from "@vue/test-utils";
import Registry from "@/components/Registry.vue";
import ImageUploader from "@/components/base/ImageUploader";

import { render, screen } from "@testing-library/vue";

import Vue from "vue";
import Vuetify from "vuetify";
import Vuex from "vuex";

Vue.use(Vuetify).use(Vuex);

const renderWithVuetify = (component, options, callback) => {
  const root = document.createElement("div");
  root.setAttribute("data-app", "true");

  return render(
    component,
    {
      container: document.body.appendChild(root),
      // for Vuetify components that use the $vuetify instance property
      vuetify: new Vuetify(),
      ...options,
    },
    callback
  );
};

describe("Registry.vue", () => {
  let data = {
    valid: false,
    date: null,
    birthdayModel: false,
    birthdayPicker: null,
    minDate: "1950-01-01",
  };

  let store = {
    state: {
      newUser: {
        firstName: "",
        lastName: "",
        name: "",
        birthday: "", //date
        avatar: "",
        email: "",
        telNr: "",
        occupation: "",
        company: "",
        school: "",
        subject: "",
        country: null,
        bio: "",
        role: [0], //define roles as numbers (enum)
        isActive: true,
      },
    },
    actions: {},
    mutations: {},
  };

  // before(() =>{

  // })

  // afterEach(()=>{

  // })

  it("Should render the component", () => {
    const { getByText } = renderWithVuetify(ImageUploader, {
      data: () => data,
      store: store,
    });
    console.log(getByText);
  });
});
