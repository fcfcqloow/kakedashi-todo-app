<script setup lang="ts">
import AppSnackbar from '@/components/organisms/app/AppSnackbar.vue';
import AppLoadingOverlay from '@/components/organisms/app/AppLoadingOverlay.vue';
import AppSettingsLoader from './components/organisms/app/AppSettingsLoader.vue';
import { useAppStore } from './stores/app';
import AppNotification from './components/organisms/app/AppNotification.vue';

const { settings } = useAppStore();
</script>

<template>
  <app-notification />
  <app-settings-loader />
  <app-loading-overlay />
  <app-snackbar />
  <v-app :theme="settings.mode === 'dark' ? 'dark' : 'white'">
    <router-view />
  </v-app>
</template>
