export async function login(params){

    let response = fetch("http://localhost:3000/login", {
        mode: 'cors',
        method: "post",
        body: JSON.stringify(params),
        headers: {
          "Content-type": "application/json"
        }
      });
      let data = await response.json();
      return data;
}