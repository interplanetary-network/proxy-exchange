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

  orderOfUserAndIndex(user, index) {
    return this.contract.methods.orderOfUserAndIndex(user, index).call();
  }

  totalOrderOfUser(user) {
    return this.contract.methods.totalOrderOfUser(user).call();
  }

  poolOfUserAndIndex(user, index) {
    return this.contract.methods.poolOfUserAndIndex(user, index).call();
  }

  totalPoolOfUser(user) {
    return this.contract.methods.totalPoolOfUser(user).call();
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

  voteOfUserAndOrder(user, orderID) {
    return this.contract.methods.voteOfUserAndOrder(user, orderID).call();
  }

  upVote(orderID, from) {
    return this.contract.methods.upVote(orderID).send({from: from});
  }

  downVote(orderID, from) {
    return this.contract.methods.downVote(orderID).send({from: from});
  }

  recallVote(orderID, from) {
    return this.contract.methods.recallVote(orderID).send({from: from});
  }

  buy(id, startAt, duration, value, from) {
    return this.contract.methods.buy(id, startAt, duration).send({from: from, value: value})
  }

  publish( pricePerMinute, validBeforeAt, proxies, from) {
    return this.contract.methods.publish(pricePerMinute, validBeforeAt, proxies).send({from: from})
  }
}

