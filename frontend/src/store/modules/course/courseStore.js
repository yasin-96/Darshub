import editCourseModule from "./modules/editCourseModule";
import generalCourseModule from "./modules/generalCourseModule";

const courseStore = {
  namespaced: true,
  modules: {
    general: generalCourseModule,
    edit: editCourseModule,
  },
};

export default courseStore;
