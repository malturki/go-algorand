// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package generated

import (
	"encoding/json"
)

const (
	Api_keyScopes = "api_key.Scopes"
)

// Defines values for AccountSigType.
const (
	Lsig AccountSigType = "lsig"
	Msig AccountSigType = "msig"
	Sig  AccountSigType = "sig"
)

// Account information at a given round.
//
// Definition:
// data/basics/userBalance.go : AccountData
type Account struct {
	// the account public key
	Address string `json:"address"`

	// \[algo\] total number of MicroAlgos in the account
	Amount uint64 `json:"amount"`

	// specifies the amount of MicroAlgos in the account, without the pending rewards.
	AmountWithoutPendingRewards uint64 `json:"amount-without-pending-rewards"`

	// \[appl\] applications local data stored in this account.
	//
	// Note the raw object uses `map[int] -> AppLocalState` for this type.
	AppsLocalState *[]ApplicationLocalState `json:"apps-local-state,omitempty"`

	// \[teap\] the sum of all extra application program pages for this account.
	AppsTotalExtraPages *uint64 `json:"apps-total-extra-pages,omitempty"`

	// Specifies maximums on the number of each type that may be stored.
	AppsTotalSchema *ApplicationStateSchema `json:"apps-total-schema,omitempty"`

	// \[asset\] assets held by this account.
	//
	// Note the raw object uses `map[int] -> AssetHolding` for this type.
	Assets *[]AssetHolding `json:"assets,omitempty"`

	// \[spend\] the address against which signing should be checked. If empty, the address of the current account is used. This field can be updated in any transaction by setting the RekeyTo field.
	AuthAddr *string `json:"auth-addr,omitempty"`

	// \[appp\] parameters of applications created by this account including app global data.
	//
	// Note: the raw account uses `map[int] -> AppParams` for this type.
	CreatedApps *[]Application `json:"created-apps,omitempty"`

	// \[apar\] parameters of assets created by this account.
	//
	// Note: the raw account uses `map[int] -> Asset` for this type.
	CreatedAssets *[]Asset `json:"created-assets,omitempty"`

	// MicroAlgo balance required by the account.
	//
	// The requirement grows based on asset and application usage.
	MinBalance uint64 `json:"min-balance"`

	// AccountParticipation describes the parameters used by this account in consensus protocol.
	Participation *AccountParticipation `json:"participation,omitempty"`

	// amount of MicroAlgos of pending rewards in this account.
	PendingRewards uint64 `json:"pending-rewards"`

	// \[ebase\] used as part of the rewards computation. Only applicable to accounts which are participating.
	RewardBase *uint64 `json:"reward-base,omitempty"`

	// \[ern\] total rewards of MicroAlgos the account has received, including pending rewards.
	Rewards uint64 `json:"rewards"`

	// The round for which this information is relevant.
	Round uint64 `json:"round"`

	// Indicates what type of signature is used by this account, must be one of:
	// * sig
	// * msig
	// * lsig
	SigType *AccountSigType `json:"sig-type,omitempty"`

	// \[onl\] delegation status of the account's MicroAlgos
	// * Offline - indicates that the associated account is delegated.
	// *  Online  - indicates that the associated account used as part of the delegation pool.
	// *   NotParticipating - indicates that the associated account is neither a delegator nor a delegate.
	Status string `json:"status"`

	// The count of all applications that have been opted in, equivalent to the count of application local data (AppLocalState objects) stored in this account.
	TotalAppsOptedIn uint64 `json:"total-apps-opted-in"`

	// The count of all assets that have been opted in, equivalent to the count of AssetHolding objects held by this account.
	TotalAssetsOptedIn uint64 `json:"total-assets-opted-in"`

	// The count of all apps (AppParams objects) created by this account.
	TotalCreatedApps uint64 `json:"total-created-apps"`

	// The count of all assets (AssetParams objects) created by this account.
	TotalCreatedAssets uint64 `json:"total-created-assets"`
}

// Indicates what type of signature is used by this account, must be one of:
// * sig
// * msig
// * lsig
type AccountSigType string

// AccountParticipation describes the parameters used by this account in consensus protocol.
type AccountParticipation struct {
	// \[sel\] Selection public key (if any) currently registered for this round.
	SelectionParticipationKey []byte `json:"selection-participation-key"`

	// \[stprf\] Root of the state proof key (if any)
	StateProofKey *[]byte `json:"state-proof-key,omitempty"`

	// \[voteFst\] First round for which this participation is valid.
	VoteFirstValid uint64 `json:"vote-first-valid"`

	// \[voteKD\] Number of subkeys in each batch of participation keys.
	VoteKeyDilution uint64 `json:"vote-key-dilution"`

	// \[voteLst\] Last round for which this participation is valid.
	VoteLastValid uint64 `json:"vote-last-valid"`

	// \[vote\] root participation public key (if any) currently registered for this round.
	VoteParticipationKey []byte `json:"vote-participation-key"`
}

