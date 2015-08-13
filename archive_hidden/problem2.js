
var total = 0;

var v1 =1;
var v2 =1;
for (var i =1;i<4000000;i++) {
    name(v1, v2, function (vv1, vv2) {

        if(vv1 % 2 === 0){
            total = total + vv1;
        }
        v1 = vv1;
        v2 = vv2;
        //console.log(v1, v2)
    });
}

function name(v1,v2,callback){
    //do here
    return callback(v1+v2,v1);
}

console.log(total)