<script setup lang="ts">
import { useLoginStore } from "@/stores/session/loginStore";
import { computed, ref } from "vue";
import { useNow, useDateFormat } from "@vueuse/core";

const loginStore = useLoginStore();

const userInfo = computed(() => {
  return loginStore.getUser;
});

const formattedDate = (dateToConvert: string) => {
  return useDateFormat(useNow(), dateToConvert);
};

const panelOption = ref(["personal", "moreInfo"]);
</script>

<template>
  <v-expansion-panels
    variant="popout"
    class="my-4"
    multiple
    v-model="panelOption"
  >
    <v-expansion-panel title="Personal Information" value="personal">
      <v-expansion-panel-text>
        <v-container>
          <v-text-field
            label="Full Name"
            :model-value="`${userInfo.first_name} ${userInfo.last_name}`"
            variant="outlined"
            disabled
            readonly
          />
          <v-text-field label="Birthday" variant="outlined" />
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
      </v-expansion-panel-text>
    </v-expansion-panel>
    <v-expansion-panel title="More Information" value="moreInfo">
      <v-expansion-panel-text>
        <v-container>
          <v-text-field
            label="Company"
            :model-value="userInfo.company"
            variant="outlined"
            readonly
            disabled
          />
          <v-text-field
            label="Occupation"
            :model-value="userInfo.occupation"
            variant="outlined"
            readonly
            disabled
          />
          <v-text-field
            label="School"
            :model-value="userInfo.school"
            variant="outlined"
            readonly
            disabled
          />
          <v-text-field
            label="Subject"
            :model-value="userInfo.subject"
            variant="outlined"
            readonly
          />
          <v-text-field
            label="Contry"
            :model-value="userInfo.country"
            variant="outlined"
            readonly
          />
        </v-container>
      </v-expansion-panel-text>
    </v-expansion-panel>
    <v-expansion-panel title="Roles" value="userRoles">
      <v-expansion-panel-text>
        <v-container>
          <v-checkbox value="basic">Basic</v-checkbox>
          <v-checkbox value="basic">Author</v-checkbox>
          <v-checkbox value="basic">Manager</v-checkbox>
          <v-checkbox value="basic">Admin</v-checkbox>
        </v-container>
      </v-expansion-panel-text>
    </v-expansion-panel>
    <v-expansion-panel title="Account Information" value="creationInfo">
      <v-expansion-panel-text> Created: logs: </v-expansion-panel-text>
    </v-expansion-panel>
    <v-expansion-panel title="Danger Zone" value="deleteAccount">
      <v-expansion-panel-text>
        Konto loeschen Konto deactivieren
      </v-expansion-panel-text>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<style scoped></style>
