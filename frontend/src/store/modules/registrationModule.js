// ROLE ={
//     1 = USER,
//     2 = Author,
//     3 = ADMIN,
//     ...
// }

// in userModel = [1,2,3]
// in userModel = [3]
// in userModel = [1]


const registrationModule = {
  state: () => ({
    newUser: {
      firstName: "",
      lastName: "",
      birthday: "", //date
      profileImage: "",
      email: "",
      telNr: "", //need to discuss for that the need the tel number
      occupation: "",
      company: "",
      school: "",
      subject: "", //stand for that?
      country: "",
      bio:"",
      role: [], //define roles as numbers (enum)
    },
  }),
  actions: {
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

    // act_createNewUser: function({commit, state}){
        
    // }
  },
  mutations: {
    MUT_SET_FIRSTNAME: function (state, newFirstName) {
      state.newUser.firstName = newFirstName;
    },
    MUT_SET_LASTNAME: function (state, newLastname) {
      state.newUser.lastname = newLastname;
    },
    MUT_SET_BIRTHDAY: function (state, newBirthday) {
      state.newUser.birthday = newBirthday;
    },
    MUT_SET_PROFILEIMAGE: function (state, newProfileimage) {
      state.newUser.profileimage = newProfileimage;
    },
    MUT_SET_EMAIL: function (state, newEmail) {
      state.newUser.email = newEmail;
    },
    MUT_SET_TELNR: function (state, newTelnr) {
      state.newUser.telnr = newTelnr;
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
  },
  getters: {},
};

export default registrationModule;
