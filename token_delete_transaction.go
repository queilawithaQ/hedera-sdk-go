package hedera

import (
	"time"

	"github.com/hashgraph/hedera-sdk-go/proto"
)

// Deletes an already created Token.
// If no value is given for a field, that field is left unchanged. For an immutable tokens
// (that is, a token created without an adminKey), only the expiry may be deleted. Setting any
// other field in that case will cause the transaction status to resolve to TOKEN_IS_IMMUTABlE.
type TokenDeleteTransaction struct {
	Transaction
	pb *proto.TokenDeleteTransactionBody
}

func NewTokenDeleteTransaction() *TokenDeleteTransaction {
	pb := &proto.TokenDeleteTransactionBody{}

	transaction := TokenDeleteTransaction{
		pb:          pb,
		Transaction: newTransaction(),
	}

	return &transaction
}

// The Token to be deleted
func (transaction *TokenDeleteTransaction) SetTokenID(tokenID TokenID) *TokenDeleteTransaction {
	transaction.pb.Token = tokenID.toProtobuf()
	return transaction
}

//
// The following methods must be copy-pasted/overriden at the bottom of **every** _transaction.go file
// We override the embedded fluent setter methods to return the outer type
//

func tokenDeleteTransaction_getMethod(request request, channel *channel) method {
	return method{
		transaction: channel.getToken().DeleteToken,
	}
}

func (transaction *TokenDeleteTransaction) IsFrozen() bool {
	return transaction.isFrozen()
}

// Sign uses the provided privateKey to sign the transaction.
func (transaction *TokenDeleteTransaction) Sign(
	privateKey PrivateKey,
) *TokenDeleteTransaction {
	return transaction.SignWith(privateKey.PublicKey(), privateKey.Sign)
}

func (transaction *TokenDeleteTransaction) SignWithOperator(
	client *Client,
) (*TokenDeleteTransaction, error) {
	// If the transaction is not signed by the operator, we need
	// to sign the transaction with the operator

	if client.operator == nil {
		return nil, errClientOperatorSigning
	}

	if !transaction.IsFrozen() {
		transaction.FreezeWith(client)
	}

	return transaction.SignWith(client.operator.publicKey, client.operator.signer), nil
}

// SignWith executes the TransactionSigner and adds the resulting signature data to the Transaction's signature map
// with the publicKey as the map key.
func (transaction *TokenDeleteTransaction) SignWith(
	publicKey PublicKey,
	signer TransactionSigner,
) *TokenDeleteTransaction {
	if !transaction.IsFrozen() {
		transaction.Freeze()
	}

	if transaction.keyAlreadySigned(publicKey) {
		return transaction
	}

	for index := 0; index < len(transaction.transactions); index++ {
		signature := signer(transaction.transactions[index].GetBodyBytes())

		transaction.signatures[index].SigPair = append(
			transaction.signatures[index].SigPair,
			publicKey.toSignaturePairProtobuf(signature),
		)
	}

	return transaction
}

// Execute executes the Transaction with the provided client
func (transaction *TokenDeleteTransaction) Execute(
	client *Client,
) (TransactionResponse, error) {
	if !transaction.IsFrozen() {
		transaction.FreezeWith(client)
	}

	transactionID := transaction.id

	if !client.GetOperatorID().isZero() && client.GetOperatorID().equals(transactionID.AccountID) {
		transaction.SignWith(
			client.GetOperatorKey(),
			client.operator.signer,
		)
	}

	resp, err := execute(
		client,
		request{
			transaction: &transaction.Transaction,
		},
		transaction_shouldRetry,
		transaction_makeRequest,
		transaction_advanceRequest,
		transaction_getNodeId,
		tokenDeleteTransaction_getMethod,
		transaction_mapResponseStatus,
		transaction_mapResponse,
	)

	if err != nil {
		return TransactionResponse{}, err
	}

	return TransactionResponse{
		TransactionID: transaction.id,
		NodeID:        resp.transaction.NodeID,
	}, nil
}

func (transaction *TokenDeleteTransaction) onFreeze(
	pbBody *proto.TransactionBody,
) bool {
	pbBody.Data = &proto.TransactionBody_TokenDeletion{
		TokenDeletion: transaction.pb,
	}

	return true
}

func (transaction *TokenDeleteTransaction) Freeze() (*TokenDeleteTransaction, error) {
	return transaction.FreezeWith(nil)
}

func (transaction *TokenDeleteTransaction) FreezeWith(client *Client) (*TokenDeleteTransaction, error) {
	transaction.initFee(client)
	if err := transaction.initTransactionID(client); err != nil {
		return transaction, err
	}

	if !transaction.onFreeze(transaction.pbBody) {
		return transaction, nil
	}

	return transaction, transaction_freezeWith(&transaction.Transaction, client)
}

func (transaction *TokenDeleteTransaction) GetMaxTransactionFee() Hbar {
	return transaction.Transaction.GetMaxTransactionFee()
}

// SetMaxTransactionFee sets the max transaction fee for this TokenDeleteTransaction.
func (transaction *TokenDeleteTransaction) SetMaxTransactionFee(fee Hbar) *TokenDeleteTransaction {
	transaction.Transaction.SetMaxTransactionFee(fee)
	return transaction
}

func (transaction *TokenDeleteTransaction) GetTransactionMemo() string {
	return transaction.Transaction.GetTransactionMemo()
}

// SetTransactionMemo sets the memo for this TokenDeleteTransaction.
func (transaction *TokenDeleteTransaction) SetTransactionMemo(memo string) *TokenDeleteTransaction {
	transaction.Transaction.SetTransactionMemo(memo)
	return transaction
}

func (transaction *TokenDeleteTransaction) GetTransactionValidDuration() time.Duration {
	return transaction.Transaction.GetTransactionValidDuration()
}

// SetTransactionValidDuration sets the valid duration for this TokenDeleteTransaction.
func (transaction *TokenDeleteTransaction) SetTransactionValidDuration(duration time.Duration) *TokenDeleteTransaction {
	transaction.Transaction.SetTransactionValidDuration(duration)
	return transaction
}

func (transaction *TokenDeleteTransaction) GetTransactionID() TransactionID {
	return transaction.Transaction.GetTransactionID()
}

// SetTransactionID sets the TransactionID for this TokenDeleteTransaction.
func (transaction *TokenDeleteTransaction) SetTransactionID(transactionID TransactionID) *TokenDeleteTransaction {
	transaction.id = transactionID
	transaction.Transaction.SetTransactionID(transactionID)
	return transaction
}

func (transaction *TokenDeleteTransaction) GetNodeAccountIDs() []AccountID {
	return transaction.Transaction.GetNodeAccountIDs()
}

// SetNodeTokenID sets the node TokenID for this TokenDeleteTransaction.
func (transaction *TokenDeleteTransaction) SetNodeAccountIDs(nodeID []AccountID) *TokenDeleteTransaction {
	transaction.Transaction.SetNodeAccountIDs(nodeID)
	return transaction
}