package hedera

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/hashgraph/hedera-sdk-go/proto"
	"google.golang.org/grpc"
)

// Default max fees and payments to 1 h-bar
var defaultMaxTransactionFee Hbar = NewHbar(1)
var defaultMaxQueryPayment Hbar = NewHbar(1)

// Client is the Hedera protocol wrapper for the SDK used by all
// transaction and query types.
type Client struct {
	maxTransactionFee Hbar
	maxQueryPayment   Hbar

	operator *operator

	networkNodes   map[AccountID]*node
	networkNodeIds []AccountID
}

type node struct {
	conn    *grpc.ClientConn
	id      AccountID
	address string
}

// TransactionSigner is a closure or function that defines how transactions will be signed
type TransactionSigner func(message []byte) []byte

type operator struct {
	accountID  AccountID
	privateKey *Ed25519PrivateKey
	publicKey  Ed25519PublicKey
	signer     TransactionSigner
}

var mainnetNodes = map[string]AccountID{
	"35.237.200.180:50211": {Account: 3},
	"35.186.191.247:50211": {Account: 4},
	"35.192.2.25:50211":    {Account: 5},
	"35.199.161.108:50211": {Account: 6},
	"35.203.82.240:50211":  {Account: 7},
	"35.236.5.219:50211":   {Account: 8},
	"35.197.192.225:50211": {Account: 9},
	"35.242.233.154:50211": {Account: 10},
	"35.240.118.96:50211":  {Account: 11},
	"35.204.86.32:50211":   {Account: 12},
}

var testnetNodes = map[string]AccountID{
	"0.testnet.hedera.com:50211": {Account: 3},
	"1.testnet.hedera.com:50211": {Account: 4},
	"2.testnet.hedera.com:50211": {Account: 5},
	"3.testnet.hedera.com:50211": {Account: 6},
}

var previewnetNodes = map[string]AccountID{
	"0.previewnet.hedera.com:50211": {Account: 3},
	"1.previewnet.hedera.com:50211": {Account: 4},
	"2.previewnet.hedera.com:50211": {Account: 5},
	"3.previewnet.hedera.com:50211": {Account: 6},
}

// ClientForMainnet returns a preconfigured client for use with the standard
// Hedera mainnet.
// Most users will want to set an operator account with .SetOperator so
// transactions can be automatically given TransactionIDs and signed.
func ClientForMainnet() *Client {
	return NewClient(mainnetNodes)
}

// ClientForTestnet returns a preconfigured client for use with the standard
// Hedera testnet.
// Most users will want to set an operator account with .SetOperator so
// transactions can be automatically given TransactionIDs and signed.
func ClientForTestnet() *Client {
	return NewClient(testnetNodes)
}

// ClientForPreviewnet returns a preconfigured client for use with the standard
// Hedera previewnet.
// Most users will want to set an operator account with .SetOperator so
// transactions can be automatically given TransactionIDs and signed.
func ClientForPreviewnet() *Client {
	return NewClient(previewnetNodes)
}

// NewClient takes in a map of node addresses to their respective IDS (network)
// and returns a Client instance which can be used to
func NewClient(network map[string]AccountID) *Client {
	client := &Client{
		maxQueryPayment:   defaultMaxQueryPayment,
		maxTransactionFee: defaultMaxTransactionFee,
		networkNodes:      map[AccountID]*node{},
		networkNodeIds:    []AccountID{},
	}

	client.ReplaceNodes(network)

	return client
}

type configOperator struct {
	AccountID  string `json:"accountId"`
	PrivateKey string `json:"privateKey"`
}

type clientConfig struct {
	Network  map[string]string `json:"network"`
	Operator *configOperator   `json:"operator"`
}

type clientConfigString struct {
	Network  string `json:"network"`
	Operator *configOperator   `json:"operator"`
}

// ClientFromJSON takes in the byte slice representation of a JSON string or
// document and returns Client based on the configuration.
func ClientFromJSON(jsonBytes []byte) (*Client, error) {
	var clientConfigString clientConfigString
	var clientConfig clientConfig
	//var s interface{}
	var client *Client
	err := json.Unmarshal(jsonBytes,&clientConfigString)
	if err != nil{
		err := json.Unmarshal(jsonBytes,&clientConfig)
		if err != nil{
			return nil, err
		}
		var network map[string]AccountID =  make(map[string]AccountID)

		for id, url := range clientConfig.Network {
			accountID, err := AccountIDFromString(id)
			if err != nil {
				return nil, err
			}

			network[url] = accountID
		}

		client = NewClient(network)

		if clientConfig.Operator == nil {
			return client, nil
		}

		operatorID, err := AccountIDFromString(clientConfig.Operator.AccountID)
		if err != nil {
			return nil, err
		}

		operatorKey, err := Ed25519PrivateKeyFromString(clientConfig.Operator.PrivateKey)
		if err != nil {
			return nil, err
		}

		operator := operator{
			accountID:  operatorID,
			privateKey: &operatorKey,
			publicKey:  operatorKey.PublicKey(),
			signer:     operatorKey.Sign,
		}

		client.operator = &operator

		return client, nil
	}
	fmt.Println("in client")
	fmt.Printf("(%v, %T)\n", clientConfigString, clientConfigString)

	switch clientConfigString.Network{
		case "mainnet":
			client = ClientForMainnet()
		case "testnet":
			client = ClientForTestnet()
		case "previewnet":
			client = ClientForPreviewnet()
	}

	// if the operator is not provided, finish here
	if clientConfig.Operator == nil {
		return client, nil
	}

	operatorID, err := AccountIDFromString(clientConfig.Operator.AccountID)
	if err != nil {
		return nil, err
	}

	operatorKey, err := Ed25519PrivateKeyFromString(clientConfig.Operator.PrivateKey)
	if err != nil {
		return nil, err
	}

	operator := operator{
		accountID:  operatorID,
		privateKey: &operatorKey,
		publicKey:  operatorKey.PublicKey(),
		signer:     operatorKey.Sign,
	}

	client.operator = &operator

	return client, nil
}

