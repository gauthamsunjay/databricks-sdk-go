package internal

import (
	"context"
	"encoding/base64"
	"path/filepath"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func myNotebookPath(t *testing.T, w *databricks.WorkspaceClient) string {
	ctx := context.Background()
	testDir := filepath.Join("/Users", me(t, w).UserName, ".sdk", RandomName("t-"))
	notebook := filepath.Join(testDir, RandomName("n-"))

	err := w.Workspace.MkdirsByPath(ctx, testDir)
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.Workspace.Delete(ctx, workspace.Delete{
			Path:      testDir,
			Recursive: true,
		})
		require.NoError(t, err)
	})

	return notebook
}

func TestAccWorkspaceIntegration(t *testing.T) {
	ctx, w := workspaceTest(t)
	notebook := myNotebookPath(t, w)

	// Import the test notebook
	err := w.Workspace.Import(ctx, workspace.Import{
		Path:      notebook,
		Format:    workspace.ExportFormatSource,
		Language:  workspace.LanguagePython,
		Content:   base64.StdEncoding.EncodeToString([]byte("# Databricks notebook source\nprint('hello from job')")),
		Overwrite: true,
	})
	require.NoError(t, err)

	// Get test notebook status
	getStatusResponse, err := w.Workspace.GetStatusByPath(ctx, notebook)
	require.NoError(t, err)
	assert.True(t, getStatusResponse.Language == workspace.LanguagePython)
	assert.True(t, getStatusResponse.Path == notebook)
	assert.True(t, getStatusResponse.ObjectType == workspace.ObjectTypeNotebook)

	// Export the notebook and assert the contents
	exportResponse, err := w.Workspace.Export(ctx, workspace.ExportRequest{
		DirectDownload: false,
		Format:         workspace.ExportFormatSource,
		Path:           notebook,
	})
	require.NoError(t, err)
	assert.True(t, exportResponse.Content == base64.StdEncoding.EncodeToString([]byte("# Databricks notebook source\nprint('hello from job')")))

	// Assert the test notebook is present in test dir using list api
	objects, err := w.Workspace.ListAll(ctx, workspace.ListWorkspaceRequest{
		Path: filepath.Dir(notebook),
	})
	require.NoError(t, err)

	paths, err := w.Workspace.ObjectInfoPathToObjectIdMap(ctx, workspace.ListWorkspaceRequest{
		Path: filepath.Dir(notebook),
	})
	require.NoError(t, err)
	assert.Equal(t, len(objects), len(paths))
	assert.Contains(t, paths, notebook)
}

func TestAccWorkspaceRecursiveListNoTranspile(t *testing.T) {
	ctx, w := workspaceTest(t)
	notebook := myNotebookPath(t, w)

	// Import the test notebook
	err := w.Workspace.Import(ctx, workspace.Import{
		Path:      notebook,
		Format:    workspace.ExportFormatSource,
		Language:  workspace.LanguagePython,
		Content:   base64.StdEncoding.EncodeToString([]byte("# Databricks notebook source\nprint('hello from job')")),
		Overwrite: true,
	})
	require.NoError(t, err)

	allMyNotebooks, err := w.Workspace.RecursiveList(ctx, filepath.Join("/Users", me(t, w).UserName))
	require.NoError(t, err)
	assert.True(t, len(allMyNotebooks) >= 1)
}
