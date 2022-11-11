export default {
  appName: "DarsHub",
  router: {
    sidebar:{
      all: {
        home: "Start",
        courseOverview: "Übersicht der Kurse",
      },
      admin:{
        userSettings: "Benutzerverwaltung",
        courseSettings: "Kursverwaltung"
      },
      author:{
        create: "Kurs erstellen",
        dashboard: "Mein Kurse verwalten"
      }
    }
    
  },
  navigation: {
    lang: "Sprache",
  },
  deploymentType: {
    local: "Lokal",
    dev: "Entwicklungsversion!",
    prod: "Productionsversion",
  },
};
