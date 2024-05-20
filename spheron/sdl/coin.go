package sdl

import (
	"errors"
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"gopkg.in/yaml.v3"
)

// v2Coin is an alias sdk.Coin to allow our custom UnmarshalYAML
// for now it supports PoC when actual pricing is specified as two fields
// aka amount and denom. we let UnmarshalYAML to deal with that and put result
// into Value field.
// discussion https://github.com/akash-network/node/issues/771
type v2Coin struct {
	Value sdk.DecCoin `yaml:"-"`
}

var errInvalidCoinAmount = errors.New("invalid coin amount")

func (sdl *v2Coin) UnmarshalYAML(node *yaml.Node) error {
	parsedCoin := struct {
		Amount string `yaml:"amount"`
		Denom  string `yaml:"denom"`
	}{}

	if err := node.Decode(&parsedCoin); err != nil {
		return err
	}

	value := new(big.Int)
	value.SetString(parsedCoin.Amount, 10) // base 10

	//set precision to 18 so our value doesnt get multiplied by 18 in conversion
	amount := sdk.NewDecFromBigIntWithPrec(value, 18) //read this function as: value has prec 18 so convert it like that (it wont add any 0)

	if amount.IsZero() {
		return fmt.Errorf("%w: amount is zero", errInvalidCoinAmount)
	}

	// Never pass negative amounts to cosmos SDK DecCoin
	if amount.IsNegative() {
		return fmt.Errorf("%w: amount %q is negative", errNegativeValue, amount.String())
	}

	coin := sdk.NewDecCoinFromDec(parsedCoin.Denom, amount)

	*sdl = v2Coin{
		Value: coin,
	}

	return nil
}
