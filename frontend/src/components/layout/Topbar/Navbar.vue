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
      <v-btn color="info" outlined class="mr-5" @click="goToRegPage()">
        Registrieren
      </v-btn>

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
        <!-- 
        <v-col cols="12">
          <v-list dense>
            <v-list-group :value="true" prepend-icon="mdi-account-group">
              <template v-slot:activator>
                <v-list-item-title>Benutzer</v-list-item-title>
              </template>

              <v-list-group :value="true" no-action sub-group>
                <template v-slot:activator>
                  <v-list-item-content>
                    <v-list-item-title>Admin</v-list-item-title>
                  </v-list-item-content>
                </template>

                <v-list-item v-for="([title, icon], i) in admins" :key="i" link>
                  <v-list-item-title v-text="title"></v-list-item-title>

                  <v-list-item-icon>
                    <v-icon v-text="icon"></v-icon>
                  </v-list-item-icon>
                </v-list-item>
              </v-list-group>

              <v-list-group no-action sub-group>
                <template v-slot:activator>
                  <v-list-item-content>
                    <v-list-item-title>Actions</v-list-item-title>
                  </v-list-item-content>
                </template>

                <v-list-item v-for="([title, icon], i) in cruds" :key="i" link>
                  <v-list-item-title v-text="title"></v-list-item-title>

                  <v-list-item-icon>
                    <v-icon v-text="icon"></v-icon>
                  </v-list-item-icon>
                </v-list-item>
              </v-list-group>
            </v-list-group>
            <v-list-group :value="true" prepend-icon="mdi-view-dashboard-edit">
              <template v-slot:activator>
                <v-list-item-content>
                  <v-list-item-title>Kurse</v-list-item-title>
                </v-list-item-content>
              </template>

              <v-list-item
                v-for="([title, icon], i) in sidebarLeftNavigationLinks.admin
                  .course"
                :key="i"
                link
              >
                <v-list-item-icon text="right">
                  <v-icon v-text="icon"></v-icon>
                  <span class="pl-2">
                    {{ title }}
                  </span>
                </v-list-item-icon>
              </v-list-item>
            </v-list-group>
          </v-list>
        </v-col>-->
        <v-col class="12">
          <v-divider></v-divider>
        </v-col>
        <v-col cols="12">
          <v-list dense>
            <v-list-item>
              <v-list-item-icon>
                <v-icon>mdi-domain</v-icon>
              </v-list-item-icon>

              <v-list-item-title>Über uns</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-col>
      </v-row>
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
  methods:{
    goToRegPage(){
      this.$router.push({ name: "Registry" });
    }
  }
};
</script>

<style scoped></style>
