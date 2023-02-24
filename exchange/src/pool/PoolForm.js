import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Web3 from 'web3';
import ProxyExchange from '../contracts/ProxyExchange';
import Loading from '../Loading';
import './PoolForm.scss'

export default function PoolForm() {
  const navigate = useNavigate();

  const [transactionHash, setTransactionHash] = useState("");
  const [inLoading, setInLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");

  async function onSubmit(e) {
    e.preventDefault();
    setTransactionHash("")
    setErrorMessage("")
    setInLoading(false)

    const formData = new FormData(e.target);
    const data = Object.fromEntries(formData.entries());
    const proxies = data.proxies.split('\n').map(proxy => {
      const [url, location] = proxy.split(',');
      return [url, Web3.utils.asciiToHex(location)];
    })
    const valiBeforeAt = new Date(data.validBeforeAt).getTime() / 1000;
    const pricePerMinute = Web3.utils.toWei(data.pricePerMinute, 'ether');
    const contract = new ProxyExchange();
    const accounts = await contract.requestAccounts();
    contract.publish(pricePerMinute,  valiBeforeAt, proxies, accounts[0])
      .on("sending", function(){
        setInLoading(true)
      })
      .on("transactionHash", function(hash){
        console.log(hash)
        setTransactionHash(hash)
      })
      .on("receipt", function(receipt){
        navigate(`/pools/${receipt.events.Publish.returnValues.poolID}`)
        console.log(receipt)
      })
      .on("error", function(error) {
        setErrorMessage(error.message)
        setInLoading(false)
        console.log(error)
      })
  }

  return (
    <form className="pool-form" action="" onSubmit={onSubmit}>
      <div className="proxies">
        <label>Proxies:</label>
        <div>
          <textarea required rows="5" placeholder="socks5://us.example.com:1080,CN&#10;http://jp.example.com:8080,JP&#10;..." name="proxies"></textarea>
        </div>
      </div>
      <div className="pool-info">
        <div>
          <label>Valid Before At:</label>
          <div>
            <input type="datetime-local" required name="validBeforeAt"/>
          </div>
        </div>
        <div>
          <label>Price Per Minute (ETH):</label>
          <div>
            <input type="number" pattern="^[0-9]" required name="pricePerMinute"/>
          </div>
        </div>
      </div>
      <div>
        <input type="reset" />&nbsp;<input type="submit" />
      </div>
      <div className={"message " + (errorMessage !== "" ? "error" : "")}>{transactionHash} {inLoading ? <Loading withoutText={true} /> : null} {errorMessage}</div>
    </form>
  )
}
