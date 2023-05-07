const { exec } = require('child_process');
const path = require('path');
const appPath = path.resolve(process.argv[2]);
const command = `open ${appPath}`;

exec(command, (err, stdout, stderr) => {
  if (err) console.error(err);
});
