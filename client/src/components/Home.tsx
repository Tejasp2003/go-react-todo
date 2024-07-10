import axios from "axios";
import { useEffect, useState } from "react";
import { AiFillEdit, AiFillDelete, AiOutlineCheck } from "react-icons/ai";
import { useNavigate } from "react-router-dom";

interface Todo {
  id: string;
  title: string;
  done: boolean;
}

const Home = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [editTodoId, setEditTodoId] = useState<string | null>(null);
  const [editTodoText, setEditTodoText] = useState<string>("");
  const [newTodoText, setNewTodoText] = useState<string>("");

  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      window.location.href = "/login";
    } else {
      fetchTodos();
    }
  }, []);

  const token = localStorage.getItem("token");

  const fetchTodos = async () => {
    try {
      const response = await axios.get("http://localhost:8000/api/todos", {
        headers: {
          Authorization: `${token}`,
        },
      });
      setTodos(response.data);
    } catch (error) {
      console.error(error);
    }
  };

  const markTodoDone = async (id: string, done: boolean) => {
    try {
      await axios.put(`http://localhost:8000/api/todos/${id}/complete`, {
        done: !done,
      }, {
        headers: {
          Authorization: `${token}`,
        },
      });
      fetchTodos();
    } catch (error) {
      console.error(error);
    }
  };

  const deleteTodo = async (id: string) => {
    try {
      await axios.delete(`http://localhost:8000/api/todos/${id}`, {
        headers: {
          Authorization: `${token}`,
        },
      });
      fetchTodos();
    } catch (error) {
      console.error(error);
    }
  };

  const editTodo = (id: string, text: string) => {
    setEditTodoId(id);
    setEditTodoText(text);
  };

  const saveTodo = async (id: string) => {
    try {
      await axios.put(`http://localhost:8000/api/todos/${id}`, {
        title: editTodoText,
      }, {
        headers: {
          Authorization: `${token}`,
        },
      });
      setEditTodoId(null);
      setEditTodoText("");
      fetchTodos();
    } catch (error) {
      console.error(error);
    }
  };

  const addTodo = async () => {
    try {
      await axios.post(`http://localhost:8000/api/todos`, {
        title: newTodoText,
        done: false,
      }, {
        headers: {
          Authorization: `${token}`,
        },
      });
      setNewTodoText("");
      fetchTodos();
    } catch (error) {
      console.error(error);
    }
  };

  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate("/login");
  };

  return (
    <div className="min-h-screen flex items-center justify-center w-full bg-gray-100">
      <div className="bg-white/60 backdrop-blur-lg rounded-lg shadow-lg drop-shadow-2xl shadow-black p-8 max-w-lg w-full border-2 border-purple-600">
        <h2 className="text-3xl font-bold mb-6 text-black text-center">Todos</h2>
        <div className="mb-4 flex justify-between items-center">
          <input
            type="text"
            className="w-2/3 p-2 mb-2 border rounded-lg focus:outline-none"
            placeholder="Add a new todo"
            value={newTodoText}
            onChange={(e) => setNewTodoText(e.target.value)}
          />
          <button
            onClick={addTodo}
            className="bg-purple-600 text-white p-2 rounded-lg hover:bg-purple-700 ml-2"
          >
            Add Todo
          </button>
          <button
            onClick={handleLogout}
            className="bg-red-500 text-white p-2 rounded-lg hover:bg-red-600 ml-2"
          >
            Logout
          </button>
        </div>
        <ul>
          {todos && todos.map((todo) => (
            <li key={todo.id} className={`flex items-center justify-between p-2 mb-2 border rounded-lg ${todo.done ? "bg-green-100" : "bg-red-100"}`}>
              <input
                type="checkbox"
                checked={todo.done}
                onChange={() => markTodoDone(todo.id, todo.done)}
                className="mr-2"
              />
              {editTodoId === todo.id ? (
                <input
                  type="text"
                  className="flex-1 p-2 border rounded-lg focus:outline-none"
                  value={editTodoText}
                  onChange={(e) => setEditTodoText(e.target.value)}
                />
              ) : (
                <span className={`flex-1 ${todo.done ? "line-through" : ""}`}>
                  {todo.title}
                </span>
              )}
              <div className="flex space-x-2">
                {editTodoId === todo.id ? (
                  <button onClick={() => saveTodo(todo.id)} className="text-green-500 hover:text-green-700">
                    <AiOutlineCheck size={24} />
                  </button>
                ) : (
                  <>
                    <button onClick={() => editTodo(todo.id, todo.title)} className="text-blue-500 hover:text-blue-700">
                      <AiFillEdit size={24} />
                    </button>
                    <button onClick={() => deleteTodo(todo.id)} className="text-red-500 hover:text-red-700">
                      <AiFillDelete size={24} />
                    </button>
                  </>
                )}
              </div>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default Home;
