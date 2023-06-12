
/**
 * 
 * @param {[number]} arr 
 */
const quickSort = (arr) => {
    const result = quickSortRecur(arr, 0, arr.length);

    return result;
}

/**
 * 
 * @param {[number]} arr 
 * @param {number} start 
 * @param {number} end 
 */
const quickSortRecur = (arr, start, end) => {

    if (start >= end) return arr;

    let pivotPosition = start;
    console.log(`pivot value ${arr[pivotPosition]}`);

    for (let i = start; i < end; i++) {
        if (arr[i] < arr[pivotPosition]) {
            [arr[pivotPosition], arr[i]] = [arr[i], arr[pivotPosition]];
            pivotPosition = i;
        }
    }

    arr = quickSortRecur(arr, 0, pivotPosition);
    arr = quickSortRecur(arr, pivotPosition+1, end);

    return arr;
}

console.log(quickSort([55,8,1,20,4,2,18,22,29,30]));