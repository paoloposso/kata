/**
 * 
 * @param {[][]} grid 
 */
function solution(grid) {
    for (let i = 0; i < grid.length; i++) {
        if (!checkSequence(grid[i])) {
            return false;
        }

        const colArr = [];

        for (let j = 0; j < grid.length; j++) {
            colArr.push(grid[j][i]);
        }

        if (!checkSequence(colArr)) {
            return false;
        }
    }

    return checkSubMatrices(grid);
}

function checkSubMatrices(grid) {
    for (let i = 0; i <= 6; i+=3) {
        if (!checkSequence(getSubmatrixAsArray(grid, i, 0))) return false;
        if (!checkSequence(getSubmatrixAsArray(grid, i, 3))) return false;
        if (!checkSequence(getSubmatrixAsArray(grid, i, 6))) return false;
    }

    return true;
}

/**
 * 
 * @param {[][]} matrix 
 * @param {number} startRow 
 * @param {number} startCol 
 * @returns 
 */
function getSubmatrixAsArray(matrix, startRow, startCol) {
    const result = [];
    
    for (let i = startRow; i < startRow + 3; i++) {
      for (let j = startCol; j < startCol + 3; j++) {
        result.push(matrix[i][j]);
      }
    }
    
    return result;
}

/**
 * 
 * @param {[string]} line 
 */
function checkSequence(line) {
    return onlyUnique(line.filter(p => p !== '.'));
}

function onlyUnique(arr) {
    const set = new Set(arr);
    return set.size === arr.length;
}

console.log(solution(
    [
        [".","1",".",".",".",".",".",".","."], 
        [".",".","4",".",".",".",".",".","."], 
        [".",".",".","1",".",".","7",".","."], 
        [".",".",".",".",".","1",".",".","."], 
        [".",".",".","3",".",".",".","6","."], 
        [".",".",".",".",".","6",".","9","."], 
        [".",".",".",".","1",".",".",".","."], 
        [".",".",".",".",".",".","2",".","."], 
        [".",".",".","8",".",".",".","2","."]
    ]
));