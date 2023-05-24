

/**
 * 
 * @param {[number]} n 
 * @param {number} k 
 * @returns 
 */
function solution(array, k) {
    const pairs = [];

    for (let i = 0; i < array.length; i++) {
        for (let j = i + 1; j < array.length; j++) {
            pairs.push([array[i], array[j]]);
        }
    }

    let result = 0;

    for (let i = 0; i < pairs.length; i++) {
        if ((pairs[i][0] + pairs[i][1]) > k && (pairs[i][0] + pairs[i][1]) % k === 0) {
            result++;
        }
    }

    return result;
}

solution([70, 13, 9, 70, 46, 14], 5);