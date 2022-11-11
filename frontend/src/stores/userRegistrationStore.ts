import type { UserRequest } from "@/types";
import axios from "axios";
import { defineStore } from "pinia";

interface Registration {
  user: UserRequest;
}

const BACKEND_API = import.meta.env.VITE_APP_BACKEND_URL;

export const useUserRegistrationStore = defineStore("userRegistrationStore", {
  state: (): Registration => ({
    user: {
      id: "",
      password: [],
      first_name: "",
      last_name: "",
      birthday: null,
      avatar: "",
      email: "",
      telNr: "",
      company: "",
      occupation: "",
      school: "",
      subject: "",
      country: "",
      isActive: true,
      bio: "",
      role: new Array<Uint8Array>(),
    },
  }),
  actions: {
    clear() {
      this.$reset();
    },
    setBirthday(newBirthday: Date) {
      this.user.birthday = newBirthday;
    },

    async addNewUser() {
      console.log(`${BACKEND_API}/user`)

      // fetch(`${BACKEND_API}/user`, {
      //   method: "POST",
      //   body: JSON.stringify(this.user)
      // }).then((resp) => {
      //   // console.log(resp)
      // })

      await axios
        .post(`${BACKEND_API}/user`, this.user)
        .then((resp) => {
          console.log(resp)
          this.$reset();
        })
        .catch((err) => {
          console.log(err)

        });

      //   RestBackendCalls.doPostRequest(api_endpoints.user.registry, null, state.newUser)
      //     .then((resp) => {
      //       if (resp && resp.data && resp.data.length && resp.status == 201) {
      //         dispatch(
      //           "ui/notification/act_setSuccessNotification",
      //           `Es wurde ein neuer Benutzer erstellt. [${resp.data}]`,
      //           { root: true }
      //         );
      //         commit("MUT_CLEAR_STATE");
      //         router.push({ name: "Rework" });
      //       }
      //     })
      //     .catch((err) => {
      //       console.log(err);
      //     });
      
    },
  },
});
