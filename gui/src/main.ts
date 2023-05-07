import path from 'path';
import { app, BrowserWindow } from 'electron';
import { runKoaServer } from './koa-server';


let win: BrowserWindow;
app.whenReady().then(() => {
  win = new BrowserWindow()
  if (process.env.VITE_DEV_SERVER_URL) {
    win.loadURL(process.env.VITE_DEV_SERVER_URL);
  } else {
    win.loadFile(path.join(__dirname, '../dist-frontend/index.html'));
  }

  runKoaServer();
});

