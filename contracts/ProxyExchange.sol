// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

contract ProxyExchange is Initializable, OwnableUpgradeable, UUPSUpgradeable {
    using Counters for Counters.Counter;

    struct Proxy {
        string url;
        // ISO 3166-1
        bytes2 location;
    }

    struct Pool {
        address provider;
        uint256[] proxies;
        // Wei
        uint256 pricePerMinute;
        // UNIX timestamp in seconds
        uint64 validBeforeAt;
        Counters.Counter vote;
    }

    struct Order {
        uint256 poolID;
        // startAt UNIX timestamp in seconds, seconds since JAN 01 1970 (UTC)
        uint64 startAt;
        // number of minute
        uint16 duration;
        address customer;
        address provider;
    }

    enum VoteType {
        Up,
        Down
    }

    struct Vote {
        VoteType t;
        bool voted;
    }

    event Publish(uint256 poolID);
    event Buy(uint256 orderID);

    mapping(address => uint256[]) private _myOrders;
    mapping(address => uint256[]) private _myPools;

    mapping(uint256 => Order) private _ordersMap;
    mapping(uint256 => Pool) private _poolsMap;
    mapping(uint256 => Proxy) private _proxiesMap;

    mapping(uint256 => mapping(address => Vote)) private _poolsVoteMap;

    uint256[] private _orders;
    uint256[] private _pools;

    // Global pool ID
    Counters.Counter private _poolID;

    // Global order ID
    Counters.Counter private _orderID;

    // Global proxy ID
    Counters.Counter private _proxyID;

    // Commission rate
    uint8 private _commissionRate;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize() public initializer {
        __Ownable_init();
        __UUPSUpgradeable_init();
    }

    function _authorizeUpgrade(address newImplementation)
        internal
        override
        onlyOwner
    {}

    function publish(
        uint256 pricePerMinute,
        uint64 validBeforeAt,
        Proxy[] calldata proxies
    ) external onlyProxy returns (uint256, Pool memory) {
        require(proxies.length > 0, "Proxies is required");

        Pool memory pool;
        pool.proxies = new uint256[](proxies.length);

        for (uint256 i = 0; i < proxies.length; i++) {
            // check location code
            require(
                proxies[i].location.length == 2 &&
                    proxies[i].location >= bytes2("AA") &&
                    proxies[i].location <= bytes2("ZZ"),
                "Invalid location code"
            );

            _proxyID.increment();
            uint256 proxyID = _proxyID.current();

            _proxiesMap[proxyID] = Proxy({
                url: proxies[i].url,
                location: proxies[i].location
            });

            pool.proxies[i] = proxyID;
        }

        pool.pricePerMinute = pricePerMinute;
        pool.provider = msg.sender;
        pool.validBeforeAt = validBeforeAt;

        _poolID.increment();
        uint256 poolID = _poolID.current();

        _myPools[msg.sender].push(poolID);
        _pools.push(poolID);
        _poolsMap[poolID] = pool;

        emit Publish(poolID);

        return (poolID, pool);
    }

    function buy(
        uint256 poolID,
        uint64 startAt,
        uint16 duration
    ) external payable onlyProxy returns (uint256, Order memory) {
        Pool memory pool = _poolsMap[poolID];
        require(pool.provider != address(0), "Invalid pool ID");

        uint256 totalPrice = (pool.pricePerMinute * duration);
        require(msg.value == totalPrice, "Paid value not match total price");

        uint256 commission = msg.value * _commissionRate / 100;

        // transfer to provider
        payable(pool.provider).transfer(msg.value - commission);

        Order memory o;
        o.poolID = poolID;
        o.duration = duration;
        o.startAt = startAt;
        o.customer = msg.sender;
        o.provider = pool.provider;

        _orderID.increment();
        uint256 oid = _orderID.current();

        _ordersMap[oid] = (o);
        _orders.push(oid);

        _myOrders[msg.sender].push(oid); // For consumer
        _myOrders[pool.provider].push(oid); // For provider

        emit Buy(oid);

        return (oid, o);
    }

    function totalOrder() external view onlyProxy returns (uint256) {
        return _orders.length;
    }

    function totalPool() external view onlyProxy returns (uint256) {
        return _pools.length;
    }

    function totalOrderOfUser(address user)
        external
        view
        onlyProxy
        returns (uint256)
    {
        return _myOrders[user].length;
    }

    function orderOfUserAndIndex(address user, uint256 index)
        external
        view
        onlyProxy
        returns (uint256, Order memory)
    {
        uint256 id = _myOrders[user][index];
        return (id, _ordersMap[id]);
    }

    function orderOf(uint256 id)
        external
        view
        onlyProxy
        returns (Order memory)
    {
        return _ordersMap[id];
    }

    function totalPoolOfUser(address user)
        external
        view
        onlyProxy
        returns (uint256)
    {
        return _myPools[user].length;
    }

    function poolOfUserAndIndex(address user, uint256 index)
        external
        view
        onlyProxy
        returns (uint256, Pool memory)
    {
        uint256 id = _myPools[user][index];
        return (id, _poolsMap[id]);
    }

    function poolOfIndex(uint256 index)
        external
        view
        onlyProxy
        returns (uint256, Pool memory)
    {
        uint256 id = _pools[index];
        return (id, _poolsMap[id]);
    }

    function poolOf(uint256 id) external view onlyProxy returns (Pool memory) {
        return _poolsMap[id];
    }

    function proxyOf(uint256 id) external view onlyProxy returns (Proxy memory) {
        return _proxiesMap[id];
    }

    function setPoolValidBeforeAt(uint256 poolID, uint64 validBeforeAt)
        external
        onlyProxy
    {
        require(_poolsMap[poolID].provider != msg.sender, "Permission denied");
        _poolsMap[poolID].validBeforeAt = validBeforeAt;
    }

    function setPoolPricePerMinute(uint256 poolID, uint256 pricePerMinute) external onlyProxy {
        require(_poolsMap[poolID].provider != msg.sender, "Permission denied");
        _poolsMap[poolID].pricePerMinute = pricePerMinute;
    }

    function commissionRate() external view onlyProxy returns (uint8) {
        return _commissionRate;
    }

    function setCommissionRate(uint8 v) external onlyOwner onlyProxy {
        require(v <= 20, "Invalid commission rate");
        _commissionRate = v;
    }

    function withdraw(uint256 amount) external onlyOwner onlyProxy {
        require((address(this).balance) >= amount, "Insufficient balance");
        payable(msg.sender).transfer(amount);
    }

    function upVote(uint256 orderID) external onlyProxy {
        require(_ordersMap[orderID].customer == msg.sender, "Permission denied");

        _poolsVoteMap[_ordersMap[orderID].poolID][msg.sender] = Vote({t: VoteType.Up, voted: true});
        _poolsMap[_ordersMap[orderID].poolID].vote.increment();
    }

    function downVote(uint256 orderID) external onlyProxy {
        require(_ordersMap[orderID].customer == msg.sender, "Permission denied");

        _poolsVoteMap[_ordersMap[orderID].poolID][msg.sender] = Vote({
            t: VoteType.Down,
            voted: true
        });
        _poolsMap[_ordersMap[orderID].poolID].vote.decrement();
    }

    function isVoted(uint256 orderID) external view onlyProxy returns (bool) {
        return _poolsVoteMap[_ordersMap[orderID].poolID][msg.sender].voted;
    }

    function recallVote(uint256 orderID) external onlyProxy {
        require(
            _poolsVoteMap[_ordersMap[orderID].poolID][msg.sender].voted == true,
            "You don't have vote for the pool of this order yet"
        );

        if (_poolsVoteMap[_ordersMap[orderID].poolID][msg.sender].t == VoteType.Up) {
            _poolsMap[_ordersMap[orderID].poolID].vote.decrement();
        } else if (_poolsVoteMap[_ordersMap[orderID].poolID][msg.sender].t == VoteType.Down) {
            _poolsMap[_ordersMap[orderID].poolID].vote.increment();
        }

        delete _poolsVoteMap[_ordersMap[orderID].poolID][msg.sender];
    }
}
