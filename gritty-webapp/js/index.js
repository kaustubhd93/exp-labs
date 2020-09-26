let reqdomainDiv = document.getElementById('reqdomain');

async function postData(url = '', data = {}) {
    const response = await fetch(url, {
      method: 'POST',
      mode: 'cors',
      cache: 'no-cache',
      headers: {
        //'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
      },
      redirect: 'follow',
      referrerPolicy: 'no-referrer',
      body: JSON.stringify(data)
    });
    return response;
  }
  
postData('http://X.X.X.X/api/', { "data": "" })
.then(res => (res.json()))
.then(resObj => {
    console.log(resObj)
    reqdomainDiv.innerHTML += `<p> ${JSON.stringify(resObj)} </p>`
});

