$(function(){
  function sendData() {
    var XHR = new XMLHttpRequest();

  var formElement = document.querySelector("form");
    // Bind the FormData object and the form element
    var FD = new FormData(formElement);

//     FD.append("firstname", document.querySelector("#firstName").value);
//     FD.append("lastName", document.querySelector("#lastName").value);
//     FD.append("nickname", document.querySelector("#nickname").value);
//     FD.append("skill", document.querySelector("#skill").value);


    // Define what happens on successful data submission
    XHR.addEventListener("onload", function(event) {
      alert(event.target.responseText);
    });

    // Define what happens in case of error
    XHR.addEventListener("error", function(event) {
      alert('Oops! Something went wrong.');
    });

    // Set up our request
    XHR.open("POST", "http://localhost:8080/players" );
    XHR.setRequestHeader("Access-Control-Allow-Origin", "*");
    XHR.setRequestHeader("Content-Type", "multipart/form-data");
    XHR.setRequestHeader("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT");

    // The data sent is what the user provided in the form
    XHR.send(FD);
  }

  // Access the form element...
  var form = document.getElementById("myForm");

  // ...and take over its submit event.
  form.addEventListener("submit", function (event) {
    event.preventDefault();

    sendData();
  });
});
