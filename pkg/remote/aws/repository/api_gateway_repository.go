package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface"
	"github.com/cloudskiff/driftctl/pkg/remote/cache"
)

type ApiGatewayRepository interface {
	ListAllRestApis() ([]*apigateway.RestApi, error)
	GetAccount() (*apigateway.Account, error)
	ListAllApiKeys() ([]*apigateway.ApiKey, error)
	ListAllRestApiAuthorizers([]*apigateway.RestApi) ([]*apigateway.Authorizer, error)
}

type apigatewayRepository struct {
	client apigatewayiface.APIGatewayAPI
	cache  cache.Cache
}

func NewApiGatewayRepository(session *session.Session, c cache.Cache) *apigatewayRepository {
	return &apigatewayRepository{
		apigateway.New(session),
		c,
	}
}

func (r *apigatewayRepository) ListAllRestApis() ([]*apigateway.RestApi, error) {
	if v := r.cache.Get("apigatewayListAllRestApis"); v != nil {
		return v.([]*apigateway.RestApi), nil
	}

	var restApis []*apigateway.RestApi
	input := apigateway.GetRestApisInput{}
	err := r.client.GetRestApisPages(&input,
		func(resp *apigateway.GetRestApisOutput, lastPage bool) bool {
			restApis = append(restApis, resp.Items...)
			return !lastPage
		},
	)
	if err != nil {
		return nil, err
	}

	r.cache.Put("apigatewayListAllRestApis", restApis)
	return restApis, nil
}

func (r *apigatewayRepository) GetAccount() (*apigateway.Account, error) {
	if v := r.cache.Get("apigatewayGetAccount"); v != nil {
		return v.(*apigateway.Account), nil
	}

	account, err := r.client.GetAccount(&apigateway.GetAccountInput{})
	if err != nil {
		return nil, err
	}

	r.cache.Put("apigatewayGetAccount", account)
	return account, nil
}

func (r *apigatewayRepository) ListAllApiKeys() ([]*apigateway.ApiKey, error) {
	if v := r.cache.Get("apigatewayListAllApiKeys"); v != nil {
		return v.([]*apigateway.ApiKey), nil
	}

	var apiKeys []*apigateway.ApiKey
	input := apigateway.GetApiKeysInput{}
	err := r.client.GetApiKeysPages(&input,
		func(resp *apigateway.GetApiKeysOutput, lastPage bool) bool {
			apiKeys = append(apiKeys, resp.Items...)
			return !lastPage
		},
	)
	if err != nil {
		return nil, err
	}

	r.cache.Put("apigatewayListAllApiKeys", apiKeys)
	return apiKeys, nil
}

func (r *apigatewayRepository) ListAllRestApiAuthorizers(apis []*apigateway.RestApi) ([]*apigateway.Authorizer, error) {
	var authorizers []*apigateway.Authorizer
	for _, api := range apis {
		a := *api
		cacheKey := fmt.Sprintf("apigatewayListAllRestApiAuthorizers_api_%s", *a.Id)
		if v := r.cache.Get(cacheKey); v != nil {
			authorizers = append(authorizers, v.([]*apigateway.Authorizer)...)
			continue
		}

		input := &apigateway.GetAuthorizersInput{
			RestApiId: a.Id,
		}
		resources, err := r.client.GetAuthorizers(input)
		if err != nil {
			return nil, err
		}

		r.cache.Put(cacheKey, resources.Items)
		authorizers = append(authorizers, resources.Items...)
	}
	return authorizers, nil
}