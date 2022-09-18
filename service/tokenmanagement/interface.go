// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tokenmanagement

import (
	"context"
)

// Enables administrators to get all tokens and delete tokens for other users.
// Admins can either get every token, get a specific token by ID, or get all
// tokens for a particular user.
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type TokenManagementService interface {

	// Create a token on behalf of a service principal.
	CreateOboToken(ctx context.Context, request CreateOboTokenRequest) (*CreateOboTokenResponse, error)

	// Delete a token, specified by its ID.
	DeleteToken(ctx context.Context, request DeleteTokenRequest) error

	// DeleteTokenByTokenId calls DeleteToken, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteTokenByTokenId(ctx context.Context, tokenId string) error

	// Get a token, specified by its ID.
	GetTokenInfo(ctx context.Context, request GetTokenInfoRequest) (*TokenInfo, error)

	// GetTokenInfoByTokenId calls GetTokenInfo, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetTokenInfoByTokenId(ctx context.Context, tokenId string) (*TokenInfo, error)

	// List all tokens belonging to a workspace or a user.
	ListAllTokens(ctx context.Context, request ListAllTokensRequest) (*ListTokensResponse, error)
}