// Application state delta.
type AccountStateDelta struct {
	Address string `json:"address"`

	// Application state delta.
	Delta StateDelta `json:"delta"`
}

// Application index and its parameters
type Application struct {
	// \[appidx\] application index.
	Id uint64 `json:"id"`

	// Stores the global information associated with an application.
	Params ApplicationParams `json:"params"`
}

// Stores local state associated with an application.
type ApplicationLocalState struct {
	// The application which this local state is for.
	Id uint64 `json:"id"`

	// Represents a key-value store for use in an application.
	KeyValue *TealKeyValueStore `json:"key-value,omitempty"`

	// Specifies maximums on the number of each type that may be stored.
	Schema ApplicationStateSchema `json:"schema"`
}

// Stores the global information associated with an application.
type ApplicationParams struct {
	// \[approv\] approval program.
	ApprovalProgram []byte `json:"approval-program"`

	// \[clearp\] approval program.
	ClearStateProgram []byte `json:"clear-state-program"`

	// The address that created this application. This is the address where the parameters and global state for this application can be found.
	Creator string `json:"creator"`

	// \[epp\] the amount of extra program pages available to this app.
	ExtraProgramPages *uint64 `json:"extra-program-pages,omitempty"`

	// Represents a key-value store for use in an application.
	GlobalState *TealKeyValueStore `json:"global-state,omitempty"`

	// Specifies maximums on the number of each type that may be stored.
	GlobalStateSchema *ApplicationStateSchema `json:"global-state-schema,omitempty"`

	// Specifies maximums on the number of each type that may be stored.
	LocalStateSchema *ApplicationStateSchema `json:"local-state-schema,omitempty"`
}

// Specifies maximums on the number of each type that may be stored.
type ApplicationStateSchema struct {
	// \[nbs\] num of byte slices.
	NumByteSlice uint64 `json:"num-byte-slice"`

	// \[nui\] num of uints.
	NumUint uint64 `json:"num-uint"`
}

// Specifies both the unique identifier and the parameters for an asset
type Asset struct {
	// unique asset identifier
	Index uint64 `json:"index"`

	// AssetParams specifies the parameters for an asset.
	//
	// \[apar\] when part of an AssetConfig transaction.
	//
	// Definition:
	// data/transactions/asset.go : AssetParams
	Params AssetParams `json:"params"`
}

// Describes an asset held by an account.
//
// Definition:
// data/basics/userBalance.go : AssetHolding
type AssetHolding struct {
	// \[a\] number of units held.
	Amount uint64 `json:"amount"`

	// Asset ID of the holding.
	AssetID uint64 `json:"asset-id"`

	// \[f\] whether or not the holding is frozen.
	IsFrozen bool `json:"is-frozen"`
}

// AssetParams specifies the parameters for an asset.
//
// \[apar\] when part of an AssetConfig transaction.
//
// Definition:
// data/transactions/asset.go : AssetParams
type AssetParams struct {
	// \[c\] Address of account used to clawback holdings of this asset.  If empty, clawback is not permitted.
	Clawback *string `json:"clawback,omitempty"`

	// The address that created this asset. This is the address where the parameters for this asset can be found, and also the address where unwanted asset units can be sent in the worst case.
	Creator string `json:"creator"`

	// \[dc\] The number of digits to use after the decimal point when displaying this asset. If 0, the asset is not divisible. If 1, the base unit of the asset is in tenths. If 2, the base unit of the asset is in hundredths, and so on. This value must be between 0 and 19 (inclusive).
	Decimals uint64 `json:"decimals"`

	// \[df\] Whether holdings of this asset are frozen by default.
	DefaultFrozen *bool `json:"default-frozen,omitempty"`

	// \[f\] Address of account used to freeze holdings of this asset.  If empty, freezing is not permitted.
	Freeze *string `json:"freeze,omitempty"`

	// \[m\] Address of account used to manage the keys of this asset and to destroy it.
	Manager *string `json:"manager,omitempty"`

	// \[am\] A commitment to some unspecified asset metadata. The format of this metadata is up to the application.
	MetadataHash *[]byte `json:"metadata-hash,omitempty"`

	// \[an\] Name of this asset, as supplied by the creator. Included only when the asset name is composed of printable utf-8 characters.
	Name *string `json:"name,omitempty"`

	// Base64 encoded name of this asset, as supplied by the creator.
	NameB64 *[]byte `json:"name-b64,omitempty"`

	// \[r\] Address of account holding reserve (non-minted) units of this asset.
	Reserve *string `json:"reserve,omitempty"`

	// \[t\] The total number of units of this asset.
	Total uint64 `json:"total"`

	// \[un\] Name of a unit of this asset, as supplied by the creator. Included only when the name of a unit of this asset is composed of printable utf-8 characters.
	UnitName *string `json:"unit-name,omitempty"`

	// Base64 encoded name of a unit of this asset, as supplied by the creator.
	UnitNameB64 *[]byte `json:"unit-name-b64,omitempty"`

	// \[au\] URL where more information about the asset can be retrieved. Included only when the URL is composed of printable utf-8 characters.
	Url *string `json:"url,omitempty"`

	// Base64 encoded URL where more information about the asset can be retrieved.
	UrlB64 *[]byte `json:"url-b64,omitempty"`
}

