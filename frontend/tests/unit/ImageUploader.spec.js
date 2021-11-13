import { shallowMount } from "@vue/test-utils";
import ImageUploader from "@/components/base/ImageUploader";

import {render, screen} from '@testing-library/vue';

import Vue from "vue"
import Vuetify from "vuetify"
import Vuex from "vuex"

Vue.use(Vuetify).use(Vuex)

const renderWithVuetify = (component, options, callback) => {
  const root = document.createElement('div')
  root.setAttribute('data-app', 'true')

  return render(
    component,
    {
      container: document.body.appendChild(root),
      // for Vuetify components that use the $vuetify instance property
      vuetify: new Vuetify(),
      ...options,
    },
    callback,
  )
}


describe("Registry.vue", () => {
  let data = {
    maxFileSize: 2 * 1024 * 1024,
    hasError: false,
    displayErrorMessage: "",
    fileReader: null,
    maxWidth: 300,
    maxHeight: 150,
  }

  let store = {
    state: {
      images: ["image/png", "image/jpeg", "image/svg"],    <%
        if (captain.hasRootSsl) {
    %>
        listen              443 ssl;
        ssl_certificate     /nginx-shared/bundlcr.crt;
        ssl_certificate_key /nginx-shared/caprover.mni.thm.de.key;

    <%
        }
    %>
    },
    actions:{

    },
    mutations:{}
  }

  // before(() =>{

  // })

  // afterEach(()=>{

  // })
  
  
  it("Should render the component", () => {
	  const { getByText } = renderWithVuetify(ImageUploader, {data: ()=>data, store: store });
    console.log(getByText)
  });
});
