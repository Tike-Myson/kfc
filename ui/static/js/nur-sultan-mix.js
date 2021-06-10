var myGeoJson = JSON.parse(Get("http://127.0.0.1:4000/api"));

console.log(myGeoJson)

function Get(yourUrl){
    var Httpreq = new XMLHttpRequest(); // a new request
    Httpreq.open("GET",yourUrl,false);
    Httpreq.send(null);
    console.log(Httpreq.responseText);
    return Httpreq.responseText;
}