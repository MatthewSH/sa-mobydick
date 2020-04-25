let lib = require('./lib');
let fs = require('fs');

let mobyText = fs.readFileSync('./data/mobydick.txt', 'utf8');
let stopWordsText = fs.readFileSync('./data/stop-words.txt', 'utf8');

let stopWordsArray = lib.getStopWordsAsArray(stopWordsText);
let mobyTextArray = lib.getMobyTextAsArray(mobyText);
mobyTextArray = lib.removeStopWords(mobyTextArray, stopWordsArray);
let counted = lib.countUniqueWords(mobyTextArray);
let sorted = lib.sortByHighest(counted);

console.log(sorted[0]);