package main

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "github.com/vulcangz/gf2-authz/api/authz/v1"
	sdk "github.com/vulcangz/gf2-authz/pkg/sdk"
	"github.com/vulcangz/gf2-authz/pkg/sdk/rule"
	"google.golang.org/grpc/codes"
)

func main() {
	// Initialize client
	authzClient, err := sdk.NewClient(&sdk.Config{
		ClientID:       "cab02d47-03f2-11f0-a218-0242ac110002",
		ClientSecret:   "bIY37jTBQR_BcB17PXVr5M-qB1CgkPfO4fTaG4zqaETY8q5g",
		GrpcAddr:       "127.0.0.1:8081",
		ServiceName:    "authz",
		RegistrySchema: "file",
	})
	if err != nil {
		panic(fmt.Sprintf("cannot initialize client: %v", err))
	}

	ctx := context.Background()

	// Create a principal named "user-123"
	principalGetResponse, err := authzClient.PrincipalGet(ctx, &v1.PrincipalGetRequest{
		Id: "user-123",
	})

	if err != nil {
		gc := gerror.Code(err)
		if gc.Code() != gconv.Int(codes.NotFound) {
			panic(err.Error())
		}

		principalResponse, err := authzClient.PrincipalCreate(ctx, &v1.PrincipalCreateRequest{
			Id: "user-123",
			Attributes: []*v1.Attribute{
				{Key: "email", Value: "johndoe@acme.tld"},
			},
		})

		if err != nil {
			gc = gerror.Code(err)
			if gc.Code() != gconv.Int(codes.AlreadyExists) {
				panic(fmt.Sprintf("cannot create principal: %v", err))
			}
		} else {
			fmt.Printf("Principal created: %s\n", principalResponse.GetPrincipal().GetId())
		}
	} else {
		fmt.Printf("Principal already exists: %s\n", principalGetResponse.GetPrincipal().GetId())
	}

	// Create a first post resource not matching attribute value
	resourceGetResponse, err := authzClient.ResourceGet(ctx, &v1.ResourceGetRequest{
		Id: "post.123",
	})

	if err != nil {
		gc := gerror.Code(err)
		if gc.Code() != gconv.Int(codes.NotFound) {
			panic(err.Error())
		}
		resourceResponse, err := authzClient.ResourceCreate(ctx, &v1.ResourceCreateRequest{
			Id:    "post.123",
			Kind:  "post",
			Value: "123",
			Attributes: []*v1.Attribute{
				{Key: "owner_email", Value: "someoneelse@acme.tld"},
			},
		})

		if err != nil {
			gc = gerror.Code(err)
			if gc.Code() != gconv.Int(codes.AlreadyExists) {
				panic(fmt.Sprintf("cannot create resource: %v", err))
			}
		} else {
			fmt.Printf("Resource created: %s\n", resourceResponse.GetResource().GetId())
		}
	} else {
		fmt.Printf("Resource already exists: %s\n", resourceGetResponse.GetResource().GetId())
	}

	// Create 2 other post resources matching attributes
	for _, identifier := range []string{"456", "789"} {
		resourceGetResponse, err := authzClient.ResourceGet(ctx, &v1.ResourceGetRequest{
			Id: "post." + identifier,
		})

		if err != nil {
			gc := gerror.Code(err)
			if gc.Code() != gconv.Int(codes.NotFound) {
				panic(err.Error())
			}

			resourceResponse, err := authzClient.ResourceCreate(ctx, &v1.ResourceCreateRequest{
				Id:    "post." + identifier,
				Kind:  "post",
				Value: identifier,
				Attributes: []*v1.Attribute{
					{Key: "owner_email", Value: "johndoe@acme.tld"},
				},
			})
			if err != nil {
				gc = gerror.Code(err)
				if gc.Code() != gconv.Int(codes.AlreadyExists) {
					panic(fmt.Sprintf("cannot create resource: %v", err))
				}
			} else {
				fmt.Printf("Resource created: %s\n", resourceResponse.GetResource().GetId())
			}
		} else {
			fmt.Printf("Resource already exists: %s\n", resourceGetResponse.GetResource().GetId())
		}
	}

	// Create a policy
	policyGetResponse, err := authzClient.PolicyGet(ctx, &v1.PolicyGetRequest{
		Id: "post-owners",
	})

	if err != nil {
		gc := gerror.Code(err)
		if gc.Code() != gconv.Int(codes.NotFound) {
			panic(err.Error())
		}

		policyResponse, err := authzClient.PolicyCreate(ctx, &v1.PolicyCreateRequest{
			Id:        "post-owners",
			Resources: []string{"post.*"},
			Actions:   []string{"edit", "delete"},
			AttributeRules: []string{
				rule.AttributeEqual(
					rule.PrincipalResourceAttribute{
						PrincipalAttribute: "email",
						ResourceAttribute:  "owner_email",
					},
				),
			},
		})
		if err != nil {
			gc = gerror.Code(err)
			if gc.Code() != gconv.Int(codes.AlreadyExists) {
				panic(fmt.Sprintf("cannot create policy: %v", err))
			}
		} else {
			fmt.Printf("Policy created: %s\n", policyResponse.GetPolicy().GetId())
		}
	} else {
		fmt.Printf("Policy already exists: %s\n", policyGetResponse.GetPolicy().GetId())
	}

	// Check if principal is allowed
	for _, identifier := range []string{"123", "456", "789"} {
		isAllowed, err := authzClient.IsAllowed(ctx, &v1.Check{
			Principal:     "user-123",
			ResourceKind:  "post",
			ResourceValue: identifier,
			Action:        "edit",
		})
		if err != nil {
			panic(fmt.Sprintf("cannot check if principal is allowed to edit post: %v", err))
		}

		fmt.Printf("Is allowed to edit post #%s? %v\n", identifier, isAllowed)
	}
}
