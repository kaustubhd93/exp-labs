# Gist of UI components:

- HTML is for adding meaning to raw content by marking it up.
- CSS is formatting that marked up content.
- JavaScript is for making that content and formatting, interactive.

# Structure of an HTML document:

- An HTML doc opens with an HTML tag and also closes with it.
- ```<!DOCTYPE html>```: The browser checks for this tag and identifies the page in HTML5.
- Between the html tag there are 2 sections: "head" and "body".
- "head" is where the Meta-data goes and "body" is where the content goes.
- "head" includes: Page Title, CSS links, Other abstract things.
- "body" includes: headings, paragraphs, other things you can see.

# Other type of content:

- The other major type of content is "inline elements" or "phrasing content".
- For example: ```<em>``` is for emphasizing and italicizing the text and ```<strong>``` makes the text bold.

# Empty HTML elements:

- Line breaks and horizontal rules are the most common empty elements.
- Line break: ```<br/>```
- Horizontal rule is used to separate a section prominently. Using this displays a line.
- Horizontal rule: ```<hr/>```

# How to create Links:

- Links are created with the ```<a>``` element, which stands for "anchor".
- They alone cannot do anything. They need an attribute "href" to successfully take a user to a desired link.

## Attributes:

- An HTML attribute adds meaning to the HTML element it is attached to.
- The extra bit of information provided by the ```href``` attribute tells the browser that this ```<a>``` element is in fact a link and the browser should render this in it's default blue text.
- The ```target``` attribute defines where to display the page once user clicks it. Usually it takes the value ```_blank``` which opens the page in a new tab.

# Links:

## Absolute links:

- Directly refers to a web resource by mentioning the entire URL.
- Starts with ```http://``` or ```https://```.

## Relative links:

- Points to a web page inside your project.
- No need for ```http://``` or ```https://```, just the path will be enough.
- Forward pointing is common. We are allowed to refer a page backwards also by using the ```../``` notation.

## Root Relative links:

- Root-Relative links are similar to the previous section, but instead of being relative to the current page, they are relative to the "root" of the entire website.
- If website is hosted on our-site.com all root relative links will be relative to our-site.com.

# Images:

- Images are included in HTML files by using the ```<img/>``` element.
- It also carries a ```src="path_of_image"``` attribute.
- Also take a note that this is an empty html element like ```<hr/>```.
- There is another attribute which is used; Dimension ```width=""```.
- There is ```height=""``` also, but only one of them is used, as adding the height attribute with width stetches the image.

## 4 main image formats in use of the web.

- jpg
- png
- gif
- svg (scalable vector graphics): as the name suggests, images in this format can be scaled up or down without any loss in quality of the image.

## Text alternatives.

- Adding ```alt``` attribute to the ```<img/>``` element is a best practice.
- Helps with search engines and text-only browsers(eg: people that use text-to-speech software.)

# CSS (Cascading style sheets.)

- A css file is a set of rules how the content will be presented.
- Syntax of Rule:
```css
selector {
    property: value;
}
```
- The block defined by the curly braces is a declaration block.

## How to link css files:

- Use the ```<link/>``` element inside the ```head``` section and add the ```rel``` attribute to ```link```.
- Add the ```href``` attribute to mention the css file with rules.

> NOTE: - There’s no direct connection between the browser and our stylesheet. It’s only through the HTML markup that the browser can find it. CSS, images, and even JavaScript all rely on an HTML page to glue everything together, making HTML the heart of most websites

## Units of measurement:

- Font is measured in **px**(pixel) or **em**(pronounced as m).
- ```em``` is based on a basic font-size sent in px.
- Recommended way is to set a basic font size and then ```em``` every where else to scale up font sizes. This way by adjusting the base every other size is scaled accordingly.

## The Cascade

- The cascading part of CSS is due to the fact that rules cascade down from multiple sources.
- The one which we write in our project is an external style sheet.

### CSS heirarchy:

- Browser's default stylesheet
- User defined stylesheet (This can be edited in web developer console.)
- External stylesheet.
- Page specific stylesheet.
- Inline styles.

> This is ordered from least to most precedence, which means styles defined in each subsequent step override previous ones. For example, inline styles will always make the browser ignore its default styles.

