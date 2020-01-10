#!/usr/bin/env bash
set -e
PASSWORD=1234567890

REPLACE_JACK_ADDR="link16xyempempp92x9hyzz9wrgf94r6j9h5f06pxxv"
REPLACE_OPERATOR_ADDR="linkoperatormpp92x9hyzz9wrgf94r6j9h5f06pxxv"
REPLACE_ALLOCATOR_ADDR="linkallocatorpp92x9hyzz9wrgf94r6j9h5f06pxxv"
REPLACE_ISSUER_ADDR="linkissuetormpp92x9hyzz9wrgf94r6j9h5f06pxxv"
REPLACE_RETURNER_ADDR="linkreturnormpp92x9hyzz9wrgf94r6j9h5f06pxxv"
REPLACE_TX_HASH="BCBE20E8D46758B96AE5883B792858296AC06E51435490FBDCAE25A72B3CC76B"
REPLACE_TOKEN_SYMBOL="conydkv"
REPLACE_COLLECTION_SYMBOL="conydk2"
REPLACE_NFT_SYMBOL="conydk3"
REPLACE_NFT_COLLECTION_SYMBOL="conydk4"
REPLACE_MSG_EXAMPLES="\"MsgExamples\""
CHAIN_ID="lcd"
HOME="/tmp/contract_tests/.linkcli"
SWAGGER='/tmp/contract_tests/swagger.yaml'
