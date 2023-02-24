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
import MyOrder from './my/MyOrder';
import MyPool from './my/MyPool';
import PoolNew from './pool/PoolNew';

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
    path: "/pools/:poolID",
    element: <PoolDetail />,
  },
  {
    path: "/pools",
    element: <MyPool />,
  },
  {
    path: "/orders",
    element: <MyOrder />,
  },
  {
    path: "/orders/:orderID",
    element: <OrderDetail />,
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
