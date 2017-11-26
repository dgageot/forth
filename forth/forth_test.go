package forth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDump(t *testing.T) {
	assert.Equal(t, "[]\n", Run("dump"))
	assert.Equal(t, "[1]\n", Run("1 dump"))
	assert.Equal(t, "[1, 2]\n", Run("1 2 dump"))
}

func TestPrint(t *testing.T) {
	assert.Equal(t, "42\n", Run("42 ."))
	assert.Equal(t, "1337\n", Run("1337 ."))
	assert.Equal(t, "1\n2\n", Run("2 1 . ."))
}

func TestPrintNothing(t *testing.T) {
	assert.Equal(t, "", Run("42"))
}

func TestEmpty(t *testing.T) {
	assert.Empty(t, Run(""))
}

func TestFailToPrint(t *testing.T) {
	defer func() {
		r := recover()
		require.EqualValues(t, r, "unable to pop")
	}()

	Run(".")
}

func TestSum(t *testing.T) {
	assert.Equal(t, "[10]\n", Run("1 9 + dump"))
	assert.Equal(t, "[10.1]\n", Run("9.1 1 + dump"))
}

func TestSub(t *testing.T) {
	assert.Equal(t, "[8]\n", Run("2 10 - dump"))
	assert.Equal(t, "[-8.2]\n", Run("10.2 2 - dump"))
}

func TestMul(t *testing.T) {
	assert.Equal(t, "[20]\n", Run("2 10 * dump"))
	assert.Equal(t, "[3]\n", Run("1.5 2 * dump"))
	assert.Equal(t, "[2.25]\n", Run("1.5 1.5 * dump"))
}

func TestDiv(t *testing.T) {
	assert.Equal(t, "[1]\n", Run("12 12 / dump"))
	assert.Equal(t, "[1.5]\n", Run("2 3 / dump"))
}

func TestMod(t *testing.T) {
	assert.Equal(t, "[0]\n", Run("3 12 mod dump"))
	assert.Equal(t, "[1]\n", Run("3 13 mod dump"))
	assert.Equal(t, "[1]\n", Run("3.1 13.1 mod dump"))
}

func TestMax(t *testing.T) {
	assert.Equal(t, "[10]\n", Run("1 10 max dump"))
	assert.Equal(t, "[10]\n", Run("10 1 max dump"))
}

func TestMin(t *testing.T) {
	assert.Equal(t, "[1]\n", Run("1 10 min dump"))
	assert.Equal(t, "[1]\n", Run("10 1 min dump"))
}

func TestNegate(t *testing.T) {
	assert.Equal(t, "[-1]\n", Run("1 negate dump"))
	assert.Equal(t, "[1.1]\n", Run("-1.1 negate dump"))
}

func TestAbs(t *testing.T) {
	assert.Equal(t, "[0]\n", Run("0 abs dump"))
	assert.Equal(t, "[0]\n", Run("-0 abs dump"))
	assert.Equal(t, "[1]\n", Run("1 abs dump"))
	assert.Equal(t, "[1]\n", Run("-1 abs dump"))
	assert.Equal(t, "[1.1]\n", Run("1.1 abs dump"))
	assert.Equal(t, "[1.1]\n", Run("-1.1 abs dump"))
}

func TestDup(t *testing.T) {
	assert.Equal(t, "[8, 8]\n", Run("8 dup dump"))
}

func TestSwap(t *testing.T) {
	assert.Equal(t, "[2, 1]\n", Run("1 2 swap dump"))
}

func TestDrop(t *testing.T) {
	assert.Equal(t, "[1]\n", Run("1 2 drop dump"))
}

func TestEqual(t *testing.T) {
	assert.Equal(t, "[0]\n", Run("1 2 = dump"))
	assert.Equal(t, "[-1]\n", Run("1 1 = dump"))
}

func TestFunction(t *testing.T) {
	assert.Equal(t, "[16]\n", Run(": square dup * ; 4 square dump"))
	assert.Equal(t, "[4, 4, 4, 4]\n", Run(": dup3 dup dup dup ; 4 dup3 dump"))
}

func TestConstant(t *testing.T) {
	assert.Equal(t, "[10]\n", Run("10 constant C C dump"))
}

func TestBye(t *testing.T) {
	assert.Equal(t, "1\n", Run("1 . bye 2 ."))
}
