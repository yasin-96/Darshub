<template>
  <div class="relative px-6 lg:px-8">
    <div class="mx-auto max-w-3xl pt-20 pb-32 sm:pt-48 sm:pb-40">
      <div>
        <h1 class="mb-5 text-5xl md:text-9xl lg:text-9xl font-bold">404</h1>
        <div class="hidden sm:mb-8 sm:flex sm:justify-center">
          <div
            class="relative overflow-hidden rounded-lg py-3.5 px-4 text-sm leading-6 ring-1 ring-gray-900/10 hover:ring-gray-900/20"
          >
            <div class="">
              <h2 class="mb-5 text-2xl md:text-4xl lg:text-4xl">TODO</h2>
              <router-link to="/courses">
                <button
                  class="w-full btn-info inline-block bg-green-500 px-3 py-3 rounded-lg text-white"
                >
                  {{ $t("sidebar.default.index") }}
                </button>
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useI18n } from "vue-i18n";
import { useMeta } from "vue-meta";

export default defineComponent({
  name: "NotFoundView",
  setup() {
    const { t } = useI18n();
    useMeta({
      title: "404 Error",
      meta: [{ name: "description", content: "" }],
    });
    return {
      t,
    };
  },
});
</script>
