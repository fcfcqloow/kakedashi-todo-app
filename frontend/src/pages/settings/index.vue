<script setup lang="ts">
import {  getCurrentInstance  } from 'vue';
import { ISettingsRepository } from '@/core/IF/repository';
import { useAppStore } from '@/stores/app';
import { Settings } from '@/core/Settings';
import MainTemplate from '@/components/template/MainTemplate.vue';
import SettingsForm from '@/components/organisms/setting/SettingsForm.vue';
import { Color } from '@/core/Color';

const { settings, setSettings, showSnackbar } = useAppStore();
const app = getCurrentInstance();
const $settingsRepository: ISettingsRepository = app?.appContext.config.globalProperties?.$settingsRepository;

const onSubmit = (newSettings: Settings) => {
  $settingsRepository.update(newSettings)
    .then(() => {
      setSettings(newSettings);
      showSnackbar('Success', Color.SUCCESS);
    })
    .catch(() => {
      showSnackbar('Failure', Color.ERROR);
    });
};
</script>

<template>
  <main-template title="Settings">
    <v-container fluid>
      <settings-form
        :settings="settings"
        @submit="onSubmit"
      />
    </v-container>
  </main-template>
</template>