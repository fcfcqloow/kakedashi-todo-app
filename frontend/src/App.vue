<script setup lang="ts">
import AppSnackbar from '@/components/organisms/app/AppSnackbar.vue';
import AppLoadingOverlay from '@/components/organisms/app/AppLoadingOverlay.vue';
import AppSettingsLoader from './components/organisms/app/AppSettingsLoader.vue';
import { useAppStore } from './stores/app';
import AppNotification from './components/organisms/app/AppNotification.vue';
import AppConfirmDialog from './components/organisms/app/AppConfirmDialog.vue';
import AppCalenderDialog from './components/organisms/app/AppCalenderDialog.vue';

const { settings } = useAppStore();
</script>

<template>
  <app-settings-loader />
  <app-loading-overlay />
  <app-confirm-dialog />
  <app-notification />
  <app-snackbar />
  <app-calender-dialog />
  <v-app :theme="settings.mode === 'dark' ? 'dark' : 'white'">
    <router-view />
  </v-app>
</template>
