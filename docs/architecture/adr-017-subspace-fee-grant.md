# ADR 017: Subspace fee grant

## Changelog

- Nov 14th, 2022: First draft;
- Nov 21th, 2022: First review.

## Status

ACCEPTED Not Implemented

## Abstract

This ADR introduces a new subspace-based fee grant method, which allows subspace owners/admins to pay the fees of subspace-related transactions on behalf of the users.

## Context

One of the major problems of current Web3-based applications that make them complicated to use by common people is the fact that they require users to have some tokens in order to pay transaction fees to perform various operations. For instance, in order to create a post on any Desmos-based social network, users have to:
1. understand what a _transaction_ is and why they need some DSM to broadcast it;
2. get some DSM either via an on-ramp service or by swapping existing funds.

The `x/feegrant` module of the Cosmos SDK allows anyone to pay fees on behalf of other users, meaning that the latter can ideally perform any kind of transaction without even knowing the concept of fees. However, the `x/feegrant` allowance is not subspace-specific: this means that subspace admins/owners might unexpectingly pay the fees for transactions that are related to other subspaces.

## Decision

We will implement a subspace-specific fee grant process, based on  the `x/feegrant` module, that allows subspace fees providers to pay fees for the users inside the specified subspace. This system will allow subspace owners/admins to issue a fee grant towards either a users group or a single user, so that they can later leverage this grant to execute subspace-related transactions without having to worry about paying fees.

### DeductFeeDecorator

Currently, the `x/auth` module provides a `DeductFeeDecorator` that deducts the fees of a transaction from the signer/feepayer. We will build a new subspace-specific `DeductFeeDecorator` based on the existing one that acts as follows:

1. check if all the messages in the transaction are messages related to the same subspace;
2. perform the following operation:
   1. if all the messages are all related to the same subspace and a fee grant exists for the fee payer inside the subspace, run the `x/subspaces` `DeductFeeDecorator`
   2. otherwise, run the `x/auth` `DeductFeeDecorator`

### Types

#### Allowance

The `x/feegrant` module already provides an interface (`FeeAllowanceI`) to represent a generic allowance, along with a set of useful implementations like `BasicAllowance`, `AllowedMsgAllowance` and `PeriodicAllowance`. Since we are not going to introduce new kinds of allowances, nor we are going to edit how an allowance is represented, we will reuse these types.

#### User Grant

The `x/feegrant` module provides a `Grant` object used to store the `granter`, `grantee` and what kind of `allowance` is granted to a specific user. Since this is enough information for us as well, we will use the same object to store user-related fee grants.

#### Group Grant

In order to represent a fee grant granted to a user group, we will implement a `GroupGrant` type. This will contain the `granter`, `group_id` and what kind of `allowance` is granted to the group:

```protobuf
message GroupGrant {
    // granter is the address of the user granting an allowance of their funds.
  string granter = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];

  // the id of the group being granted an allowance of another user's funds.
  uint32 group_id = 2;

  // allowance can be any of basic, periodic, allowed fee allowance.
  google.protobuf.Any allowance = 3 [(cosmos_proto.accepts_interface) = "FeeAllowanceI"];
}
```

### Store

#### Group grant

In order to simplify the management of group allowances, each granted group allowance will be stored using the following key:
```
SubspaceGroupAllowancePrefix | SubspaceID | GranterAddress | GroupID | -> Protobuf(GroupGrant)
```

This structure allows granters to easily manage the group allowance inside a subspace by iterating over all allowances for the granters, which will be the most used query. On the other hand, it allows to perform the following queries: 
1. finding a specific grant given the granter address.   
   This can be done in O(N), where N is the number of granted groups by the granter.
2. finding all the grants given to a user based on the groups they are part of.
   This can be done in O(N * log(M)), where N is the number of all the grants inside the subspace, and M is the number of members in the user group.

Considering that: 
- only users with permissions will be able to create grants in each subspace
- there will be a small number of grants each granter will create

then the time complexity of searching if a user has been granted a fee grant should not become too large.

