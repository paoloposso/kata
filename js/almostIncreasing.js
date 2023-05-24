
module.exports.isAlmostIncreasing = function isAlmostIncreasing(sequence) {
    let start = itsIncreasingUntil(sequence);

    for (let i = start; i < sequence.length; i++) {
        let isSorted = isIncreasing(sequence, i);

        if (isSorted) return true;
    }
    return false;
}

function isIncreasing(arr, excepIx) {
    for (let i = 0; i < arr.length; i++) {
        if (i === excepIx) continue;
        let next = i+1 === excepIx ? i+2 : i+1;
        if (arr[i] >= arr[next]) {
            return false;
        }
    }

    return true;
}

function itsIncreasingUntil(arr) {
    let result = 0;
    
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] >= arr[i+1]) {
            return i;
        }
        result = i;
    }

    return result;
}
