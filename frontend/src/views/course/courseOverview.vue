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
  <v-containter>
    <div class="text-center py-10">
      <h2
        class="sm:text-xl md:text-2xl lg:text-4xl font-bold tracking-tight text-gray-900"
      >
        Course (TODO)
      </h2>
    </div>
  </v-containter>
  <v-container>
    <div>
      <v-row>
        <v-col
          cols="12"
          xs="12"
          sm="6"
          md="4"
          lg="3"
          xl="2"
          xxl="2"
          v-for="c in coursesWithMinimalDetails"
        >
          <CourseQuickInfo
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
        </v-col>
      </v-row>
    </div>
  </v-container>
</template>
