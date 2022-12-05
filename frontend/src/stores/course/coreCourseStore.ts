import { defineStore } from "pinia";
import { useFetch } from "@vueuse/core";

import type { CoreCourseState } from "@/models/course/interfaces";
import type { Course } from "@/models/course/types";

const BACKEND_API = import.meta.env.VITE_APP_BACKEND_URI;

export const useCoreCourseStore = defineStore("coreCourseStore", {
  state: (): CoreCourseState => ({
    courses: [] as Array<Course>,
    coursePreview: {} as Course,
  }),
  actions: {
    async act_loadAllCourseAsQuickInfo() {
      const { isFetching, error, data } = await useFetch<Array<Course>>(
        `${BACKEND_API}/courses`
        // { timeout: BACKEND_API_TIMEOUT }
      ).json();

      if (error.value) {
        //TODO: Message popup
        console.log("err", error.value);
        this.$reset();
        return;
      }

      if (data.value) {
        console.log(data.value);
        this.courses = data.value;
      }
    },
    async act_loadCourseById(courseId: string) {
      const { isFetching, error, data } = await useFetch<Array<Course>>(
        `${BACKEND_API}/course/${courseId}`
        // { timeout: BACKEND_API_TIMEOUT }
      ).json();

      if (error.value) {
        //TODO: Message popup
        console.log("err", error.value);
        this.$reset();
        return;
      }

      if (data.value) {
        console.log(data.value);
        this.coursePreview = data.value;
      }
    },
  },
  getters: {
    allCourseInMinimalForm(): Array<Course> {
      return this.courses;
    },
    getCoursePreview(): Course {
      return this.coursePreview;
    },
  },
});
