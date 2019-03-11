$(function(){
    var XHR = new XMLHttpRequest();
    var XHRImage  = new XMLHttpRequest();

        // Define what happens on successful data submission
        XHR.addEventListener("load", function(event) {
          document.getElementById("teamName").innerHTML = event.name;
          for (team in event) {
            document.getElementById("teams").load("teams.html");
            for(player in team.member) {
              XHRImage.onload=function(event) {
                    document.getElementById('image').src=event.image;
                    document.getElementById('image').alt=player.nickname;
              });
              XHRImage.open("GET", "https://api.adorable.io/avatars/285/"+player.nickname );
              XHRImage. send();

              document.getElementById('nickname').innerHTML = player.nickname;
            }
          }
        });

        // Define what happens in case of error
        XHR.addEventListener("error", function(event) {
          alert('Oops! Something went wrong.');
        });

        // Set up our request
        XHR.open("GET", "/localhost/teams");

        // The data sent is what the user provided in the form
        XHR.send(FD);
});
