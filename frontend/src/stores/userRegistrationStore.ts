import type { UserRegistrationCheck, UserRequest } from "@/models/user/types";
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
    setFirstName(newFN: string) {
      this.user.first_name = newFN;
    },
    setLastName(newLN: string) {
      this.user.last_name = newLN;
    },
    setBirthday(newBirthday: string) {
      const parsedUnixTime = Date.parse(newBirthday);
      this.user.birthday = parsedUnixTime;
    },

    setAvatar(newAvatar: string | null) {
      this.user.avatar = newAvatar;
    },

    async addNewUser() {
      console.log(`${BACKEND_API}/user`);

      await axios
        .post(`${BACKEND_API}/user`, this.user)
        .then((resp) => {
          console.log(resp);
          this.$reset();
        })
        .catch((err) => {
          // console.log(err)
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
  getters: {
    getAvatar(): string {
      return String(this.user.avatar);
    },
  },
});
