import axios from "axios";

export const getAllTodos = async (token: string) => {
  await axios
    .get("http://localhost:8000/api/todos", {
      headers: {
        Authorization: `${token}`,
      },
    })
    .then((response) => {
    
      return response.data;
    })
    .catch((error) => {
      console.error(error);
    });
};
