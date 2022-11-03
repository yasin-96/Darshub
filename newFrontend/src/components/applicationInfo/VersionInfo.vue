<template>
  <div>
    <div class="container">
      <div class="mx-auto" max-width="344">
        <div-title>
          <h1 class="mx-auto headline">Appversion</h1>
        </div-title>
        <div-text>
          <div class="container">
            <v-row v-if="onlyDate">
              <v-col>
                <span class="subtitle-1 mr-5">Git Date:</span>
              </v-col>
              <v-col>
                <v-chip label class="text-right">
                  <strong>{{ onlyDate }}</strong>
                </v-chip>
              </v-col>
            </v-row>
            <v-row v-if="minimalHashCode">
              <v-col>
                <span class="subtitle-1 mr-5">Git Commit Id:</span>
              </v-col>
              <v-col>
                <v-chip label class="text-right">
                  <strong>{{ minimalHashCode }}</strong>
                </v-chip>
              </v-col>
            </v-row>

            <v-row v-if="checkVersionTag">
              <v-col>
                <span class="subtitle-1 mr-5">Git Tag:</span>
              </v-col>
              <v-col>
                <v-chip label class="text-right">
                  <strong>{{ checkVersionTag }}</strong>
                </v-chip>
              </v-col>
            </v-row>
          </div>
        </div-text>
      </div>
    </div>

    <div class="container">
      <h2>Über diese App</h2>
      <h4>TODO</h4>
    </div>
  </div>
</template>

<script>
// import { mapGetters } from "vuex";

export default {
  name: "VersionInfo",
  computed: {
    ...mapGetters({
      onlyDate: "gitData/onlyDate",
      minimalHashCode: "gitData/minimalHashCode",
      checkVersionTag: "gitData/checkVersionTag",
    }),
  },
};
</script>
