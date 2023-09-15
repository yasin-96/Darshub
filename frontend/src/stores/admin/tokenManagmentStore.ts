import { defineStore } from "pinia";
import { useFetch } from "@vueuse/core";
import type { Auth0TokenDetails } from "@/models/admin/types";
import axios, { AxiosHeaders } from "axios";

const AUTH0_API = import.meta.env.VITE_A0_DOMAIN_URL;
const AUTH0_API_CLIENT_ID = import.meta.env.VITE_A0_CLIENT_ID;
const AUTH0_API_CLIENT_SECRET = import.meta.env.VITE_A0_CLIENT_SECRET;

export const useTokenManagementStore = defineStore("tokenManagementStore", {
  state: () => ({
    tokenDetails: {} as Auth0TokenDetails,
    lastCreation: 0 as number,
  }),
  actions: {
    async getManagementToken() {
      const { isFetching, error, data } = useFetch<Auth0TokenDetails>(
        `${AUTH0_API}/api/oauth/token`,
        {
          async beforeFetch({ url, options, cancel }) {
            options.method = "POST";

            options.headers = {
              ...options.headers,
              "Content-Type": "application/x-www-form-urlencoded",
            };

            options.body = new URLSearchParams({
              grant_type: "client_credentials",
              client_id: AUTH0_API_CLIENT_ID,
              client_secret: AUTH0_API_CLIENT_SECRET,
              audience: `${AUTH0_API}/api/v2/`,
            });

            return {
              options,
            };
          },
          timeout: 30000,
        }
      );

      if (error.value) {
        console.log("Token was not created: ", error);
      }

      if (data.value) {
        this.tokenDetails = data.value;
        this.lastCreation = new Date().setSeconds(data.value.expires_in, 0);
      }
    },
    async getManagementTokenV2() {
      //let headers = new
      let payload = new FormData();
      payload.append("grant_type", "client_credentials");
      payload.append("client_id", AUTH0_API_CLIENT_ID);
      payload.append("client_secret", AUTH0_API_CLIENT_SECRET);
      payload.append("audience", `${AUTH0_API}/api/v2/`);

      var hPayload = new AxiosHeaders();
      hPayload.set("Content-Type", "application/x-www-form-urlencoded");

      console.log("Start token v2");

      let response = await axios.post<Auth0TokenDetails>(
        `${AUTH0_API}/api/oauth/token`,
        payload,
        {
          headers: hPayload,
        }
      );

      console.log(response.data);
      if (response.status == 200) {
        this.tokenDetails = response.data;
        this.lastCreation = new Date().setSeconds(response.data.expires_in, 0);
      }

      console.log("Ende token v2");
    },
  },
  getters: {
    getToken(): String {
      return this.tokenDetails.access_token;
    },
    getExpiredTimeInTicks(): number {
      return this.tokenDetails.expires_in + this.lastCreation;
    },
  },
});
