# Introduction

[Proxy Exchange](https://proxyex.xyz) is a platform for proxy service exchange based on blockchain & smart contract. 

It is using smart contract to do transactions between proxy pool provider and user, using the user's signature as authentication info.

# The authentication rule

```
Username: ORDER ID
Password: sign(ORDER ID)
```

# The directory introduction

contracts: Smart Contract source code for exchange transaction

exchange: The web app to interact with exchange Smart Contract

proxyd: A proxy server, which provides SOCKS5 and HTTP protocol proxy service, accepts the above authentication rule. It is just a demonstrate, in theory any proxy server would work with Proxy Exchange if only they were accepted the authentication rule.

# How to become a proxy pool provider ?

Go to exchange website, publish pool. This will create a new POOL ID, then you could download proxyd and run it with `--pool=POOL ID` flag.

# How to use published pool as a normal user ?

Go to exchange website, checkout from a pool, an order wil be generated after making a purchase. In the order page, there is a `Generate Authentication` button to use fetch authentication info (username & password). The authentication could be used for all proxies of the pool.

Example: 

A pool includes proxies:

```
socks5://example.com:10800
http://example.com:10801
```

Make an order with it, the order id is `123`, generated authentication would be like this:

```
Username: 123
Password: 0x7dfa8192972143c83f055c5549ba627f7339...ae7e7a8bd59721222a4bb2c069834178c1b
```

To use the HTTP protocol proxy in curl

```
curl --proxy http://123:0x7dfa8192972143c83f055c5549ba627f7339...ae7e7a8bd59721222a4bb2c069834178c1b@example.com:10801 https://ifconfig.me
```

Check the responsed IP address.
