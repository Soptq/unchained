/*
Tendermint RPC

Tendermint supports the following RPC protocols:  * URI over HTTP * JSONRPC over HTTP * JSONRPC over websockets  ## Configuration  RPC can be configured by tuning parameters under `[rpc]` table in the `$TMHOME/config/config.toml` file or by using the `--rpc.X` command-line flags.  Default rpc listen address is `tcp://0.0.0.0:26657`. To set another address, set the `laddr` config parameter to desired value. CORS (Cross-Origin Resource Sharing) can be enabled by setting `cors_allowed_origins`, `cors_allowed_methods`, `cors_allowed_headers` config parameters.  ## Arguments  Arguments which expect strings or byte arrays may be passed as quoted strings, like `\"abc\"` or as `0x`-prefixed strings, like `0x616263`.  ## URI/HTTP  A REST like interface.      curl localhost:26657/block?height=5  ## JSONRPC/HTTP  JSONRPC requests can be POST'd to the root RPC endpoint via HTTP.      curl --header \"Content-Type: application/json\" --request POST --data '{\"method\": \"block\", \"params\": [\"5\"], \"id\": 1}' localhost:26657  ## JSONRPC/websockets  JSONRPC requests can be also made via websocket. The websocket endpoint is at `/websocket`, e.g. `localhost:26657/websocket`. Asynchronous RPC functions like event `subscribe` and `unsubscribe` are only available via websockets.  Example using https://github.com/hashrocket/ws:      ws ws://localhost:26657/websocket     > { \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewBlock'\"], \"id\": 1 } 

API version: Master
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// BroadcastTxCommitResponseResult struct for BroadcastTxCommitResponseResult
type BroadcastTxCommitResponseResult struct {
	Height string `json:"height"`
	Hash string `json:"hash"`
	DeliverTx BroadcastTxCommitResponseResultDeliverTx `json:"deliver_tx"`
	CheckTx BroadcastTxCommitResponseResultDeliverTx `json:"check_tx"`
}

// NewBroadcastTxCommitResponseResult instantiates a new BroadcastTxCommitResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastTxCommitResponseResult(height string, hash string, deliverTx BroadcastTxCommitResponseResultDeliverTx, checkTx BroadcastTxCommitResponseResultDeliverTx) *BroadcastTxCommitResponseResult {
	this := BroadcastTxCommitResponseResult{}
	this.Height = height
	this.Hash = hash
	this.DeliverTx = deliverTx
	this.CheckTx = checkTx
	return &this
}

// NewBroadcastTxCommitResponseResultWithDefaults instantiates a new BroadcastTxCommitResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastTxCommitResponseResultWithDefaults() *BroadcastTxCommitResponseResult {
	this := BroadcastTxCommitResponseResult{}
	return &this
}

// GetHeight returns the Height field value
func (o *BroadcastTxCommitResponseResult) GetHeight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *BroadcastTxCommitResponseResult) GetHeightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *BroadcastTxCommitResponseResult) SetHeight(v string) {
	o.Height = v
}

// GetHash returns the Hash field value
func (o *BroadcastTxCommitResponseResult) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *BroadcastTxCommitResponseResult) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *BroadcastTxCommitResponseResult) SetHash(v string) {
	o.Hash = v
}

// GetDeliverTx returns the DeliverTx field value
func (o *BroadcastTxCommitResponseResult) GetDeliverTx() BroadcastTxCommitResponseResultDeliverTx {
	if o == nil {
		var ret BroadcastTxCommitResponseResultDeliverTx
		return ret
	}

	return o.DeliverTx
}

// GetDeliverTxOk returns a tuple with the DeliverTx field value
// and a boolean to check if the value has been set.
func (o *BroadcastTxCommitResponseResult) GetDeliverTxOk() (*BroadcastTxCommitResponseResultDeliverTx, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliverTx, true
}

// SetDeliverTx sets field value
func (o *BroadcastTxCommitResponseResult) SetDeliverTx(v BroadcastTxCommitResponseResultDeliverTx) {
	o.DeliverTx = v
}

// GetCheckTx returns the CheckTx field value
func (o *BroadcastTxCommitResponseResult) GetCheckTx() BroadcastTxCommitResponseResultDeliverTx {
	if o == nil {
		var ret BroadcastTxCommitResponseResultDeliverTx
		return ret
	}

	return o.CheckTx
}

// GetCheckTxOk returns a tuple with the CheckTx field value
// and a boolean to check if the value has been set.
func (o *BroadcastTxCommitResponseResult) GetCheckTxOk() (*BroadcastTxCommitResponseResultDeliverTx, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckTx, true
}

// SetCheckTx sets field value
func (o *BroadcastTxCommitResponseResult) SetCheckTx(v BroadcastTxCommitResponseResultDeliverTx) {
	o.CheckTx = v
}

func (o BroadcastTxCommitResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["height"] = o.Height
	}
	if true {
		toSerialize["hash"] = o.Hash
	}
	if true {
		toSerialize["deliver_tx"] = o.DeliverTx
	}
	if true {
		toSerialize["check_tx"] = o.CheckTx
	}
	return json.Marshal(toSerialize)
}

type NullableBroadcastTxCommitResponseResult struct {
	value *BroadcastTxCommitResponseResult
	isSet bool
}

func (v NullableBroadcastTxCommitResponseResult) Get() *BroadcastTxCommitResponseResult {
	return v.value
}

func (v *NullableBroadcastTxCommitResponseResult) Set(val *BroadcastTxCommitResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastTxCommitResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastTxCommitResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastTxCommitResponseResult(val *BroadcastTxCommitResponseResult) *NullableBroadcastTxCommitResponseResult {
	return &NullableBroadcastTxCommitResponseResult{value: val, isSet: true}
}

func (v NullableBroadcastTxCommitResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastTxCommitResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


