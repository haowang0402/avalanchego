// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package x

import (
	"github.com/haowang0402/avalanchego/codec"
	"github.com/haowang0402/avalanchego/codec/linearcodec"
	"github.com/haowang0402/avalanchego/utils/wrappers"
	"github.com/haowang0402/avalanchego/vms/avm"
	"github.com/haowang0402/avalanchego/vms/nftfx"
	"github.com/haowang0402/avalanchego/vms/propertyfx"
	"github.com/haowang0402/avalanchego/vms/secp256k1fx"
)

const (
	// CodecVersion is the current default codec version
	CodecVersion = 0

	// TODO: verify that these are correct
	SECP256K1FxIndex = 0
	NFTFxIndex       = 1
	PropertyFxIndex  = 2
)

// Codecs do serialization and deserialization
var (
	Codec codec.Manager
)

func init() {
	c := linearcodec.NewDefault()
	Codec = codec.NewDefaultManager()

	errs := wrappers.Errs{}
	errs.Add(
		c.RegisterType(&avm.BaseTx{}),
		c.RegisterType(&avm.CreateAssetTx{}),
		c.RegisterType(&avm.OperationTx{}),
		c.RegisterType(&avm.ImportTx{}),
		c.RegisterType(&avm.ExportTx{}),
		c.RegisterType(&secp256k1fx.TransferInput{}),
		c.RegisterType(&secp256k1fx.MintOutput{}),
		c.RegisterType(&secp256k1fx.TransferOutput{}),
		c.RegisterType(&secp256k1fx.MintOperation{}),
		c.RegisterType(&secp256k1fx.Credential{}),
		c.RegisterType(&nftfx.MintOutput{}),
		c.RegisterType(&nftfx.TransferOutput{}),
		c.RegisterType(&nftfx.MintOperation{}),
		c.RegisterType(&nftfx.TransferOperation{}),
		c.RegisterType(&nftfx.Credential{}),
		c.RegisterType(&propertyfx.MintOutput{}),
		c.RegisterType(&propertyfx.OwnedOutput{}),
		c.RegisterType(&propertyfx.MintOperation{}),
		c.RegisterType(&propertyfx.BurnOperation{}),
		c.RegisterType(&propertyfx.Credential{}),

		Codec.RegisterCodec(CodecVersion, c),
	)
	if errs.Errored() {
		panic(errs.Err)
	}
}
