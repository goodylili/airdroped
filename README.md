# Airdroped

Airdroped is a free to use, cheap tool for airdropping on any on Ethereum-compatible blockchains. Pay gas once, airdrop
to as many accounts as possible with one command.


## Features

- [x] Batch Transfers in one transactions.
- [x] CSV File based because it could all be so simple
- [x] ERC-20 Support: Fully compatible with ERC-20 tokens.


## Prerequisites
- Go 1.20+
- A CSV File that works
- Enough tokens to go round for all users in your CSV File
- Enough gas for just one transaction


## How to Airdrop?
First, execute this command to clone the repository and move into it:

```shell
git clone https://github.com/your-repo/airdroped.git
cd airdroped
```

Execute this command to install necessary dependencies:
```shell
go mod tidy
```

Now, create a `.env` file in the directory and fill these details:

```shell
RPC_URL=
PRIVATE_KEY=
PUBLIC_KEY=
TOKEN_ADDRESS=
CONTRACT_DEPLOYMENT_ADDRESS=
CHAIN_ID=
CSV_FILE_PATH=
ADDRESS_COLUMN=
AMOUNTS_COLUMN=
```

That's all the details the program needs for the airdrop and they're all private. There's a `.env` specification  in the `.gitignore` file, so it's safe if you want to reuse the code.


[//]: # (TODO: Add tutorial Link)


> ⚠️ You'll need to deploy the contract on the network before proceeding. You can open an issue and I'll do that for you, or follow this tutorial I wrote to deploy the con tract.


I deployed the contract on Base network at `0x145b7982a83cb864be2ab4f1a3dfd1f920ff2954`, you can use it as the `CONTRACT_DEPLOYMENT_ADDRESS` if you're working with Base.

Finally, execute the script with this command:

```shell
go run cmd/server/main.go
```

Here's an output similar to what you should expect after executing the command.

<img title="airdroped output" alt="airdroped output" src="/internal/services/img.png">

Your sequel action should be viewing the transaction hash on a block explorer.


## Contributing
This repository is open to contributions. Feel free to open an issue/PR for any changes.