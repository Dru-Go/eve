var nameErr = document.querySelector("#username-err");
var emailErr = document.querySelector("#email-err");
var phoneErr = document.querySelector("#phone-err");
var formUser = document.querySelector("#form_create");

var userName = document.querySelector("#icon_prefix");

//    username must be unique
userName.addEventListener("input", function() {
  console.log(userName.value);
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/checkusername/" + userName.value);
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
  xhr.send();
  // xhr1.send()
});

var email = document.querySelector("#email");

//    email must be unique
email.addEventListener("input", function() {
  console.log(email.value);
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/checkemail/" + email.value);
  // xhr.send(email.value);
  xhr.addEventListener("readystatechange", function() {
    if (xhr.readyState === 4) {
      var item = xhr.responseText;
      console.log(item);
      if (item == "true") {
        emailErr.textContent = "Email taken - Try another mail!";
      } else {
        emailErr.textContent = "";
      }
    }
  });
  xhr.send();
});

var phone = document.querySelector("#phone-no");

phone.addEventListener("input", function(e) {
  console.log(phone.value);
  var xhr = new XMLHttpRequest();
  xhr.open("GET", "/api/checkphone/" + phone.value);
  xhr.addEventListener("readystatechange", function() {
    if (xhr.readyState === 4) {
      var item = xhr.responseText;
      console.log(item);
      if (item == "true") {
        phoneErr.textContent = "Phone No already exists!";
      } else {
        phoneErr.textContent = "";
      }
    }
  });
  xhr.send();
});

formUser.addEventListener("submit", function(e) {
  console.log(nameErr.textContent);
  if (nameErr.textContent + emailErr.textContent + phoneErr.textContent == "") {
    return;
  }
  e.preventDefault();
});

(function($) {
  $(function() {
    $(".datepicker").datepicker();
    $(".modal").modal();
    $("select").formSelect();
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
  var lastname = document.getElementById("LastName");

  var email = document.getElementById("email");
  var phone = document.getElementById("phone-no");
  var role = document.getElementById("role");

  var roles;
  switch (parseInt(role.value)) {
    case 1:
      roles = "Guest";
      break;
    case 2:
      roles = "Organizing Commitee";
      break;
    case 3:
      roles = "Speaker";
      break;
  }

  function nonempty(fullname, email, phone, role) {
    var names;
    if (fullname != "") {
      names = "Name:" + fullname;
    }
    if (email != "") {
      names += "  Email:" + email;
    }
    if (phone != "") {
      names += "  Phone Number:" + phone;
    }
    if (role != "") {
      names += "  Role:" + role;
    }
    return names;
  }

  console.log(nonempty(fullname.value, email.value, phone.value, roles));
  var col = nonempty(fullname.value, email.value, phone.value, roles);

  qrcode.makeCode(col);
}
var btn = document.getElementById("button");

  $(".right").on("load",function () {
    
  })



$("#button")
  .on("click", function() {
    makeCode();
  })
  .on("keydown", function(e) {
    if (e.keyCode == 13) {
      makeCode();
    }
  });

$("#icon_prefix")
  .on("blur", function() {
    makeCode();
  })
  .on("keydown", function(e) {
    if (e.keyCode == 13) {
      makeCode();
    }
  });

$("#LastName")
  .on("blur", function () {
    makeCode();
  })
  .on("keydown", function (e) {
    if (e.keyCode == 13) {
      makeCode();
    }
  });

$("#email")
  .on("blur", function() {
    makeCode();
  })
  .on("keydown", function(e) {
    if (e.keyCode == 13) {
      makeCode();
    }
  });

$("#phone-no")
  .on("blur", function() {
    makeCode();
  })
  .on("keydown", function(e) {
    if (e.keyCode == 13) {
      makeCode();
    }
  });
