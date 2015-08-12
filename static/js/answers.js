var answer = {};
answer['1-blue']='e69cd6f97fc7da2f39d39519d1ccf3c3957ff6be88a6ea4bdf40fbed1356b96412a71b25199d3b6f7140b30c936b8a86afaccf8759b232191d2327f648f77055';
answer['1-lord']='e69cd6f97fc7da2f39d39519d1ccf3c3957ff6be88a6ea4bdf40fbed1356b96412a71b25199d3b6f7140b30c936b8a86afaccf8759b232191d2327f648f77055';
answer['2-blue']='da51d72322c22a824a1c50ec65a03393c3da6c7e6996fdb9d64f19a90713258d891181a45abb35b376a6f8750aa17ad0e1cb484689b52c1a8daa950aaa0478f0';
answer['2-lord']='da51d72322c22a824a1c50ec65a03393c3da6c7e6996fdb9d64f19a90713258d891181a45abb35b376a6f8750aa17ad0e1cb484689b52c1a8daa950aaa0478f0';


// To encrypt
console.log("answer['1-blue']='" + CryptoJS.SHA3("233168") + "';")
console.log("answer['1-lord']='" + CryptoJS.SHA3("233168") + "';")

console.log("answer['2-blue']='" + CryptoJS.SHA3("4613732") + "';")
console.log("answer['2-lord']='" + CryptoJS.SHA3("4613732") + "';")
