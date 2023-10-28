<script setup lang="ts">
import { useLoginStore } from "@/stores/session/loginStore";
import { computed, ref } from "vue";
import { useNow, useDateFormat } from "@vueuse/core";

const loginStore = useLoginStore();

const userInfo = computed(() => {
  return loginStore.getUser;
});

const userId = computed(() => loginStore.getUserId);

const isAdmin = computed(() => {
  return loginStore.isUserAdmin;
});

const isAuthor = computed(() => {
  return loginStore.isUserAuthor;
});

const isUserCourseManager = computed(() => {
  return loginStore.isUserCourseManager;
});

const isAccountManager = computed(() => {
  return loginStore.IsUserAccountManager;
});

const formattedDate = (dateToConvert: string) => {
  return useDateFormat(useNow(), dateToConvert);
};

const isGoogleLoginEnable = computed(() => false);
const isFacebookLoginEnable = computed(() => false);
const isAppleLoginEnable = computed(() => false);
const isGithubLoginEnable = computed(() => false);

const tabNavigation = ref(null);

const editUserGeneralModal = ref(false);
const newUserGeneralData = ref({
  nickName: userInfo.value.nickname,
  name: userInfo.value.name,
});
</script>

<template>
  <v-row>
    <v-col cols="12" md="3">
      <v-card>
        <v-tabs v-model="tabNavigation" direction="vertical" color="primary">
          <v-tab value="details">
            <v-icon start> mdi-badge-account</v-icon>
            Details
          </v-tab>
          <v-tab value="roles">
            <v-icon start> mdi-archive-cog </v-icon>
            Roles
          </v-tab>
          <v-tab value="meta">
            <v-icon start> mdi-application-cog-outline </v-icon>
            META
          </v-tab>
          <v-tab value="logs">
            <v-icon start>mdi-protocol</v-icon>
            Logs
          </v-tab>
          <v-tab value="devices">
            <v-icon start>mdi-tablet-cellphone</v-icon>
            Devices
          </v-tab>
          <v-tab value="account">
            <v-icon start>mdi-cog</v-icon>
            Account
          </v-tab>
        </v-tabs>
      </v-card>
    </v-col>

    <v-col cols="12" md="9">
      <v-window v-model="tabNavigation">
        <v-window-item value="details">
          <v-container class="pt-0">
            <v-card elevation="0">
              <v-card-title>
                <v-containter align="center">
                  <span class="text-h4">General</span>
                </v-containter>
              </v-card-title>
              <v-card-text>
                <v-card>
                  <v-card-text>
                    <v-row>
                      <v-col cols="12" sm="2" md="1">
                        <v-avatar class="ml-2 mr-2" size="large">
                          <v-img :src="userInfo.picture" alt="John"></v-img>
                        </v-avatar>
                      </v-col>
                      <v-col cols="12" sm="10" md="11">
                        <v-text-field
                          label="ID"
                          :model-value="`${userId}`"
                          variant="outlined"
                          readonly
                        >
                          <template v-slot:append-inner>
                            <v-icon
                              v-if="userInfo.appIsEmailVerfiyed"
                              icon="mdi-check-decagram"
                              color="success"
                              size="large"
                            />
                            <v-icon
                              v-else
                              icon="mdi-alert-decagram"
                              color="warning"
                              size="large"
                            />
                          </template>
                        </v-text-field>
                      </v-col>
                      <v-col cols="12" md="6">
                        <v-text-field
                          label="Nickname"
                          :model-value="`${userInfo.nickname}`"
                          variant="outlined"
                          disabled
                          readonly
                        />
                      </v-col>
                      <v-col cols="12" md="6">
                        <v-text-field
                          label="Name"
                          :model-value="`${userInfo.name}`"
                          variant="outlined"
                          disabled
                          readonly
                        />
                      </v-col>
                    </v-row>
                  </v-card-text>
                  <v-card-action>
                    <v-dialog
                      v-model="editUserGeneralModal"
                      persistent
                      width="auto"
                    >
                      <template v-slot:activator="{ props }">
                        <v-btn
                          block
                          color="warning"
                          variant="outlined"
                          v-bind="props"
                          prepend-icon="mdi-pencil"
                        >
                          Edit
                        </v-btn>
                      </template>
                      <v-card>
                        <v-card-title class="text-h4 ma-5">
                          Change Name
                        </v-card-title>
                        <v-card-text>
                          <v-row>
                            <v-col cols="12">
                              <v-text-field
                                label="Nickname"
                                :model-value="`${newUserGeneralData.nickName}`"
                                variant="outlined"
                              />
                            </v-col>
                            <v-col cols="12">
                              <v-text-field
                                label="Name"
                                :model-value="`${newUserGeneralData.name}`"
                                variant="outlined"
                              />
                            </v-col>
                          </v-row>
                        </v-card-text>
                        <v-card-actions>
                          <v-spacer></v-spacer>
                          <v-btn
                            color="red-darken-1"
                            variant="outlined"
                            @click="editUserGeneralModal = false"
                          >
                            Abbord
                          </v-btn>
                          <v-btn
                            color="green-darken-1"
                            variant="text"
                            @click="editUserGeneralModal = false"
                          >
                            Change
                          </v-btn>
                        </v-card-actions>
                      </v-card>
                    </v-dialog>
                  </v-card-action>
                </v-card>
              </v-card-text>
            </v-card>
          </v-container>
          <v-container>
            <v-card elevation="0">
              <v-card-title>
                <v-container>
                  <span class="text-h4"> Contact </span>
                </v-container>
              </v-card-title>
              <v-card-text>
                <v-card>
                  <v-card-title>
                    <v-container>
                      <span class="text-h5">Accessibility</span>
                    </v-container>
                  </v-card-title>
                  <v-card-text>
                    <v-row>
                      <v-col cols="12" md="6">
                        <!-- <v-text-field label="Birthday" variant="outlined" /> -->
                        <v-text-field
                          label="Email"
                          :model-value="userInfo.email"
                          variant="outlined"
                          readonly
                          disabled
                          prepend-inner-icon="mdi-email"
                        />
                      </v-col>
                      <v-col cols="12" md="6">
                        <v-text-field
                          label="Mobile"
                          :model-value="userInfo.telNr"
                          variant="outlined"
                          readonly
                          disabled
                          prepend-inner-icon="mdi-phone"
                        />
                      </v-col>
                    </v-row>
                  </v-card-text>
                </v-card>
              </v-card-text>
            </v-card>
          </v-container>
          <v-container>
            <v-card elevation="0">
              <v-card-title>
                <v-container>
                  <span class="text-h4"> Security </span>
                </v-container>
              </v-card-title>
              <v-card-text>
                <v-row>
                  <v-col>
                    <v-card elevation="1">
                      <v-card-title>
                        <v-container>
                          <span class="text-h5">Social options</span>
                        </v-container>
                      </v-card-title>
                      <v-card-text>
                        <v-list lines="one">
                          <v-list-item>
                            <template v-slot:prepend>
                              <v-icon large> mdi-lock-check</v-icon>
                            </template>
                            <template v-slot:append>
                              <v-icon large color="success">
                                mdi-checkbox-marked-circle</v-icon
                              >
                            </template>
                            <v-list-item-title>
                              Auth0-Default-Login
                            </v-list-item-title>
                          </v-list-item>
                          <v-list-item>
                            <template v-slot:prepend>
                              <v-icon large> mdi-google</v-icon>
                            </template>
                            <template v-slot:append>
                              <v-icon
                                v-if="isGoogleLoginEnable"
                                large
                                color="success"
                              >
                                mdi-checkbox-marked-circle</v-icon
                              >
                              <v-icon v-else large color="error">
                                mdi-checkbox-blank-circle</v-icon
                              >
                            </template>
                            <v-list-item-title>
                              Google-Login
                            </v-list-item-title>
                          </v-list-item>
                          <v-list-item>
                            <template v-slot:prepend>
                              <v-icon large> mdi-facebook</v-icon>
                            </template>
                            <template v-slot:append>
                              <v-icon
                                v-if="isFacebookLoginEnable"
                                large
                                color="success"
                              >
                                mdi-checkbox-marked-circle</v-icon
                              >
                              <v-icon v-else large color="error">
                                mdi-checkbox-blank-circle</v-icon
                              >
                            </template>
                            <v-list-item-title>
                              Facebook-Login
                            </v-list-item-title>
                          </v-list-item>
                          <v-list-item>
                            <template v-slot:prepend>
                              <v-icon large> mdi-apple</v-icon>
                            </template>
                            <template v-slot:append>
                              <v-icon
                                v-if="isAppleLoginEnable"
                                large
                                color="success"
                              >
                                mdi-checkbox-marked-circle</v-icon
                              >
                              <v-icon v-else large color="error">
                                mdi-checkbox-blank-circle</v-icon
                              >
                            </template>
                            <v-list-item-title> Apple-Login </v-list-item-title>
                          </v-list-item>
                          <v-list-item>
                            <template v-slot:prepend>
                              <v-icon large> mdi-github</v-icon>
                            </template>
                            <template v-slot:append>
                              <v-icon
                                v-if="isGithubLoginEnable"
                                large
                                color="success"
                              >
                                mdi-checkbox-marked-circle</v-icon
                              >
                              <v-icon v-else large color="error">
                                mdi-checkbox-blank-circle</v-icon
                              >
                            </template>
                            <v-list-item-title>
                              Github-Login
                            </v-list-item-title>
                          </v-list-item>
                        </v-list>
                      </v-card-text>
                    </v-card>
                  </v-col>
                </v-row>
              </v-card-text>
            </v-card>
          </v-container>
          <v-container>
            <v-card elevation="0">
              <v-card-title>
                <v-container>
                  <span class="text-h4"> Roles </span>
                </v-container>
              </v-card-title>
              <v-card-text>
                <v-row>
                  <v-col>
                    <v-card>
                      <v-card-text>
                        <v-list>
                          <v-list-item color="#42A5F5">
                            <template v-slot:prepend>
                              <v-icon large> mdi-account-check</v-icon>
                            </template>
                            <template v-slot:append>
                              <v-icon large color="success">
                                mdi-checkbox-marked-circle</v-icon
                              >
                            </template>
                            <v-list-item-title> Basic </v-list-item-title>
                          </v-list-item>

                          <v-list-item color="#5C6BC0">
                            <template v-slot:prepend>
                              <v-icon large> mdi-book-clock</v-icon>
                            </template>
                            <template v-slot:append>
                              <v-icon v-if="isAuthor" large color="success">
                                mdi-checkbox-marked-circle</v-icon
                              >
                              <v-icon v-else large color="error">
                                mdi-checkbox-blank-circle</v-icon
                              >
                            </template>
                            <v-list-item-title> Autor </v-list-item-title>
                          </v-list-item>

                          <v-list-item color="#FF8F00">
                            <template v-slot:prepend>
                              <v-icon large> mdi-book-multiple</v-icon>
                            </template>
                            <template v-slot:append>
                              <v-icon
                                v-if="isUserCourseManager"
                                large
                                color="success"
                              >
                                mdi-checkbox-marked-circle</v-icon
                              >
                              <v-icon v-else large color="error">
                                mdi-checkbox-blank-circle</v-icon
                              >
                            </template>
                            <v-list-item-title>
                              Course Manager
                            </v-list-item-title>
                          </v-list-item>

                          <v-list-item color="#D84315">
                            <template v-slot:prepend>
                              <v-icon large> mdi-account-group</v-icon>
                            </template>
                            <template v-slot:append>
                              <v-icon
                                v-if="isAccountManager"
                                large
                                color="success"
                              >
                                mdi-checkbox-marked-circle</v-icon
                              >
                              <v-icon v-else large color="error">
                                mdi-checkbox-blank-circle</v-icon
                              >
                            </template>
                            <v-list-item-title>
                              Account Manager
                            </v-list-item-title>
                          </v-list-item>

                          <v-list-item color="#263238">
                            <template v-slot:prepend>
                              <v-icon large> mdi-shield-crown</v-icon>
                            </template>
                            <template v-slot:append>
                              <v-icon v-if="isAdmin" large color="success">
                                mdi-checkbox-marked-circle</v-icon
                              >
                              <v-icon v-else large color="error">
                                mdi-checkbox-blank-circle</v-icon
                              >
                            </template>
                            <v-list-item-title> Admin </v-list-item-title>
                          </v-list-item>
                        </v-list>
                      </v-card-text>
                    </v-card>
                  </v-col>
                </v-row>
              </v-card-text>
            </v-card>
          </v-container>
        </v-window-item>

        <v-window-item value="roles">
          <v-container class="pt-0">
            <v-card elevation="2">
              <v-card-title class="text-h4">
                <v-container> Rollen </v-container>
              </v-card-title>
              <v-card-text>
                <v-row>
                  <v-col cols="12">
                    <v-expansion-panels expanded>
                      <v-expansion-panel
                        title="Online-Lernender: Die Entdeckung des Wissens im digitalen Zeitalter"
                      >
                        <template v-slot:text>
                          <p>
                            Als Online-Lernender ist es meine Mission,
                            kontinuierlich zu lernen, zu wachsen und mein Wissen
                            zu erweitern. Die Flexibilität des Online-Lernens
                            ermöglicht es mir, meinen eigenen Weg zu wählen und
                            in meinem eigenen Tempo zu lernen, was eine
                            aufregende und bereichernde Bildungserfahrung
                            darstellt.
                          </p>
                          <v-list lines="two">
                            <v-list-item>
                              Kursauswahl: Mein erster Schritt ist die Auswahl
                              von Kursen, die meinen Interessen und Lernzielen
                              entsprechen. Ich durchsuche eine Vielzahl von
                              Online-Plattformen und wähle Kurse aus, die mir
                              helfen, meine Fähigkeiten zu erweitern oder neue
                              Kenntnisse zu erwerben.
                            </v-list-item>
                            <v-list-item>
                              Kurskonsum: Ich konsumiere Kursinhalte in Form von
                              Videos, Texten, Aufgaben und Diskussionen. Ich
                              arbeite mich durch den Lehrplan und bearbeite
                              Übungen, um das Gelernte zu festigen.
                            </v-list-item>
                            <v-list-item>
                              Selbstmotivation: Online-Lernen erfordert
                              Selbstmotivation und Disziplin. Ich setze mir
                              Ziele und bleibe motiviert, um den Kurs
                              erfolgreich abzuschließen.
                            </v-list-item>
                            <v-list-item>
                              Zusammenarbeit: In einigen Kursen gibt es
                              Möglichkeiten zur Zusammenarbeit mit anderen
                              Lernenden. Wir diskutieren, teilen Wissen und
                              helfen uns gegenseitig, Herausforderungen zu
                              bewältigen.
                            </v-list-item>
                            <v-list-item>
                              Feedback und Fragen: Wenn ich Fragen habe oder
                              zusätzliche Unterstützung benötige, wende ich mich
                              an den Kursleiter oder die Community, um Antworten
                              zu erhalten.
                            </v-list-item>
                            <v-list-item>
                              Evaluierung: Am Ende des Kurses nehme ich an
                              Prüfungen oder Evaluierungen teil, um mein
                              Verständnis und meine Fähigkeiten zu überprüfen.
                            </v-list-item>
                            <v-list-item>
                              Weiterbildung: Nach Abschluss eines Kurses schaue
                              ich nach neuen Lernmöglichkeiten und Kursen, um
                              meine Bildungskurve weiter nach oben zu treiben.
                            </v-list-item>
                          </v-list>
                        </template>
                      </v-expansion-panel>
                    </v-expansion-panels>
                  </v-col>
                  <v-col v-show="isAuthor" cols="12">
                    <v-expansion-panels variant="inset">
                      <v-expansion-panel
                        title="Online-Kursgestalter: Das Schaffen von Bildungserfahrungen im digitalen Zeitalter"
                      >
                        <template v-slot:text>
                          <p>
                            Die Rolle des Online-Kursgestalters erfordert
                            Kreativität, pädagogisches Know-how und technische
                            Fähigkeiten. Mein Ziel ist es, inspirierende
                            Bildungserfahrungen zu schaffen, die Lernende in die
                            Lage versetzen, neue Fähigkeiten zu erwerben und
                            Wissen zu erweitern, unabhängig von ihrem Standort
                            und ihrem Zeitplan.
                          </p>
                          <v-list lines="two">
                            <v-list-item>
                              Kursplanung und -entwicklung: Mein erster Schritt
                              ist die Planung des Kurses, einschließlich der
                              Festlegung von Lernzielen, Lehrmaterialien und
                              Evaluationsmethoden. Ich erstelle den Lehrplan und
                              gestalte den Kurs so, dass er benutzerfreundlich
                              und effektiv ist.
                            </v-list-item>
                            <v-list-item>
                              Multimedia-Inhalte: Ich arbeite mit verschiedenen
                              Medienformaten wie Videos, Audio, Text und
                              interaktiven Elementen, um ansprechende und
                              effektive Lerninhalte zu erstellen. Dies erfordert
                              kreative Gestaltung und die Auswahl geeigneter
                              Tools.
                            </v-list-item>
                            <v-list-item>
                              Plattformintegration: Die Integration des Kurses
                              in eine Lernplattform oder Website ist
                              entscheidend. Ich sorge dafür, dass die
                              Benutzeroberfläche intuitiv ist und dass Benutzer
                              leicht auf Materialien zugreifen können.
                            </v-list-item>
                            <v-list-item>
                              Interaktivität und Engagement: Online-Lernen
                              sollte interaktiv und ansprechend sein. Ich
                              entwickle Übungen, Quizfragen und
                              Diskussionsforen, um die Beteiligung der Lernenden
                              zu fördern.
                            </v-list-item>
                            <v-list-item>
                              Begleitung und Unterstützung: Ich stehe den
                              Lernenden zur Seite, um Fragen zu beantworten,
                              Feedback zu geben und Unterstützung anzubieten, um
                              sicherzustellen, dass sie den Kurs erfolgreich
                              abschließen.
                            </v-list-item>
                            <v-list-item>
                              Kontinuierliche Verbesserung: Nach jedem Kurs
                              sammle ich Feedback von den Teilnehmern,
                              analysiere die Erfolge und Herausforderungen und
                              passe den Kurs entsprechend an, um die Qualität
                              der Bildung zu verbessern.
                            </v-list-item>
                            <v-list-item>
                              Technologie und Trends: In einer sich ständig
                              weiterentwickelnden digitalen Welt halte ich mich
                              über neue Technologien und Bildungstrends auf dem
                              Laufenden, um sicherzustellen, dass meine Kurse
                              relevant und aktuell sind.
                            </v-list-item>
                          </v-list>
                        </template>
                      </v-expansion-panel>
                    </v-expansion-panels>
                  </v-col>

                  <v-col v-show="isUserCourseManager" cols="12">
                    <v-expansion-panels variant="inset">
                      <v-expansion-panel
                        title="Bildung gestalten und organisieren: Die Rolle des Kursverwalters"
                      >
                        <template v-slot:text>
                          <p>
                            Die Rolle eines Kursverwalters erfordert
                            Organisationstalent, Kommunikationsfähigkeiten und
                            die Fähigkeit, Bildungsziele effektiv umzusetzen. Es
                            ist meine Mission, die Bildungserfahrung für alle
                            Teilnehmer so reibungslos und wertvoll wie möglich
                            zu gestalten und dazu beizutragen, dass sie ihr
                            volles Potenzial entfalten können.
                          </p>
                          <v-list lines="two">
                            <v-list-item>
                              Kursplanung: Mein erster Schritt als Kursverwalter
                              besteht darin, den Lehrplan zu entwickeln und
                              sicherzustellen, dass er den Bedürfnissen der
                              Lernenden entspricht. Dies umfasst die Auswahl der
                              relevanten Lehrmaterialien, die Festlegung von
                              Lehrzielen und die Erstellung eines Stundenplans.
                            </v-list-item>
                            <v-list-item>
                              Teilnehmerverwaltung: Die Anmeldung, Zuweisung von
                              Klassen und die Kommunikation mit den Teilnehmern
                              gehören ebenfalls zu meinen Verantwortlichkeiten.
                              Ich sorge dafür, dass alle Informationen zu den
                              Kursen für die Lernenden leicht zugänglich sind.
                            </v-list-item>
                            <v-list-item>
                              Ressourcenmanagement: Als Kursverwalter muss ich
                              sicherstellen, dass alle benötigten Ressourcen,
                              wie Lehrbücher, technische Ausrüstung oder
                              Räumlichkeiten, rechtzeitig verfügbar sind.
                            </v-list-item>
                            <v-list-item>
                              Kommunikation: Eine klare Kommunikation mit den
                              Lernenden, Dozenten und anderen beteiligten
                              Parteien ist von größter Bedeutung. Ich beantworte
                              Fragen, löse Probleme und halte alle Beteiligten
                              auf dem Laufenden.
                            </v-list-item>
                            <v-list-item>
                              Evaluation und Verbesserung: Nach jedem Kurs
                              analysiere ich die Ergebnisse und sammle Feedback
                              von Teilnehmern, um den Lehrplan und die Methoden
                              zu verbessern und sicherzustellen, dass die
                              Bildungsziele erreicht werden.
                            </v-list-item>
                          </v-list>
                        </template>
                      </v-expansion-panel>
                    </v-expansion-panels>
                  </v-col>

                  <v-col v-show="isAccountManager">
                    <v-expansion-panels>
                      <v-expansion-panel
                        title="Benutzerverwalter: Die Schaltzentrale für Zugriffssteuerung und Benutzererfahrung"
                      >
                        <template v-slot:text>
                          <p>
                            Die Rolle eines Benutzerverwalters erfordert ein
                            gutes Verständnis für Zugriffskontrollen,
                            Sicherheitspraktiken und Benutzererfahrung. Es ist
                            meine Mission, die Plattformnutzung sicher und
                            benutzerfreundlich zu gestalten und gleichzeitig die
                            Integrität und Vertraulichkeit von Daten zu wahren.
                          </p>
                          <v-list lines="two">
                            <v-list-item>
                              Kontenverwaltung: Mein erster Schritt als
                              Benutzerverwalter besteht darin, Benutzerkonten zu
                              erstellen, zu aktualisieren oder zu deaktivieren.
                              Ich sorge dafür, dass alle Benutzer Zugriff auf
                              die benötigten Ressourcen haben.
                            </v-list-item>
                            <v-list-item>
                              Berechtigungsmanagement: Die Zuweisung und
                              Überwachung von Zugriffsrechten ist ein zentraler
                              Aspekt meiner Arbeit. Ich stelle sicher, dass
                              Benutzer die richtigen Berechtigungen haben, um
                              auf bestimmte Funktionen oder Daten zuzugreifen,
                              und verhindere unbefugten Zugriff.
                            </v-list-item>
                            <v-list-item>
                              Passwortverwaltung: Die Sicherheit der
                              Benutzerkonten hat oberste Priorität. Ich erzwinge
                              starke Passwortrichtlinien und helfe Benutzern bei
                              der Wiederherstellung ihrer Konten, wenn sie ihr
                              Passwort vergessen haben.
                            </v-list-item>
                            <v-list-item>
                              Kommunikation: Eine klare Kommunikation mit
                              Benutzern und Administratoren ist unerlässlich.
                              Ich stehe für Fragen zur Verfügung, kläre Benutzer
                              über Richtlinien auf und sorge für eine
                              reibungslose Kommunikation bei Anmeldeproblemen
                              oder Sicherheitswarnungen.
                            </v-list-item>
                            <v-list-item>
                              Sicherheitsüberwachung: Die Überwachung von
                              Benutzeraktivitäten und die Erkennung verdächtiger
                              Aktivitäten sind ein wichtiger Bestandteil meiner
                              Rolle, um Sicherheitsverletzungen zu verhindern.
                            </v-list-item>
                            <v-list-item>
                              Datenverwaltung: Als Benutzerverwalter habe ich
                              oft Zugriff auf Benutzerdaten und muss
                              sicherstellen, dass sie gemäß den
                              Datenschutzbestimmungen behandelt werden.
                            </v-list-item>
                          </v-list>
                        </template>
                      </v-expansion-panel>
                    </v-expansion-panels>
                  </v-col>

                  <v-col v-show="isAdmin">
                    <v-expansion-panels>
                      <v-expansion-panel
                        title="Administrator der Online-Kursplattform: Die Steuerzentrale für reibungsloses E-Learning"
                      >
                        <template v-slot:text>
                          <p>
                            Die Rolle eines Administrators der
                            Online-Kursplattform erfordert technisches Know-how,
                            Sicherheitsbewusstsein und die Fähigkeit zur
                            effizienten Verwaltung von Ressourcen. Mein Ziel ist
                            es, eine stabile, sichere und benutzerfreundliche
                            Plattform zu schaffen, auf der effektives E-Learning
                            stattfinden kann, und dabei die Bedürfnisse der
                            Lernenden und Kursleiter im Blick zu behalten.
                          </p>
                          <v-list lines="two">
                            <v-list-item>
                              Plattformverwaltung: Mein erster Schritt als
                              Administrator besteht darin, die
                              Online-Kursplattform einzurichten und zu
                              konfigurieren. Ich verwalte Benutzerkonten, Rollen
                              und Berechtigungen, um sicherzustellen, dass alle
                              Zugriff auf die benötigten Ressourcen haben.
                            </v-list-item>

                            <v-list-item>
                              Technische Wartung: Die regelmäßige Wartung der
                              Plattform, einschließlich Updates und
                              Fehlerbehebung, liegt in meiner Verantwortung. Ich
                              sorge dafür, dass die Plattform sicher und auf dem
                              neuesten Stand ist.
                            </v-list-item>
                            <v-list-item>
                              Sicherheit: Die Sicherheit der Plattform und der
                              Benutzerdaten hat oberste Priorität. Ich
                              implementiere Sicherheitsrichtlinien,
                              Überwachungssysteme und Datenschutzbestimmungen,
                              um die Integrität und Vertraulichkeit der Daten zu
                              gewährleisten.
                            </v-list-item>
                            <v-list-item>
                              Ressourcenmanagement: Ich verwalte Ressourcen wie
                              Serverkapazität, Bandbreite und Lizenzen, um
                              sicherzustellen, dass die Plattform den
                              Anforderungen gerecht wird und effizient läuft.
                            </v-list-item>
                            <v-list-item>
                              Benutzerunterstützung: Ich bin die erste
                              Anlaufstelle für technische Unterstützung und
                              Fragen von Benutzern. Ich stelle sicher, dass
                              Benutzer schnell und effizient Hilfe erhalten.
                            </v-list-item>
                            <v-list-item>
                              Berichterstattung und Analyse: Ich erstelle
                              Berichte über die Nutzung der Plattform,
                              Benutzeraktivitäten und den Erfolg der Kurse.
                              Diese Analysen helfen bei der kontinuierlichen
                              Verbesserung der Plattform.
                            </v-list-item>
                            <v-list-item>
                              Schulung: Ich biete Schulungen und Ressourcen für
                              Kursleiter und Benutzer an, um sicherzustellen,
                              dass sie die Plattform optimal nutzen können.
                            </v-list-item>
                            <v-list-item>
                              Kommunikation: Die Kommunikation mit Kursleitern,
                              Administratoren und Benutzern ist entscheidend.
                              Ich halte alle Parteien über Updates, Änderungen
                              und Problemlösungen auf dem Laufenden.
                            </v-list-item>
                          </v-list>
                        </template>
                      </v-expansion-panel>
                    </v-expansion-panels>
                  </v-col>
                </v-row>
              </v-card-text>
              <v-card-actions>
                <v-row>
                  <v-col cols="12" md="6">
                    <v-btn disabled block color="warning" variant="outlined">
                      Verringern</v-btn
                    >
                  </v-col>
                  <v-col cols="12" md="6">
                    <v-btn
                      disabled
                      block
                      color="success"
                      variant="outlined"
                      prepend-icon="mdi-plus"
                    >
                      Erweitern</v-btn
                    >
                  </v-col>
                </v-row>
              </v-card-actions>
            </v-card>
          </v-container>
        </v-window-item>

        <v-window-item value="meta">
          <v-container class="pt-0">
            <v-card elevation="2">
              <v-card-title class="text-h4">
                <v-container> Course </v-container>
              </v-card-title>
              <v-card-text> </v-card-text>
              <v-card-actions>
                <v-btn block color="warning" variant="outlined">
                  Bearbeiten</v-btn
                >
              </v-card-actions>
            </v-card>
          </v-container>
        </v-window-item>

        <v-window-item value="logs">
          <v-container class="pt-0">
            <v-card elevation="2">
              <v-card-title class="text-h4">
                <v-container> Logs </v-container>
              </v-card-title>
              <v-card-text> 

                


              </v-card-text>
              <v-card-actions>
                <v-btn block color="warning" variant="outlined">
                  Download</v-btn
                >
              </v-card-actions>
            </v-card>
          </v-container>
        </v-window-item>

        <v-window-item value="account">
          <v-container class="pt-0">
            <v-card elevation="2">
              <v-card-title class="text-h4">
                <v-container> Benutzerdaten herunterladen </v-container>
              </v-card-title>
              <v-card-text> TODO</v-card-text>
              <v-card-actions> </v-card-actions>
            </v-card>
          </v-container>

          <v-container>
            <v-card elevation="2">
              <v-card-title class="text-h4">
                <v-container> Konto-Löschung bestätigen </v-container>
              </v-card-title>
              <v-card-text>
                <v-list>
                  <v-list-item>
                    Bitte beachten Sie, dass die Löschung Ihres Kontos
                    irreversibel ist. Nach der Bestätigung wird Ihr Konto
                    vollständig und unwiderruflich gelöscht. Alle Daten,
                    Informationen und Zugriffsrechte werden entfernt, und Sie
                    verlieren jeglichen Zugang zu Ihrem Konto und seinen
                    Inhalten.</v-list-item
                  >
                  <v-list-item>
                    Wenn Sie sicher sind, dass Sie Ihr Konto löschen möchten,
                    klicken Sie auf "Bestätigen". Andernfalls können Sie diesen
                    Vorgang abbrechen, indem Sie auf "Abbrechen" klicken. Wir
                    schätzen Ihr Feedback und sind hier, um Ihnen zu helfen,
                    falls Sie Fragen oder Bedenken haben.
                  </v-list-item>
                  <v-list-item>
                    Bitte stellen Sie sicher, dass Sie eventuelle wichtige Daten
                    oder Informationen vor der Löschung Ihres Kontos sichern, da
                    sie nach der Bestätigung nicht wiederhergestellt werden
                    können.
                  </v-list-item>
                  <v-list-item
                    >Wir bedanken uns für Ihre Nutzung unserer Plattform und
                    hoffen, dass Sie in Zukunft wieder zurückkehren. Falls Sie
                    Ihre Meinung ändern, stehen wir Ihnen gerne zur Verfügung,
                    um Ihnen erneut unsere Dienste anzubieten.</v-list-item
                  >
                </v-list>
              </v-card-text>
              <v-card-action>
                <v-btn large color="error" block prepend-icon="mdi-trash-can">
                  loeschen
                </v-btn>
              </v-card-action>
            </v-card>
            <v-container> </v-container>
          </v-container>
        </v-window-item>
      </v-window>
    </v-col>
  </v-row>
</template>
