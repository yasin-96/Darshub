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
      courses: "Kurse",
    },
    admin: {
      sidebarMenuTitel: "Admin",
      userManagement: "Benutzerverwaltung",
      courseManagement: "Kursverwaltung",
      authorManagement: "Autorenverwaltung",
    },
    user: {
      sidebarMenuTitel: "Benutzer",
      course: "Meine Kurse",
      dashboard: "Übersicht",
      settings: "Einstellungen",
    },
    footer: {
      // impress: "Impressum",
      team: "Team",
      contact: "Kontakt",
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
    admin: {
      courseManagement: {
        table: {
          header: {
            id: "ID",
            name: "Titel",
            description: "Beschreibung",
            duration: "Dauer",
            level: "Schwierigkeitsgrad",
            content: "Inhalt",
            author: "Autor",
            released: "Erstellt",
            lastupdate: "Verändert",
            action: "Aktion",
          },
        },
      },
    },
    error: {
      notFound: "",
    },
  },
};
