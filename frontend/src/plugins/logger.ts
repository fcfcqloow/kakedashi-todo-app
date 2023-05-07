import type { App, Plugin } from 'vue';
import axios from 'axios';
import { debug } from '@/aop/logger';

export const RepositoryPlugin: Plugin = {
  install: (app: App, options?: any) => {
    axios.interceptors.request.use(
      (request) => {
        debug(request.url, request.method);
        return Promise.resolve(request);
      });
    
  },
};

export default RepositoryPlugin;