#### User grant

Each subspace fee granted allowance will be stored using keys with the following structure:
```
SubspaceUserAllowancePrefix | SubspaceID | GranterAddress | GranteeAddress |-> Protobuf(Grant)
```

This structure allows granters to easily manage their grants inside a subspace by iterating over all grants for the granters, which will be the most used query. On the other hand, it allows the following queries:
1. retrieving a specific grant.  
   This will require grantees to know who the granter address, and can be done in O(1).
2. retrieving all the grants that a user has been granted. 
   This can be done by iterating over all the granters' grants in O(N), where N is the number of all the grants inside the given subspace.

### Expired Allowance

The subspace owners may want to set the allowance to expire after a certain amount of time. In order to reduce the database size, we will use a custom `BeginBlocker` to remove all expired allowances.

#### Store

In order to make it easier to remove expired allowances, we will implement queues sorted by expiration time. Then, during the `BeginBlock` phase, we iterate over all expired keys before the block time and remove them. The structure of the queues keys will be the following:

```
ExpiringAllowanceQueuePrefix | Expired Time | Subspace ID | Granter Address | (User Address or Group ID) -> 0x01
```

### `Msg` Service

In order to allow subspace fees providers to grant an allowance for the users, we will have the following operations:

- grant an allowance to a user
- revoke an allowance to a user
- grant an allowance to a group
- revoke an allowance to a group

```protobuf
service Msg {
    // GrantUserAllowance allows the granter to grant a fee allowance to the grantee.
    rpc GrantUserAllowance(MsgGrantUserAllowance) returns(MsgGrantUserAllowanceResponse);

    // RevokeUserAllowance allows a granter to revoke any existing allowance that has to been granted to the grantee.
    rpc RevokeUserAllowance(MsgRevokeUserAllowance) returns(MsgRevokeUserAllowanceResponse);

    // GrantGroupAllowance allows the granter to grant a fee allowance to the group.
    rpc GrantGroupAllowance(MsgGrantGroupAllowance) returns(MsgGrantGroupAllowanceResponse);

    // RevokeGroupAllowance allows a granter to revoke any existing allowance that has to been granted to the group.
    rpc RevokeGroupAllowance(MsgRevokeGroupAllowance) returns(MsgRevokeGroupAllowanceResponse);
}

// MsgGrantUserAllowance adds permissions for the grantee to spend up allowance of fees from the granter inside the given subspace.
message MsgGrantUserAllowance {
    // the id of the subspace where the granter grants the allowance to the grantee.
    uint64 subspace_id = 1;

    // the address of the user granting an allowance of their funds.
    string granter = 2;

    // the address of the user being granted an allowance of another user's funds.
    string grantee = 3;

    // allowance can be any of fee allowances which implements FeeAllowanceI.
    google.protobuf.Any allowance = 4 [(cosmos_proto.accepts_interface) = "FeeAllowanceI"];
}

// MsgGrantUserAllowanceResponse defines the Msg/GrantAllowanceResponse response type.
message MsgGrantUserAllowanceResponse {}

// MsgRevokeUserAllowance removes any existing allowance from granter to the grantee inside the subspace.
message MsgRevokeUserAllowance {
    // the id of the subspace where the granter grants the allowance to the grantee.
    uint64 subspace_id = 1;

    // the address of the user granting an allowance of their funds.
    string granter = 2;

    // the address of the user being granted an allowance of another user's funds.
    string grantee = 3;
}

// MsgRevokeUserAllowanceResponse defines the Msg/RevokeAllowanceResponse response type.
message MsgRevokeUserAllowanceResponse {}

// MsgGrantGroupAllowance adds permissions for the group to spend up allowance of fees from the granter inside the given subspace.
message MsgGrantGroupAllowance {
    // the id of the subspace where the granter grants the allowance to the grantee.
    uint64 subspace_id = 1;

    // the id of the group being granted an allowance of another user's funds.  
    uint32 group_id = 2;

    // the address of the user granting an allowance of their funds.
    string granter = 3;

    // allowance can be any of fee allowances which implements FeeAllowanceI.
    google.protobuf.Any allowance = 4 [(cosmos_proto.accepts_interface) = "FeeAllowanceI"];
}

// MsgGrantGroupAllowanceResponse defines the Msg/GrantAllowanceResponse response type.
message MsgGrantGroupAllowanceResponse {}

// MsgRevokeGroupAllowance removes any existing allowance from granter to the group inside the subspace.
message MsgRevokeGroupAllowance {
    // the id of the subspace where the granter grants the allowance to the group.
    uint64 subspace_id = 1;

    // the id of the group being granted an allowance of another user's funds.
    uint32 group_id = 2;

    // the address of the user granting an allowance of their funds.
    string granter = 3;
}

// MsgRevokeGroupAllowanceResponse defines the Msg/RevokeAllowanceResponse response type.
message MsgRevokeUserAllowanceResponse {}
```

