package function_executor

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/orted-org/isdn/internal/lang_handler"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestFunctionExecutor(t *testing.T) {
	file, err := os.Open("./test_codes/code.cpp")
	require.NoError(t, err)
	require.NotEmpty(t, file)

	lh, err := lang_handler.New()
	require.NoError(t, err)
	require.NotEmpty(t, lh)

	fe, err := New(lh, FunctionExecutorParams{
		RequestID: "1",
		Code:      file,
		Language:  "C++ 11",
		Input:     FunctionInput{},
	})
	require.NoError(t, err)
	require.NotEmpty(t, fe)

	output := fe.Run(context.Background())
	require.NotEmpty(t, output)
	fmt.Println(output)
}
