
/**
 * 
 * @param {string} s 
 */
function solution(s) {

    const seen = [];

    for (let i = 0; i < s.length; i++) {
        if (seen.includes(s[i])) continue;
        
        seen.push(s[i]);
        if (s.slice(i+1).indexOf(s[i]) < 0) {
            return s[i];
        }
    }

    return '_';
}

console.log(solution('abacabad'));