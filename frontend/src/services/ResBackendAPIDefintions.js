const API_BASE_ENDPOINTS = {
  backend: {
    user: `${process.env.VUE_APP_BACKEND_URI}/user`,
    session: `${process.env.VUE_APP_BACKEND_URI}/session`,
    course: `${process.env.VUE_APP_BACKEND_URI}/course`,
    courses: `${process.env.VUE_APP_BACKEND_URI}/courses`,
    courseCategory: `${process.env.VUE_APP_BACKEND_URI}/courseCategory`,
    chapter: `${process.env.VUE_APP_BACKEND_URI}/chapter`,
    subchapter: `${process.env.VUE_APP_BACKEND_URI}/subchapter`,
  },
};

const api_endpoints = {
  user: {
    registry: `${API_BASE_ENDPOINTS.backend.user}`,
    login: `${API_BASE_ENDPOINTS.backend.session}`,
  },
  course: {
    getAll: `${API_BASE_ENDPOINTS.backend.courses}`,
    createOne: `${API_BASE_ENDPOINTS.backend.course}/getAll`,
    updateOne: `${API_BASE_ENDPOINTS.backend.course}/getAll`,
    deleteOne: `${API_BASE_ENDPOINTS.backend.course}/getAll`,
  },
};

export default api_endpoints;
