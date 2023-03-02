import { useEffect, useState } from 'react';
import ProxyExchange from '../contracts/ProxyExchange.js';
import Layout from '../Layout.js';
import {toLocal, isPast } from '../utils/date.js';
import Loading from '../Loading.js';
import './MyOrders.scss';
import { Link } from 'react-router-dom';

export default function MyOrders() {
  const [orders, setOrders] = useState([]);
  const [inLoading, setInLoading] = useState(false);

  useEffect(() => {
    async function fetchMyOrders() {
      setInLoading(true);
      const contract = new ProxyExchange();
      const accounts = await contract.requestAccounts();
      const total = await contract.totalOrderOfUser(accounts[0])

      const orders = []
      for (let i = (total - 1); i >= 0; i--) {
        const rawOrder = await contract.orderOfUserAndIndex(accounts[0], i)
        const order = {...rawOrder[1], ...{id: rawOrder[0]}}
        orders.push(order) 
      }
      setInLoading(false);
      setOrders(orders)
    }

    fetchMyOrders()
  }, [])

  return (
    <Layout>
      <div className="my-order">
        <h1>My Order</h1>
        <div className="list">
          {inLoading ? (
            <Loading />
          ) : (
            orders.length > 0 ? (
              orders.map((order, index) => {
                return (
                  <div className={"item " + (isPast(Number(order.startAt) + (Number(order.duration) * 60)) ? 'expired' : 'active')} key={index} >
                    <div className="id">{order.id}</div>
                    <div className="start-at">{toLocal(order.startAt)}</div>
                    <div className="duration">{order.duration}</div>
                    <div className="status">{isPast(Number(order.startAt) + (Number(order.duration) * 60)) ? 'expired' : 'active'}</div>
                    <div className="actions">
                      <Link to={`/my/orders/${order.id}`}>Detail</Link>
                    </div>
                  </div>
                )
              })
            ) : (
              <div>No order found</div>
            )
          )}
        </div>
      </div>
    </Layout>
  )
}
