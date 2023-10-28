import { createAuth0 } from "@auth0/auth0-vue";

export const authClient = createAuth0({
  domain: import.meta.env.VITE_A0_DOMAIN,
  clientId: import.meta.env.VITE_A0_CLIENT_ID,
  authorizationParams: {
    redirect_uri: `${window.location.origin}`,
    scope: "openid profile email roles meta",
  },
  useRefreshTokens: true,
  legacySameSiteCookie: true,
  authorizeTimeoutInSeconds: 60000,
});

export const userClient = createAuth0({
  domain: import.meta.env.VITE_A0_DOMAIN,
  clientId: import.meta.env.VITE_A0_USER_API_CLIENT_ID,
  authorizationParams: {
    redirect_uri: `${window.location.origin}`,
    scope: "openid profile email roles meta",
    audience: import.meta.env.VITE_A0_USER_API_CLIENT_AUD,
  },
  useRefreshTokens: true,
  authorizeTimeoutInSeconds: 60000,
});

export const accountManagementClient = createAuth0({
  domain: import.meta.env.VITE_A0_DOMAIN,
  clientId: import.meta.env.VITE_A0_ACCOUNT_MANAGER_API_CLIENT_ID,
  authorizationParams: {
    redirect_uri: `${window.location.origin}`,
    scope: "openid profile email roles meta",
    audience: import.meta.env.VITE_A0_ACCOUNT_MANAGER_API_CLIENT_AUD,
  },
  useRefreshTokens: true,
  authorizeTimeoutInSeconds: 60000,
});
