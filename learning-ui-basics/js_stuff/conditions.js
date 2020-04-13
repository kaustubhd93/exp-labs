"use strict";

// let result = prompt("In what year was covid19 declared a Pandemic?", 2019);
// if (result === "2020") {
//     alert("You are correct.");
// } else {
//     alert("Wrong answer.");
//     let choice = confirm("Do you want to know the correct answer?");
//     if (choice) {
//         alert("Covid19 was declared as a Pandemic in 2020 by WHO.");
//     } else {
//         alert("No problem, you have chosen to be ignorant.");
//     }
// }

// let result = Number(prompt("How many zeros are present in a thousand?", ""));
// if (result > 3) {
//     alert("No, that is too large!");
// } else if (result < 3) {
//     alert("No, that is too little!");
// } else if (result === 3) {
//     alert("Perfect");
// } else {
//     alert("Please enter a number.");
// }

// let age = prompt("What is your age?","");
// //let accessAllowed = (age >= 18) ? true : false;
// let accessAllowed = age >= 18 ? true : false;
// accessAllowed ? alert("Access granted") : alert("Acccess denied");

// Special "?" example

let login = prompt("What is your designation", "");
message = (login == "employee") ? "Hello employee" :
    (login == "director") ? "Greetings direcor!" :
    (login == "") ? "Nothing entered" :
    "Invalid choice";
alert(message);
