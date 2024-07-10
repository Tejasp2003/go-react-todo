import axios from "axios";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

const Login = () => {

    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')

    const navigate = useNavigate();

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault()
        axios
        .post('http://localhost:8000/login', { email, password })
        .then((response) => {
          localStorage.setItem('token', response.data.token);
          navigate('/');
        })
        .catch((error) => {
          console.error(error);
        });

    }


    useEffect(() => {
        const token = localStorage.getItem('token')
      if(token){
        axios.get('http://localhost:8000/api/user', {
          headers: {
            Authorization: `${token}`
          }

        }).then((response) => {
          console.log(response)
          navigate('/')
        }
        ).catch((error) => {
          console.error(error)
        })

      }
    }
    , [])

  return (
    <div className="min-h-screen flex items-center justify-center w-full ">
      <div className="bg-white/40  backdrop-blur-lg rounded-lg shadow-lg drop-shadow-2xl shadow-black p-8 max-w-sm w-full border-2 border-purple-600">
        <h2 className="text-2xl font-bold mb-6 text-black text-center">
          Login
        </h2>
        <form
            className="space-y-4"
            onSubmit={handleSubmit}
        >
          <div className="mb-4">
            <label
              htmlFor="email"
              className="block text-black text-sm font-medium mb-2"
            >
              Email
            </label>
            <input
              type="email"
              id="email"
              className="w-full px-4 py-2 rounded-lg bg-white bg-opacity-20 text-black placeholder-black/80 focus:outline-none focus:ring-2 focus:ring-purple-500"
              placeholder="Enter your email"
                onChange={(e) => setEmail(e.target.value)}
                required
            />
          </div>
          <div className="mb-6">
            <label
              htmlFor="password"
              className="block text-black text-sm font-medium mb-2"
            >
              Password
            </label>
            <input
              type="password"
              id="password"
              className="w-full px-4 py-2 rounded-lg bg-white bg-opacity-20 text-black placeholder-black/80 focus:outline-none focus:ring-2 focus:ring-purple-500"
              placeholder="Enter your password"
                onChange={(e) => setPassword(e.target.value)}
                required
            />
          </div>
          <button
            type="submit"
            className="w-full py-2 bg-purple-400 hover:bg-purple-500 text-black font-bold rounded-lg transition duration-300 ease-in-out"
          >
            Login
          </button>
        </form>
      </div>
    </div>
  );
};

export default Login;
