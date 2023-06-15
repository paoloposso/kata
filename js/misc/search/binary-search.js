/**
 * 
 * @param {[number]} arr 
 * @param {number} value 
 */
module.exports.search = search;

/**
 * 
 * @param {[number]} arr 
 * @param {number} value 
 * @returns 
 */
function search(arr, value) {
    arr = arr.sort();
    return searchRecur(arr, value, 0, arr.length);
}

const searchRecur = (arr, value, start, end) => {
    if (start > end) return false;

    const ix = start + Math.round((end-start)/2);

    if (arr[ix] === value) return true;

    if (value < arr[ix]) {
        return searchRecur(arr, value, start, ix-1);
    }
    return searchRecur(arr, value, ix+1, end);
}

console.log(search([1,2,3,5,6,7,9,10,15,18], 2));