
/**
 * 
 * @param {[number]} arr 
 */
module.exports.quickSort = quickSort = (arr) => {
    if (arr.length < 1) return arr;
    const result = quickSortRecur(arr, 0, arr.length-1);

    return arr;
}

/**
 * 
 * @param {[number]} arr 
 * @param {number} start 
 * @param {number} end 
 */
const quickSortRecur = (arr, start, end) => {

    if (start >= end) return;

    let pivotNeedle = start;
    let pivotValue = arr[start];

    for (let j = start; j <= end; j++) {
        if (arr[j] < pivotValue) {
            pivotNeedle++;

            let a = arr[pivotNeedle];
            arr[pivotNeedle] = arr[j];
            arr[j] = a;

            // [arr[pivotNeedle], arr[j]] = [arr[j], arr[pivotNeedle]];
        }
    }

    [arr[start], arr[pivotNeedle]] = [arr[pivotNeedle], arr[start]];
    
    quickSortRecur(arr, start, pivotNeedle-1);
    quickSortRecur(arr, pivotNeedle+1, end);
}

console.log(quickSort([5,18,1,23,44,8,2]));