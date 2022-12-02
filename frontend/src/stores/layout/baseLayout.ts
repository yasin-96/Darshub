import { defineStore } from "pinia";

export const useBaseLayoutStore = defineStore("baseLayoutStore", {
  state: () => ({
    infoPanel: {},
    navBarTop: {
      drawerForSidebarLeft: {
        showPanel: false,
      },
      loginModal: {
        showPanel: false,
      },
    },
    sidebarLeft: {},
    mainContent: {},
    footer: {},
  }),
  actions: {
    act_toggleSidebarLeft(tValue: boolean) {
      this.navBarTop.drawerForSidebarLeft.showPanel = tValue;
    },
    act_toggleLoginWindow(tValue: boolean) {
      this.navBarTop.loginModal.showPanel = tValue;
    },
  },
  getters: {
    currentStateDrawerSidebarLeft(): boolean {
      return this.navBarTop.drawerForSidebarLeft.showPanel;
    },
    getCurrentStateOfLoginWindow(): boolean {
      return this.navBarTop.loginModal.showPanel;
    },
  },
});
