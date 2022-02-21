const api_url = "http://localhost:8080/api/hits";

async function getHits(url) {
console.log("Getting Hits,")

// Storing response
const response = await fetch(url);

// Storing data in form of JSON
var data = await response.json();
console.log(data.Hits);  
show(data.Hits);
}

function show(hits) {
let tab = 
  ``;

// Loop to access all rows 
for (let r of hits) {
  if (r.Member != 'favicon.ico') {
      tab += `<tr> 
      <td>${r.Member}</td>
      <td>${r.Score} </td>

      </tr>`;

  }


}
// Setting innerHTML as tab variable
document.getElementById("hits").innerHTML = tab;
}






