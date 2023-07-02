/**
 * 
 * @param {number} value 
 * @param {*} left 
 * @param {*} right 
 * @returns 
 */
function createNode(value, left = null, right = null) {
    return {
      value: value,
      left: left,
      right: right
    };
}

/**
 * 
 * @param {*} rootNode 
 * @param {number} value 
 */
function insert(rootNode, value) {
    insertRecur(rootNode, value);
    return rootNode;
}

function insertRecur(curNode, value) {
    if (value > curNode.value) {
        if (curNode.right) {
            return insertRecur(curNode.right, value);
        }
        curNode.right = createNode(value);
        return;
    } else {
        if (curNode.left) {
            return insertRecur(curNode.left, value);
        }
        curNode.left = createNode(value);
        return;
    }
}

function preOrderTraversal(rootNode) {
    let res = [];

    preOrderTraversalRecur(rootNode, res);

    return res;
}

/**
 * 
 * @param {*} curNode 
 * @param {[number]} res 
 */
function preOrderTraversalRecur(curNode, res) {

    if (!curNode) return;

    res.push(curNode.value);

    preOrderTraversalRecur(curNode.left, res);
    preOrderTraversalRecur(curNode.right, res);
}

/**
 * 
 * @param {string} input 
 */
function processData(input) {
    const values = input.split('\n')[1].split(' ');

    let rootNode = createNode(values[0]);

    values.slice(1).forEach(v => {
        insertRecur(rootNode, v);
    });

    let a = preOrderTraversal(rootNode);

    console.log(a.join(' '));
}

module.exports = { createNode, insert }