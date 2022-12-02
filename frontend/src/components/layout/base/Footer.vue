<script setup>
import { useI18n } from "vue-i18n";
const { t } = useI18n();

const logoFileName = "";
</script>

<template>
  <footer class="p-10 bg-slate-500">
    <div class="">
      <div class="avatar">
        <div class="w-24 rounded-full">
          <img width="128" src="logoUrl" alt="Logo" />
        </div>
      </div>
      <p>
        Copyright &copy; {{ new Date().getFullYear() }} - {{ $t("footer.cr") }}
      </p>
    </div>
    <div></div>
  </footer>
</template>
