import { defineStore } from 'pinia'
import { Settings } from '@/core/Settings';
import { copyObjectKV } from '@/util/object';
import { emptyFunc } from '@/util/func';

export const useAppStore = defineStore('app', {
  state: () => {
    return {
      snackbar : {
        isShow : false,
        value  : '',
        color  : '',
      },
      confirm : {
        isShow : false,
        title  : 'Confirm',
        ok     : emptyFunc,
        cancel : emptyFunc,
      },
      isGlobalLoading : false,
      settings        : <Settings>{},
      calenderDialog  : {
        initDate     : new Date(),
        isShow       : false,
        dateCallback : (date: Date) => {},
      },
    }
  },
  actions : {
    showCalender (initDate?: Date, dateCallback?: (date: Date) => void) {
      initDate && (this.calenderDialog.initDate = new Date(initDate));
      dateCallback && (this.calenderDialog.dateCallback = dateCallback);
      this.calenderDialog.isShow = true;
    },
    closeCalender () {
      this.calenderDialog.isShow = false;
      this.calenderDialog.dateCallback = emptyFunc;
    },
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
    showConfirmDialog (args: { title: string, ok?: () => void; cancel?: () => void; }) {
      this.confirm.isShow = true;
      this.confirm.title = args.title;
      this.confirm.ok = args.ok ?? emptyFunc;
      this.confirm.cancel = args.cancel ?? emptyFunc;
    },
    closeConfirmDialog() {
      this.confirm.isShow = false;
      setTimeout(() => this.confirm.title = 'Confirm', 150);
      this.confirm.ok = emptyFunc;
      this.confirm.cancel = emptyFunc;
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
