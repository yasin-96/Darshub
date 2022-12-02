// export type Gender = {
//   MALE;
//   FEMALE;
//   DIVERS;
// };

export enum GenderOption {
  MALE = 0,
  FEMALE,
  DIVERS,
}

export enum GenderOptionRO {
  MALE = "MALE",
  FEMALE = "FEMALE",
  DIVERS = "DIVERS",
}

export type UserRegistrationCheck = {
  first_name: string;
  last_name: string;
  birthday: number | null;
  avatar: string;
  email: string;
  telNr: string;
  bio: string;
};

export type UserLoginData = {
  password: string;
  email: string;
};

export type User = {
  id: string;
  password: Array<string>;
  first_name: string;
  last_name: string;
  birthday: Date | null;
  avatar: string;
  email: string;
  telNr: string;
  company: string;
  occupation: string;
  school: string;
  subject: string;
  country: string;
  isActive: boolean;
  bio: string;
  role: Array<Uint8Array>;
};

export type UserInfo = {
  id: string;
  first_name: string;
  last_name: string;
  birthday: Date | null;
  avatar: string;
  email: string;
  telNr: string;
  company: string;
  occupation: string;
  school: string;
  subject: string;
  country: string;
  isActive: boolean;
  bio: string;
  role: Array<Uint8Array>;
};

export type UserRequest = {
  id: string;
  password: string;
  first_name: string;
  last_name: string;
  birthday: Date | null;
  avatar: string | null;
  email: string;
  telNr: string;
  company: string;
  occupation: string;
  school: string;
  subject: string;
  country: string;
  isActive: boolean;
  bio: string;
  role: Array<Uint8Array>;
};

export type UpdateUserRequest = {
  password: Array<string>;
  first_name: string;
  last_name: string;
  birthday: string;
  avatar: string;
  email: string;
  telNr: string;
  company: string;
  occupation: string;
  school: string;
  subject: string;
  country: string;
};
