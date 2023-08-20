/**
 * Handler that will be called during the execution of a PostLogin flow.
 * Append user role onto to claim part
 *
 * @param {Event} event - Details about the user and the context in which they are logging in.
 * @param {PostLoginAPI} api - Interface whose methods can be used to change the behavior of the login.
 */
exports.onExecutePostLogin = async (event, api) => {
    const { roles } = event.user.app_metadata;
    
    api.idToken.setCustomClaim("appUID", event.user.user_id);
  
    if (!roles) {
      api.idToken.setCustomClaim("appRoles", [1]);
    } else {
      api.idToken.setCustomClaim("appRoles", roles);
    }
  };
  