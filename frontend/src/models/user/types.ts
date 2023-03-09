// export type Gender = {
//   MALE;
//   FEMALE;
//   DIVERS;
// };

export enum UserRoles {
  USER = 1,
  AUTHOR = 2,
  USER_MANAGER = 3,
  COURSE_MANAGER = 3,
  ADMIN = 99,
}

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
  role: Array<UserRoles>;
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
  role: Array<UserRoles>;
};

export type UserRequest = {
  id: string;
  password: string;
  first_name: string;
  last_name: string;
  birthday: string;
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
  role: Array<UserRoles>;
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