> NOTE: Javascript gets executed before the html part.

# Comments in Javascript

- Use this for inline comments ```// Comment```
- Use this for multi-line comments ```/* Comment */```

# Including Javascript in HTML.

- This can be done using the ```<script>``` element.
- Referring external js and adding code in the script block doesn't work.
- Only the code in the external js file will be executed. The code in the script file will be ignored.

# Modern Javascript:

- For a long time, JavaScript evolved without compatibility issues. New features were added to the language while old functionality didn’t change.
- That had the benefit of never breaking existing code. But the downside was that any mistake or an imperfect decision made by JavaScript’s creators got stuck in the language forever.
- This was the case until 2009 when ECMAScript 5 (ES5) appeared. It added new features to the language and modified some of the existing ones. To keep the old code working, most such modifications are off by default. You need to explicitly enable them with a special directive: ```"use strict"```.

# Variables and DataTypes.

## Variables
- There are two ways to define a variable.
```
let variable_name;
var variable_name;
```
- ```var``` is an old way to define variables. More on this later.
- When you use ```const``` to define a variable. You cannot reassign a new value to it.

### Limitations on variable names in JavaScript:

- Name must contain only letters, digits, symbols $ and _
- The first character must not be a digit.
- camelCase is commonly used. (Not exacly a limitation but a recommendation.)
- non-latin letters are also allowed but not recommended.
- List of reserved words: let, class, return, function.

## DataTypes

- Number: can be int, float
- BigInt: use this while reperesenting huge numbers like microsecond precision timestamp as the number type cannot represent integer values larger than 2^53 or less than -2^53.
```javascript
const hugeNo = 1234567890145686789012345678901234567890n;
```
- String: There are 3 ways to quote a string.
```javascript
let str = "Hello";
let strTwo = 'Single quotes are okay too';
let phrase = `can embed another ${str}`;
```
As you can see backticks are "extended functionality" quotes. They allow us to embed variables and expressions into a string by wrapping them in ```${}```
- Boolean: has two values true/false
- null: same as None is Python
- undefined: Declare a variable and don't assign a value to it then it becomes undefined, not null.
- object: Can be used to store collections of data. More on this later.
- symbol: used to create unique identites on objects.

> NOTE: Doing maths is “safe” in JavaScript. We can do anything: divide by zero, treat non-numeric strings as numbers, etc. The script will never stop with a fatal error (“die”). At worst, we'll get NaN as the result.

# Operators

## unary, binary, operands

- In 5*2. 5 and 2 are operands.
- an operator is binary if it has two operands.
- operator is unary if it has one operand. (In -1 the negative operator makes the operand negative.)

## String concatenation, binary +

- An expression like ("1" + 2) in Python would result in an error but JavaScript processes this differenty.
- ("1" + 2) => "12" (Note this is a string)
- (2 + "1") => "21"
- (4 + 5 + "7") => "97" (Since the expression is evaluated left to right so 4 and 5 will be added.)
- But this is only special with the "+" operator. For all the others, the if the number is represented as a string then the end result will be a number and the mathematical operation will take place normally.

> NOTE : An operator always returns a value. That’s obvious for most of them like addition + or multiplication ```*```. But the assignment operator follows this rule too. The call ```x = value``` writes the value into ```x``` and then returns it.

## Exponentiation

- The expression below gives the answer as 32.
```javascript
alert(2 ** 5)
```

## Increment/Decrement

- If we use the normal form ```counter = counter + 1```, the result will always be a number greater than the previous one. But counter++ and ++counter have a different meaning.
- This will result in ```3```.
```javascript
let counter = 2;
counter++;
alert(counter);
```
- This will result in ```3```.
```javascript
let counter = 2;
++counter;
alert(counter);
```
- The difference is that the postfix form (var_name++) returns the previous value and the prefix form (++var_name) returns the calculated value. You will notice this in the developer console.
- This will result in 1.
```javascript
let counter = 1;
alert(counter++);
```
- This will result in 2.
```javascript
let counter = 1;
alert(++counter);
```

## Some tricky operators in javascript.

```javascript
"" + 1 + 0 = "10" // (1)
"" - 1 + 0 = -1 // (2)
true + false = 1
6 / "3" = 2
"2" * "3" = 6
4 + 5 + "px" = "9px"
"$" + 4 + 5 = "$45"
"4" - 2 = 2
"4px" - 2 = NaN
7 / 0 = Infinity
" -9  " + 5 = " -9  5" // (3)
" -9  " - 5 = -14 // (4)
null + 1 = 1 // (5)
undefined + 1 = NaN // (6)
" \t \n" - 2 = -2 // (7)
```

1. The addition with a string "" + 1 converts 1 to a string: "" + 1 = "1", and then we have "1" + 0, the same rule is applied.
2. The subtraction - (like most math operations) only works with numbers, it converts an empty string "" to 0.
3. The addition with a string appends the number 5 to the string.
4. The subtraction always converts to numbers, so it makes " -9 " a number -9 (ignoring spaces around it).
5. null becomes 0 after the numeric conversion.
6. undefined becomes NaN after the numeric conversion.
7. Space characters, are trimmed off string start and end when a string is converted to a number. Here the whole string consists of space characters, such as \t, \n and a “regular” space between them. So, similarly to an empty string, it becomes 0

