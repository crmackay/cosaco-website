var buttons = document.getElementsByClassName("showData");


// function to reveal data in the row associaed with a button
function revealData(event)
{
  var targetElement = event.target || event.srcElement;

  var id = targetElement.getAttribute("data-id");
    console.log(id)
  var toToggle = document.getElementById(id);
    console.log(toToggle)
  toToggle.classList.toggle("hidden");
}

//add the event listener to each button

for (var i=0; i< buttons.length; i++)
{
  buttons[i].addEventListener("click", revealData, false);  
}