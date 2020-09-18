let reqdomainDiv = document.getElementById('reqdomain');

async function postData(url = '', data = {}) {
    // Default options are marked with *
    const response = await fetch(url, {
      method: 'POST', // *GET, POST, PUT, DELETE, etc.
      mode: 'cors', // no-cors, *cors, same-origin
      cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
      //credentials: 'same-origin', // include, *same-origin, omit
      headers: {
        //'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
      },
      redirect: 'follow', // manual, *follow, error
      referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
      body: JSON.stringify(data) // body data type must match "Content-Type" header
    });
    return response; // parses JSON response into native JavaScript objects
  }
  
postData('http://192.168.1.3/api/', { "url": "https://www.hotstar.com/ads.txt" })
.then(res => (res.json()))
.then(adstxt => {
    console.log(adstxt)
    reqdomainDiv.innerHTML += `<p> ${JSON.stringify(adstxt)} </p>`
});

