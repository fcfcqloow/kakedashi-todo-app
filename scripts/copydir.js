const fs = require('fs');
const path = require('path');

const inputPath = path.resolve(process.argv[2]);
const outputPath = path.resolve(process.argv[3]);

const copyFile = (inpath, outpath) => fs.readFile(inpath, {}, (err, data) => {
    console.debug(`Copy`, inpath, outpath);
    return err
        ? console.warn(err)
        : fs.writeFile(outpath, data, {}, (err) => err && console.warn(err));
});

const copyDir = (inpath, outpath) => {
    if (!fs.statSync(inpath).isDirectory()) {
        copyFile(inpath, outpath);
        return;
    }

    fs.mkdirSync(outpath, { recursive: true });
    fs.readdirSync(inpath).forEach((file) => 
        copyDir(path.join(inpath, file), path.join(outpath, file)),
    );
};

console.debug(`\ninputPath: ${inputPath}, \noutputPath: ${outputPath}\n`);
copyDir(inputPath, outputPath);
