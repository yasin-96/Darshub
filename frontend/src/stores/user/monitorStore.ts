import { defineStore } from "pinia";
import { useFetch } from "@vueuse/core";

import type { Auth0UserExtendedObject } from "@/models/user/types";
import type { Auth0UserManagementDetails } from "@/models/admin/interfaces";
import type { Auth0UserStats } from "@/models/admin/types";
import { useTokenManagementStore } from "../admin/tokenManagmentStore"; 
import { useLoginStore } from "../session/loginStore";


const AUTH0_API = import.meta.env.VITE_A0_DOMAIN_URL;
const AUTH0_API_CLIENT_ID = import.meta.env.VITE_A0_CLIENT_ID;
const AUTH0_API_CLIENT_SECRET = import.meta.env.VITE_A0_CLIENT_SECRET;

export const useUserMonitorStore = defineStore("userMonitorStore", {
  state: () => ({
    logs: [] as any
  }),
  actions: {
    async getStats() {

    },

    async getUserLogs() {
      const tokenStore = useTokenManagementStore();
      const userId = useLoginStore().getUser;
      console.log("Start fetching logs data from user");
      let currentDateInTicks: number = Date.parse(new Date().toString());
      await tokenStore.getManagementTokenFromUserApi();
      if (currentDateInTicks > tokenStore.getExpiredTimeInTicks) {
        await tokenStore.getManagementTokenFromUserApi();
      }
        
      console.log(`${AUTH0_API}/api/v2/users/${userId}/logs`)
      const { isFetching, error, data } = await useFetch<
        Array<any>
      >(`${AUTH0_API}/api/v2/users/${userId}/logs`, {
        async beforeFetch({ url, options, cancel }) {
          console.log("STATE of token :"+tokenStore.getToken);
          if (!tokenStore.getToken) cancel();
          console.log(tokenStore.getToken);
          options.method = "GET";
          options.redirect = 'follow';
          options.headers = {
            ...options.headers,
            Accept: "application/json",
            Authorization: `Bearer ${tokenStore.getToken}`,
          };

          return {
            options,
          };
        },
      });


      if (error.value) {
        console.log("Error: Data from auth0 was not fetched", error);
        this.$reset();
        return;
      }

      if (data.value) {
        console.log(data.value);
        this.logs = data.value;
      }
    },
  },
});
