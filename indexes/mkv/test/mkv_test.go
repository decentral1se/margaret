// SPDX-License-Identifier: MIT

package test

import (
	"testing"

	"go.cryptoscope.co/margaret/indexes/test"
)

func TestMKV(t *testing.T) {
	t.Run("SetterIndex", test.RunSetterIndexTests)
	t.Run("SeqSetterIndex", test.RunSeqSetterIndexTests)
}
