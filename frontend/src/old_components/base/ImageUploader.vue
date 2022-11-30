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
    Current profile photo Choose profile photo

    <form class="flex items-center space-x-6">
      <div class="shrink-0">
        <img
          class="h-16 w-16 object-cover rounded-full"
          src="https://images.unsplash.com/photo-1580489944761-15a19d654956?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1361&q=80"
          alt="Current profile photo"
        />
      </div>
      <label class="block">
        <span class="sr-only">Choose profile photo</span>
        <input
          type="file"
          class="block w-full text-sm text-slate-500 file:mr-4 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold file:bg-violet-50 file:text-violet-700 hover:file:bg-violet-100"
        />
      </label>
    </form>
  </div>
</template>
//
<script>
import { mapState } from "vuex";
export default {
  name: "ImageUploader",
  data: () => ({
    maxFileSize: 2 * 1024 * 1024,
    hasError: false,
    displayErrorMessage: "",
    fileReader: null,
    maxWidth: 300,
    maxHeight: 150,
  }),
  created() {
    this.fileReader = new FileReader();
    this.fileReader.addEventListener(
      "load",
      () => {
        this.avatar = this.fileReader.result;
      },
      false
    );
  },
  destroyed() {
    this.fileReader = null;
  },
  methods: {
    checkFileSize(fileSizeToCheck) {
      if (
        fileSizeToCheck == null ||
        fileSizeToCheck == 0 ||
        fileSizeToCheck == ""
      ) {
        this.hasError = true;
        this.displayErrorMessage = `File is empty!`;
        return false;
      } else if (fileSizeToCheck > this.maxFileSize) {
        this.hasError = true;
        this.displayErrorMessage = `Max file size is 2MB!`;
        return false;
      } else {
        this.hasError = false;
        return true;
      }
    },
    checkFileType(typeToCheck) {
      console.log("this.allowedTypes", this.allowedTypes);
      console.log("typeToCheck", typeToCheck);
      const result = this.allowedImagesTypes.includes(typeToCheck);
      console.log("result", result);
      if (!result) {
        this.hasError = true;
        this.displayErrorMessage = `File type [${
          typeToCheck.split("/")[1]
        }] is not allowed.`;

        return false;
      } else {
        return true;
      }
    },
    selectFile(newAvatarFile) {
      if (newAvatarFile) {
        const { size, type } = newAvatarFile;
        console.log(size, type);
        console.log(this.maxFileSize);
        if (this.checkFileType(type)) {
          if (this.checkFileSize(size)) {
            this.fileReader.readAsDataURL(newAvatarFile);
          }
        }
      }
    },
    resetInput() {
      this.avatar = null;
      this.hasError = false;
    },
  },
  computed: {
    ...mapState({
      allowedImagesTypes: (state) => state.types.images,
      newUser: (state) => state.registry.newUser,
    }),
    avatar: {
      get() {
        return this.newUser.avatar;
      },
      set(newValue) {
        this.$store.dispatch("registry/act_setProfileImage", newValue);
      },
    },
    allowedTypes() {
      // return this.allowedImagesTypes.map((t) => t.split("image/")[1]).toString();
      return this.allowedImagesTypes
        .map((t) => t.split("image/")[1])
        .toString();
    },
  },
};
</script>

<style></style>
