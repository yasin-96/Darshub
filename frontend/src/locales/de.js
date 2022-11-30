export default {
  appName: "",
  navbar: {
    default: {
      lang: "Sprachen",
    },
  },
  sidebar: {
    default: {
      index: "Start",
      gallery: "",
      courses: "",
    },
    admin: {},
    footer: {
      impress: "",
      team: "",
      contact: "",
    },
  },
  footer: {
    c: "Urheberrecht",
    cr: "Alle Rechte vorbehalten.",
  },
  deploymentType: {
    local: "Lokal",
    dev: "Entwicklungsversion!",
    prod: "Productionsversion",
  },
  languages: {
    de: "Deutsch",
    en: "Englisch",
  },
  gender: {
    MALE: "Männlich",
    FEMALE: "Weiblich",
    DIVERS: "",
  },
  sites: {
    registry: {
      title: "",
      steps: {
        first: {
          title: "Basis Daten",
          input: {
            firstName: {
              label: "Vorname",
              placeholder: "Vorname wie Max",
              hints: "",
            },
            lastName: {
              label: "Nachname",
              placeholder: "Nachname wie Müller",
              hints: "",
            },
            birthday: {
              label: "Geburtstag",
              placeholder: "",
              hints: "",
            },
            gender: {
              label: "Geschlecht",
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
    error:{
      notFound:""
    }
  },
};
