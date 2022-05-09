<template>
  <v-container fill-height align="center" justify="center">
    <v-row>
      <v-col cols="12" sm="12" md="6" lg="6" xl="6" class="mx-auto">
        <v-card elevation="5" outlined>
          <v-card-title>
            <h1 class="display pa-5">Umbau arbeiten</h1>
          </v-card-title>
          <v-img :src="require('../assets/underConstruction.png')" alt="Rework">
            <v-container class="ma-5">
              <v-row>
                <v-col class="mx-auto" cols="5">
                  <p class="title">
                    Die Website ist aktuell im Aufbau und wird in Kürze
                    freigegeben. Wir bitten um etwas Geduld.
                  </p>
                </v-col>
              </v-row>
            </v-container>
          </v-img>
        </v-card>
        <v-card-actions>
          <v-progress-linear
            color="red accent-4"
            indeterminate
            rounded
            height="4"
          ></v-progress-linear>
        </v-card-actions>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "Rework",
  data: () => ({}),
};
</script>
