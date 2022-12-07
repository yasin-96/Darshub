<script setup lang="ts">
import { useLoginStore } from "@/stores/session/loginStore";
import { computed } from "vue";
import { useNow, useDateFormat } from "@vueuse/core";

const loginStore = useLoginStore();

const userInfo = computed(() => {
  return loginStore.getUser;
});

const formattedDate = (dateToConvert: string) => {
  return useDateFormat(useNow(), dateToConvert);
};
</script>

<template>
  <div
    class="mx-auto max-w-2xl py-16 px-4 sm:py-24 sm:px-6 lg:max-w-7xl lg:px-8"
  >
    <div
      class="grid sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-4 xl:grid-cols-6 gap-9"
    >
      <div
        class="sm:col-span-1 md:col-span-2 lg:col-span-2 xl:col-span-2 overflow-hidden shadow-xl rounded-lg"
      >
        <div class="px-4 py-5 sm:px-6">
          <h3 class="text-lg font-medium leading-6 text-gray-900">
            Personal Information
          </h3>
          <p class="mt-1 text-sm text-gray-500">
            <!-- TODO -->
          </p>
        </div>
        <div class="border-t border-gray-200">
          <dl>
            <div
              class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">Full name</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
                <span>
                  {{ userInfo.first_name }}
                </span>
                <span class="pl-2">{{ userInfo.last_name }}</span>
              </dd>
            </div>
            <div
              class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">Birthday</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
                <!-- {{ formattedDate(userInfo.birthday) }} -->
              </dd>
            </div>
            <div
              class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">Email address</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
                {{ userInfo.email }}
              </dd>
            </div>
            <div
              class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">Mobile</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
                {{ userInfo.telNr }}
              </dd>
            </div>
            <div
              class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">Gender</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
                {{ userInfo.telNr }}
              </dd>
            </div>
          </dl>
        </div>
      </div>

      <div
        class="sm:col-span-1 md:col-span-2 lg:col-span-2 xl:col-span-2 overflow-hidden shadow-xl rounded-lg"
      >
        <div class="px-4 py-5 sm:px-6">
          <h3 class="text-lg font-medium leading-6 text-gray-900">
            More Information
          </h3>
          <p class="mt-1 max-w-2xl text-sm text-gray-500">
            <!-- //TODO -->
          </p>
        </div>
        <div class="border-t border-gray-200">
          <dl>
            <div
              class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">company</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
                <span>
                  {{ userInfo.company }}
                </span>
              </dd>
            </div>
            <div
              class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">occupation</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
                {{ userInfo.occupation }}
              </dd>
            </div>
            <div
              class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">school</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
                {{ userInfo.school }}
              </dd>
            </div>
            <div
              class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">subject</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
                {{ userInfo.subject }}
              </dd>
            </div>
            <div
              class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">country</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
                {{ userInfo.country }}
              </dd>
            </div>
          </dl>
        </div>
      </div>

      <div
        class="sm:col-span-1 md:col-span-2 lg:col-span-2 xl:col-span-2 overflow-hidden shadow-xl rounded-lg"
      >
        <div class="px-4 py-5 sm:px-6">
          <h3 class="text-lg font-medium leading-6 text-gray-900">Roles</h3>
          <p class="mt-1 max-w-2xl text-sm text-gray-500">
            <!-- //TODO -->
          </p>
        </div>
        <div class="border-t border-gray-200">
          <dl>
            <div
              class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">Basic</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
                <i class="i-heroicons-check block"></i>
              </dd>
            </div>
            <div
              class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">Author</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
                <i class="i-heroicons-check block"></i>
              </dd>
            </div>
            <div
              class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">Manager</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
                <i class="i-heroicons-check block"></i>
              </dd>
            </div>
          </dl>
        </div>
      </div>

      <div
        class="sm:col-span-1 md:col-span-2 lg:col-span-3 xl:col-span-4 overflow-hidden shadow-xl rounded-lg"
      >
        <div class="px-4 py-5 sm:px-6">
          <h3 class="text-lg font-medium leading-6 text-gray-900">
            Account Information
          </h3>
          <p class="mt-1 max-w-2xl text-sm text-gray-500">
            <!-- //TODO -->
          </p>
        </div>
        <div class="border-t border-gray-200">
          <dl>
            <div
              class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium text-gray-500">Created</dt>
              <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0"></dd>
            </div>

            <div
              class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6"
            >
              <dt class="text-sm font-medium py-1 px-1">Logs</dt>
              <dd
                class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0 px-1 py-1"
              >
                <ul>
                  <li>asd</li>
                  <li>asd</li>
                  <li>asd</li>
                </ul>
              </dd>
            </div>
          </dl>
        </div>
      </div>

      <div
        class="sm:col-span-1 md:col-span-2 lg:col-span-1 xl:col-span-2 overflow-hidden shadow-xl rounded-lg"
      >
        <div class="px-4 py-5 sm:px-6">
          <h3 class="text-lg font-medium leading-6 text-gray-900">
            Dangers Zone
          </h3>
          <p class="mt-1 max-w-2xl text-sm text-gray-500">
            <!-- //TODO -->
          </p>
        </div>
        <div class="border-t border-gray-200">
          <dl>
            <div
              class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 hover:bg-red-300 text-center"
            >
              <dt class="text-sm font-medium py-1 px-1">Konto Deaktivieren</dt>
              <dd
                class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0 px-1 py-1 text-end"
              >
                <!-- todo -->
                <button
                  class="hover:bg-gray-200 rounded-lg py-1 px-1"
                  @click="() => {}"
                >
                  <i class="i-heroicons-archive-box block text-end"></i>
                </button>
              </dd>
            </div>
          </dl>
          <dl>
            <div
              class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 hover:bg-red-600 text-center"
            >
              <dt class="text-sm font-medium py-1 px-1">Konto loeschen</dt>
              <dd
                class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0 px-1 py-1 text-end"
              >
                <!-- todo -->
                <button
                  class="hover:bg-gray-200 rounded-lg py-1 px-1"
                  @click="() => {}"
                >
                  <i class="i-heroicons-trash block text-end"></i>
                </button>
              </dd>
            </div>
          </dl>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
