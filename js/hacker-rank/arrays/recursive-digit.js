/**
 * 
 * @param {string} n 
 * @param {number} k 
 */
function superDigit(n, k) {
    return superDigitRecur(n.repeat(k));
}

/**
 * 
 * @param {string} n 
 */
function superDigitRecur(n) {
    if (n.length === 1) return n;

    let res = 0;
    for (let i = 0; i < n.length; i++) {
        res += parseInt(n.charAt(i));
    }

    return superDigitRecur(res.toString());
}

console.log(superDigit('9875', 4));