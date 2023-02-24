abi: clean node_modules
	solc --abi --base-path '/' --include-path 'node_modules/' --optimize -o build/abi ProxyExchange.sol

node_modules:
	npm ci

remixd:
	./node_modules/.bin/remixd -s $${PWD} -u https://remix.ethereum.org

clean:
	rm -rf build
