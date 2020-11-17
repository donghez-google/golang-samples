// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START asset_quickstart_analyze_iam_policy]

// Sample analyze_iam_policy analyzes accessible IAM policies that match a request.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	asset "cloud.google.com/go/asset/apiv1"
	assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1"
)

func main() {
	scope := flag.String("scope", "", "Scope of the analysis.")
	identity := flag.String("identity", "", "Query identity.")
	expandResources := flag.Bool("expandResources", false, "Query option of expanding resources.")
	outputResourceEdges := flag.Bool("outputResourceEdges", false, "Query option of outputing resource edges.")
	flag.Parse()
	ctx := context.Background()
	client, err := asset.NewClient(ctx)
	if err != nil {
		log.Fatalf("asset.NewClient: %v", err)
	}
	defer client.Close()

	req := &assetpb.AnalyzeIamPolicyRequest{
		AnalysisQuery: &assetpb.IamPolicyAnalysisQuery{
			Scope: *scope,
			IdentitySelector: &assetpb.IamPolicyAnalysisQuery_IdentitySelector{
				Identity: *identity,
			},
			Options: &assetpb.IamPolicyAnalysisQuery_Options{
				ExpandResources:     *expandResources,
				OutputResourceEdges: *outputResourceEdges,
			},
		},
	}

	op, err := client.AnalyzeIamPolicy(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	for index, result := range op.MainAnalysis.AnalysisResults {
		fmt.Println(index, result)
	}
}

// [END asset_quickstart_analyze_iam_policy]

