import RestBackendCalls from "@/services/RestBackendCalls";
import api_endpoints from "../../../../services/ResBackendAPIDefintions";

const editCourseModule = {
  namespaced: true,
  state: {
    courseId: null,
  },
  actions: {
    act_setCourseIdForEditing({commit}, data){
        if(data){
            commit("MUT_SET_COURSE_ID", data)
        }
    },

    async act_loadAllCourses({ commit }) {
      RestBackendCalls.doGetRequest(api_endpoints.course.getAll, null)
        .then((resp) => {
          if (resp && resp.data && resp.data.length && resp.status == 200) {
            commit("MUT_SET_ALL_COURSES", resp.data);
          }
        })
        .catch((err) => {
          commit("MUT_SET_ALL_COURSES", []);

          console.log(err);
        });
    },
  },
  mutations: {
    MUT_SET_ALL_COURSES(state, data) {
      state.allCourses = data;
    },
    MUT_SET_COURSE_ID(state, data){
        state.courseId = data
    }
  },
};

export default editCourseModule;
