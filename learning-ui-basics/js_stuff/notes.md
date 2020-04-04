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
