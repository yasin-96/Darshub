import type { Auth0VueClient } from "@auth0/auth0-vue";
import type { UserRequest } from "../user/types";
export interface LoginDetails {
  //user: UserRequest;
  authClient: Auth0VueClient | null
}
export interface UserDetails {
  user: UserRequest;
}

export interface Registration {
  user: UserRequest;
}
