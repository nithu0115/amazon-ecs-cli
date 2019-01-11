// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package secretsmanager

//go:generate mockgen.sh github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface SecretsManagerAPI mock/sdk/secretsmanageriface_mock.go
//go:generate mockgen.sh github.com/aws/amazon-ecs-cli/ecs-cli/modules/clients/aws/secretsmanager SMClient mock/client.go
