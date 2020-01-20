import store from "@/store/user";

export async function login(params){

    let response = await fetch("http://localhost:3000/login/", {
        method: "post",
        body: JSON.stringify(params),
        headers: {
          "Content-type": "application/json"
        }
      });
      let data = await response;
      return data;
}

export async function createUser(params){

  let response = await fetch("http://localhost:3000/users/", {
    method: "post",
    body: JSON.stringify(params),
    headers: {
      "Authorization" : 'Bearer ' + store.state.token,
      "Content-type": "application/json"
    }
  });
  let data = await response;
  return data;
}

export async function getUsers(){
  
  let response = await fetch("http://localhost:3000/users/", {
    method: "get",
    headers: {
      "Authorization" : 'Bearer ' + store.state.token,
      "Content-type": "application/json"
    }
  });
  let data = await response;
  return data;
}