package util

import "github.com/platformercloud/platformer-cli/internal/config"

var ContextUrl = config.GetDefaultContextURL()

const MizzenPath = "/mizzen"
var MizzenClusterRegistrationURL = ContextUrl + MizzenPath + "/api/v1/cluster"
var MizzenYAMLGenerationURL = ContextUrl + MizzenPath + "/generate/v1/agent/"

const AuthPath = "/auth"
var AuthOrganizationListURL = ContextUrl + AuthPath + "/api/v1/organization/list"
var AuthProjectListURL = ContextUrl + AuthPath + "/api/v1/organization/project/list"
var AuthValidTokenURL = ContextUrl + AuthPath + "/api/v1/user/logintime"
var AuthTokenCreateURL = ContextUrl + AuthPath + "/api/v1/serviceaccount/token/create"