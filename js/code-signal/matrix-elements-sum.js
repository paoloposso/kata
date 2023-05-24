
/**
 * 
 * @param {[[number]]} matrix 
 */
function solution(matrix) {
    let emptiesMap = new Map();
    for (let line = 0; line < matrix.length; line++) { //check lines
        for (let column=0; column < matrix[line].length; column++) {
            if (matrix[line][column] === 0 && !emptiesMap.has(column)) {
                emptiesMap.set(column, line);
            }
        }
    }

    let total = 0;

    for (let line = 0; line < matrix.length; line++) { //check lines
        for (let column=0; column < matrix[line].length; column++) {
            console.log(`${line}x${column}`);
            if (emptiesMap.has(column) && emptiesMap.get(column) < line) {
                continue;
            }
            total += matrix[line][column];
        }
    }

    return total;
}

module.exports.solution = solution;