export type Course = {
  id: String;
  name: String;
  description: String | null;
  duration: Date | null;
  level: Number | null;
  content: Array<String> | null;
  author: String | null;
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

export type Chapter = {
  id: String;
  name: String;
  description: String;
  skills: String;
  subchapters: Array<String>;
};

export type SubChapter = {
  id: String;
  content: String;
  listing: String;
};
