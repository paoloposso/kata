/**
 * 
 * @param {[][]} a 
 */
function solution(a) {
    let result = [];

    for (let i = 0; i < a.length; i++) {
        let row = [];
        for (let x = a.length - 1; x >= 0; x--) {
            row.push(a[x][i]);
        }
        result.push(row);
    }

    return result;
}

console.log(solution(
    [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9]
    ]
));
