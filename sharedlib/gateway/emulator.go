package gateway

import (
	"fmt"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-cli/flow/cli"
	emulator "github.com/onflow/flow-emulator"
	"github.com/onflow/flow-go-sdk"
)

type EmulatorGateway struct {
	emulator *emulator.Blockchain
}

func NewEmulatorGateway() *EmulatorGateway {
	return &EmulatorGateway{
		emulator: newEmulator(),
	}
}

func newEmulator() *emulator.Blockchain {
	b, err := emulator.NewBlockchain()
	if err != nil {
		panic(err)
	}
	return b
}

func (g *EmulatorGateway) GetAccount(address flow.Address) (*flow.Account, error) {
	return g.emulator.GetAccount(address)
}

func (g *EmulatorGateway) SendTransaction(tx *flow.Transaction, signer *cli.Account) (*flow.Transaction, error) {
	return nil, fmt.Errorf("Not Supported Yet")
}

func (g *EmulatorGateway) GetTransactionResult(tx *flow.Transaction) (*flow.TransactionResult, error) {
	return nil, fmt.Errorf("Not Supported Yet")
}

func (g *EmulatorGateway) ExecuteScript(script []byte, arguments []cadence.Value) (cadence.Value, error) {
	return nil, fmt.Errorf("Not Supported Yet")
}
