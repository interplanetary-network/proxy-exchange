import abi from './ProxyExchange.json';
import Web3 from 'web3';

export default class ProxyExchange {
  constructor() {
    this.web3 = new Web3(Web3.givenProvider || process.env.REACT_APP_GATEWAY);  
    this.contract = new this.web3.eth.Contract(abi, process.env.REACT_APP_CONTRACT_ADDRESS);
  }

  sign(message, from) {
    return this.web3.eth.personal.sign(message, from);
  }

  requestAccounts() {
    return this.web3.eth.requestAccounts();
  }

  orderOfUserAndIndex(address, index) {
    return this.contract.methods.orderOfUserAndIndex(address, index).call();
  }

  totalOrderOfUser(address) {
    return this.contract.methods.totalOrderOfUser(address).call();
  }

  poolOfUserAndIndex(address, index) {
    return this.contract.methods.poolOfUserAndIndex(address, index).call();
  }

  totalPoolOfUser(address) {
    return this.contract.methods.totalPoolOfUser(address).call();
  }

  proxyOf(id) {
    return this.contract.methods.proxyOf(id).call();
  }

  orderOf(id) {
    return this.contract.methods.orderOf(id).call();
  }

  totalPool() {
    return this.contract.methods.totalPool().call();
  }

  poolOfIndex(index) {
    return this.contract.methods.poolOfIndex(index).call();
  }

  poolOf(id) {
    return this.contract.methods.poolOf(id).call();
  }

  buy(id, startAt, duration, value, from) {
    return this.contract.methods.buy(id, startAt, duration).send({from: from, value: value})
  }

  publish( pricePerMinute, validBeforeAt, proxies, from) {
    return this.contract.methods.publish(pricePerMinute, validBeforeAt, proxies).send({from: from})
  }
}

