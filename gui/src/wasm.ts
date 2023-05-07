import fs from 'fs';
import util from 'util';
import crypto from 'crypto';

const globalForWasm: typeof globalThis & Record<string, any> = globalThis;

globalForWasm.require = require;
globalForWasm.fs = fs;
globalForWasm.TextEncoder = util.TextEncoder;
globalForWasm.TextDecoder = util.TextDecoder;
globalForWasm.performance = {
	now() {
		const [sec, nsec] = process.hrtime();
		return sec * 1000 + nsec / 1000000;
	},
} as any;
globalForWasm.crypto = {
	getRandomValues : crypto.getRandomValues,
	subtle          : crypto.subtle,
	randomUUID      : crypto.randomUUID,
};
require('./wasm_exec');



const writeTodo = (value) => {
	console.log(value);
	fs.writeFileSync('/Users/kiyotakazuhiro/.todoapp/data/todo.db', value);
};
const go = new globalForWasm.Go();
go.env = Object.assign({ TMPDIR: require("os").tmpdir() }, process.env);
go.exit = process.exit;
WebAssembly.instantiate(
    fs.readFileSync('./dist/main.wasm'),
    go.importObject,
)
    .then(async (result) => {
		go.run(result.instance);
		const response = globalForWasm.addTask(JSON.stringify({
			id         : '1234',
			value      : 'danskcaio',
			deadLine   : 123456,
			createDate : 12345,
		}));
		writeTodo(response);

		
    });