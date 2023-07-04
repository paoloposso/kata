/**
 * 
 * @param {[string]} crypt 
 * @param {[][]} solution 
 */
function solution(crypt, solution) {
    const solutionMap = new Map(solution.map(item => [item[0], item[1]]));

    let val1 = getCryptoValue(crypt[0], solutionMap);
    let val2 = getCryptoValue(crypt[1], solutionMap);
    let val3 = getCryptoValue(crypt[2], solutionMap);

    if ((val1.length > 1 && hasLeadingZeros(val1)) 
        || (val2.length > 1 && hasLeadingZeros(val2))
        || (val3.length > 1 && hasLeadingZeros(val3))) {
        return false;
    }

    return parseInt(val1) + parseInt(val2) === parseInt(val3);
}

/**
 * 
 * @param {string} word 
 * @param {Map} solutionMap 
 */
function getCryptoValue(word, solutionMap) {
    let result = '';
    for (let i = 0; i < word.length; i++) {
        result += solutionMap.get(word.charAt(i));
    }
    return result;
}

function hasLeadingZeros(numericString) {
    // Regex pattern to match numeric string with leading zeros
    const regex = /^0/;
    
    // Test if the numeric string matches the regex pattern
    return regex.test(numericString);
}

let crypt = 
["WASD", 
 "IJKL", 
 "AOPAS"]
let sol= 
[["W","2"], 
 ["A","0"], 
 ["S","4"], 
 ["D","1"], 
 ["I","5"], 
 ["J","8"], 
 ["K","6"], 
 ["L","3"], 
 ["O","7"], 
 ["P","9"]]

 solution(crypt, sol);