// Request data type for dryrun endpoint. Given the Transactions and simulated ledger state upload, run TEAL scripts and return debugging information.
type DryrunRequest struct {
	Accounts []Account     `json:"accounts"`
	Apps     []Application `json:"apps"`

	// LatestTimestamp is available to some TEAL scripts. Defaults to the latest confirmed timestamp this algod is attached to.
	LatestTimestamp uint64 `json:"latest-timestamp"`

	// ProtocolVersion specifies a specific version string to operate under, otherwise whatever the current protocol of the network this algod is running in.
	ProtocolVersion string `json:"protocol-version"`

	// Round is available to some TEAL scripts. Defaults to the current round on the network this algod is attached to.
	Round   uint64            `json:"round"`
	Sources []DryrunSource    `json:"sources"`
	Txns    []json.RawMessage `json:"txns"`
}

// DryrunSource is TEAL source text that gets uploaded, compiled, and inserted into transactions or application state.
type DryrunSource struct {
	AppIndex uint64 `json:"app-index"`

	// FieldName is what kind of sources this is. If lsig then it goes into the transactions[this.TxnIndex].LogicSig. If approv or clearp it goes into the Approval Program or Clear State Program of application[this.AppIndex].
	FieldName string `json:"field-name"`
	Source    string `json:"source"`
	TxnIndex  uint64 `json:"txn-index"`
}

// Stores the TEAL eval step data
type DryrunState struct {
	// Evaluation error if any
	Error *string `json:"error,omitempty"`

	// Line number
	Line uint64 `json:"line"`

	// Program counter
	Pc      uint64       `json:"pc"`
	Scratch *[]TealValue `json:"scratch,omitempty"`
	Stack   []TealValue  `json:"stack"`
}

// DryrunTxnResult contains any LogicSig or ApplicationCall program debug information and state updates from a dryrun.
type DryrunTxnResult struct {
	AppCallMessages *[]string      `json:"app-call-messages,omitempty"`
	AppCallTrace    *[]DryrunState `json:"app-call-trace,omitempty"`

	// Budget added during execution of app call transaction.
	BudgetAdded *uint64 `json:"budget-added,omitempty"`

	// Budget consumed during execution of app call transaction.
	BudgetConsumed *uint64 `json:"budget-consumed,omitempty"`

	// Net cost of app execution. Field is DEPRECATED and is subject for removal. Instead, use `budget-added` and `budget-consumed.
	Cost *uint64 `json:"cost,omitempty"`

	// Disassembled program line by line.
	Disassembly []string `json:"disassembly"`

	// Application state delta.
	GlobalDelta *StateDelta          `json:"global-delta,omitempty"`
	LocalDeltas *[]AccountStateDelta `json:"local-deltas,omitempty"`

	// Disassembled lsig program line by line.
	LogicSigDisassembly *[]string      `json:"logic-sig-disassembly,omitempty"`
	LogicSigMessages    *[]string      `json:"logic-sig-messages,omitempty"`
	LogicSigTrace       *[]DryrunState `json:"logic-sig-trace,omitempty"`
	Logs                *[][]byte      `json:"logs,omitempty"`
}

// An error response with optional data field.
type ErrorResponse struct {
	Data    *map[string]interface{} `json:"data,omitempty"`
	Message string                  `json:"message"`
}

// Represents a TEAL value delta.
type EvalDelta struct {
	// \[at\] delta action.
	Action uint64 `json:"action"`

	// \[bs\] bytes value.
	Bytes *string `json:"bytes,omitempty"`

	// \[ui\] uint value.
	Uint *uint64 `json:"uint,omitempty"`
}

