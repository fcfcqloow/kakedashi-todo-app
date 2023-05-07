require('fs').rm(process.argv[2], { recursive: true }, (err) => {
    if (err && err.errno !== -2) throw err;
});
