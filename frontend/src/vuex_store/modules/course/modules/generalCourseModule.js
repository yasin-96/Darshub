import RestBackendCalls from "@/services/RestBackendCalls";
import api_endpoints from "../../../../services/ResBackendAPIDefintions";

const generalCourseModule = {
  namespaced: true,
  state: {
    allCourses: [],
  },
  actions: {
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
  },
};

export default generalCourseModule;
