// we want to get the data dynamically and add to the page
// 
// first we want to start with get

// send the request for all the users
// when we get the request 
// we parse it into an object 
// we iterate through this object
// we write the data to a table

// 

//---------------------- util ---------------------------
function hide(id)
{
    document.getElementById(id).style = "display: none;";
}

function show(id)
{
    document.getElementById(id).style = "display: block;";
}

function textContent(id, value)
{
    document.getElementById(id).textContent=value;
}

const nelson_sec = 500;

// --------------------- users --------------------------
function wrapIntr(val)
{
    return "<tr>" + val + "</tr>";
}

function wrapIntd(val)
{
    return "<td>" + val + "</td>";
}

function fillUsersTable(users)
{
    // we are assumning that the parameter is an array
    console.log(users.length)
    for(i=0; i<users.length; i++)
    {
        var tr = wrapIntr
        (
            wrapIntd(users[i].id) + 
            wrapIntd(users[i].email) + 
            wrapIntd(users[i].password)
        );

        document.getElementById("usersTableBody").innerHTML += tr;
    }
}

function gUsersPanel(usersresponse)
{
    let panel="guserspanel";
    let panelbody="guserspanelbody";

    show(panel);
    textContent(panelbody,usersresponse);

    setTimeout(
        function(){
            hide(panel);
        }, 
        nelson_sec
    );
}

function getusers()
{
    var url = "http://localhost:8080/users/";
    var method = "GET";
    var xhttp = new XMLHttpRequest();

    xhttp.onreadystatechange = function() {
        var var_responseText = xhttp.responseText;
        gUsersPanel(var_responseText);

        const var_users = JSON.parse(var_responseText);
        fillUsersTable(var_users);
    };

    xhttp.open(method, url);
    xhttp.send();
}

function createuser()
{
    let email = document.getElementById("email").value;
    let password = document.getElementById("password").value;

    let url = "http://localhost:8080/users/";
    let method = "POST";

    var xhttp = new XMLHttpRequest();

    xhttp.onreadystatechange = function() {
        var var_responseText = xhttp.responseText;
        gUsersPanel(var_responseText);

        const var_users = JSON.parse(var_responseText);
        fillNotesTable(var_users);
    };

    let data = "?email="+email+"&&password="+password;
    url += data;
    xhttp.open(method, url);
    xhttp.send();
}

//------------------------ notes ----------------------
function createnotestatus(status)
{
    let panel="cnotespanel";
    let panelBody="cnotespanelbody";

    show(panel);
    textContent(panelBody,status);

    setTimeout(
        function(){
            hide(panel);
        },
        nelson_sec
    );
}

function createnote()
{
    let url="http://localhost:8080/notes/";
    let method="POST";

    let title=document.getElementById("title").value;
    let body=document.getElementById("body").value; 
    let data="?title="+title+"&&body="+body;

    let xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function(){
        let response = xhttp.responseText;
        createnotestatus(response);
    };

    xhttp.open(method, url+data);
    xhttp.send();
}

