import RestBackendCalls from "../../services/RestBackendCalls";

import router from "@/router/index";
import api_endpoints from "../../services/ResBackendAPIDefintions";

const registrationModule = {
  namespaced: true,
  state: () => ({
    newUser: {
      id: "",
      password: "",
      first_name: "",
      last_name: "",
      name: "",
      birthday: "", //date
      avatar: "",
      email: "",
      telNr: "",
      company: "",
      occupation: "",
      school: "",
      subject: "",
      country: null,
      bio: "",
      role: [0], //define roles as numbers (enum)
      isActive: true,
    },
  }),
  actions: {
    act_setPassword: function ({ commit }, newPassword) {
      commit("MUT_SET_PW", newPassword);
    },
    act_setFirstName: function ({ commit }, newFirstName) {
      commit("MUT_SET_FIRSTNAME", newFirstName);
    },
    act_setLastName: function ({ commit }, newLastName) {
      commit("MUT_SET_LASTNAME", newLastName);
    },
    act_setBirthday: function ({ commit }, newBirthday) {
      commit("MUT_SET_BIRTHDAY", newBirthday);
    },
    act_setProfileImage: function ({ commit }, newProfileImage) {
      commit("MUT_SET_PROFILEIMAGE", newProfileImage);
    },
    act_setEmail: function ({ commit }, newEmail) {
      commit("MUT_SET_EMAIL", newEmail);
    },
    act_setTelNr: function ({ commit }, newTelNr) {
      commit("MUT_SET_TELNR", newTelNr);
    },
    act_setOccupation: function ({ commit }, newOccupation) {
      commit("MUT_SET_OCCUPATION", newOccupation);
    },
    act_setCompany: function ({ commit }, newCompany) {
      commit("MUT_SET_COMPANY", newCompany);
    },
    act_setSchool: function ({ commit }, newSchool) {
      commit("MUT_SET_SCHOOL", newSchool);
    },
    act_setSubject: function ({ commit }, newSubject) {
      commit("MUT_SET_SUBJECT", newSubject);
    },
    act_setCountry: function ({ commit }, newCountry) {
      commit("MUT_SET_COUNTRY", newCountry);
    },
    act_setBio: function ({ commit }, newBio) {
      commit("MUT_SET_BIO", newBio);
    },
    act_setRole: function ({ commit }, newRole) {
      commit("MUT_SET_ROLE", newRole);
    },
    act_createNewUser: function ({ commit, dispatch, state }) {
      RestBackendCalls.doPostRequest(api_endpoints.user.registry, null, state.newUser)
        .then((resp) => {
          if (resp && resp.data && resp.data.length && resp.status == 201) {
            dispatch(
              "ui/notification/act_setSuccessNotification",
              `Es wurde ein neuer Benutzer erstellt. [${resp.data}]`,
              { root: true }
            );
            commit("MUT_CLEAR_STATE");
            router.push({ name: "Rework" });
          }
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
  mutations: {
    MUT_SET_PW: function (state, newPassword) {
      state.newUser.password = newPassword;
    },
    MUT_SET_FIRSTNAME: function (state, newFirstName) {
      state.newUser.first_name = newFirstName;
    },
    MUT_SET_LASTNAME: function (state, newLastname) {
      state.newUser.last_name = newLastname;
    },
    MUT_SET_BIRTHDAY: function (state, newBirthday) {
      state.newUser.birthday = newBirthday;
    },
    MUT_SET_PROFILEIMAGE: function (state, newProfileimage) {
      state.newUser.avatar = newProfileimage;
    },
    MUT_SET_EMAIL: function (state, newEmail) {
      state.newUser.email = newEmail;
    },
    MUT_SET_TELNR: function (state, newTelnr) {
      state.newUser.telNr = newTelnr;
    },
    MUT_SET_OCCUPATION: function (state, newOccupation) {
      state.newUser.occupation = newOccupation;
    },
    MUT_SET_COMPANY: function (state, newCompany) {
      state.newUser.company = newCompany;
    },
    MUT_SET_SCHOOL: function (state, newSchool) {
      state.newUser.school = newSchool;
    },
    MUT_SET_SUBJECT: function (state, newSubject) {
      state.newUser.subject = newSubject;
    },
    MUT_SET_COUNTRY: function (state, newCountry) {
      state.newUser.country = newCountry;
    },
    MUT_SET_BIO: function (state, newBio) {
      state.newUser.bio = newBio;
    },
    MUT_SET_ROLE: function (state, newRole) {
      state.newUser.role.add(newRole);
    },
    MUT_CLEAR_STATE(state) {
      state.newUser = null;
      state.newUser = {
        id: "",
        password: "",
        first_name: "",
        last_name: "",
        name: "",
        birthday: "", //date
        avatar: "",
        email: "",
        telNr: "",
        company: "",
        occupation: "",
        school: "",
        subject: "",
        country: null,
        bio: "",
        role: [0], //define roles as numbers (enum)
        isActive: true,
      };
    },
  },
  getters: {},
};

export default registrationModule;