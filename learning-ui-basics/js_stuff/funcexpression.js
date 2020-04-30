"use strict";
// ------ A function expression ------
// let sayHi = function() {
//     alert("Hey, what up?");
// }; // Do not forget to add semicolon here.
//
// // Notice, there are no parentheses for "sayHi"
// // The value of sayHi is assigned to func.
// let func = sayHi;
//
// func();
// sayHi();

// ---- Pass a function itslelf as an argument -----
// function ask(question, yes, no) {
//     if (confirm(question)) {
//         yes();
//     } else {
//         no();
//     }
// }

// showOk and showCancel are called callback functions or callbacks.
// function showOk() {
//     alert("You agreed.");
// }
//
// function showCancel() {
//     alert("You canceled.");
// }

// ----- using function expressions -----
// ask(
//     "Do you agree?",
//     function() { alert("You agreed");},
//     function() { alert("You canceled");});

// ------ Funcion decalration ------
let age = 16;
if (age < 18) {
    welcome();

    function welcome() {
        alert("Hello there!");
    }

    welcome();
} else {
    function welcome() {
        alert("Greetings toddler");
    }
}

welcome();
