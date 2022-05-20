<template>
  <div>
    <v-container fluid>
      <v-row dense>
        <v-col cols="4" class="mx-auto">
          <v-text-field
            solo
            v-model="searchValue"
            placeholder="Suche ..."
            :append-icon="
              isFilterEnabled ? 'mdi-filter-off-outline' : 'mdi-filter-outline'
            "
            @click:append="isFilterEnabled = !isFilterEnabled"
          >
          </v-text-field>
        </v-col>
      </v-row>
      <v-row dense>
        <v-col cols="4" class="mx-auto">
          <v-expand-transition>
            <v-card v-show="isFilterEnabled" class="mx-auto">
              <v-card-text>
                <v-row>
                  <v-col cols="12" class="mx-auto">
                    <v-checkbox
                      v-model="searchFilter.title"
                      label="Titel einbinden"
                      color="primary"
                      hide-details
                    ></v-checkbox>
                  </v-col>
                  <v-col cols="12" class="mx-auto">
                    <v-checkbox
                      v-model="searchFilter.author"
                      label="Autor einbinden"
                      color="primary"
                      hide-details
                    ></v-checkbox>
                  </v-col>
                  <v-col cols="12" class="mx-auto">
                    <v-checkbox
                      v-model="searchFilter.description"
                      label="Beschreibung einbinden"
                      color="primary"
                      hide-details
                    ></v-checkbox>
                  </v-col>
                  <v-col cols="12" class="mx-auto">
                    <v-range-slider
                      disabled
                      :tick-labels="rangeDescriptions"
                      :value="[0]"
                      dense
                      max="3"
                      min="0"
                    ></v-range-slider>
                    <span>{{ searchFilter.rangeOfLevel }}</span>
                  </v-col>
                </v-row></v-card-text
              >
            </v-card>
          </v-expand-transition>
        </v-col>
      </v-row>
    </v-container>
    <v-container>
      <v-row>
        <v-col
          cols="12"
          sm="12"
          md="6"
          lg="4"
          xl="3"
          class="mx-auto"
          v-for="course in loadedCourses"
          :key="course.ID"
        >
          <mcd
            :cTitle="course.Name"
            :cDescription="course.Description"
            :cAuthor="course.Author"
            :cLevel="course.Level"
          />
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import { mapState } from "vuex";
import MinimalCourseDetails from "@/components/course/MinimalCourseDetails";
export default {
  name: "View_CourseOverview",
  components: {
    mcd: MinimalCourseDetails,
  },
  data: () => ({
    isFilterEnabled: false,
    searchValue: null,
    searchFilter: {
      title: null,
      author: null,
      description: null,
      rangeOfLevel: null,
    },
    rangeDescriptions: ["Low", "Normal", "Medium", "Hard"],
  }),
  methods: {
    checkContentOfStrings(item, search) {
      console.log(
        "item.toLowerCase().includes(search.toLowerCase():",
        item.toLowerCase().includes(search.toLowerCase())
      );

      return item.toLowerCase().includes(search.toLowerCase());
    },
  },
  computed: {
    ...mapState({
      allCourses: (state) => state.courseStore.general.allCourses,
    }),
    loadedCourses() {
      if (this.searchValue) {
        if (!this.isFilterEnabled) {
          return this.allCourses.filter((item) =>
            item.Name.includes(this.searchValue)
          );
        }

        return this.allCourses.filter((item) => {
          if (
            this.searchFilter.title &&
            this.searchFilter.author &&
            this.searchFilter.description
          ) {
            return (
              this.checkContentOfStrings(item.Name, this.searchValue) ||
              this.checkContentOfStrings(item.Author, this.searchValue) ||
              this.checkContentOfStrings(item.Description, this.searchValue)
            );
          } else if (this.searchFilter.title && this.searchFilter.author) {
            return (
              this.checkContentOfStrings(item.Name, this.searchValue) ||
              this.checkContentOfStrings(item.Author, this.searchValue)
            );
          } else if (this.searchFilter.title && this.searchFilter.description) {
            return (
              this.checkContentOfStrings(item.Name, this.searchValue) ||
              this.checkContentOfStrings(item.Description, this.searchValue)
            );
          } else if (this.searchFilter.title) {
            return this.checkContentOfStrings(item.Name, this.searchValue);
          } else if (this.searchFilter.author) {
            return this.checkContentOfStrings(item.Author, this.searchValue);
          } else if (this.searchFilter.description) {
            return this.checkContentOfStrings(
              item.Description,
              this.searchValue
            );
          }
        });
      }
      return this.allCourses;
    },
  },
};
</script>

<style></style>
