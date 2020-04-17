"use strict";
let prime_set = [2,3,5,7];
let max_limit = +prompt("You want me to display prime numbers upto?", 10);
notprime_1: for (let i = 2; i <= max_limit; i++) {
    if (i % 2 === 0 || i % 3 === 0 || i % 5 === 0 || i % 7 === 0) {
        continue;
    } else if (Math.sqrt(i) % 1 === 0) {
        continue;
    } else {
        // prime_set.push(i);
        let temp_i = (i - 1)/2;
        for (let j = 11; j < temp_i ; j++) {
            if ( i % j === 0) continue notprime_1;
        }
        prime_set.push(i)
    }
}
// console.log(prime_set);

// ---- Alternate method ----
let prime_set_2 = [];
let max_limit_2 = +prompt("You want me to display prime numbers upto?", 10);

notprime: for (let i = 2; i <= max_limit_2; i++) {
    for (let j = 2; j < i; j++) {
        if ( i % j === 0) continue notprime;
    }

    prime_set_2.push(i);
}
// console.log(prime_set_2)

if (prime_set.length != prime_set_2.length) {
    let mismatch = [];
    for (let i=0; i < prime_set.length; i++) {
        if (prime_set[i] != prime_set_2[i]) {
            mismatch.push(prime_set[i]);
        }
    }
    console.log(mismatch);
} else {
    alert("Your algorithm works!");
}
