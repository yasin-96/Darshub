export default {
  appName: "",
  navbar: {
    default: {
      lang: "Languages",
    },
  },
  sidebar: {
    default: {
      index: "Home",
      courses: "Courses",
    },
    admin: {
      sidebarMenuTitel: "Admin",
      userManagement: "User Management",
      courseManagement: "Course Management",
    },
    user: {
      sidebarMenuTitel: "User",
      course: "Course",
      dashboard: "Dashboard",
      settings: "Settings",
    },
    footer: {
      // impress: "Impress",
      team: "Team",
      contact: "Contact",
    },
  },
  footer: {
    c: "Urheberrecht",
    cr: "Alle Rechte vorbehalten.",
  },
  deploymentType: {
    local: "Local",
    dev: "DEV - VERSION",
    prod: "PRODUCTION - VERSION",
  },
  languages: {
    de: "Germany",
    en: "English",
  },
  gender: {
    MALE: "Man",
    FEMALE: "Woman",
    DIVERS: "undefined",
  },
  sites: {
    registry: {
      title: "",
      steps: {
        first: {
          title: "Basis Daten",
          input: {
            firstName: {
              label: "First Name",
              placeholder: "",
              hints: "",
            },
            lastName: {
              label: "Last Name",
              placeholder: "",
              hints: "",
            },
            birthday: {
              label: "Birthday",
              placeholder: "",
              hints: "",
            },
            gender: {
              label: "Gender",
              placeholder: "",
              hints: "",
            },
          },
        },
        second: {
          title: "Basis Daten",
          input: {
            email: {
              label: "",
              placeholder: "",
              hints: "",
            },
            mobileNumber: {
              label: "",
              placeholder: "",
              hints: "",
            },
          },
        },
        third: {
          title: "Basis Daten",
          input: {
            password: {
              label: "",
              placeholder: "",
              hints: "",
            },
            checkPassword: {
              label: "",
              placeholder: "",
              hints: "",
            },
          },
        },
      },
    },
    admin: {
      courseManagement: {
        table: {
          header: {
            id: "ID",
            name: "Name",
            description: "Description",
            duration: "Duration",
            level: "Level",
            content: "Content",
            author: "Author",
            released: "Released",
            lastupdate: "Lastupdate",
            action: "Action",

          },
        },
      },
    },
  },
};
