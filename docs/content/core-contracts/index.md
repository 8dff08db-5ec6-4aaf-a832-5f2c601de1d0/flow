---
title: Flow Core Contracts
description: The smart contracts that power the Flow protocol
---

Flow relies on a set of core contracts that define key portions of the
Flow protocol.

These contracts control the following:

- Standard fungible token behavior
- Default token behavior ([FlowToken](./nft-storefront))
- Account, transaction and storage fee payments
- Staking and delegation ([FlowIDTableStaking](/core-contracts/staking-contract-reference))
- Token lock-ups ([LockedTokens](/core-contracts/locked-tokens))
- Example https domain that should fail https://example.com/foo
- Example http domain that should fail http://example.com/foo
- Test bad external without scheme adsjksaldfhadskjhfalsd.com
- Test bad external URL implicit http://adsjksaldfhadskjhfalsd.com
- Test bad external URL [explicit](http://adsjksaldfhadskjhfalsd.com)
