fetch("http://localhost:3000/api/v1/auth/login", {
    method: "POST",
    body: JSON.stringify({
      email: "hooman@test.com",
      password: "hooman@test.com"
    }),
    headers: {
      "Content-type": "application/json; charset=UTF-8"
    }
  })
  .then((response) => response.json())
  .then((json) => console.log(json));