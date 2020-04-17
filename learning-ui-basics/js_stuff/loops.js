"use strict";

// ------ Example of break statement -------
// let init_val = 1;
//
// while (true) {
//     let input_val = +prompt("Enter a number", "");
//     if (!input_val) break;
//     init_val += input_val;
// }
// alert("The sum is :" + init_val);

// ------ Example of labeling -------
outer: for (let i=0; i < 3; i++) {
    for (let j=0; j < 3; j++) {
        let input_val = prompt(`Value at coords (${i},${j})`, "");
        if (!input_val) break outer;
    }
}
alert("Done!");
