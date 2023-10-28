<script setup lang="ts">
import type {
  Course,
  Chapter,
  SubChapter,
  NewChapterWithSubChapters,
} from "@/models/course/types";
import { useCoreCourseStore } from "@/stores/course/coreCourseStore";
import { useLoginStore } from "@/stores/session/loginStore";
import { reactive } from "vue";
import { computed } from "vue";
import { ref } from "vue";

const markdownEditorConfig = ref({
  leftToolbar:
    "default： undo redo clear | h bold italic strikethrough quote | ul ol table hr | link image code",
  rightToolbar: "preview toc sync-scroll",
});

const loginStore = useLoginStore();
const coreCourseStore = useCoreCourseStore();

const userId = computed(() => {
  return loginStore.getUserId;
});

const userName = computed(() => {
  return loginStore.authDetails.user?.nickname;
});

const courseDetails = reactive({} as Course);
courseDetails.author = userId.value;

const courseChapters = ref([] as Array<NewChapterWithSubChapters>);
courseChapters.value[0] = {} as NewChapterWithSubChapters;
courseChapters.value[0].subChapter = [] as Array<SubChapter>;
courseChapters.value[0].subChaptersCounter = 0;

const newChapter = ref({} as Chapter);
const newSubChapter = ref({} as SubChapter);

const indexChapter = ref(0);

function addNewSubChapter() {
  let currentPos = courseChapters.value[indexChapter.value].subChaptersCounter;
  courseChapters.value[indexChapter.value].subChapter[currentPos] =
    newSubChapter.value;
  courseChapters.value[indexChapter.value].subChaptersCounter += 1;
}

function addNewChapter() {
  indexChapter.value += 1;
  courseChapters.value[indexChapter.value] = {} as NewChapterWithSubChapters;
}
</script>

