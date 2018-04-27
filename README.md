# IPFS CoreAPI Sample App

It's sample apps using IPFS CoreAPI.
If you want to implement embedded IPFS apps, It can help.

## ToC

- Install
- Usage
- License


## Install

You can download on [release page](https://github.com/camelmasa/ipfs-coreapi-sample-app/releases).
Or, You can build it from source.

*Notice*: If you want to build from source, you need to install [gx](https://github.com/whyrusleeping/gx).

```sh
git clone git@github.com:camelmasa/ipfs-coreapi-sample-app.git
cd ipfs-coreapi-sample-app
gx i
go build
```


## Usage

*Notice*: If you didn't create IPFS's config to your home directory like `/Home/camelmasa/.ipfs`, you need to create that first.

You can call API.
```sh
./ipfs-coreapi-sample-app

curl -X POST http://localhost:4002/api/articles -d "content=Hello Decentralized World"
# => Added! You can check on gateway. https://ipfs.io/ipfs/QmRnKYfpw4VKyMmrWcGzZEd17oU6NgiLTvHQEyGxcjc4EZ
```

And then, you can access the gateway URL.
Let's try to change the content param !


## License

MIT License
