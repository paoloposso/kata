const { expect } = require('expect');

const { quickSort } = require('./quick-sort.js');

test('Should sort correctly', () => {
    const result = quickSort([5,18,1,23,44,8,2]);
    expect(result[result.length-1]).toBe(44);
    expect(result[0]).toBe(1);
});

test('Should sort correctly when totally reversed', () => {
    const result = quickSort([10,9,8,7,6,5,4,3,2,1]);
    expect(result[result.length-1]).toBe(10);
    expect(result[0]).toBe(1);
});

test('Should not break with empty array', () => {
    const result = quickSort([]);
    expect(result.length).toBe(0);
});