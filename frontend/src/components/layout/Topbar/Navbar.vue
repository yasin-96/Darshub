<template>
  <div>
    <!-- App-Bar which is shown on top of whole App -->
    <v-app-bar app hide-on-scroll dense>
      <v-app-bar-nav-icon
        @click="showSidebarLeftDrawer = !showSidebarLeftDrawer"
      ></v-app-bar-nav-icon>

      <!-- <router-link to="/" tag="img" :src="require('../../assets/logo.png')" alt="SEF Logo" position="left" contain height="100%" class="mx-1" /> -->

      <v-spacer></v-spacer>

      <v-toolbar-title>
        <div class="title">{{ $t("appName") }}</div>
      </v-toolbar-title>

      <div class="mx-2"></div>

      <v-spacer></v-spacer>
      <v-menu
      transition="slide-x-transition"
      bottom
      right
      
    >
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          dark
          v-bind="attrs"
          v-on="on"
          color="info" outlined
          
        >
          Anmelden/ Registrieren
        </v-btn>
      </template>

      <v-list>
        <v-list-item @click="goToRegPage()"
        >
         <v-list-item-icon>
            <v-icon>mdi-account-plus</v-icon>
          </v-list-item-icon>
          <v-list-item-title @click="goToRegPage()">
            Registrieren</v-list-item-title>
        </v-list-item>
        <v-list-item  @click="goToLoginPage()"
        >
          <v-list-item-icon>
            <v-icon>mdi-login-variant</v-icon>
          </v-list-item-icon>
            
          <v-list-item-title @click="goToLoginPage()">
            Anmelden</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>


      <!-- <v-chip v-if="user && user.info.logId">
        <v-avatar left>
          <v-icon>mdi-account-circle</v-icon>
        </v-avatar>
        {{ userFullName }}
      </v-chip> -->
      <LocaleSwitcher />
    </v-app-bar>

    <!-- Navigation drawer on left side of UI -->
    <v-navigation-drawer app v-model="showSidebarLeftDrawer" temporary>
      <v-row>
        <v-col cols="12">
          <v-list dense>
            <v-list-item
              v-for="(aItem, i) in sidebarLeftNavigationLinks.all"
              :key="i"
              :link="aItem.route ? true : false"
              :to="aItem.route"
            >
              <v-list-item-icon>
                <v-icon>{{ aItem.icon }}</v-icon>
              </v-list-item-icon>

              <v-list-item-title>{{ $t(aItem.title) }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-col>
        <v-col class="12">
          <v-divider></v-divider>
        </v-col>

        <v-col cols="12">
          <v-list dense>
            <v-list-item
              v-for="(aItem, i) in sidebarLeftNavigationLinks.admin"
              :key="i"
              :link="aItem.route ? true : false"
              :to="aItem.route"
            >
              <v-list-item-icon>
                <v-icon>{{ aItem.icon }}</v-icon>
              </v-list-item-icon>

              <v-list-item-title>{{ $t(aItem.title) }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-col>
        <v-col class="12">
          <v-divider></v-divider>
        </v-col>
      </v-row>
      <template v-slot:append>
        <v-divider></v-divider>
        <v-list dense>
          <v-list-item>
            <v-list-item-icon>
              <v-icon>mdi-domain</v-icon>
            </v-list-item-icon>

            <v-list-item-title>Über uns</v-list-item-title>
          </v-list-item>
        </v-list>
      </template>
    </v-navigation-drawer>
  </div>
</template>

<script>
import LocaleSwitcher from "@/components/locales/LocaleSwitcher";

import { mapState } from "vuex";
export default {
  components: {
    LocaleSwitcher,
  },
  data: () => ({
    admins: [
      ["Management", "mdi-account-multiple-outline"],
      ["Settings", "mdi-cog-outline"],
    ],
    cruds: [
      ["Create", "mdi-plus-outline"],
      ["Read", "mdi-file-outline"],
      ["Update", "mdi-update"],
      ["Delete", "mdi-delete"],
    ],
  }),
  computed: {
    ...mapState({
      drawerPanelLeft: (state) => state.ui.sidebarLeft.showSidebarLeftDrawer,
      sidebarLeftNavigationLinks: (state) =>
        state.ui.sidebarLeft.sidebarLeftNavigationLinks,
    }),

    showSidebarLeftDrawer: {
      get() {
        return this.drawerPanelLeft;
      },
      set(toggleValue) {
        this.$store.dispatch(
          "ui/sidebarLeft/act_toggleSidebarLeftPanel",
          toggleValue
        );
      },
    },
  },
  methods: {
    goToRegPage() {
      this.$router.push({ name: "Registry" });
    },
    goToLoginPage(){
      this.$router.push({ name: "Login" });
    }
  },
};
</script>

<style scoped></style>
