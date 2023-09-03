<script setup lang="ts">
import { useLoginStore } from "@/stores/session/loginStore";
import { computed, ref } from "vue";
import { useNow, useDateFormat } from "@vueuse/core";

const loginStore = useLoginStore();

const userInfo = computed(() => {
  return loginStore.getUser;
});

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

const tabNavigation = ref(null);
</script>

<template>
  <v-container>
    <v-row>
      <v-col> </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" md="2">
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
              <v-icon start> </v-icon>
              Logs
            </v-tab>
            <v-tab value="meta">
              <v-icon start>mdi-tablet-cellphone</v-icon>
              Devices
            </v-tab>
          </v-tabs>
        </v-card>
      </v-col>

      <v-col cols="12" md="10">
        <v-card>
          <v-card-text>
            <v-window v-model="tabNavigation">
              <v-window-item value="details">
                <v-container>
                  <v-row>
                    <v-col cols="1">
                      <v-avatar class="ml-2 mr-2" size="large">
                        <v-img :src="userInfo.picture" alt="John"></v-img>
                      </v-avatar>
                    </v-col>
                    <v-col>
                      <v-text-field
                        label="ID"
                        :model-value="`${userInfo.appUID}`"
                        variant="solo"
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
                  </v-row>
                </v-container>
                <v-container>
                  <v-row>
                    <v-col>
                      <v-text-field
                        label="Nickname"
                        :model-value="`${userInfo.nickname}`"
                        variant="outlined"
                        disabled
                        readonly
                      />
                    </v-col>
                    <v-col>
                      <v-text-field
                        label="Name"
                        :model-value="`${userInfo.name}`"
                        variant="outlined"
                        disabled
                        readonly
                      />
                    </v-col>
                  </v-row>
                  <!-- <v-text-field label="Birthday" variant="outlined" /> -->
                  <v-text-field
                    label="Email"
                    :model-value="userInfo.email"
                    variant="outlined"
                    readonly
                    disabled
                  />
                  <v-text-field
                    label="Mobile"
                    :model-value="userInfo.telNr"
                    variant="outlined"
                    readonly
                    disabled
                  />
                </v-container>
                <pre>
                  {{ userInfo }}
                </pre>
              </v-window-item>

              <v-window-item value="roles">
                <v-container>
                  <v-row>
                    <v-col>
                      <v-card color="#42A5F5" theme="dark">
                        <v-card-title class="text-h5"> BASIC </v-card-title>

                        <v-card-subtitle
                          >Listen to your favorite artists and albums whenever
                          and wherever, online and offline.</v-card-subtitle
                        >
                      </v-card>
                    </v-col>
                    <v-col v-show="isAuthor">
                      <v-card color="#5C6BC0" theme="dark">
                        <v-card-title class="text-h5"> AUTHOR </v-card-title>

                        <v-card-subtitle
                          >Listen to your favorite artists and albums whenever
                          and wherever, online and offline.</v-card-subtitle
                        >
                      </v-card>
                    </v-col>

                    <v-col v-show="isUserCourseManager">
                      <v-card color="#FF8F00" theme="dark">
                        <v-card-title class="text-h5">
                          COURSE MANAGER
                        </v-card-title>

                        <v-card-subtitle
                          >Listen to your favorite artists and albums whenever
                          and wherever, online and offline.</v-card-subtitle
                        >
                      </v-card>
                    </v-col>

                    <v-col v-show="isAccountManager">
                      <v-card color="#D84315" theme="dark">
                        <v-card-title class="text-h5">
                          USER MANAGER
                        </v-card-title>

                        <v-card-subtitle
                          >Listen to your favorite artists and albums whenever
                          and wherever, online and offline.</v-card-subtitle
                        >
                      </v-card>
                    </v-col>

                    <v-col v-show="isAdmin">
                      <v-card color="#263238" theme="dark">
                        <v-card-title class="text-h5"> ADMIN </v-card-title>

                        <v-card-subtitle
                          >Listen to your favorite artists and albums whenever
                          and wherever, online and offline.</v-card-subtitle
                        >
                      </v-card>
                    </v-col>
                  </v-row>
                </v-container>
              </v-window-item>

              <v-window-item value="meta"> Three </v-window-item>
            </v-window>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-container>
        <v-card color="red">
          <v-card-title> DANGER ZONE </v-card-title>
          <v-container> </v-container>
        </v-card>
      </v-container>
    </v-row>
  </v-container>
</template>
