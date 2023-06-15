const { expect } = require('expect');

const { search } = require('./binary-search.js');

test('Test true', () => {
    const result = search([5,18,1,23,44,8,2], 2);
    expect(result).toBe(true);
});

test('Test true 2', () => {
    const result = search([10,9,8,7,6,5,4,3,2,1], 3);
    expect(result).toBe(true);
});

test('Test true 2', () => {
    const result = search([10,9,8,7,6,5,4,3,2,1], 20);
    expect(result).toBe(false);
});
