# mq
Query mnemonic fields by location indices.

Entering mnemonic on hardware wallets often requires entering
words of mnemonic sentence in a specific order. This tool lets
you query mnemonic words by their location indices hiding what
you query and clearing out screen for security purposes.

## disclaimer
> The use of this tool does not guarantee security or usability for any
> particular purpose. Please review the code and use at your own risk.

> Don't trust, verify

## installation
This step assumes you have [Go compiler toolchain](https://go.dev/dl/)
installed on your system.

```bash
go install github.com/kubetrail/mq@latest
```

## usage
```bash
mq ${MNEMONIC}
```
