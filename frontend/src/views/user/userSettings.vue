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
          <v-tab value="meta">
            <v-icon start>mdi-protocol</v-icon>
            Logs
          </v-tab>
          <v-tab value="meta">
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
              <v-card-text>
                <v-containter align="center">
                  <span class="text-h4">General</span>
                </v-containter>
              </v-card-text>
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
            <v-row>
              <v-col>
                <v-card color="#42A5F5" theme="dark">
                  <v-card-title class="text-h5"> BASIC </v-card-title>

                  <v-card-subtitle
                    >Listen to your favorite artists and albums whenever and
                    wherever, online and offline.</v-card-subtitle
                  >
                </v-card>
              </v-col>
              <v-col v-show="isAuthor">
                <v-card color="#5C6BC0" theme="dark">
                  <v-card-title class="text-h5"> AUTHOR </v-card-title>

                  <v-card-subtitle
                    >Listen to your favorite artists and albums whenever and
                    wherever, online and offline.</v-card-subtitle
                  >
                </v-card>
              </v-col>

              <v-col v-show="isUserCourseManager">
                <v-card color="#FF8F00" theme="dark">
                  <v-card-title class="text-h5"> COURSE MANAGER </v-card-title>

                  <v-card-subtitle
                    >Listen to your favorite artists and albums whenever and
                    wherever, online and offline.</v-card-subtitle
                  >
                </v-card>
              </v-col>

              <v-col v-show="isAccountManager">
                <v-card color="#D84315" theme="dark">
                  <v-card-title class="text-h5"> USER MANAGER </v-card-title>

                  <v-card-subtitle
                    >Listen to your favorite artists and albums whenever and
                    wherever, online and offline.</v-card-subtitle
                  >
                </v-card>
              </v-col>

              <v-col v-show="isAdmin">
                <v-card color="#263238" theme="dark">
                  <v-card-title class="text-h5"> ADMIN </v-card-title>

                  <v-card-subtitle
                    >Listen to your favorite artists and albums whenever and
                    wherever, online and offline.</v-card-subtitle
                  >
                </v-card>
              </v-col>
            </v-row>
          </v-container>
        </v-window-item>

        <v-window-item value="meta">
          <v-container class="pt-0">
            <v-card>
              <v-card-title> Theme </v-card-title>
              <v-card-text></v-card-text>
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
