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
          <Link to="/orders">My Order</Link>
          <Link to="/pools">My Pool</Link>
        </div>
      </div>
      <div className="content">
        {children}
      </div>
    </div>
  )
}
