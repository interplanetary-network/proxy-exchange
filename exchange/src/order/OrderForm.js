import { useState } from "react";
import { useNavigate } from "react-router-dom";
import Web3 from "web3";
import ProxyExchange from "../contracts/ProxyExchange";
import Loading from "../Loading";
import './OrderForm.scss';

export default function OrderForm({ pool }) {
  const [transactionHash, setTransactionHash] = useState("");
  const [error, setError] = useState("");
  const [inLoading, setInLoading] = useState(false);

  const contract = new ProxyExchange();
  const navigate = useNavigate();

  async function onSubmit(e) {
    e.preventDefault();
    setInLoading(true);
    setError("");
    setTransactionHash("");

    const formData = new FormData(e.target);
    const data = Object.fromEntries(formData.entries());

    const startAt = (new Date(data.startAt)).getTime() / 1000;
    const accounts = await contract.requestAccounts();
    const value = Web3.utils.toBN(pool.pricePerMinute).mul(Web3.utils.toBN(data.duration)).toString();

    contract.buy(pool.id, startAt, data.duration, value, accounts[0])
      .on("transactionHash", (hash) => {
        setTransactionHash(hash);
      })
      .on("receipt", (receipt) => {
        console.log(receipt);
        setInLoading(false);
        navigate(`/orders/${receipt.events.Buy.returnValues.orderID}`)
      })
      .on("error",  (error) => {
        setInLoading(false);
        setError(error.message);
      })
  }

  return (
    <form className="order-form" action="" onSubmit={onSubmit}>
      <label>Start At:
        <input name="startAt" type="datetime-local" required />
      </label>
      <label>Duration (minute):
        <input name="duration" type="number" required min="1" step="1" />
      </label>
      <div className="action">
        <input type="submit" value="Checkout" />
      </div>
      <div className={"message"}>
        <div className="error">{error}</div>
        <div>{transactionHash}</div>
        {inLoading ? <Loading withoutText={true} /> : null}
      </div>
    </form>
  )
}
