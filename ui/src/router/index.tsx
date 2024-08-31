import React from "react";
import { createBrowserRouter } from "react-router-dom";

import RootLayout from "@/layouts/RootLayout";
import ChatLayout from "@/layouts/ChatLayout";

import Home from "../pages/home";
import Login from "@/pages/login";
import Chat from "@/pages/chat";

const ErrorPage = React.lazy(() => import("../pages/error"))
const NotFound = React.lazy(() => import("../pages/404"))

const routes = createBrowserRouter([
  {
    element: <RootLayout />,
    children: [
      {
        index: true,
        path: "/",
        element: <Home />
      }
    ]
  },
  {
    element: <ChatLayout />,
    children: [
      {
        path: "/chat",
        element: <Chat />
      }
    ]
  },
  {
    path: "/login",
    element: <Login />
  },
  {
    path: "/error",
    element: <ErrorPage />
  },
  {
    path: "*",
    element: <NotFound />
  }
])


export default routes