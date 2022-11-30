import { defineStore } from "pinia";

export const useBaseLayoutStore = defineStore("baseLayoutStore", {
  state: () => ({
    infoPanel: {},
    navBarTop: {
      drawerForSidebarLeft: {
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
  },
  getters: {
    currentStateDrawerSidebarLeft(): boolean {
      return this.navBarTop.drawerForSidebarLeft.showPanel;
    },
  },
});
