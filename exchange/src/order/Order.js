import { useEffect, useState } from "react";
import Web3 from "web3";
import ProxyExchange from "../contracts/ProxyExchange";
import Loading from "../Loading";
import Pool from "../pool/Pool";
import "./Order.scss";

export default function Order({ order }) {
  const startAt = new Date(order.startAt * 1000);
  const [pool, setPool] = useState();
  const [sign, setSign] = useState("");

  useEffect(() => {
    async function fetchPool() {
      const contract = new ProxyExchange();
      const rawPool = await contract.poolOf(order.poolID);
      const pool = {
        id: order.poolID,
        pricePerMinute: rawPool.pricePerMinute,
        proxies: [],
        validBeforeAt: rawPool.validBeforeAt,
        provider: rawPool.provider,
        vote: rawPool.vote,
      }

      // fetch proxies
      for(let ii=0; ii<rawPool.proxies.length; ii++) {
        const proxy = await contract.proxyOf(rawPool.proxies[ii]);
        pool.proxies.push({
          location: Web3.utils.hexToAscii(proxy.location),
          url: proxy.url,
        })
      }

      setPool(pool);
    }

    fetchPool();
  }, [order])

  async function onGenerateAuthentication() {
    const contract = new ProxyExchange();
    const message = `${order.id}`
    const accounts = await contract.requestAccounts();
    const sign = await contract.sign(message, accounts[0])
    setSign(sign)
  }

  return (
    <div className="order">
      <div className="info">
        <div className="row">
          <div className="col">
            <label>Start At:</label>
            <div>{`${startAt.getFullYear()}-${startAt.getMonth()}-${startAt.getDate()}T${startAt.getHours()}:${startAt.getMinutes()}`}</div>
          </div>
          <div className="col">
            <label>Duration (minute):</label>
            <div>{order.duration}</div>
          </div>
        </div>
        <div className="row">
          <div className="col">
            <label>Provider:</label>
            <div>{order.provider}</div>
          </div>
          <div className="col">
            <label>Customer:</label>
            <div>{order.customer}</div>
          </div>
        </div>
        <div className="row">
          <div className="col">
            <label>Pool:</label>
            {pool !== undefined ? <Pool pool={pool} showDetail={true} withoutCheckoutForm={true} /> : <Loading />}
          </div>
        </div>
      </div>
      <div className="actions">
        <button onClick={onGenerateAuthentication}>Generate Authentication</button>
      </div>
      <div className="authentication">
        {sign !== "" ? (
          <>
            <pre>{`Username: ${order.id}\nPassword: ${sign}`}</pre>
            <span><strong>DO NOT</strong> share the password!</span>
          </>
        ) : null
        }
      </div>
    </div>
  )
}
