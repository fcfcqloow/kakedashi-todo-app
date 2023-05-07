<script lang="ts" setup>
import { Topic } from '@/core/Topic';

type Props = {
  topics: Array<Topic>;
  height: string;
  width: string;
};

type Emit = {
  (e: 'remove', topic: Topic): void;
};

const props = defineProps<Props>();
const emits = defineEmits<Emit>();
</script>

<template>
  <v-table
    :height="props.height"
    :width="props.width"
    fixed-header
  >
    <thead>
      <tr>
        <th class="text-left">Name</th>
        <th class="text-left">Created</th>
        <th class="text-left">Color</th>
        <th style="width: 180px; padding: 10px;" class="text-left">Action</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="item in props.topics" :key="item.id">
        <td>{{ item.value }}</td>
        <td>{{ item.created.toLocaleString() }}</td>
        <td :style="`background-color: ${item.color};`">{{ item.color }}</td>
        <td>
          <v-row>
            <v-col cols="4">
              <v-btn icon="mdi-check" color="green" size="35"></v-btn>
            </v-col>
            <v-col cols="4">
              <v-btn icon="mdi-pencil" color="yellow" size="35"></v-btn>
            </v-col>
            <v-col cols="4">
              <v-btn icon="mdi-delete" color="red" size="35" @click="emits('remove', item)"></v-btn>
            </v-col>
          </v-row>
        </td>
      </tr>
    </tbody>
  </v-table>
</template>