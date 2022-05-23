<template>
  <div>
    <!-- App-Bar which is shown on top of whole App -->
    <v-app-bar app elevate-on-scroll dense color="white">
      <v-app-bar-nav-icon
        @click="showSidebarLeftDrawer = !showSidebarLeftDrawer"
      ></v-app-bar-nav-icon>
      <v-toolbar-title>
        <div class="title">{{ $t("appName") }}</div>
      </v-toolbar-title>

      <v-spacer></v-spacer>
      <v-btn
        small
        rounded
        color="info"
        outlined
        @click="openSignInDialog"
        v-if="!loggedInUser || !loggedInUser.ID"
      >
        Anmelden
      </v-btn>
      <SignInDialog v-if="!loggedInUser || !loggedInUser.ID" />

      <UserInfo v-if="loggedInUser && loggedInUser.ID" />
      <v-divider
        vertical
        inset
        class="ml-3"
        v-if="loggedInUser && loggedInUser.ID"
      ></v-divider>

      <LocaleSwitcher cssStyle="ml-3" />
      <v-divider
        vertical
        inset
        class="ml-3"
        v-if="loggedInUser && loggedInUser.ID"
      ></v-divider>
      <v-btn
        small
        rounded
        color="info"
        outlined
        @click="logUserOut"
        class="ml-3"
        v-if="loggedInUser && loggedInUser.ID"
      >
        <v-icon>mdi-logout-variant</v-icon>
      </v-btn>
    </v-app-bar>

    <!-- Navigation drawer on left side of UI -->
    <v-navigation-drawer app v-model="showSidebarLeftDrawer" temporary>
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

      <v-divider> </v-divider>
      <div v-if="userHasAuthorRights">
        <v-chip class="ml-3 mt-2" label>AUTHOR</v-chip>
        <v-list dense>
          <v-list-item
            v-for="(aItem, i) in sidebarLeftNavigationLinks.author"
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
      </div>
      <v-divider> </v-divider>
      <div v-if="userHasCourseManamengtRights">
        <v-chip class="ml-3 mt-2" label>COURSE MANAGEMENT</v-chip>
        <v-list dense>
          <v-list-item
            v-for="(aItem, i) in sidebarLeftNavigationLinks.courseManagement"
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
      </div>
      <v-divider> </v-divider>
      <div v-if="userHasUserManamengtRights">
        <v-chip class="ml-3 mt-2" label>USER MANAGEMENT</v-chip>
        <v-list dense>
          <v-list-item
            v-for="(aItem, i) in sidebarLeftNavigationLinks.userManagement"
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
      </div>
      <v-divider> </v-divider>
      <div v-if="userHasAdminRights">
        <v-chip class="ml-3 mt-2" label>ADMIN</v-chip>
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
      </div>

      <template v-slot:append>
        <v-divider></v-divider>
        <v-list-item
          v-for="(aItem, i) in sidebarLeftNavigationLinks.app"
          :key="i"
          :link="aItem.route ? true : false"
          :to="aItem.route"
        >
          <v-list-item-icon>
            <v-icon>{{ aItem.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-title>{{ $t(aItem.title) }}</v-list-item-title>
        </v-list-item>
      </template>
    </v-navigation-drawer>
  </div>
</template>

<script>
import LocaleSwitcher from "@/components/locales/LocaleSwitcher";
import SignInDialog from "@/components/signin_up/SignInDialog";

import { mapGetters, mapState } from "vuex";
import UserInfo from "@/components/user/UserInfo.vue";
import VueRouter from "vue-router";
const { isNavigationFailure, NavigationFailureType } = VueRouter;
export default {
  components: {
    LocaleSwitcher,
    SignInDialog,
    UserInfo,
  },
  data: () => ({}),
  computed: {
    ...mapGetters(["userStore/store", "getFullName"]),
    ...mapState({
      drawerPanelLeft: (state) => state.ui.sidebarLeft.showSidebarLeftDrawer,
      sidebarLeftNavigationLinks: (state) =>
        state.ui.sidebarLeft.sidebarLeftNavigationLinks,
      loggedInUser: (state) => state.userStore.user.info,
    }),
    userHasReadRights() {
      return this.checkUserRights(0);
    },
    userHasAuthorRights() {
      return this.checkUserRights(1);
    },
    userHasCourseManamengtRights() {
      return this.checkUserRights(2);
    },
    userHasUserManamengtRights() {
      return this.checkUserRights(3);
    },
    userHasAdminRights() {
      return this.checkUserRights(99);
    },

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
      this.$router.push({ name: "Registry" }).catch((failure) => {
        if (isNavigationFailure(failure, NavigationFailureType.duplicated)) {
          return;
        }
      });
    },
    goToLoginPage() {
      this.$router.push({ name: "Login" }).catch((failure) => {
        if (isNavigationFailure(failure, NavigationFailureType.duplicated)) {
          return;
        }
      });
    },
    openSignInDialog() {
      this.$store.dispatch("ui/dialogs/act_toggleSignInDialog", true);
    },
    logUserOut() {
      this.$router.push({ name: "Logout" }).catch((failure) => {
        if (isNavigationFailure(failure, NavigationFailureType.duplicated)) {
          return;
        }
      });
    },
    checkUserRights(roleNumber) {
      if (!!this.loggedInUser && !!this.loggedInUser.Role) {
        let check = this.loggedInUser.Role.includes(roleNumber);
        if (check) {
          return true;
        }
      }
      return false;
    },
  },
};
</script>

<style scoped></style>
