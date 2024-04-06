async function submitForm(event) {
    event.preventDefault();
    let input = $("#search_string").val();
    input = input.replaceAll(" ", "-");
    $("#loading").attr("hidden", false);
  
    await fetch(
      `https://brand-monitor-api.azurewebsites.net/?search_string=${input}`
    )
      .then((response) => {
        if (!response.ok) throw new Error("HTTP status " + response.status);
        return response.json();
      })
      .then((json) => {
        $("#results").empty();
        $("#loading").attr("hidden", true);
        json.forEach((element) => {
          let row = createListItem(element);
          $("#results").append(row);
        });
      });
}

function createListItem(obj) {
    let li = $("<li>").addClass("results__item").html(`
          <p class="title" >Título: ${obj.Title}</p>
          <p>Descrição: ${obj.Description}</p>
          <a href="${obj.URL}" >link</a>
      `);
  
    return li;
}
  