var xmlhttp = new XMLHttpRequest();
xmlhttp.open("GET","localhost/8081");

xmlhttp.onreadystatechange = function(){
    println("Hello welcome to my page");
}

cmlhttp.onreadystatechange = function(){
    console.log("Ready State: "+ xmlhttp.readyState)
}

xmlhttp.send();