<script setup lang="ts">
import { useCoreCourseStore } from "@/stores/course/coreCourseStore";
import CourseQuickInfo from "@/components/course/CourseQuickInfo.vue";

import { computed } from "vue";

const courseStore = useCoreCourseStore();

const coursesWithMinimalDetails = computed(() => {
  return courseStore.allCourseInMinimalForm;
});
</script>

<template>
  <div
    class="mx-auto max-w-2xl py-16 px-4 sm:py-24 sm:px-6 lg:max-w-7xl lg:px-8"
  >
    <div class="text-center py-10">
      <h2
        class="sm:text-xl md:text-2xl lg:text-4xl font-bold tracking-tight text-gray-900"
      >
        Course (TODO)
      </h2>
    </div>
    <div
      class="grid sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 justify-between gap-y-10 gap-x-6 sm:flex-row"
    >
      <CourseQuickInfo
        v-for="c in coursesWithMinimalDetails"
        :key="`course-min-det-${c.id}`"
        :id="c.id"
        :name="c.name"
        :description="c.description"
        :duration="c.duration"
        :level="c.level"
        :author="c.author"
        :released="c.released"
        :lastupdate="c.lastupdate"
      />
    </div>
  </div>
</template>