// Key-value pairs for StateDelta.
type EvalDeltaKeyValue struct {
	Key string `json:"key"`

	// Represents a TEAL value delta.
	Value EvalDelta `json:"value"`
}

// Proof of membership and position of a light block header.
type LightBlockHeaderProof struct {
	// The index of the light block header in the vector commitment tree
	Index uint64 `json:"index"`

	// The encoded proof.
	Proof []byte `json:"proof"`

	// Represents the depth of the tree that is being proven, i.e. the number of edges from a leaf to the root.
	Treedepth uint64 `json:"treedepth"`
}

// Details about a pending transaction. If the transaction was recently confirmed, includes confirmation details like the round and reward details.
type PendingTransactionResponse struct {
	// The application index if the transaction was found and it created an application.
	ApplicationIndex *uint64 `json:"application-index,omitempty"`

	// The number of the asset's unit that were transferred to the close-to address.
	AssetClosingAmount *uint64 `json:"asset-closing-amount,omitempty"`

	// The asset index if the transaction was found and it created an asset.
	AssetIndex *uint64 `json:"asset-index,omitempty"`

	// Rewards in microalgos applied to the close remainder to account.
	CloseRewards *uint64 `json:"close-rewards,omitempty"`

	// Closing amount for the transaction.
	ClosingAmount *uint64 `json:"closing-amount,omitempty"`

	// The round where this transaction was confirmed, if present.
	ConfirmedRound *uint64 `json:"confirmed-round,omitempty"`

	// Application state delta.
	GlobalStateDelta *StateDelta `json:"global-state-delta,omitempty"`

	// Inner transactions produced by application execution.
	InnerTxns *[]PendingTransactionResponse `json:"inner-txns,omitempty"`

	// \[ld\] Local state key/value changes for the application being executed by this transaction.
	LocalStateDelta *[]AccountStateDelta `json:"local-state-delta,omitempty"`

	// \[lg\] Logs for the application being executed by this transaction.
	Logs *[][]byte `json:"logs,omitempty"`

	// Indicates that the transaction was kicked out of this node's transaction pool (and specifies why that happened).  An empty string indicates the transaction wasn't kicked out of this node's txpool due to an error.
	PoolError string `json:"pool-error"`

	// Rewards in microalgos applied to the receiver account.
	ReceiverRewards *uint64 `json:"receiver-rewards,omitempty"`

	// Rewards in microalgos applied to the sender account.
	SenderRewards *uint64 `json:"sender-rewards,omitempty"`

	// The raw signed transaction.
	Txn map[string]interface{} `json:"txn"`
}

// Application state delta.
type StateDelta = []EvalDeltaKeyValue

// Represents a state proof and its corresponding message
type StateProof struct {
	// Represents the message that the state proofs are attesting to.
	Message StateProofMessage `json:"Message"`

	// The encoded StateProof for the message.
	StateProof []byte `json:"StateProof"`
}

// Represents the message that the state proofs are attesting to.
type StateProofMessage struct {
	// The vector commitment root on all light block headers within a state proof interval.
	BlockHeadersCommitment []byte `json:"BlockHeadersCommitment"`

	// The first round the message attests to.
	FirstAttestedRound uint64 `json:"FirstAttestedRound"`

	// The last round the message attests to.
	LastAttestedRound uint64 `json:"LastAttestedRound"`

	// An integer value representing the natural log of the proven weight with 16 bits of precision. This value would be used to verify the next state proof.
	LnProvenWeight uint64 `json:"LnProvenWeight"`

	// The vector commitment root of the top N accounts to sign the next StateProof.
	VotersCommitment []byte `json:"VotersCommitment"`
}

// Represents a key-value pair in an application store.
type TealKeyValue struct {
	Key string `json:"key"`

	// Represents a TEAL value.
	Value TealValue `json:"value"`
}

// Represents a key-value store for use in an application.
type TealKeyValueStore = []TealKeyValue

// Represents a TEAL value.
type TealValue struct {
	// \[tb\] bytes value.
	Bytes string `json:"bytes"`

	// \[tt\] value type. Value `1` refers to **bytes**, value `2` refers to **uint**
	Type uint64 `json:"type"`

	// \[ui\] uint value.
	Uint uint64 `json:"uint"`
}

// AccountInformationParams defines parameters for AccountInformation.
type AccountInformationParams struct {
	// Configures whether the response object is JSON or MessagePack encoded.
	Format *AccountInformationParamsFormat `form:"format,omitempty" json:"format,omitempty"`

	// When set to `all` will exclude asset holdings, application local state, created asset parameters, any created application parameters. Defaults to `none`.
	Exclude *AccountInformationParamsExclude `form:"exclude,omitempty" json:"exclude,omitempty"`
}

