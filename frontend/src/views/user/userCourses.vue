<script setup lang="ts">
import type { Course, Chapter, SubChapter } from "@/models/course/types";
import { useCoreCourseStore } from "@/stores/course/coreCourseStore";
import { useLoginStore } from "@/stores/session/loginStore";
import { reactive } from "vue";
import { computed } from "vue";
import { ref } from "vue";

const loginStore = useLoginStore();
const coreCourseStore = useCoreCourseStore();

const userId = computed(() => {
  return loginStore.getUserId;
});

const userName = computed(() => {
  return loginStore.authDetails.user?.nickname;
});

const tabs = ref();
const chapterTabs = ref(["1.)"]);

const courseDetails = reactive({} as Course);
courseDetails.author = userId.value;

const courseChapter = reactive([] as Array<Chapter>);
const indexChapter = ref(0);

const courseSubChapter = reactive([] as Array<SubChapter>);
const indexSubChapter = ref(0);
</script>

<template>
  <v-container fluid>
    <v-row>
      <v-col cols="12" md="2">
        <v-card>
          <v-tabs v-model="tabs" direction="vertical" color="primary">
            <v-tab value="option-1">
              <v-icon start> mdi-book-cog-outline </v-icon>
              General
            </v-tab>
            <v-tab value="option-2">
              <v-icon start> mdi-lock </v-icon>
              Chapter
            </v-tab>
            <v-tab value="option-3">
              <v-icon start> mdi-access-point </v-icon>
              
            </v-tab>
          </v-tabs>
        </v-card>
      </v-col>
      <v-col>
        <v-card>
          <v-window v-model="tabs">
            <v-window-item value="option-1">
              <v-card>
                <v-form>
                  <v-card-title>
                    <v-text-field
                      clearable
                      label="Course Title"
                      variant="outlined"
                      :v-model="courseDetails.name"
                    ></v-text-field>
                  </v-card-title>
                  <v-card-text>
                    <v-container>
                      <v-text-field
                        label="Author"
                        variant="filled"
                        v-model="userName"
                        readonly
                      />
                      <v-textarea label="Description" variant="outlined" />

                      <v-rating
                        :v-model="courseDetails.level"
                        :item-labels="['easy', '', '', '', 'hard']"
                        item-label-position="top"
                      />
                      <!-- <v-time-picker v-model="courseDetails.duration"></v-time-picker> -->
                    </v-container>
                  </v-card-text>
                </v-form>
              </v-card>
            </v-window-item>
            <v-window-item value="option-2">
              <v-card flat>
                <v-card-text>
                </v-card-text>
              </v-card>
            </v-window-item>
            <v-window-item value="option-3">
              <v-card flat>
                <v-card-text>
                  <p>
                    Fusce a quam. Phasellus nec sem in justo pellentesque
                    facilisis. Nam eget dui. Proin viverra, ligula sit amet
                    ultrices semper, ligula arcu tristique sapien, a accumsan
                    nisi mauris ac eros. In dui magna, posuere eget, vestibulum
                    et, tempor auctor, justo.
                  </p>

                  <p class="mb-0">
                    Cras sagittis. Phasellus nec sem in justo pellentesque
                    facilisis. Proin sapien ipsum, porta a, auctor quis, euismod
                    ut, mi. Donec quam felis, ultricies nec, pellentesque eu,
                    pretium quis, sem. Nam at tortor in tellus interdum
                    sagittis.
                  </p>
                </v-card-text>
              </v-card>
            </v-window-item>
          </v-window>
        </v-card>
      </v-col>
    </v-row>

    <v-row>
      <v-col>
        <v-btn icon color="green">
          <v-icon large size="large">mdi-book-education-outline</v-icon>
        </v-btn>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12"> </v-col>
      <v-col>
        <v-btn color="green">
          <v-icon large>mdi-plus-circle-outline</v-icon>
          <span>Create Chapter</span>
        </v-btn>
        <v-md-editor v-model="courseChapter[indexChapter]" height="400px"></v-md-editor>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12">
        <v-btn color="green">
          <v-icon large>mdi-plus-circle-outline</v-icon>
          <span>More Chapter</span>
        </v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped></style>
