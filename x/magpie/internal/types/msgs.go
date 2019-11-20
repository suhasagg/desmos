package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgCreateSession defines the MsgCreateSession message
type MsgCreateSession struct {
	Owner         sdk.AccAddress `json:"owner"`
	Namespace     string         `json:"namespace"`
	ExternalOwner string         `json:"external_owner"`
	PubKey        string         `json:"pub_key"`
	Signature     string         `json:"signature"`
}

// NewMsgCreateSession is the constructor of MsgCreateSession
func NewMsgCreateSession(owner sdk.AccAddress, namespace string, externalOwner string, pubkey string, signature string) MsgCreateSession {
	return MsgCreateSession{
		Owner:         owner,
		Namespace:     namespace,
		ExternalOwner: externalOwner,
		PubKey:        pubkey,
		Signature:     signature,
	}
}

// Route should return the name of the module
func (msg MsgCreateSession) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateSession) Type() string { return ActionCreationSession }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateSession) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(fmt.Sprintf("Invalid session owner: %s", msg.Owner))
	}

	if len(strings.TrimSpace(msg.Namespace)) == 0 {
		return sdk.ErrUnknownRequest("Session namespace cannot be empty")
	}

	if len(strings.TrimSpace(msg.PubKey)) == 0 {
		return sdk.ErrUnknownRequest("Signer public key cannot be empty")
	}

	// The external signer address doesn't have to exist on Desmos
	if len(strings.TrimSpace(msg.ExternalOwner)) == 0 {
		return sdk.ErrUnknownRequest("Session external owner cannot be empty")
	}

	if len(strings.TrimSpace(msg.Signature)) == 0 {
		return sdk.ErrUnknownRequest("Session signature cannot be empty")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateSession) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateSession) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
