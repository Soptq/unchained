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

// ABCIQueryResponseResultResponse struct for ABCIQueryResponseResultResponse
type ABCIQueryResponseResultResponse struct {
	Log string `json:"log"`
	Height string `json:"height"`
	Proof string `json:"proof"`
	Value string `json:"value"`
	Key string `json:"key"`
	Index string `json:"index"`
	Code string `json:"code"`
}

// NewABCIQueryResponseResultResponse instantiates a new ABCIQueryResponseResultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewABCIQueryResponseResultResponse(log string, height string, proof string, value string, key string, index string, code string) *ABCIQueryResponseResultResponse {
	this := ABCIQueryResponseResultResponse{}
	this.Log = log
	this.Height = height
	this.Proof = proof
	this.Value = value
	this.Key = key
	this.Index = index
	this.Code = code
	return &this
}

// NewABCIQueryResponseResultResponseWithDefaults instantiates a new ABCIQueryResponseResultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewABCIQueryResponseResultResponseWithDefaults() *ABCIQueryResponseResultResponse {
	this := ABCIQueryResponseResultResponse{}
	return &this
}

// GetLog returns the Log field value
func (o *ABCIQueryResponseResultResponse) GetLog() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Log
}

// GetLogOk returns a tuple with the Log field value
// and a boolean to check if the value has been set.
func (o *ABCIQueryResponseResultResponse) GetLogOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Log, true
}

// SetLog sets field value
func (o *ABCIQueryResponseResultResponse) SetLog(v string) {
	o.Log = v
}

// GetHeight returns the Height field value
func (o *ABCIQueryResponseResultResponse) GetHeight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *ABCIQueryResponseResultResponse) GetHeightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *ABCIQueryResponseResultResponse) SetHeight(v string) {
	o.Height = v
}

// GetProof returns the Proof field value
func (o *ABCIQueryResponseResultResponse) GetProof() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Proof
}

// GetProofOk returns a tuple with the Proof field value
// and a boolean to check if the value has been set.
func (o *ABCIQueryResponseResultResponse) GetProofOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proof, true
}

// SetProof sets field value
func (o *ABCIQueryResponseResultResponse) SetProof(v string) {
	o.Proof = v
}

// GetValue returns the Value field value
func (o *ABCIQueryResponseResultResponse) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ABCIQueryResponseResultResponse) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ABCIQueryResponseResultResponse) SetValue(v string) {
	o.Value = v
}

// GetKey returns the Key field value
func (o *ABCIQueryResponseResultResponse) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ABCIQueryResponseResultResponse) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ABCIQueryResponseResultResponse) SetKey(v string) {
	o.Key = v
}

// GetIndex returns the Index field value
func (o *ABCIQueryResponseResultResponse) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ABCIQueryResponseResultResponse) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *ABCIQueryResponseResultResponse) SetIndex(v string) {
	o.Index = v
}

// GetCode returns the Code field value
func (o *ABCIQueryResponseResultResponse) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ABCIQueryResponseResultResponse) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ABCIQueryResponseResultResponse) SetCode(v string) {
	o.Code = v
}

func (o ABCIQueryResponseResultResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["log"] = o.Log
	}
	if true {
		toSerialize["height"] = o.Height
	}
	if true {
		toSerialize["proof"] = o.Proof
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["index"] = o.Index
	}
	if true {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableABCIQueryResponseResultResponse struct {
	value *ABCIQueryResponseResultResponse
	isSet bool
}

func (v NullableABCIQueryResponseResultResponse) Get() *ABCIQueryResponseResultResponse {
	return v.value
}

func (v *NullableABCIQueryResponseResultResponse) Set(val *ABCIQueryResponseResultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableABCIQueryResponseResultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableABCIQueryResponseResultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableABCIQueryResponseResultResponse(val *ABCIQueryResponseResultResponse) *NullableABCIQueryResponseResultResponse {
	return &NullableABCIQueryResponseResultResponse{value: val, isSet: true}
}

func (v NullableABCIQueryResponseResultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableABCIQueryResponseResultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