#### You can always refer [MDN CSS reference](https://developer.mozilla.org/en-US/docs/Web/CSS/Reference)

# CSS Box model

- It is a set of rules how every web page on the internet is rendered.
- CSS treats each element in your html document as a "box" with a bunch of different properties that determine where it appears on page.
- It gives each box(both inline and block) four properties.
    - content: The text, image or other media content in the element.
    - padding: space between boxes contents and it's border.
    - border: line between boxes padding and margin.
    - margin: space between box and surrounding boxes.
![css-box-model-73a525](https://user-images.githubusercontent.com/9679436/78031172-cc5ee280-7380-11ea-857b-c14881d650d2.png)
> IMG source: https://www.internetingishard.com/html-and-css/css-box-model/

## Things about margin and padding:

- Padding of a box has a background, while margins are always transparent.
- Padding is included in the click area of an element, while margins aren't.
- Margins collapse vertically, while padding does not.
- One has to choose whether to use margin or padding.

# div and span:

- ```div``` and ```span``` are container elements. They do not affect the semantic structure of an html document.
- However they do provide a hook for adding CSS styles to arbitrary sections of a web page.

# Class-selectors:

- They let you apply css styles to a specific HTML element.
- They let you differentiate between elements of the same type like div.
- They require 2 things:
    - A class attribute in the html element.
    - A matching css selector in your stylesheet.
- The standard naming convention is **lowercase with hyphens** if the name is too big.
- You can declare one more class inside the element. Ensure the css block comes after the main class attribute. In this class you can add new properties or even override existing ones.
- The order in which we specify the class attribute in the html element does not matter.

## Descendant selectors:

- They let you target only those elements who are inside another element like ```<em>``` or ```<strong>```.
- You can nest Descendant selectors as deep as you want but do not overdo it. It will compromise the code readability.

# Pseudo classes

## For Links

- We can set a different property for a link when it is visted, clicked or someone hovers over it.
- It has to be in this order: a:link, a:visited, a:hover, a:active.

## For Button

- It has a similar format as above; .class_name:link/visited/hover/active.

> NOTE: There is a bunch of pseudo classes at our disposal: Check this link  https://developer.mozilla.org/en-US/docs/Web/CSS/Pseudo-classes

## CSS Specificity

- Unfortunately, not all CSS selectors are created equal. “CSS specificity” is the weight given to different categories of selectors.
- This means that certain selectors will always override other ones, regardless of where they appear in the stylesheet.
- This can get very confusing. It’s such a big problem that an entire methodology called ![BEM](http://getbem.com/introduction/) has evolved. BEM attempts to make CSS rules more reusable by making everything a class selector. This completely eliminates the potential for specificity issues. BEM : Block, Element Modifier.
- The takeaway here is that CSS rules are not necessarily applied in sequential order, but you should try to make the browser do so by writing CSS that uses the same specificity.

# Floats

- "Floats" let you put block-level elements side-by-side instead of on top each other.
- Syntax
```css
.selector {
    /* Which direction would you like this block to float*/
    float: direction; /*left or right*/
}
```
> NOTE: To align a block in the centre use "margin: 0 auto". For left or right use "float"

- When you float multiple elements in the same direction, they'll stack horizontally, much like the default vertical layout algorithm, except rotated 90 degrees.
- Percentages in CSS are relative to their parent element.

# Flexbox

- This gives you an alternative to floats, but gives you more power by giving full control of designing the web page horizontally and vertically.
- Flexbox uses 2 types of boxes: flex containers and flex items.
    - The flex container groups a bunch of flex items together and define how they are positioned.
    - Every HTML element that's a direct child of a flex container is an "item".
- Syntax for making an element a flexbox.
```css
.selector {
    /* some property here ;*/
    /* some property here ;*/
    /* some property here ;*/
    display: flex;
    /* align horizontally */
    justify-content: positions; /* flex-end, flex-start, space-between, space-around, center*/
    /* align vertically */
    align-items: positions; /* flex-end, flex-start, center, stretch, baseline */
    flex-wrap: wrap_style; /* wrap, nowrap, wrap-reverse */
    flex-direction: direction; /* row, column, row-reverse, column-reverse*/
}
```
