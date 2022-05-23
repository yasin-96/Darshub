<template>
  <v-container fill-height align="center" justify="center">
    <v-row>
      <v-col cols="12" sm="12" md="12" lg="8" xl="6" class="mx-auto">
        <v-card elevation="3" shaped>
          <v-card-text>
            <SignIn />
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import SignIn from "@/components/signin_up/SignIn";
export default {
  name: "LoginView",
  components: {
    SignIn,
  },
};
</script>

<style></style>
