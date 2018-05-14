function loadData() {

    var myURL = 'http://api.openweathermap.org/data/2.5/weather?zip=94040,us';
    //var myURL = 'https://google.com'
    fetch(myURL, {
        mode: 'cors',
        credentials: 'include',
    }).then(function(data){

        console.log(data.ok);

    });

    console.log(myURL);

};

$('#button-action1').submit(loadData);
