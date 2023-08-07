
/**
 * 
 * @param {*} l 
 * @param {number} k 
 */
function solution(l, k) {
  while (l && l.value === k) {
    let temp = l.next;
    l = null;
    l = temp;
  }

  let curr = l;

  while (curr && curr.next) {
    if (curr.next.value === k) {
      curr.next = curr.next.next;
    }
    curr = curr.next;
  }
  return l;
}


let o = {
  value: 3,
  next: {
    value: 2,
    next: {
      value: 8,
      next: null
    }
  }
}

solution(o, 8);