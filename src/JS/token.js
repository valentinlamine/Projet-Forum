var token = document.querySelector(".tok").innerHTML ;

localStorage.setItem("token",token);

function Deconnect(){
    localStorage.clear("token");
    
}
