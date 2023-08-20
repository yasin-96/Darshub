/**
 * Handler that will be called during the execution of a PreUserRegistration flow.
 * Extra: Add default user roles with read only mode to appMetaData
 *
 * @param {Event} event - Details about the context and user that is attempting to register.
 * @param {PreUserRegistrationAPI} api - Interface whose methods can be used to change the behavior of the signup.
 */
exports.onExecutePreUserRegistration = async (event, api) => {
  const roles = [0];
  api.user.setAppMetadata("roles", roles);
};
