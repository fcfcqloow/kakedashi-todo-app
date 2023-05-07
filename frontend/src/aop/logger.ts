const isDev = true;
export const debug = (...args: any) => {
  if (isDev) console.debug('[DEBUG]:', ...args);
};
