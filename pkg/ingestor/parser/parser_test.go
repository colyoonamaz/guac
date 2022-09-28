//
// Copyright 2022 The GUAC Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package parser

import (
	"reflect"
	"testing"

	"github.com/guacsec/guac/internal/testing/ingestor/testdata"
	"github.com/guacsec/guac/pkg/assembler"
	"github.com/guacsec/guac/pkg/handler/processor"
	"github.com/guacsec/guac/pkg/ingestor/verifier"
)

var (
	docTree = processor.DocumentNode{
		Document: &testdata.Ite6DSSEDoc,
		Children: []*processor.DocumentNode{
			{
				Document: &testdata.Ite6SLSADoc,
				Children: []*processor.DocumentNode{},
			},
		},
	}

	graphInput = []assembler.AssemblerInput{{
		Nodes: testdata.DsseNodes,
		Edges: testdata.DsseEdges,
	}, {
		Nodes: testdata.SlsaNodes,
		Edges: testdata.SlsaEdges,
	}}
)

func TestParseDocumentTree(t *testing.T) {
	err := verifier.RegisterVerifier(testdata.NewMockSigstoreVerifier(), "sigstore")
	if err != nil {
		t.Errorf("verifier.RegisterVerifier() failed with error: %v", err)
	}
	tests := []struct {
		name    string
		tree    processor.DocumentTree
		want    []assembler.AssemblerInput
		wantErr bool
	}{{
		name:    "testing",
		tree:    processor.DocumentTree(&docTree),
		want:    graphInput,
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDocumentTree(tt.tree)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDocumentTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDocumentTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
