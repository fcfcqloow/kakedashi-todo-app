<script lang="ts" setup>
import Calender from '@/components/atoms/Calender.vue';
import DropDeleteIcon from '@/components/atoms/DropDeleteIcon.vue';
import MonthPicker from '@/components/atoms/MonthPicker.vue';
import PlusButton from '@/components/atoms/PlusButton.vue';
import RestoreButton from '@/components/atoms/RestoreButton.vue';
import SearchTextField from '@/components/atoms/SearchTextField.vue';
import SortButton from '@/components/atoms/SortButton.vue';
import YearPicker from '@/components/atoms/YearPicker.vue';
import RadioButtons from '@/components/atoms/RadioButtons.vue';
import TimeSelector from '@/components/atoms/TimeSelector.vue';

</script>

<template>
  <v-container>
    <v-row no-gutters justify="space-between"  style="padding-top: 30px; padding-bottom: 50px; border-top: solid">
      <v-col align-self="center">Time Selector</v-col>
      <v-col align-self="center"><time-selector /></v-col>
    </v-row>

    <v-row no-gutters justify="space-between"  style="padding-top: 30px; padding-bottom: 50px; border-top: solid">
      <v-col align-self="center">Calender</v-col>
      <v-col align-self="center"><calender :value="new Date()" /></v-col>
    </v-row>

    <v-row no-gutters justify="space-between" style="padding-top: 50px; padding-bottom: 50px; border-top: solid">
      <v-col align-self="center">Drop Delete Icon</v-col>
      <v-col align-self="center"><drop-delete-icon/></v-col>
    </v-row>

    <v-row no-gutters justify="space-between" style="padding-top: 50px; padding-bottom: 50px; border-top: solid">
      <v-col align-self="center">Year Picker</v-col>
      <v-col align-self="center"><year-picker /></v-col>
    </v-row>

    <v-row no-gutters justify="space-between" style="padding-top: 50px; padding-bottom: 50px; border-top: solid">
      <v-col align-self="center">Month Picker</v-col>
      <v-col align-self="center"><month-picker /></v-col>
    </v-row>

    <v-row no-gutters justify="space-between" style="padding-top: 50px; padding-bottom: 50px; border-top: solid">
      <v-col align-self="center">Plus Button</v-col>
      <v-col align-self="center"><plus-button /></v-col>
    </v-row>

    <v-row no-gutters justify="space-between" style="padding-top: 50px; padding-bottom: 50px; border-top: solid">
      <v-col align-self="center">Restore Button</v-col>
      <v-col align-self="center"><restore-button /></v-col>
    </v-row>

    <v-row no-gutters justify="space-between" style="padding-top: 50px; padding-bottom: 50px; border-top: solid">
      <v-col align-self="center">Restore Button</v-col>
      <v-col align-self="center"><sort-button /></v-col>
    </v-row>

    <v-row no-gutters justify="space-between" style="padding-top: 50px; padding-bottom: 50px; border-top: solid">
      <v-col align-self="center">Search Text Field</v-col>
      <v-col align-self="center"><search-text-field /></v-col>
    </v-row>

    <v-row no-gutters justify="space-between" style="padding-top: 50px; padding-bottom: 50px; border-top: solid">
      <v-col align-self="center">Search Text Field</v-col>
      <v-col align-self="center"><radio-buttons :labels="[1, 2, 3]" /></v-col>
    </v-row>

  </v-container>
</template>
