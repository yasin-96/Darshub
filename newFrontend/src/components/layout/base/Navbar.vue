<script setup lang="ts">
import { ref } from "vue";
import LocaleSwitcher from "@/components/locale/LocaleSwitcher.vue";
import { useRoute } from "vue-router";
import { useI18n } from "vue-i18n";

const r = useRoute();
const drawer = ref(false);
const local = useI18n();
const { t } = useI18n();
const generalSidebar = ref([
  {
    name: "sidebar.default.index",
    icon: "M2.25 12l8.954-8.955c.44-.439 1.152-.439 1.591 0L21.75 12M4.5 9.75v10.125c0 .621.504 1.125 1.125 1.125H9.75v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21h4.125c.621 0 1.125-.504 1.125-1.125V9.75M8.25 21h8.25",
    disabled: false,
    href: "/",
  },
  {
    name: "",
    icon: "M15 9h3.75M15 12h3.75M15 15h3.75M4.5 19.5h15a2.25 2.25 0 002.25-2.25V6.75A2.25 2.25 0 0019.5 4.5h-15a2.25 2.25 0 00-2.25 2.25v10.5A2.25 2.25 0 004.5 19.5zm6-10.125a1.875 1.875 0 11-3.75 0 1.875 1.875 0 013.75 0zm1.294 6.336a6.721 6.721 0 01-3.17.789 6.721 6.721 0 01-3.168-.789 3.376 3.376 0 016.338 0z",
    disabled: false,
    href: "",
  },
  {
    name: "sidebar.default.books",
    icon: "M12 6.042A8.967 8.967 0 006 3.75c-1.052 0-2.062.18-3 .512v14.25A8.987 8.987 0 016 18c2.305 0 4.408.867 6 2.292m0-14.25a8.966 8.966 0 016-2.292c1.052 0 2.062.18 3 .512v14.25A8.987 8.987 0 0018 18a8.967 8.967 0 00-6 2.292m0-14.25v14.25",
    disabled: false,
    href: "",
  },
  // {
  //   name: "sidebar.default.news",
  //   icon: "mdi-newspaper",
  //   disabled: true,
  //   href: "news",
  // },
  {
    name: "sidebar.default.gallery",
    icon: "  M2.25 15.75l5.159-5.159a2.25 2.25 0 013.182 0l5.159 5.159m-1.5-1.5l1.409-1.409a2.25 2.25 0 013.182 0l2.909 2.909m-18 3.75h16.5a1.5 1.5 0 001.5-1.5V6a1.5 1.5 0 00-1.5-1.5H3.75A1.5 1.5 0 002.25 6v12a1.5 1.5 0 001.5 1.5zm10.5-11.25h.008v.008h-.008V8.25zm.375 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z",
    disabled: false,
    href: "",
  },
]);
const generalFooter = ref([
  // {
  //   name: "sidebar.default.contact",
  //   icon: "M21.75 6.75v10.5a2.25 2.25 0 01-2.25 2.25h-15a2.25 2.25 0 01-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0019.5 4.5h-15a2.25 2.25 0 00-2.25 2.25m19.5 0v.243a2.25 2.25 0 01-1.07 1.916l-7.5 4.615a2.25 2.25 0 01-2.36 0L3.32 8.91a2.25 2.25 0 01-1.07-1.916V6.75",
  //   disabled: false,
  //   href: "/contact",
  // },
  { name: "sidebar.default.impressum", icon: "mdi-contact", href: "/impressum" },
  { name: "sidebar.default.team", icon: "mdi-contact", href: "/team" },
  { name: "sidebar.default.contact", icon: "mdi-contact", href: "/contact" },
]);
const currentSideName = ref(r.name ? t(`sidebar.default.${String(r.name)}`) : "");

// watch: {
//   currentSideName() {
//     this.currentSideName = this.$t(`sidebar.default.${String(this.r.name)}`);
//   },
// },
</script>

<template>
  <nav>
    <div class="drawer">
      <input id="leftSidebarDrawer" v-model="drawer" type="checkbox" class="drawer-toggle" />
      <div class="drawer-content flex flex-col">
        <!-- Navbar -->
        <div class="`w-full navbar bg-base-100 fixed top-0 z-50 pt-2 pb-2`">
          <div class="flex-none">
            <label for="leftSidebarDrawer" class="btn btn-square btn-ghost">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="inline-block w-6 h-6 stroke-current">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
              </svg>
            </label>
          </div>
          <div class="flex-1 px-2 mx-2">{{}}</div>
          <!-- <div class="flex-none hidden lg:block">
          <ul class="menu menu-horizontal">
            <li v-for="(link) in generalSidebar" :key="link.name">

              <a :href="link.href">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                  stroke="currentColor" class="w-6 h-6">
                  <path stroke-linecap="round" stroke-linejoin="round" :d="link.icon" />

                </svg>
                {{$t(link.name)}}
              </a>
            </li>
          </ul>
        </div> -->

          <LocaleSwitcher />
          <router-link to="/registry">
            <button class="btn gap-2">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
              </svg>
              Button
            </button>
          </router-link>
        </div>
        <main class="">
          <slot></slot>
        </main>
      </div>

      <div class="drawer-side">
        <label for="leftSidebarDrawer" class="drawer-overlay"></label>
        <ul class="menu p-4 overflow-y-auto w-80 bg-base-100">
          <li v-for="link in generalSidebar" :key="link.name">
            <router-link :to="link.href" @click="drawer = !drawer">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                <path stroke-linecap="round" stroke-linejoin="round" :d="link.icon" />
              </svg>
              {{ $t(link.name) }}
            </router-link>
          </li>
          <li class="mt-auto"><div class="divider"></div></li>
          <li v-for="link in generalFooter" :key="link.name">
            <router-link v-if="link.href" :to="link.href" @click="drawer = !drawer">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                <path stroke-linecap="round" stroke-linejoin="round" :d="link.icon" />
              </svg>
              {{ $t(link.name) }}
            </router-link>
            <a v-else :href="link.link" target="_blank">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                <path stroke-linecap="round" stroke-linejoin="round" :d="link.icon" />
              </svg>
              {{ $t(link.name) }}
            </a>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>
