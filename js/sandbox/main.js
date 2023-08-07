/**
 * 
 * @param {[number]} arr 
 * @param {number} targetNum 
 */
function findTargetSumPair(arr, targetNum) {
    arr = arr.sort((a,b) => a < b);

    let left = 0;
    let right = arr.length-1;

    while (left < right) {
        const sum = arr[left] + arr[right];

        if (sum === targetNum) {
            return [arr[left], arr[right]];
        }

        if (sum < targetNum) {
            left++;
        } else {
            right--;
        }
    }

    return [-1, -1];
}

findTargetSumPair([1,2,4,6,8,9,10,20,23,27,31], 17);