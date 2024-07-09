import { useEffect, useState } from "react";

import axios from "axios";

const Home = () => {
  const [todos, setTodos] = useState();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      window.location.href = "/login";
    }
  }, []);

  const token = localStorage.getItem("token");

  const fetchTodos = async () => {
    await axios
      .get("http://localhost:8000/api/todos", {
        headers: {
          Authorization: `${token}`,
        },
      })
      .then((response) => {
        console.log(response);
        console.log(response.data);
        setTodos(response.data);
      })
      .catch((error) => {
        console.error(error);
      });
  };

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      fetchTodos();
    }
  }, []);

  console.log(todos);
  return (
    <div className="min-h-screen flex items-center justify-center w-full ">
      <div className="bg-white/40  backdrop-blur-lg rounded-lg shadow-lg drop-shadow-2xl shadow-black p-8 max-w-sm w-full border-2 border-purple-600">
        <h2 className="text-2xl font-bold mb-6 text-black text-center">
          Todos
        </h2>
        <ul>
          {todos && todos.map((todo: any) => (
            <li key={todo.id}>{todo.title}</li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default Home;
