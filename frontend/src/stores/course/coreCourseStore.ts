import { defineStore } from "pinia";
import { useFetch } from "@vueuse/core";

import type { CourseQuickInfos } from "@/models/course/interfaces";
import type { Course } from "@/models/course/types";

const BACKEND_API = import.meta.env.VITE_APP_BACKEND_URI;

export const useCoreCourseStore = defineStore("coreCourseStore", {
  state: (): CourseQuickInfos => ({
    // courses: [] as Array<Course>,

    courses: [
      {
        id: "628606ec41653a2a46afda61",
        name: "testkurs",
        description: "this is a course",
        duration: new Date("0001-01-01T00:00:00.000Z"),
        level: "hard",
        content: [],
        author: "sssdadsa",
        released: new Date("0001-01-01T00:00:00.000Z"),
        lastupdate: new Date("0001-01-01T00:00:00.000Z"),
      },

      {
        id: "628d61e989a417e1e48c3214",
        name: "testkurs",
        description: "this is a course",
        duration: new Date("0001-01-01T00:00:00.000Z"),
        level: "hard",
        content: [],
        author: "sssdadsa",
        released: new Date("2022-05-24T22:53:29.567Z"),
        lastupdate: new Date("2022-05-24T22:53:29.567Z"),
      },

      {
        id: "628606f341653a2a46afda64",
        // _id: "628606f341653a2a46afda63",
        name: "lex",
        description: "ist sehr besonders",
        duration: new Date("0001-01-01T00:00:00.000Z"),
        level: "hard",
        content: [],
        author: "sssdadsa",
        released: new Date("0001-01-01T00:00:00.000Z"),
        lastupdate: new Date("0001-01-01T00:00:00.000Z"),
      },

      {
        id: "628cf1af9b55bc5ec7cba02e",
        name: "new course",
        description: "",
        duration: new Date("0001-01-01T00:00:00.000Z"),
        level: "",
        content: [],
        author: "",
        released: new Date("0001-01-01T00:00:00.000Z"),
        lastupdate: new Date("0001-01-01T00:00:00.000Z"),
      },
      {
        id: "628d62ce89a417e1e48c3216",
        name: "testkurso",
        description: "this is a course",
        duration: new Date("0001-01-01T00:00:00.000Z"),
        level: "",
        content: [],
        author: "",
        released: new Date("2022-05-24T22:57:18.407Z"),
        lastupdate: new Date("2022-05-24T22:57:18.407Z"),
      },
    ],
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
  },
  getters: {
    allCourseInMinimalForm(): Array<Course> {
      return this.courses;
    },
  },
});
