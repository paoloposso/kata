
function xorSequence(l, r) {
    let result = 0;

    let previous = 0;

    for (let i = 0; i <= l; i++) {
        previous = previous^i;
    }

    result = previous;

    for (let i = l+1; i <= r; i++) {
        previous = previous^i;
        result = result^previous;
    }

    return result;
}

xorSequence(2, 4);