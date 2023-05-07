import { onMounted } from 'vue';
<script lang="ts" setup>
import { onMounted } from 'vue';

type Props = {
  logs: string[];
};

type Emit = {
  (e: 'on:mount'): void;
};

const props = defineProps<Props>();
const emits = defineEmits<Emit>();
onMounted(() => {
  emits('on:mount');
});
</script>

<template>
  <div v-for="item in props.logs.filter(Boolean)">
    <div v-for=" v in Object.entries(JSON.parse(item))" style="padding: 20px">
      <v-row>
        <v-col cols="2">{{ v[0] }}</v-col>
        <v-col><pre><code>{{ v[1] }}</code></pre></v-col>
      </v-row>
    </div>
    <v-divider></v-divider>
  </div>
</template>

<style>
.json-viewer {
  display: flex;
  flex-direction: row;
  height: 100%;
}

.json-string {
  font-family: monospace;
  overflow: auto;
  height: 100%;
  width: 50%;
  padding: 20px;
  margin: 0;
  box-sizing: border-box;
  border-right: 1px solid gray;
}

.json-tree {
  width: 50%;
  height: 100%;
  overflow: auto;
  padding: 20px;
  margin: 0;
  box-sizing: border-box;
  font-size: 16px;
  font-family: Arial, Helvetica, sans-serif;
}
</style>