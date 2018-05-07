Docker

To build base geth image:

This is the normal geth image, we may want to just modify this image but I'm not sure.

`docker build . -t shyft_geth`

Then, build Dockerfile.shyft:

`docker build -f Dockerfile.shyft . -t shyft_geth_init`


To test this image, run:

`docker run shyft_geth_init`

For an image that runs `startShyftGeth.sh`, without running the init script, (ie a non-genesis node, or a syncing node)

For an image without running the ./initShyftGeth.sh script, in the shyftDocker directory, run:

`docker build . -t shyft_geth_start`

To test, run `docker run shyft_geth_start`