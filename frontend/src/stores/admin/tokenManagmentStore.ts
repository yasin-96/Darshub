import { defineStore } from "pinia";
import { useFetch } from "@vueuse/core";
import type { Auth0TokenDetails } from "@/models/admin/types";

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
      console.log(AUTH0_API);
      const { isFetching, error, data } = await useFetch<Auth0TokenDetails>(`https://${AUTH0_API}/api/oauth/token`, {
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
      });

      if (error.value) {
        console.log("Token was not created: ", error);
      }

      if (data.value) {
        this.tokenDetails = data.value;
        this.lastCreation = new Date().setSeconds(data.value.expires_in, 0);
      }
    },
  },
  getters:{
    getToken(): String {
        return this.tokenDetails.access_token;
    },
    getExpiredTimeInTicks(): number{
        return this.tokenDetails.expires_in + this.lastCreation
    }
  }
});
