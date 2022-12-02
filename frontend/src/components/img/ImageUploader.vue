<script lang="ts" setup>
import { reactive, computed, onMounted, onUnmounted } from "vue";
import { useRegistrationStore } from "@/stores/registrationStore";

const state = reactive({
  maxFileSize: 2 * 1024 * 1024,
  hasError: false,
  displayErrorMessage: "",
  fileReader: new FileReader(),
  maxWidth: 300,
  maxHeight: 150,
  allowedImagesTypes: ["image/png", "image/jpeg", "image/jpg", "image/svg"],
});

const registrationStore = useRegistrationStore();

const checkFileSize = (fileSizeToCheck: number | string): boolean => {
  if (
    fileSizeToCheck == null ||
    fileSizeToCheck == 0 ||
    fileSizeToCheck == ""
  ) {
    state.hasError = true;
    state.displayErrorMessage = `File is empty!`;
    return false;
  } else if (fileSizeToCheck > state.maxFileSize) {
    state.hasError = true;
    state.displayErrorMessage = `Max file size is 2MB!`;
    return false;
  } else {
    state.hasError = false;
    return true;
  }
};
const checkFileType = (typeToCheck: any) => {
  console.log("this.allowedTypes", allowedTypes.value);
  console.log("typeToCheck", typeToCheck);
  const result = state.allowedImagesTypes.includes(typeToCheck);
  console.log("result", result);
  if (!result) {
    state.hasError = true;
    state.displayErrorMessage = `File type [${
      typeToCheck.split("/")[1]
    }] is not allowed.`;

    return false;
  } else {
    return true;
  }
};
const selectFile = (newAvatarFile: Event) => {
  console.log("newAvatarFile", newAvatarFile)
  if (newAvatarFile) {
    const { size, type } = newAvatarFile;
    console.table({ size, type });
    console.log("state.maxFileSize", state.maxFileSize);
    if (checkFileType(type)) {
      if (checkFileSize(size)) {
        state.fileReader.readAsText(newAvatarFile);
      }
    }
  }
};

const resetInput = () => {
  avatar.value = null;
  state.hasError = false;
};

const avatar = computed({
  get() {
    return registrationStore.getAvatar;
  },
  set(newValue: string | null): void {
    registrationStore.setAvatar(newValue);
  },
});

const allowedTypes = computed(() => {
  return state.allowedImagesTypes.map((t) => t.split("image/")[1]).toString();
});

// onMounted(() => {
//   state.fileReader.addEventListener(
//     "load",
//     () => {
//       registrationStore.setAvatar(state.fileReader.result?.toString());

//       // avatar.value = state.fileReader.result;
//     },
//     false
//   );
// });

onUnmounted(() => {
  Object.assign({}, state.fileReader);
});
</script>

<template>
  <!-- <v-container class="pl-0 pr-0 mr-0">
    <v-row>
      <v-col cols="12" sm="12" md="6" lg="6" xl="6">
        <v-file-input
          show-size
          label="Avatar "
          @change="selectFile"
          @input="hasError = false"
          prepend-icon="mdi-badge-account"
          @click:clear="resetInput()"
          :error="hasError"
          :error-messages="hasError ? displayErrorMessage : ''"
          :persistent-hint="true"
          :hint="allowedTypes"
        ></v-file-input>
      </v-col>
      <v-col cols="12" sm="12" md="6" lg="6" xl="6">
        <v-card
          class="mx-auto"
          :width="maxWidth"
          :height="maxHeight"
          v-if="avatar"
        >
          <v-img
            :src="avatar"
            :max-width="maxWidth"
            :max-height="maxHeight"
          ></v-img>
        </v-card>
        <v-skeleton-loader
          class="mx-auto"
          v-else
          :max-width="maxWidth"
          :max-height="maxHeight"
          type="card"
        ></v-skeleton-loader>
      </v-col>
    </v-row>
  </v-container> -->
  <div>
    <form class="flex">
      <div class="shrink-0">
        <img
          v-if="avatar"
          class="h-16 w-16 object-cover rounded-full overflow-visible"
          src="avatar"
          alt="Profile"
        />
        <img else class="h-16 w-16 object-cover rounded-full" alt="Profile" />
      </div>
      <label class="block">
        <input
          type="file"
          class="rounded-lg form-input block w-full text-sm text-slate-500 file:mr-4 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold file:bg-violet-50 file:text-blue-500 hover:file:bg-violet-100"
          @change="selectFile"
        />
        <span v-if="!!state.displayErrorMessage" class="text-sm text-red-500">{{
          state.displayErrorMessage
        }}</span>
        <span v-else class="text-sm text-gray-500">
          Erlaubte formate [{{ allowedTypes }}]
        </span>
      </label>
    </form>
  </div>
</template>

<style></style>