// ClientFromFile takes a filename string representing the path to a JSON encoded
// Client file and returns a Client based on the configuration.
func ClientFromFile(filename string) (*Client, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = file.Close()
	}()

	configBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return ClientFromJSON(configBytes)
}

// Close is used to disconnect the Client from the network
func (client *Client) Close() error {
	for _, node := range client.networkNodes {
		if node.conn != nil {
			err := node.conn.Close()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ReplaceNodes replaces all nodes in the Client with a new set of nodes.
// (e.g. for an Address Book update).
func (client *Client) ReplaceNodes(network map[string]AccountID) *Client {
	for address, id := range network {
		client.networkNodeIds = append(client.networkNodeIds, id)
		client.networkNodes[id] = &node{
			id:      id,
			address: address,
		}
	}

	return client
}

// SetOperator sets that account that will, by default, be paying for
// transactions and queries built with the client and the associated key
// with which to automatically sign transactions.
func (client *Client) SetOperator(accountID AccountID, privateKey Ed25519PrivateKey) *Client {
	client.operator = &operator{
		accountID:  accountID,
		privateKey: &privateKey,
		publicKey:  privateKey.PublicKey(),
		signer:     privateKey.Sign,
	}

	return client
}

// SetOperatorWith sets that account that will, by default, be paying for
// transactions and queries built with the client, the account's Ed25519PublicKey
// and a callback that will be invoked when a transaction needs to be signed.
func (client *Client) SetOperatorWith(accountID AccountID, publicKey Ed25519PublicKey, signer TransactionSigner) *Client {
	client.operator = &operator{
		accountID:  accountID,
		privateKey: nil,
		publicKey:  publicKey,
		signer:     signer,
	}

	return client
}

// GetOperatorID returns the ID for the operator
func (client *Client) GetOperatorID() AccountID {
	return client.operator.accountID
}

// GetOperatorKey returns the Key for the operator
func (client *Client) GetOperatorKey() Ed25519PublicKey {
	return client.operator.publicKey
}

// SetMaxTransactionFee sets the maximum fee to be paid for the transactions
// executed by the Client.
// Because transaction fees are always maximums the actual fee assessed for
// a given transaction may be less than this value, but never greater.
func (client *Client) SetMaxTransactionFee(fee Hbar) *Client {
	client.maxTransactionFee = fee
	return client
}

// SetMaxQueryPayment sets the default maximum payment allowable for queries.
func (client *Client) SetMaxQueryPayment(payment Hbar) *Client {
	client.maxQueryPayment = payment
	return client
}

// Ping sends an AccountBalanceQuery to the specified node returning nil if no
// problems occur. Otherwise, an error representing the status of the node will
// be returned.
func (client *Client) Ping(nodeID AccountID) error {
	node := client.networkNodes[nodeID]
	if node == nil {
		return fmt.Errorf("node with ID %s not registered on this client", nodeID)
	}

	pingQuery := NewAccountBalanceQuery().
		SetAccountID(nodeID)

	pb := pingQuery.QueryBuilder.pb

	resp := new(proto.Response)

	err := node.invoke(methodName(pb), pb, resp)

	if err != nil {
		return newErrPingStatus(err)
	}

	respHeader := mapResponseHeader(resp)

	if respHeader.NodeTransactionPrecheckCode == proto.ResponseCodeEnum_BUSY {
		return newErrPingStatus(fmt.Errorf("%s", Status(respHeader.NodeTransactionPrecheckCode).String()))
	}

	if isResponseUnknown(resp) {
		return newErrPingStatus(fmt.Errorf("unknown"))
	}

	return nil
}

func (client *Client) randomNode() *node {
	nodeIndex := rand.Intn(len(client.networkNodeIds))
	nodeID := client.networkNodeIds[nodeIndex]

	return client.networkNodes[nodeID]
}

func (client *Client) node(id AccountID) *node {
	return client.networkNodes[id]
}

func (node *node) invoke(method string, in interface{}, out interface{}) error {
	if node.conn == nil {
		conn, err := grpc.Dial(node.address, grpc.WithInsecure())
		if err != nil {
			return newErrHederaNetwork(err)
		}

		node.conn = conn
	}

	err := node.conn.Invoke(context.TODO(), method, in, out)

	if err != nil {
		return newErrHederaNetwork(err)
	}

	return nil
}