// AccountInformationParamsFormat defines parameters for AccountInformation.
type AccountInformationParamsFormat string

// AccountInformationParamsExclude defines parameters for AccountInformation.
type AccountInformationParamsExclude string

// AccountApplicationInformationParams defines parameters for AccountApplicationInformation.
type AccountApplicationInformationParams struct {
	// Configures whether the response object is JSON or MessagePack encoded.
	Format *AccountApplicationInformationParamsFormat `form:"format,omitempty" json:"format,omitempty"`
}

// AccountApplicationInformationParamsFormat defines parameters for AccountApplicationInformation.
type AccountApplicationInformationParamsFormat string

// AccountAssetInformationParams defines parameters for AccountAssetInformation.
type AccountAssetInformationParams struct {
	// Configures whether the response object is JSON or MessagePack encoded.
	Format *AccountAssetInformationParamsFormat `form:"format,omitempty" json:"format,omitempty"`
}

// AccountAssetInformationParamsFormat defines parameters for AccountAssetInformation.
type AccountAssetInformationParamsFormat string

// GetPendingTransactionsByAddressParams defines parameters for GetPendingTransactionsByAddress.
type GetPendingTransactionsByAddressParams struct {
	// Truncated number of transactions to display. If max=0, returns all pending txns.
	Max *uint64 `form:"max,omitempty" json:"max,omitempty"`

	// Configures whether the response object is JSON or MessagePack encoded.
	Format *GetPendingTransactionsByAddressParamsFormat `form:"format,omitempty" json:"format,omitempty"`
}

// GetPendingTransactionsByAddressParamsFormat defines parameters for GetPendingTransactionsByAddress.
type GetPendingTransactionsByAddressParamsFormat string

// GetBlockParams defines parameters for GetBlock.
type GetBlockParams struct {
	// Configures whether the response object is JSON or MessagePack encoded.
	Format *GetBlockParamsFormat `form:"format,omitempty" json:"format,omitempty"`
}

// GetBlockParamsFormat defines parameters for GetBlock.
type GetBlockParamsFormat string

// GetTransactionProofParams defines parameters for GetTransactionProof.
type GetTransactionProofParams struct {
	// The type of hash function used to create the proof, must be one of:
	// * sha512_256
	// * sha256
	Hashtype *GetTransactionProofParamsHashtype `form:"hashtype,omitempty" json:"hashtype,omitempty"`

	// Configures whether the response object is JSON or MessagePack encoded.
	Format *GetTransactionProofParamsFormat `form:"format,omitempty" json:"format,omitempty"`
}

// GetTransactionProofParamsHashtype defines parameters for GetTransactionProof.
type GetTransactionProofParamsHashtype string

// GetTransactionProofParamsFormat defines parameters for GetTransactionProof.
type GetTransactionProofParamsFormat string

// TealCompileParams defines parameters for TealCompile.
type TealCompileParams struct {
	// When set to `true`, returns the source map of the program as a JSON. Defaults to `false`.
	Sourcemap *bool `form:"sourcemap,omitempty" json:"sourcemap,omitempty"`
}

// TealDryrunJSONBody defines parameters for TealDryrun.
type TealDryrunJSONBody = DryrunRequest

// GetPendingTransactionsParams defines parameters for GetPendingTransactions.
type GetPendingTransactionsParams struct {
	// Truncated number of transactions to display. If max=0, returns all pending txns.
	Max *uint64 `form:"max,omitempty" json:"max,omitempty"`

	// Configures whether the response object is JSON or MessagePack encoded.
	Format *GetPendingTransactionsParamsFormat `form:"format,omitempty" json:"format,omitempty"`
}

// GetPendingTransactionsParamsFormat defines parameters for GetPendingTransactions.
type GetPendingTransactionsParamsFormat string

// PendingTransactionInformationParams defines parameters for PendingTransactionInformation.
type PendingTransactionInformationParams struct {
	// Configures whether the response object is JSON or MessagePack encoded.
	Format *PendingTransactionInformationParamsFormat `form:"format,omitempty" json:"format,omitempty"`
}

// PendingTransactionInformationParamsFormat defines parameters for PendingTransactionInformation.
type PendingTransactionInformationParamsFormat string

// TealDryrunJSONRequestBody defines body for TealDryrun for application/json ContentType.
type TealDryrunJSONRequestBody = TealDryrunJSONBody
