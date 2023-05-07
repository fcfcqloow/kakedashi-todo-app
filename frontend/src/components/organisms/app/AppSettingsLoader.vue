<script setup lang="ts">
import { onMounted, getCurrentInstance } from 'vue';
import { ISettingsRepository } from '@/core/IF/repository';
import { useAppStore } from '@/stores/app';
import { debug } from '@/aop/logger';

const app = getCurrentInstance();
const $settingsRepository: ISettingsRepository= app?.appContext.config.globalProperties?.$settingsRepository;
const { setSettings } = useAppStore();

onMounted(() => {
  $settingsRepository.get().then(response => {
    debug('Settings', JSON.stringify(response));
    setSettings(response);
  });
});
</script>

<template></template>

