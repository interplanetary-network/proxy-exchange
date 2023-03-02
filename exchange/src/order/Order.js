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
  const [isVoted, setIsVoted] = useState(false);
  const [myVote, setMyVote] = useState(0);
  const [voteChanged, setVoteChanged] = useState(0);
  const [voteInLoading, setVoteInLoading] = useState(false);
  const [voteError, setVoteError] = useState("");

  useEffect(() => {
    async function fetchPool() {
      // reset my vote state
      setMyVote(0);
      const contract = new ProxyExchange();
      const accounts = await contract.requestAccounts();
      const vote = await contract.voteOfUserAndOrder(accounts[0], order.id);
      setIsVoted(vote.voted);
      if (vote.voted) {
        setMyVote(vote.t > 0 ? -1 : 1);
      }

      const rawPool = await contract.poolOf(order.poolID);
      const pool = {
        id: order.poolID,
        pricePerMinute: rawPool.pricePerMinute,
        proxies: [],
        validBeforeAt: rawPool.validBeforeAt,
        provider: rawPool.provider,
        votes: rawPool.votes,
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
  }, [order, voteChanged])

  async function onDownVote() {
    setVoteInLoading(true);
    setVoteError("");
    const contract = new ProxyExchange();
    const accounts = await contract.requestAccounts();
    await contract.downVote(order.id, accounts[0])
      .on("transactionHash", (hash) => {console.log(hash)})
      .on("receipt", (receipt) => {
        console.log(receipt);
        setVoteChanged((c) => ++c);
      })
      .on("error", (error) => {
        console.log(error);
        setVoteError(error.message);
      })
      .finally(() => {
        setVoteInLoading(false);
      })
  }

  async function onUpVote() {
    setVoteInLoading(true);
    setVoteError("");
    const contract = new ProxyExchange();
    const accounts = await contract.requestAccounts();
    await contract.upVote(order.id, accounts[0])
      .on("transactionHash", (hash) => {console.log(hash)})
      .on("receipt", (receipt) => {
        console.log(receipt);
        setVoteChanged((c) => ++c);
      })
      .on("error", (error) => {
        console.log(error);
        setVoteError(error.message);
      })
      .finally(() => {
        setVoteInLoading(false);
      })
  }

  async function onRecallVote() {
    setVoteInLoading(true);
    setVoteError("");
    const contract = new ProxyExchange();
    const accounts = await contract.requestAccounts();
    await contract.recallVote(order.id, accounts[0])
      .on("transactionHash", (hash) => {console.log(hash)})
      .on("receipt", (receipt) => {
        console.log(receipt);
        setVoteChanged((c) => ++c);
      })
      .on("error", (error) => {
        console.log(error);
        setVoteError(error.message);
      })
      .finally(() => {
        setVoteInLoading(false);
      })
  }

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
            {pool !== undefined ? <Pool pool={pool} showDetail={true} withoutCheckoutForm={true} myVote={myVote}/> : <Loading />}
          </div>
        </div>
      </div>
      <div className="actions">
        <div className="vote">
          { voteInLoading ? (
            <Loading withoutText={true} />
          ) : (
            isVoted ? (
              <button onClick={onRecallVote}>Recall Vote</button>
            ) : (
              <>
                <button onClick={onDownVote}>Down Vote</button>
                <button onClick={onUpVote}>Up Vote</button>
              </>
            )
          )}
          <div className="error">{voteError}</div>
        </div>
        <div>
          <button onClick={onGenerateAuthentication}>Generate Authentication</button>
        </div>
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
