export type Auth0UserDailyState = {
  date: String;
  created_at: String;
  leaked_passwords: number;
  leaked_passwords_signup: number;
  logins: number;
  signups: number;
  updated_at: String;
};

export type Auth0UserStats = {
 activeUser: Number,
 dailyState: Auth0UserDailyState
}

export type Auth0TokenDetails = {
    access_token:String;
    token_type: String;
    expires_in: number;
}

export type TokenResponseDetails = {
  tokenDetails: Auth0TokenDetails,
  lastCreation: number,
}