### `Query` Service

In order to allow clients to easily query for allowances we will implement the following new queries:

```protobuf
service Query {
    // UserAllowances returns all the grants for users.
    rpc UserAllowances(QueryUserAllowancesRequest) returns (QueryAllowancesResponse) {
        option (google.api.http).get = "/desmos/subspaces/v3/subspaces/{subspace_id}/allowances/users";
    }
    // GroupAllowances returns all the grants for groups.
    rpc GroupAllowances(QueryGroupAllowancesRequest) returns(QueryGroupAllowancesResponse) {
        option (google.api.http).get = "/desmos/subspaces/v3/subspaces/{subspace_id}/allowances/groups";
    }
}

// QueryUserAllowancesRequest is the request type for the Query/UserAllowances RPC method.
message QueryUserAllowancesRequest {
    // the id of the subspace where the granter grants the allowance to the grantee.
    uint64 subspace_id = 1;

    // (optional) the address of the user granting an allowance of their funds.
    string granter = 2;

    // (Optional) the address of the user being granted an allowance of another user's funds.
    string grantee = 3;

    // pagination defines an pagination for the request.
    cosmos.base.query.v1beta1.PageRequest pagination = 4;
}

// QueryUserAllowancesResponse is the response type for the Query/UserAllowances RPC method.
message QueryUserAllowancesResponse {
    // allowances are allowance's granted for grantee by granter.
    repeated cosmos.feegrant.v1beta1.Grant allowances = 1;

    // pagination defines an pagination for the response.
    cosmos.base.query.v1beta1.PageResponse pagination = 2;
}

// QueryGroupAllowancesRequest is the request type for the Query/GroupAllowances RPC method.
message QueryGroupAllowancesRequest {
    // the id of the subspace where the granter grants the allowance to the grantee.
    uint64 subspace_id = 1;

    // (optional) the address of the user granting an allowance of their funds.
    string granter = 2;

    // (Optional) the id of the group being granted an allowance of another user's funds.
    uint32 group_id = 3;

    // pagination defines an pagination for the request.
    cosmos.base.query.v1beta1.PageRequest pagination = 4;
}

// QueryGroupAllowancesResponse is the response type for the Query/GroupAllowances RPC method.
message QueryGroupAllowancesResponse {
    // allowances are allowance's granted for grantee by granter.
    repeated cosmos.subspace.v3.GroupGrant allowances = 1;

    // pagination defines an pagination for the response.
    cosmos.base.query.v1beta1.PageResponse pagination = 2;
}
```

## Consequences

### Backwards Compatibility

The proposed solution introduces a new set of store keys, thus it is completely backward compatible.

### Positive

- Allow subspace fee providers to grant fees allowances to their users.

### Negative

- Storing extra subspace grant info takes up more storage space;
- Having more checks during the fee deduction phase slows down the transaction processing operation;
- Removing expired allowances at the block begin phase increases the block production time.  

### Neutral

- Not known

## References
- [Cosmos SDK fee grant concepts](https://docs.cosmos.network/v0.46/modules/feegrant/01_concepts.html)