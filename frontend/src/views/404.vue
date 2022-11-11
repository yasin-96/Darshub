<template>
  <div class="hero min-h-screen" >
    <div class="hero-overlay bg-opacity-60"></div>
    <div class="hero-content text-center text-neutral-content">
      <div class="">
        <h1 class="mb-5 text-5xl md:text-9xl lg:text-9xl font-bold">404</h1>
        <h2 class="mb-5 text-2xl md:text-4xl lg:text-4xl">{{ $t("pages.error.notFound") }}</h2>
        <router-link to="/books">
          <button class="btn btn-info">
            {{ $t("pages.index.allBooks") }}
          </button>
        </router-link>
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
