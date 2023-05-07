<script lang="ts" setup>
import MainTemplate from '@/components/template/MainTemplate.vue';
import { ref } from 'vue';

const years = ref([
  {
    color: 'cyan',
    year: '1960',
  },
  {
    color: 'green',
    year: '1970',
  },
  {
    color: 'pink',
    year: '1980',
  },
  {
    color: 'amber',
    year: '1990',
  },
  {
    color: 'orange',
    year: '2000',
  },
]);
</script>

<template>
  <main-template title="Summary">
    <v-container fluid>
      <v-sheet height="calc(100vh - 100px)" style="overflow: scroll;">
        <v-timeline align="start">
          <v-timeline-item>
            WIP
          </v-timeline-item>
        </v-timeline>
      </v-sheet>
    </v-container>
  </main-template>
</template>