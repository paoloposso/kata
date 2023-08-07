/**
 * 
 * @param {*} l 
 */
function solution(l) {
    const arr = [];

    let cur = l;

    while (cur) {
        arr.push(cur.value);

        cur = cur.next;
    }

    let left = 0;
    let right = arr.length-1;

    while (left < right) {
        if (arr[left] !== arr[right]) return false;
        left++;
        right--;
    }

    return true;
}


let ll = {
    value: 0,
    next: {
        value: 1,
        next: {
            value: 0,
            next: null
        }
    }
};

solution(ll);