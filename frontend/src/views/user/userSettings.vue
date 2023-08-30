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

const panelOption = ref(["personal"]);
const tabNavigation = ref(null);
</script>

<template>
  <v-container>
    <v-avatar class="ml-2 mr-2" size="large">
      <v-img :src="userInfo.picture" alt="John"></v-img>
    </v-avatar>

    <v-row>
      <v-container>
        <v-card>
          <v-tabs v-model="tabNavigation" bg-color="primary">
            <v-tab value="one">Details</v-tab>
            <v-tab value="two">Roles</v-tab>
            <v-tab value="three">Appdata</v-tab>
            <v-tab value="three">Metadata</v-tab>
          </v-tabs>

          <v-card-text>
            <v-window v-model="tabNavigation">
              <v-window-item value="one">
                <v-container>
                  <v-text-field
                    label="Name"
                    :model-value="`${userInfo.name}`"
                    variant="outlined"
                    disabled
                    readonly
                  />
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
              </v-window-item>

              <v-window-item value="two">
                <v-container>
                  <v-row>
                    <v-col cols="6">
                      <v-container>
                        <v-checkbox :v-model="true">Basic</v-checkbox>
                        <v-checkbox value="isAuthor">Author</v-checkbox>
                        <v-checkbox :v-model="isAdmin">Admin</v-checkbox>
                      </v-container>
                    </v-col>
                    <v-col cols="6">
                      <v-container>
                        <v-checkbox :v-model="isUserCourseManager"
                          >Course Manager</v-checkbox
                        >
                        <v-checkbox :v-model="isAccountManager"
                          >Account Manager</v-checkbox
                        >
                      </v-container>
                    </v-col>
                  </v-row>
                </v-container>
              </v-window-item>

              <v-window-item value="three"> Three </v-window-item>
            </v-window>
          </v-card-text>
        </v-card>
      </v-container>
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
