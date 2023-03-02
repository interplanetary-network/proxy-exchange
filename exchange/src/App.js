import './App.scss';
import Web3 from 'web3';
import { useEffect, useState } from 'react';
import Pool from './pool/Pool';
import ProxyExchange from './contracts/ProxyExchange.js';
import Layout from './Layout';
import Loading from './Loading';

export default function App() {
  const [pools, setPools] = useState([]);
  const [inLoading, setInLoading] = useState(false);

  useEffect(() => {
    async function fetchPools() {
      setInLoading(true);
      const contract = new ProxyExchange();
      let pools = [];
      const totalPool = await contract.totalPool();
      for (let i=0; i<totalPool; i++) {
        const rawPool = await contract.poolOfIndex(i);
        const pool = {
          id: rawPool[0],
          pricePerMinute: rawPool[1].pricePerMinute,
          proxies: [],
          validBeforeAt: rawPool[1].validBeforeAt,
          provider: rawPool[1].provider,
          votes: rawPool[1].votes,
        }

        // fetch proxies
        for(let ii=0; ii<rawPool[1].proxies.length; ii++) {
          const proxy = await contract.proxyOf(rawPool[1].proxies[ii]);
          pool.proxies.push({
            location: Web3.utils.hexToAscii(proxy.location),
            url: proxy.url,
          })
        }

        pools.push(pool);
      }

      setInLoading(false);
      setPools(pools);
    }

    fetchPools();
  }, []);

  return (
    <Layout>
      <div className="home">
        <div className="pools">
          {inLoading ? (
            <Loading />
          ) : (
            pools.length > 0 ? pools.map((pool, i) => <Pool pool={pool} key={i} />) : <div>No pools</div>
          )}
        </div>
      </div>
    </Layout>
  );
}
