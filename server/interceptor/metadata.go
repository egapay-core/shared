package interceptor

import (
	"context"
	"github.com/egapay-core/shared/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

var (
	errNoLanguageID  = status.Error(codes.InvalidArgument, "language id is required")
	errNoCountryCode = status.Error(codes.InvalidArgument, "country code is required")
)

func MetadataUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if err := extractRequiredMetadataProps(ctx, info.FullMethod); err != nil {
		return nil, err
	}
	
	return handler(ctx, req)
}

func MetadataStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if err := extractRequiredMetadataProps(ss.Context(), info.FullMethod); err != nil {
		return err
	}
	
	return handler(srv, ss)
}

func extractRequiredMetadataProps(ctx context.Context, currentRoute string) error {
	excludedRoutes := []string{"GetCountries", "ServerReflection", "Check", "Watch"}
	
	var allow bool
	for _, route := range excludedRoutes {
		if strings.Contains(currentRoute, route) {
			allow = true
			break
		}
	}
	
	if allow {
		return nil
	}
	
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		countryCode, passed := md[string(utils.ContextKeyCountryCode)]
		if !passed || len(countryCode) == 0 {
			return errNoCountryCode
		}
		
		if err := utils.ValidateCountryCode(countryCode[0]); err != nil {
			return err
		} else {
			context.WithValue(ctx, utils.ContextKeyCountryCode, countryCode)
		}
		
		languageCode, passed := md[string(utils.ContextKeyLanguageID)]
		if !passed || len(languageCode) == 0 {
			return errNoLanguageID
		}
		
		if err := utils.ValidateLanguageId(languageCode[0]); err != nil {
			return err
		} else {
			context.WithValue(ctx, utils.ContextKeyLanguageID, languageCode)
		}
		
		if userAccessToken, passed := md[string(utils.ContextKeyUserAccessToken)]; passed && len(userAccessToken) > 0 {
			context.WithValue(ctx, utils.ContextKeyUserAccessToken, userAccessToken)
		}
	}
	
	return nil
}