<template>
  <v-row>
    <v-col>
      <v-stepper :items="['Course Details', 'Chapters', 'Recheck']">
        <template v-slot:item.1>
          <v-card>
            <v-card-title>
              <span class="text-h5">Course Details</span>
            </v-card-title>
            <v-card-text>
              <v-container>
                <v-row dense>
                  <v-col cols="12" md="3">
                    <v-text-field
                      label="Erstellt"
                      variant="filled"
                      v-model="userName"
                      readonly
                      prepend-inner-icon="mdi-clock"
                    /> </v-col
                  ><v-col cols="12" md="9">
                    <v-text-field
                      label="Author"
                      variant="filled"
                      v-model="userName"
                      readonly
                      prepend-inner-icon="mdi-account"
                    />
                  </v-col>
                  <v-col cols="12">
                    <v-text-field
                      clearable
                      label="Course Title"
                      variant="outlined"
                      v-model="courseDetails.name"
                    ></v-text-field> </v-col
                  ><v-col cols="12">
                    <v-md-editor
                      :left-toolbar="markdownEditorConfig.leftToolbar"
                      :right-toolbar="markdownEditorConfig.rightToolbar"
                      v-model="courseDetails.description"
                      height="400px"
                    ></v-md-editor>
                  </v-col>
                  <v-col cols="6">
                    <v-card>
                      <v-card-title>
                        <span class="text-body-1">LEVEL</span>
                      </v-card-title>
                      <v-card-text align="center">
                        <v-rating
                          aria-placeholder="Level"
                          v-model="courseDetails.level"
                          :item-labels="['easy', '', '', '', 'hard']"
                          item-label-position="top"
                          color="error"
                          size="x-large"
                        />
                      </v-card-text>
                    </v-card>
                  </v-col>
                  <v-col>
                    <v-time-picker
                      v-model="courseDetails.duration"
                    ></v-time-picker>
                  </v-col>
                </v-row>
              </v-container>
            </v-card-text>
          </v-card>
        </template>

        <template v-slot:item.2>
          <v-card>
            <v-card-title>
              <span class="text-h5">Chapters</span>
            </v-card-title>
            <v-card-text>
              <v-row dense>
                <v-col cols="12">
                  <v-expansion-panels>
                    <v-expansion-panel
                      v-for="(item, pos) in courseChapters"
                      :key="pos"
                      :title="`${pos + 1}.) Chapter`"
                    >
                      <template v-slot:text>
                        <v-container class="">
                          <v-row>
                            <v-col cols="12">
                              <v-md-editor
                                :left-toolbar="markdownEditorConfig.leftToolbar"
                                :right-toolbar="
                                  markdownEditorConfig.rightToolbar
                                "
                                v-model="
                                  courseChapters[indexChapter].chapter
                                    .description
                                "
                                height="400px"
                              ></v-md-editor>
                              <v-btn
                                block
                                color="success"
                                @click="addNewSubChapter()"
                              >
                                Add SubChapter
                              </v-btn>
                            </v-col>
                          </v-row>

                          <v-row>
                            <v-divider></v-divider>
                          </v-row>
                          <v-row>
                            <v-col cols="12">
                              <v-container>
                                <span class="text-h5"> Subchapters</span>
                              </v-container>
                            </v-col>
                            <v-col cols="12">
                              <v-expansion-panels variant="inset" class="my-4">
                                <v-expansion-panel
                                  v-for="(subItem, subPos) in courseChapters[
                                    indexChapter
                                  ].subChapter"
                                  :key="subPos"
                                  :title="`${subPos + 1} SubChapter`"
                                >
                                  <template v-slot:text>
                                    <v-md-editor
                                      :left-toolbar="
                                        markdownEditorConfig.leftToolbar
                                      "
                                      :right-toolbar="
                                        markdownEditorConfig.rightToolbar
                                      "
                                      v-model="subItem.content"
                                      height="400px"
                                    ></v-md-editor>
                                  </template>
                                </v-expansion-panel>
                              </v-expansion-panels>
                            </v-col>
                          </v-row>
                        </v-container>
                      </template>
                    </v-expansion-panel>
                  </v-expansion-panels>
                </v-col>
              </v-row>
              <v-row>
                <v-col> </v-col>
              </v-row>
            </v-card-text>
            <v-card-actions> </v-card-actions>
          </v-card>
        </template>

        <template v-slot:item.3>
          <v-card>
            <v-card-title align="center">
              <v-container>
                <span class="text-h3">
                  {{ courseDetails.name }}
                </span>
              </v-container>
            </v-card-title>
            <v-card-text>
              <v-row class="mb-5">
                <v-col cols="12">
                  <span class="text-h5"> Description </span>
                </v-col>
                <v-col cols="12" md="9">
                  <v-rating
                    aria-placeholder="Level"
                    v-model="courseDetails.level"
                    item-label-position="bottom"
                    color="error"
                    readonly
                    size="small"
                  />
                </v-col>
                <v-col cols="12">
                  <v-card>
                    <v-card-text>
                      <v-md-editor
                        left-toolbar=""
                        right-toolbar=""
                        v-model="courseDetails.description"
                        mode="preview"
                      />
                    </v-card-text>
                  </v-card>
                </v-col>
              </v-row>

              <v-expansion-panels v-for="(details, dpos) in courseChapters">
                <v-expansion-panel>
                  <v-row v-for="(subchapter, cpos) in details.subChapter" :key="cpos">
                    <v-col cols="12">
                     {{ subchapter.content }}
                    </v-col>
                  </v-row>
                </v-expansion-panel>
              </v-expansion-panels>

              <!-- <v-expansion-panels>
                <v-row v-for="(chapter, cpos) in courseChapters">
                  <v-col cols="12" v-for="(chap, cpos) in chapter">
                    <v-expansion-panel
                      :title="`${cpos}.) `"
                      text="Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi, ratione debitis quis est labore voluptatibus! Eaque cupiditate minima"
                    >
                    </v-expansion-panel>
                  </v-col>
                </v-row>
              </v-expansion-panels> -->
            </v-card-text>
          </v-card>
        </template>
      </v-stepper>
    </v-col>
  </v-row>
</template>

<style scoped></style>
