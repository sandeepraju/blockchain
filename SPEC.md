# Specification

## Block

```json
{
    "Kind": "Block",
    "Version": "v0.0.1-dev",
    "Header": {
        "HashPrev": "",
        "HashMerkleRoot": "",
        "Time": "",
        "Target": 0,
        "Nonce": 0
    },
    "Transactions": [
        {
            "Kind": "Transaction",

        }
    ]
}
```

## Transaction

```json
{
    "Kind": "Transaction",
    "Version": "v0.0.1-dev",
    "Inputs": [],
    "Outputs": [],
    "LockTime": ""
}
```

## Blockchain

```json
{
    "Kind": "Blockchain",
    "Version": "v0.0.1-dev",
    "Blocks": []
}
```

## Math

- Total Number of coins: `100,000,000C`
- Divisibility Limit of a coin: 10^-9 (nano; billionth of a coin)
- `1C = 1,000,000,000 nC`
- Standard unit is nC across the codebase