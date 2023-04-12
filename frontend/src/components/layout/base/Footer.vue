<script setup>
import { useI18n } from "vue-i18n";
const { t } = useI18n();

const logoFileName = "";
</script>

<template>
   <v-footer app>
     <div class="px-4 py-2 bg-black text-center w-100">
        Copyright &copy; {{ new Date().getFullYear() }} - {{ $t("footer.cr") }}
     </div>
   </v-footer>
</template>
