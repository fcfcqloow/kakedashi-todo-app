import Koa from 'koa';
import Router from 'koa-router';
import bodyParser from 'koa-bodyparser';
import fs from 'fs';

const koaApp = new Koa();
const router = new Router();
const port = 3000;

export const runKoaServer = () => koaApp
  .use(bodyParser())
  .use(router // TODO...
    .routes(),
  )
  .listen(port, undefined, undefined, () => console.log('Koa server listening on port ' + port));

