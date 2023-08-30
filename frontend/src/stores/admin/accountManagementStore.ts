import { defineStore } from "pinia";
import { useFetch } from "@vueuse/core";

import type { Auth0UserExtendedObject } from "@/models/user/types";
import type { Auth0UserManagementDetails } from "@/models/admin/interfaces";
import type { Auth0UserStats } from "@/models/admin/types";
import { useTokenManagementStore } from "./tokenManagmentStore";

const AUTH0_API = import.meta.env.VITE_A0_DOMAIN_URL;
const AUTH0_API_CLIENT_ID = import.meta.env.VITE_A0_CLIENT_ID;
const AUTH0_API_CLIENT_SECRET = import.meta.env.VITE_A0_CLIENT_SECRET;

export const useAccountManagementStore = defineStore("accountManagementStore", {
  state: (): Auth0UserManagementDetails => ({
    users: [] as Array<Auth0UserExtendedObject>,
    stats: {} as Auth0UserStats,
  }),
  actions: {
    async getStats() {

    },

    async getAllUsersFromAuth0() {
      const tokenStore = useTokenManagementStore();
      console.log("Start fetching data for users");
      let currentDateInTicks: number = Date.parse(new Date().toString());
      if (currentDateInTicks > tokenStore.getExpiredTimeInTicks) {
        await tokenStore.getManagementToken();
      }

      const { isFetching, error, data } = await useFetch<
        Array<Auth0UserExtendedObject>
      >(`${AUTH0_API}/api/v2/users`, {
        async beforeFetch({ url, options, cancel }) {
          if (!tokenStore.getToken) cancel();
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

      console.log("Result of fetch");
      console.log([error, data])

      if (error.value) {
        console.log("err", error);
        this.$reset();
        return;
      }

      if (data.value) {
        console.log(data.value);
        this.users = data.value;
      }
    },
    async updatePersonalInformation() {},
    async updateContactInformation() {},
    async updateRole() {},
    async updateAppData() {},
    async updateMetaData() {},
  },
});
