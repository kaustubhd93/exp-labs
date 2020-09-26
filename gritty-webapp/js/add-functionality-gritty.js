// DOM elements
const rotateLeftButton = document.getElementById('rotate-left-button');
const rotateRightButton = document.getElementById('rotate-right-button');
const grittyBorder = document.getElementById('gritty-border');
const gritterForm = document.getElementById('gritter-form');
const getAdsTxt = document.getElementById('get-ads-txt');

// global variables
let rotation = 0;

// event handlers
function toggleGritty() {
    console.log('FIRING')
    const gritty = document.getElementById('toggle-gritty');
    toggleVisibility(gritty);
}

function rotateGritty(event) {
    const gritty = document.getElementById('small-gritty');
    rotate(gritty, event);
}

function addGreet(event) {
    event.preventDefault();
    const greet = gritterForm.greet.value;
    const newGreetLi = document.createElement('li');
    const newGreetAvatar = document.createElement('div');
    newGreetAvatar.className = "avatar";
    const newGreetText = document.createElement('span');
    newGreetText.innerText = greet;
    const greets = document.getElementById('greets').getElementsByTagName('ul')[0];
    const newerGreetLi = greets.appendChild(newGreetLi);
    newerGreetLi.appendChild(newGreetAvatar);
    newerGreetLi.appendChild(newGreetText);
    gritterForm.greet.value = "";
}


// helpers
// given a DOM element, change it's visibilty style property from hidden to visible
function toggleVisibility(element) {
    if (element.style.visibility === 'hidden') {
        element.style.visibility = 'visible'
    } else {
        element.style.visibility = 'hidden'
    }
}

// given a DOM element and a direction, rotate the element that direction
function rotate(element, event) {
    if (event.target.id === 'rotate-left-button') {
        rotation = rotation - 15;
    } else {
        rotation = rotation + 15;
    }
    element.style.transform = 'rotate(' + rotation + 'deg)';
} 

async function postData(url = '', data = {}) {
    // Default options are marked with *
    const response = await fetch(url, {
      method: 'POST', // *GET, POST, PUT, DELETE, etc.
      mode: 'cors', // no-cors, *cors, same-origin
      cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
      //credentials: 'same-origin', // include, *same-origin, omit
      headers: {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
      },
      redirect: 'follow', // manual, *follow, error
      referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
      body: JSON.stringify(data) // body data type must match "Content-Type" header
    });
    //console.log(response.json())
    return response.json(); // parses JSON response into native JavaScript objects
  }


// getAdsTxt.onclick = postData('http://192.168.1.3/dev/get_content', { "url": "https://www.hotstar.com/ads.txt" });
// alert(getAdsTxt)
// .then(res => res.json())
// console.log(res.json())
// .then(adstxt => {
//     console.log(adstxt.message)
// })

// add event handlers when mouse events are triggered
grittyBorder.onmouseenter = toggleGritty;
grittyBorder.onmouseleave = toggleGritty;
rotateLeftButton.onclick = rotateGritty;
rotateRightButton.onclick = rotateGritty;
gritterForm.onsubmit = addGreet;