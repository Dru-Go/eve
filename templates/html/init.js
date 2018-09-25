
var nameErr = document.querySelector('#username-err');
var emailErr = document.querySelector('#email-err');



var userName = document.querySelector("#icon_prefix");

//    username must be unique
userName.addEventListener("input", function() {
  console.log(userName.value);
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/checkUserName" + userName.value);
  // xhr.send(userName.value);
  xhr.addEventListener("readystatechange", function() {
    if (xhr.readyState === 4) {
      var item = xhr.responseText;
      console.log(item);
      if (item == "true") {
        nameErr.textContent = "Username taken - Try another name!";
      } else {
        nameErr.textContent = "";
      }
    }
  });
});


var email = document.querySelector("#email");

//    email must be unique
email.addEventListener("input", function() {
  console.log(email.value);
  var xhr = new XMLHttpRequest();
  xhr.open("POST", "/api/checkemail");
  xhr.send(email.value);
  xhr.addEventListener("readystatechange", function() {
    if (xhr.readyState === 4) {
      var item = xhr.responseText;
      console.log(item);
      if (item == "true") {
        emailErr.textContent = "email taken - Try another name!";
      } else {
        emailErr.textContent = "";
      }
    }
  });
});



(function($){
  $(function(){
    $('.datepicker').datepicker();
    $('.modal').modal();
    $('select').formSelect();
  }); // end of document ready
})(jQuery); // end of jQuery name space



var qrcode = new QRCode(document.getElementById("qrcode"), {
  width: 200,
  height: 200,
  colorDark: "black",
  colorLight: "white",
  correctLevel: QRCode.CorrectLevel.L
});



function makeCode() {
  var fullname = document.getElementById("icon_prefix");
  var email = document.getElementById("email");
  var phone = document.getElementById("phone-no");
  var role = document.getElementById("role");

  var roles
  switch (parseInt(role.value)) {
    case 1:
      roles = "Guest"
      break;
    case 2:
      roles = "Organizing Commitee"
      break
    case 3:
      roles = "Speaker"
      break;
  }

  function nonempty(fullname, email, phone, role) {
    var names
    if (fullname != "") {
      names = "Name:" + fullname
    }
    if (email != "") {
      names += "  Email:" + email
    }
    if (phone != "") {
      names += "  Phone Number:" + phone
    }
    if (role != "") {
      names += "  Role:" + role
    }
    return names
  }


  console.log(nonempty(fullname.value, email.value, phone.value, roles))
  var col = nonempty(fullname.value, email.value, phone.value, roles)

  qrcode.makeCode(col);
}
var btn = document.getElementById("button")

$("#button").
  on("click", function () {
    makeCode();
  }).
  on("keydown", function (e) {
    if (e.keyCode == 13) {
      makeCode();
    }
  });

$('#icon_prefix').
  on("blur", function () {
    makeCode();
  }).
  on("keydown", function (e) {
    if (e.keyCode == 13) {
      makeCode();
    }
  });


$('#email').
  on("blur", function () {
    makeCode();
  }).
  on("keydown", function (e) {
    if (e.keyCode == 13) {
      makeCode();
    }
  });

$('#phone-no').
  on("blur", function () {
    makeCode();
  }).
  on("keydown", function (e) {
    if (e.keyCode == 13) {
      makeCode();
    }
  });



