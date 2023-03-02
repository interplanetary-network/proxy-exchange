import './Pool.scss';
import Web3 from 'web3';
import { useState } from 'react';
import OrderForm from '../order/OrderForm';
import { toLocal } from '../utils/date.js';

export default function Pool(props) {
  const pricePerMinute = Web3.utils.fromWei(props.pool.pricePerMinute); 
  const [showDetail, setShowDetail] = useState(props.showDetail || false);

  function onInfoClick() {
    setShowDetail(!showDetail);
  }

  return (
    <div className="pool">
      <div className="key" onClick={props.showDetail ? null : onInfoClick}>
        <div className="counter">
          <div className="id" title="Pool ID">{props.pool.id}</div>
          <div className="vote" title="Votes">{props.pool.votes} {props.myVote !== undefined ? <span className="my-vote" title="My vote">{`(${props.myVote})`}</span> : null}</div>
        </div>
        <div className="proxy">
          <div className="location">{props.pool.proxies.map((p,i) => <span key={i}>{p.location}</span>)}</div>
          <div className="protocol">{props.pool.proxies.map((p,i) => <span key={i}>{p.url.substring(0, p.url.indexOf('://'))}</span>)}</div>
        </div>
        <div className="action">
          <div className="pricing"><span>{pricePerMinute} ETH</span><span className="unit">/min</span></div>
          <div className="valid-before-at" title="Valid before this time">{toLocal(props.pool.validBeforeAt)}</div>
        </div>
      </div>
      <div className={"detail " + (showDetail ? "show" : "hidden")}>
        <div className="proxies">
          {props.pool.proxies.map((p,i) => {
            return (
              <div key={i} className={"proxy " + (i%2 ? "odd" : "even") }>
                <span className="url">{p.url}</span>
                <span className="location">{p.location}</span>
                {/*<span className="ping">{"53ms"}</span>*/}
              </div>
            )
          })}
        </div>
        <div className={"checkout-form"}>
          {props.withoutCheckoutForm ? null : <OrderForm pool={props.pool} />}
        </div>
      </div>
    </div>
  );
}
