import { createAuth0 } from "@auth0/auth0-vue";

export const authClient = createAuth0({
  domain: import.meta.env.VITE_A0_DOMAIN,
  clientId: import.meta.env.VITE_A0_CLIENT_ID,
  authorizationParams: {
    redirect_uri: `${window.location.origin}/login`,
    scope: "openid profile email roles",
  },
});