import './OrderDetail.scss';
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import ProxyExchange from '../contracts/ProxyExchange';
import Layout from '../Layout';
import Order from './Order';
import Loading from '../Loading';

export default function OrderDetail() {
  const { orderID } = useParams();
  const [order, setOrder] = useState();

  useEffect(() => {
    async function fetchOrder() {
      const contract = new ProxyExchange();
      const rawOrder = await contract.orderOf(orderID);
      console.log(rawOrder)
      setOrder({...rawOrder, ...{id: orderID}})
    }

    fetchOrder();
  }, [orderID])

  return (
    <Layout> 
      <div className="order-detail">
        {order !== undefined ? (
          <>
            <h2>{`Order ${order.id}`}</h2>
            <Order order={order}/>
          </>
        ) : (
          <Loading />
        )}
      </div>
    </Layout>
  );
}
