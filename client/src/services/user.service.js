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