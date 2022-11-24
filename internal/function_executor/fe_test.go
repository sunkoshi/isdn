package function_executor

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/orted-org/dambda/internal/lang_handler"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestFunctionExecutor(t *testing.T) {
	file, err := os.ReadFile("./code.py")
	require.NoError(t, err)
	require.NotEmpty(t, file)

	lh, err := lang_handler.New()
	require.NoError(t, err)
	require.NotEmpty(t, lh)

	fe, err := New(lh, FunctionExecutorParams{
		RequestID: "1",
		Code:      string(file),
		Language:  "Python 3",
		Input:     FunctionInput{},
	})
	require.NoError(t, err)
	require.NotEmpty(t, fe)

	output, err := fe.Run(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, output)
	fmt.Println(output)
}
