import generalCourseModule from "./modules/generalCourseModule";

const courseStore = {
  namespaced: true,
  modules: {
    general: generalCourseModule,
  },
};

export default courseStore;
