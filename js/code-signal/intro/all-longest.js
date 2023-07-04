
/**
 * 
 * @param {[string]} inputArray 
 */
function solution(inputArray) {
    
    const sorted = inputArray.sort(function(a, b) {
        return b.length - a.length;
    });

    const maxLength = sorted[0].length;

    const result = [];

    for (let i = 0; i < sorted.length; i++) {
        if (sorted[i].length < maxLength) break;
        
        result.push(sorted[i]);
    }

    return result;
}

let res = solution(["aba", "aa", "ad", "vcd", "aba"]);
res = solution(["aa"]);

console.log(res);