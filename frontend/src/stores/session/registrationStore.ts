import { defineStore } from "pinia";
import { useFetch } from "@vueuse/core";

import type { Registration } from "@/models/user/interfaces";

const BACKEND_API = import.meta.env.VITE_APP_BACKEND_URI;

export const useRegistrationStore = defineStore("userRegistrationStore", {
  state: (): Registration => ({
    user: {
      id: "",
      password: "",
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
      role: new Array<Uint8Array>(),
    },
  }),
  actions: {
    clear() {
      this.$reset();
    },
    setFirstName(newFN: string) {
      this.user.first_name = newFN;
    },
    setLastName(newLN: string) {
      this.user.last_name = newLN;
    },
    setBirthday(newBirthday: Date) {
      // const parsedUnixTime = Date.parse(newBirthday);
      this.user.birthday = newBirthday;
    },

    setAvatar(newAvatar: string | null) {
      this.user.avatar = newAvatar;
    },
    setEmail(newEmailAddr: string) {
      this.user.email = newEmailAddr;
    },
    setMobileNumber(newMobileNumber: string) {
      this.user.telNr = newMobileNumber;
    },

    setPasswd(newPasswd: string) {
      this.user.password = newPasswd;
    },

    async addNewUser() {
      console.log(JSON.stringify(this.user));
      const { isFetching, error, data } = await useFetch(
        `${BACKEND_API}/user`
        // { timeout: BACKEND_API_TIMEOUT }
      ).post(this.user);

      if (error) {
        console.log("err", error.value);
        // this.$reset();
      } else {
        console.table("data", data);
        // this.user = data;
      }

      // await axios
      //   .post(`${BACKEND_API}/user`, this.user)
      //   .then((resp) => {
      //     console.log(resp);
      //     this.$reset();
      //   })
      //   .catch((err) => {
      //     // console.log(err)
      //   });

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
  getters: {
    getAvatar(): string {
      return String(this.user.avatar);
    },
  },
});
