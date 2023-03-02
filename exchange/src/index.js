import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.scss';
import App from './App';
import PoolDetail from './pool/PoolDetail';
import reportWebVitals from './reportWebVitals';
import {
  createHashRouter,
  RouterProvider,
} from "react-router-dom";
import OrderDetail from './order/OrderDetail';
import MyOrders from './my/MyOrders';
import MyPools from './my/MyPools';
import PoolNew from './pool/PoolNew';
import About from './About';

const router = createHashRouter([
  {
    path: "/",
    element: <App />,
  },
  {
    path: "/pools/new",
    element: <PoolNew />,
  },
  {
    path: "/my/pools/:poolID",
    element: <PoolDetail />,
  },
  {
    path: "/my/pools",
    element: <MyPools />,
  },
  {
    path: "/my/orders",
    element: <MyOrders />,
  },
  {
    path: "/my/orders/:orderID",
    element: <OrderDetail />,
  },
  {
    path: "/about",
    element: <About />,
  },
]);

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
