var myInt = 0;

var canvas = document.getElementById('canvas1');
var ctx = canvas.getContext('2d');

var text1 = '_';
var text2 = '_';
var text3 = '_';
var text4 = '_';

var myAPIKey = "&appid=21b662a3cdb633b4ed2724c9c76856c4";
var myURL = "https://api.openweathermap.org/data/2.5/weather?q=Denver,us";

function buttonOne(){
    
    myInt++;
    text4 = text3;
    text3 = text2;
    text2 = text1;
    text1 = "Number " + myInt;
    $('#text1').text(text1);
    $('#text2').text(text2);
    $('#text3').text(text3);
    $('#text4').text(text4);
    console.log("Number " + myInt);

    
    if(myInt%2 === 0 ){
        ctx.fillStyle = 'red';
        ctx.fillRect(10, 10, 100, 100);
    }else{
        ctx.fillStyle = 'green';
        ctx.fillRect(10, 10, 100, 100);
    }
    
    
}


           

function myFetch(){
    fetch(myURL + myAPIKey, {
    }).then(function (response) {
        console.log(response.ok)
        return response.json();
        }).then(function (myJson) {
            console.log('weather: ' + JSON.stringify(myJson.main));
        });
         
}
$(function() {
    myFetch();
  });
$('#button1').on('click', buttonOne);
$('#button2').on('click', myFetch); 
