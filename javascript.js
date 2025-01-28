let formData = new FormData();
formData.append('email', 'hooman@test.com');
formData.append('password', 'hooman@test.com');

const params = new URLSearchParams(formData);
console.log('Serialized data:', params.toString())
JSON.stringify(Object.fromEntries(formData));
//const formdata = new FormData(document.querySelector('form'));


register:
fetch("http://localhost:3000/api/v1/auth/register", {
  method: "POST",
  body: JSON.stringify({
    email: "hooman@test.com",
    name: "hooman@test.com",
    password: "hooman@test.com",
  }),
  headers: {
    "Content-type": "application/json; charset=UTF-8"
  }
})
.then((response) => response.json())
.then((json) => console.log(json));



login:
fetch("http://localhost:8080/api/v1/auth/login", {
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


 user:
  fetch("http://localhost:8080/api/v1/auth/user", {
    method: "GET",
    headers: {
      "Content-type": "application/json; charset=UTF-8",
      "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzM0Nzc5MTQsInN1YiI6eyJJRCI6NiwiSW1hZ2UiOiJodHRwczovL3VpLWF2YXRhcnMuY29tL2FwaS8_bmFtZT1ob29tYW5AdGVzdC5jb20iLCJOYW1lIjoiaG9vbWFuQHRlc3QuY29tIiwiRW1haWwiOiJob29tYW5AdGVzdC5jb20ifX0.AP3WB7ri7JiouUSyV0yhn3CwFAn523hbGSuXerHHzP8"
    }
  })
  .then((response) => response.json())
  .then((json) => console.log(json));


add department:
  fetch("http://localhost:8080/api/v1/department/new", {
    method: "POST",
      body: JSON.stringify({
          title: "hooman@test.com",
          departmenttypeid: "1,
          parentid: 1,
      }),
    headers: {
      "Content-type": "application/json; charset=UTF-8",
      "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVBdCI6MTczNTc0NDMzMSwiSXNzdWVkQXQiOjE3MzU3NDI1MzEsInN1YiI6eyJJRCI6NiwiSW1hZ2UiOiJodHRwczovL3VpLWF2YXRhcnMuY29tL2FwaS8_bmFtZT1ob29tYW5AdGVzdC5jb20iLCJOYW1lIjoiaG9vbWFuQHRlc3QuY29tIiwiRW1haWwiOiJob29tYW5AdGVzdC5jb20ifX0.37Y1k1Qy8y-2k_T1B0Zj_M1MKJOJ8QVy0mT00EMo4HA"
    }
  })
  .then((response) => response.json())
  .then((json) => console.log(json));



//go build -ldflags="-s -w" -o myapp

// CompileDaemon -log-prefix=false -command="go run . serve"
//


