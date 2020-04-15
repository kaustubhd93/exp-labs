"use strict";

let username = prompt("Enter username", "");
if (username === "Admin") {
    let passwd =  prompt("Enter password", "");
    if (passwd === "TheMaster") {
        alert("Welcome!");
    } else if (!passwd) {
        alert("You have not entered anything!");
    } else {
        alert("wrong passwd");
    }
} else if (username && username != "Admin") {
    alert("I don't know you!");
} else if (!username) {
    alert("You have not entered anything!");
}
