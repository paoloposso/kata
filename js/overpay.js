/**
 * 
 * @param {[number]} prices 
 * @param {[string]} notes 
 * @param {number} x 
 */
function solution(prices, notes, x) {
    let totalOverpay = 0;
    for (let i = 0; i < prices.length; i++) {
        totalOverpay += getOverpay(prices[i], notes[i]);
    }

    return totalOverpay <= x || totalOverpay.toFixed(5) <= x;
}

/**
 * 
 * @param {number} price 
 * @param {string} note 
 */
function getOverpay(price, note) {
    if (note.includes('Same')) {
        return 0;
    }
    let perc = Number.parseFloat(note.split('%')[0])/100;
    if (note.includes('higher')) {
        return price - (price / (1+perc)).toFixed(5);
    }
    return (price - (price / (1-perc)).toFixed(5));
}

solution([110, 95, 70], ["10.0% higher than in-store", 
                        "5.0% lower than in-store", 
                        "Same as in-store"], 5)

solution([40, 40, 40, 40], ["0.001% higher than in-store", 
                            "0.0% lower than in-store", 
                            "0.0% higher than in-store", 
                            "0.0% lower than in-store"], 0)