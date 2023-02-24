import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Web3 from "web3";
import ProxyExchange from "../contracts/ProxyExchange";
import Layout from "../Layout";
import Loading from "../Loading";
import Pool from "./Pool";

export default function PoolDetail() {
  const { poolID }= useParams()
  const [pool, setPool] = useState();

  useEffect(() => {
    async function fetchPool() {
      const contract = new ProxyExchange();
      const rawPool = await contract.poolOf(poolID);
      const pool = {
        id: poolID,
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
  }, [poolID])

  return (
    <Layout>
      <div className="pool-detail">
        {pool !== undefined ? (
          <>
            <h2>{`Pool ${pool.id}`}</h2>
            <Pool pool={pool} showDetail={true} />
          </>
        ) : (
          <Loading />
        )}
      </div>
    </Layout>
  )
}
