// Copyright (C) 2017, Beijing Bochen Technology Co.,Ltd.  All rights reserved.
//
// This file is part of L0
// 
// The L0 is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// 
// The L0 is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// 
// GNU General Public License for more details.
// 
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package accounts

import (
	"github.com/bocheninc/L0/components/crypto"

	"github.com/bocheninc/L0/components/utils"
)

// account type
const (
	AccountTypeCommon uint32 = iota
	AccountTypeChain
	AccountTypeHot
	AccountTypeIssue
)

// AccountVariety max index of account type
const AccountVariety = 3

// Account the definition of common account struct
type Account struct {
	URL         URL `json:"url"`
	AccountType uint32
	//AhainCoords []uint32
	PublicKey *crypto.PublicKey
	Address   Address
}

// Serialize returns the serialized res of an account var
func (a *Account) Serialize() []byte {
	return utils.Serialize(a)
}

// Deserializ restore an account var from an serialized bytes
func (a *Account) Deserialize(accountBytes []byte) {
	utils.Deserialize(accountBytes, a)
}
