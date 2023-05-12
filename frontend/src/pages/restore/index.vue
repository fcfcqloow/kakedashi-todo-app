<script setup lang="ts">
import { getCurrentInstance, ref, onMounted, reactive } from 'vue';
import { ILogsRepository } from '@/core/IF/repository';
import LogsExpandList from '@/components/organisms/log/LogsExpandList.vue';
import MainTemplate from '@/components/template/MainTemplate.vue';

</script>

<template>
  <main-template title="Logs Monitor">
  </main-template>
</template>