export type Course = {
  id: String;
  name: String;
  description: String;
  duration: Date | null;
  level: String;
  content: Array<String>;
  author: String;
  released: Date | null;
  lastupdate: Date | null;
};

export type CreateCourseRequest = {
  id: String;
  name: String;
  description: String;
  duration: Date;
  level: String;
  content: Array<String>;
  author: String;
  released: Date;
  lastupdate: Date;
};

export type UpdateCourseRequest = {
  name: String;
  description: String;
  duration: Date;
  level: String;
  content: Array<String>;
  author: String;
  released: Date;
  lastupdate: Date;
};
