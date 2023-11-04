getData = async () => {
    let data = await fetch("https://v2.jokeapi.dev/joke/Programming?lang=es&amount=10")
    let res = await data.json()

    var chistes = document.getElementById("chistes");

    res.jokes.forEach(function (joke) {
        var listItem = document.createElement("li");

        // Check if the joke is a "single" type or "twopart" type and display accordingly
        if (joke.type === "single") {
            listItem.textContent = joke.joke;
        } else if (joke.type === "twopart") {
            listItem.textContent = `${joke.setup} ${joke.delivery}`;
        }

        // Append the li element to the ul
        chistes.appendChild(listItem);
    });
}

getData()