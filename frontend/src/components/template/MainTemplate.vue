<script setup lang="ts">
import { ref } from 'vue';
import { useAppStore } from '../../stores/app';

type Props = {
  title: string;
};

const props = defineProps<Props>();
const drawer = ref(false);
const { settings } = useAppStore();
</script>

<template>
  <v-layout class="height100">
    <v-app-bar :color="settings.myColor" >
      <v-app-bar-nav-icon variant="text" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>{{ props.title }}</v-toolbar-title>
      <slot name="header" />
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" temporary :color="settings.myColor">
      <v-list nav>
        <v-list-item prepend-icon="mdi-home" title="Todo List" value="TodoList" @click="$router.push('/')" />
        <v-list-item disabled prepend-icon="mdi-view-dashboard" title="Dashboard" value="Dashboard" @click="" />
        <v-list-item prepend-icon="mdi-monitor-dashboard" title="Summary" value="Summary" @click="$router.push('/summary')" />
        <v-list-item prepend-icon="mdi-message" title="Topic Editor" value="TopicEditor" @click="$router.push('/topic/editor')" />
        <v-list-item prepend-icon="mdi-console" title="Log Monitor" value="Log" @click="$router.push('/log')" />
        <v-list-item prepend-icon="mdi-backup-restore" title="Backup" value="Backup" @click="$router.push('/restore')" />
        <v-list-item prepend-icon="mdi-cog" title="Settings" value="Settings" @click="$router.push('/settings')" />
      </v-list>
    </v-navigation-drawer>
    <v-main>
      <slot></slot>
    </v-main>
  </v-layout>
</template>

<style>
.height100 {
  height: 100%;
}
</style>
