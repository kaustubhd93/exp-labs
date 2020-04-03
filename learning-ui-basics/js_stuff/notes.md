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
