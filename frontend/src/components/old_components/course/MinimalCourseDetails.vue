<template>
  <v-container>
    <v-row>
      <v-col>
        <v-card height="200" shaped elevation="3">
          <v-row dense class="pb-0">
            <v-col class="text-end">
              <div class="pr-2">
                <v-rating
                  :value="4.5"
                  color="amber"
                  dense
                  half-increments
                  readonly
                  :size="ratingSize"
                ></v-rating>
              </div>
            </v-col>
          </v-row>
          <v-card-title class="pt-0">
            <div class="ma-2">
              <h1 class="headline">
                {{ cTitle }}
              </h1>
            </div>
          </v-card-title>
          <v-card-subtitle>
            <v-icon small class="pr-2">
              mdi-badge-account-horizontal-outline
            </v-icon>
            <span class="caption font-italic"> {{ cAuthor }}</span>
            <div class="pa-1"></div>
            <v-divider class="mx-2"></v-divider>
          </v-card-subtitle>

          <v-card-text class="OVERLINE">
            {{ cDescription }}
          </v-card-text>
          <v-card-actions class="pa-5">
            <v-row>
              <v-col class="text-end">
                <v-tooltip bottom>
                  <template v-slot:activator="{ on, attrs }">
                    <v-btn
                      x-small
                      fab
                      color="primary"
                      dark
                      v-bind="attrs"
                      v-on="on"
                    >
                      <v-icon> mdi-school </v-icon>
                    </v-btn>
                  </template>
                  <span>Kurs anschauen</span>
                </v-tooltip>
              </v-col>
            </v-row>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "MinimalCourseDetails",
  props: {
    cTitle: {
      type: String,
      required: true,
    },
    cAuthor: {
      type: String,
      required: true,
    },
    cDescription: {
      type: String,
      required: true,
    },
    // cLevel: {
    //   type: [Number],
    //   required: true,
    // },
  },
  data: () => ({
    ratingSize: 10,
  }),
};
</script>

<style></style>
