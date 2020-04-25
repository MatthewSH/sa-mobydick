const _ = require('lodash');

exports.getStopWordsAsArray = function(stopWordsText) {
    let stopWords = stopWordsText.split(new RegExp('[\n\r]+')).filter((word) => {
        return (word.indexOf('#') <= -1);
    });

    return stopWords;
}

exports.getMobyTextAsArray = function(mobyWordsText) {
    let mobyText = mobyWordsText.split(new RegExp('[ \n\r]+')).filter((word) => {
        return (word.length > 0);
    });

    let fixedWord;

    mobyText.forEach((word, index) => {
        //fixedWord = word.replace(/[.,\/#!$%\^&*;:{}=_`~()\"\“\[\]\?]/g, '');
        //fixedWord = word.replace('’s', '');

        word = word.replace(/'s/g, '').replace(/,/g, '').replace(/[.,\/#!$%\^&*;:{}=_`~()\"\“\[\]\?]/g, '');

        if (word.split('—').length > 1) {
            word.split('—').forEach(subWord => {
                mobyText.push(subWord);
            })

            mobyText.slice(index, 1);
        } else if (word.split('-').length > 1) {
            word.split('-').forEach(subWord => {
                mobyText.push(subWord);
            })

            mobyText.slice(index, 1);
        } else {
            mobyText[index] = word;
        }


        // if (fixedWord.split('—').length > 1) {
        //     mobyText.splice(index, 1);
        //     fixedWord.split('—').forEach(subWord => {
        //         mobyText.push(subWord);
        //     })
        // } else if (fixedWord.split('-').length > 1) {
        //     mobyText.splice(index, 1);
        //     fixedWord.split('-').forEach(subWord => {
        //         mobyText.push(subWord);
        //     })
        // } else {
        //     mobyText[index] = fixedWord;
        // }
    });

    return mobyText;
}

exports.removeStopWords = function(mobyTextArray, stopWordsArray) { 
    return mobyTextArray.filter((word) => (stopWordsArray.indexOf(word.toLowerCase()) <= -1));
}

exports.countUniqueWords = function(mobyTextArray) {
    let wordMap = new Map();

    mobyTextArray.forEach((word, index) => {
        word = word.toLowerCase().trim();

        if (word.length > 0) {
            if (wordMap.has(word)) {
                wordMap.set(word, (wordMap.get(word)+1));
            } else {
                wordMap.set(word, 1);
            }
        }
    });

    return wordMap;
}

exports.sortByHighest = function(wordMap) {
    return Array.from(wordMap).sort((first, second) => (second[1] - first[1]));
}