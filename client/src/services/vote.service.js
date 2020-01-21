import store from "@/store/user";

export async function createLaw(params) {
  let response = await fetch("http://localhost:3000/votes/", {
    method: "post",
    body: JSON.stringify(params),
    headers: {
      Authorization: "Bearer " + store.state.token,
      "Content-type": "application/json"
    }
  });
  let data = await response;
  return data;
}

export async function getLaws() {
  let response = await fetch("http://localhost:3000/votes/", {
    method: "get",
    headers: {
      Authorization: "Bearer " + store.state.token,
      "Content-type": "application/json"
    }
  });
  let data = await response;
  return data;
}

export async function getLaw(lawId){

    let response = await fetch(`http://localhost:3000/votes/${lawId}`, {
        method: "get",
        headers: {
            "Authorization" : 'Bearer ' + store.state.token,
            "Content-type": "application/json"
        }
    });
    let data = await response;
    return data;
}
