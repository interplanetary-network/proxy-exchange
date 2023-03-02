import { Link } from 'react-router-dom';
import './Layout.scss';

export default function Layout({children}) {
  return (
    <div className="main">
      <div className="header">
        <div className="title">
          <h1><Link to="/">Proxy Exchange</Link></h1>
        </div>
        <div className="nav">
          <Link to="/pools/new">Publish Pool</Link>
          <Link to="/my/orders">My Orders</Link>
          <Link to="/my/pools">My Pools</Link>
          <Link to="/about">About</Link>
        </div>
      </div>
      <div className="content">
        {children}
      </div>
    </div>
  )
}