# Comparisons

## String Comparison

- Strings are comparable in JavaScript. It uses "dictionary" or "lexicographical" order.
- The algorithm to compare two strings is simple:
    - Compare the first character of both strings.
    - If the first character from the first string is greater (or less) than the other string’s, then the first string is greater (or less) than the second. We’re done.
    - Otherwise, if both strings’ first characters are the same, compare the second characters the same way.
    - Repeat until the end of either string.
    - If both strings end at the same length, then they are equal. Otherwise, the longer string is greater.
- Check these examples.
```javascript
alert( 'Z' > 'A' ); // true
alert( 'Glow' > 'Glee' ); // true
alert( 'Bee' > 'Be' ); // true
alert("A" > "a"); // false
/* true because  the lowercase character has a greater index in the internal encoding table JavaScript uses (Unicode).*/
alert("a" > "A");
alert("Glowb" > "Glowa"); // true
```
- While strings which are numbers in form of strings are treated as numbers while comparing it with actual numbers or comparison with a string itself.

## Strict equality

- A regular equality check (```==```) has a problem. It cannot differentiate ```0``` from ```false```.
- The problem:
```javascript
alert( 0 == false ); // true
alert( '' == false ); // true
```
- This happens because operands of different types are converted to numbers by the equality operator ==. An empty string, just like false, becomes a zero.
- **A strict equality operator === checks the equality without type conversion.**
```javascript
alert( 0 === false ); // false, because the types are different
```
- There is also a "strict non-equality" operator ```!==``` analogous to ```!=```.
- Special case: This equality check returns false because undefined only equals null.
```javascript
alert( undefined == 0 ); // false
```

# Browser specific functions to interact with other visitors

```javascript
alert("any_message");
prompt("Input_question_what_is_your_age_weight_etc", default_value);
confirm("Are you the head of engineering?");
```

- These methods mentioned above are all browser specific.
- You have to use modal elements to alter their look and feel.
- Ensure you use default_value in prompt Otherwise it fails on IE.
- When you use ```prompt```. It's default data type is string.

# Conditional Operators

## if-else-elseif

- A number ```0```, an empty string ```""```, ```null```, ```undefined```, and ```NaN``` all become false. Because of that they are called "falsy" values.
- Other values become true, so they are called "truthy".
- Syntax:
```javascript
if (condition) {
    //some code here;
} else {
    //some code here;
}

// if else if
if (condition) {
    //some code here;
} else if (condition) {
    //some code here;
} else {
    //some code here;
}
```

## ```?``` (Yes! that's exactly how it looks.)

- This is it's syntax:
```javascript
let result = condition ? val1 : val2;
```
- If the condition is ```true```, result is val1. If ```false``` then it is val2.
- Using this makes the code shorter but compromises readability.
- Special case of ```?``` operator :
```javascript
let login = prompt("What is your designation", "");
message = (login == "employee") ? "Hello employee" :
    (login == "director") ? "Greetings direcor!" :
    (login == "") ? "Nothing entered" :
    "Invalid choice";
alert(message);
```
# Logical Operators

- There are 3 logical operators in JavaScript:
```javascript
|| (or)
&& (and)
! (not)
```
## Short-circuit evaluation (OR)

- Operands can be not only values, but arbitrary expressions. OR evaluates and tests them from left to right. The evaluation stops when a truthy value is reached, and the value is returned. This process is called "a short-circuit evaluation" because it goes as short as possible from left to right.
- Check this:
```javascript
let x;
true || (x=1);
alert(x); // undefined because (x=1) was never evaluated.

//---------------------------------------------------------
let x;
false || (x=1);
alert(x); // 1 because the expression was evaluated here.
```

## Returning of value (```||``` and ```&&```)

- An actual value is returned in these logical expressions. It does not return a boolean value.
- In case of OR, the first truthy value is returned, if all are false, then the last falsy value is returned.
- In case of AND, the first falsy value is returned, if all are true, then the last truthy value is returned.

## NOT operator

- Syntax is pretty simple:
```javascript
let result = !value;
```
- Returns a boolean value.

## Some special cases.

```javascript
alert(alert(1) && alert(2)); //This will result in 1 first and then undefined.
```

# while and for

- Syntax for ```while```
```javascript
while (condition) {
    // code
    // some more code here
}
```
- Syntax for ```for```
```javascript
for (begin, condition, step) {
    // some code here
}
```
- This is an infinite loop in ```for```.
```javascript
for(;;) {
    //code here repeats without limits.
}
```
- Do not use the ```?``` operator. This results into ```Uncaught SyntaxError: Unexpected token 'continue'```. This is another example why one should not use the ```?``` operator.
```javascript
let i=4;
for (let j=0; j<8; j++) {
    (i < 5) ? alert(i) : continue;
    i++;
}
```
- One can label loops and we can ```break``` out of them or use ```continue``` to go over to the next iterator.
- A call to ```break/continue``` is only possible from inside a loop and the label must be somewhere above the directive.

# The "switch-case"

- Replaces multiple ```if``` blocks.
- Syntax:
```javascript
switch (x) {
    case expression:
        //some code here
        break;
    case expression:
        //some code here
        break;
    default:
        //some code here
        break;
}
```
- If there is no ```break``` statement, the condition that evaluates to be ```true```, that and the one's after it are executed.
