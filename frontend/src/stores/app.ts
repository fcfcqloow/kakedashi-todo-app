import { defineStore } from 'pinia'
import { Settings } from '@/core/Settings';
import { copyObjectKV } from '@/util/object';

export const useAppStore = defineStore('app', {
  state: () => {
    return {
      snackbar : {
        isShow : false,
        value  : '',
        color  : '',
      },
      isGlobalLoading: false,
      settings : <Settings>{
        mode                    : 'dark',
        taskLimit               : 1000,
        sortable                : true,
        topicColorFlag          : true,
        notificationIntervalSec : 30,
        topicListType           : 'list',
      },
    }
  },
  actions : {
    showSnackbar (text: string, color?: string) {
      this.snackbar.value = text;
      this.snackbar.isShow = true;
      this.snackbar.color = color ?? '';
    },
    closeSnackbar () {
      this.snackbar.isShow = false;
      this.snackbar.value  = '';
      this.snackbar.color = '';
    },
    setGlobalLoading (value: boolean) {
      this.isGlobalLoading = value;
    },
    subscribeGlobalLoadingPromise(promise?: Promise<any>, callback?: Function) {
      this.isGlobalLoading = true;
      promise?.finally(() => {
        this.isGlobalLoading = false;
        callback && callback();
      });
    },
    setSettings(newSettings: Settings) {
      copyObjectKV(this.settings, newSettings);
    }
  },
});
