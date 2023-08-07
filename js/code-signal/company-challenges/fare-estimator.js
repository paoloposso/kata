
/**
 * 
 * @param {number} ride_time 
 * @param {number} ride_distance 
 * @param {[number]} cost_per_minute 
 * @param {[number]} cost_per_mile 
 */
function solution(ride_time, ride_distance, cost_per_minute, cost_per_mile) {
    const result = [];

    for (let i = 0; i < cost_per_minute.length; i++) {
        result.push(cost_per_minute[i]*ride_time + cost_per_mile[i]*ride_distance);
    }

    return result;
}
