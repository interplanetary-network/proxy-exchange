import { useEffect, useState } from 'react';
import ProxyExchange from '../contracts/ProxyExchange.js';
import Layout from '../Layout.js';
import Loading from '../Loading.js';
import './MyOrder.scss';
import Pool from '../pool/Pool.js';
import Web3 from 'web3';

export default function MyPool() {
  const [pools, setPools] = useState([]);
  const [inLoading, setInLoading] = useState(false);

  useEffect(() => {
    async function fetchMyPools() {
      setInLoading(true);
      const contract = new ProxyExchange();
      const accounts = await contract.requestAccounts();
      const total = await contract.totalPoolOfUser(accounts[0])

      const pools = []
      for (let i = (total - 1); i >= 0; i--) {
        const rawPool = await contract.poolOfUserAndIndex(accounts[0], i)
        const pool = {
          id: rawPool[0],
          pricePerMinute: rawPool[1].pricePerMinute,
          proxies: [],
          validBeforeAt: rawPool[1].validBeforeAt,
          provider: rawPool[1].provider,
          vote: rawPool[1].vote,
        }

        // fetch proxies
        for(let ii=0; ii<rawPool[1].proxies.length; ii++) {
          const proxy = await contract.proxyOf(rawPool[1].proxies[ii]);
          pool.proxies.push({
            location: Web3.utils.hexToAscii(proxy.location),
            url: proxy.url,
          })
        }
        pools.push(pool) 
      }
      setInLoading(false);
      setPools(pools)
    }

    fetchMyPools()
  }, [])

  return (
    <Layout>
      <div className="my-pool">
        <h1>My Pool</h1>
        <div className="list">
          {inLoading ? (
            <Loading />
          ) : (
            pools.length > 0 ? (
              pools.map((pool, index) => {
                return (
                  <Pool pool={pool} key={index}/>
                )
              })
            ) : (
              <div>No pool found.</div>
            )
          )}
        </div>
      </div>
    </Layout>
  )
}
