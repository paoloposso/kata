/**
 * 
 * @param {[number]} a 
 */
function solution(a) {
    const s = new Map();
    let result = -1;

    for (let i = 0; i < a.length; i++) {
        if (s.has(a[i])) {
            if (i < result || result < 0) {
                result = i;
            }
        } else {
            s.set(a[i], i);
        }
    }

    return a[result] || -1;
}

solution([2, 1, 3, 5, 3, 2]);