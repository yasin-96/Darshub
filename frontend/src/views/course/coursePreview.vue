<script lang="ts" setup>
import { useCoreCourseStore } from "@/stores/course/coreCourseStore";
import { computed } from "vue";

const coreCourseStore = useCoreCourseStore();

const coursePreview = computed(() => {
  return coreCourseStore.getCoursePreview;
});
</script>

<template>
  <div class="hero min-h-screen bg-base-200">
    <div class="hero-content flex-col lg:flex-row">
      <img
        src="/images/stock/photo-1635805737707-575885ab0820.jpg"
        class="max-w-sm rounded-lg shadow-2xl"
      />
      <div>
        <h1 class="text-5xl font-bold">
           {{ coursePreview.name }}
        </h1>
        <p class="py-6">
          {{ coursePreview.description }}
        </p>
      </div>
    </div>
  </div>

  <div class="py-10">
    <div class="mx-auto max-w-2xl px-4 pt-10 pb-16 flex flex-col">
      <div class="">
        <h1 class="text-2xl font-bold text-gray-900 sm:text-3xl">
         
        </h1>
      </div>

      <div class="mt-4">
        <h2 class="">Course information</h2>

        <div class="mt-6">
          <h3 class="">Reviews</h3>
          <div class="flex items-center">
            <div class="flex items-center">
              <i
                v-for="i in coursePreview.level"
                :key="`course-${coursePreview.id}-raing-${i}`"
                class="i-heroicons-star py-3 px-3 block w-full text-black"
              ></i>
            </div>
            <p class="pl-2 text-right uppercase">{{ coursePreview.level }}</p>
            <a
              href="#"
              class="ml-3 text-sm font-medium text-indigo-600 hover:text-indigo-500"
              >117 reviews (TODO)</a
            >
          </div>
        </div>
      </div>

      <div class="py-10">
        <div>
          <h3 class="sr-only">Description</h3>

          <div class="space-y-6">
            <p class="text-base text-gray-900">
              {{ coursePreview.description }}
            </p>
          </div>
        </div>
      </div>

      <pre>{{ coursePreview }}</pre>
    </div>
  </div>
</template>
