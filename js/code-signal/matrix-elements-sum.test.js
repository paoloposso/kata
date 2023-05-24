const { expect } = require('expect');

const { solution } = require('./matrix-elements-sum.js');

test('Test OK', () => {
    const result = solution([[4,0,1], 
        [10,7,0], 
        [0,0,0], 
        [9,1,2]]);
    expect(result).toBe(15);
});
