import {
  createBrowserRouter,
  createRoutesFromElements,
  Route,
  RouterProvider,
} from "react-router-dom";
import Login from "./components/auth/Login";
import Home from "./components/Home";
import Register from "./components/auth/Register";

const App = () => {
  const router = createBrowserRouter(
    createRoutesFromElements(
      <>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/" element={<Home />} />
      </>
    )
  );

  return (
    <div className="min-h-screen w-full flex items-center justify-center bg-pink-200">
      <RouterProvider router={router} />
    </div>
  );
};

export default App;
