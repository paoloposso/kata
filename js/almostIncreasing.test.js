const { isAlmostIncreasing } =  require('./almostIncreasing.js');

test('Test false', () => {
    let result = isAlmostIncreasing([1,2,1,2])
  
    expect(result).toBe(false);
});

test('Test true', () => {
    let result = isAlmostIncreasing([1,2,3,2])
  
    expect(result).toBe(true);
});