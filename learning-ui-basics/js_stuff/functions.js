"use strict";

// let username = "john";
//
// function showMessage() {
//     // global variable "username" is altered after the function is called.if variable is modified.
//     // username = "bob";
//     // if a variable is explicitly declared using usual syntax, then global variable remains same.
//     let username = "bob";
//     let message = "Inside funcion : " + username;
//     alert(message);
// }
//
// alert("Before function called : " + username);
// showMessage();
// alert("After function called : " + username);

function showMessage(from, text = "no text available") {
    from = "*" + from + "*"
    alert(from + ": " + text);
}

// let from = "Ann";

showMessage("Ann", "What up?");
showMessage("Ann", "Yo");
// showMessage("Ann");
