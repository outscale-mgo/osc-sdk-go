/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// LoadBalancerApiService LoadBalancerApi service
type LoadBalancerApiService service

// CreateLoadBalancerOpts Optional parameters for the method 'CreateLoadBalancer'
type CreateLoadBalancerOpts struct {
    CreateLoadBalancerRequest optional.Interface
}

/*
CreateLoadBalancer Method for CreateLoadBalancer
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CreateLoadBalancerOpts - Optional Parameters:
 * @param "CreateLoadBalancerRequest" (optional.Interface of CreateLoadBalancerRequest) - 
@return CreateLoadBalancerResponse
*/
func (a *LoadBalancerApiService) CreateLoadBalancer(ctx _context.Context, localVarOptionals *CreateLoadBalancerOpts) (CreateLoadBalancerResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateLoadBalancerResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/CreateLoadBalancer"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.CreateLoadBalancerRequest.IsSet() {
		localVarOptionalCreateLoadBalancerRequest, localVarOptionalCreateLoadBalancerRequestok := localVarOptionals.CreateLoadBalancerRequest.Value().(CreateLoadBalancerRequest)
		if !localVarOptionalCreateLoadBalancerRequestok {
			return localVarReturnValue, nil, reportError("createLoadBalancerRequest should be CreateLoadBalancerRequest")
		}
		localVarPostBody = &localVarOptionalCreateLoadBalancerRequest
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["Authorization"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v CreateLoadBalancerResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// DeleteLoadBalancerOpts Optional parameters for the method 'DeleteLoadBalancer'
type DeleteLoadBalancerOpts struct {
    DeleteLoadBalancerRequest optional.Interface
}

/*
DeleteLoadBalancer Method for DeleteLoadBalancer
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *DeleteLoadBalancerOpts - Optional Parameters:
 * @param "DeleteLoadBalancerRequest" (optional.Interface of DeleteLoadBalancerRequest) - 
@return DeleteLoadBalancerResponse
*/
func (a *LoadBalancerApiService) DeleteLoadBalancer(ctx _context.Context, localVarOptionals *DeleteLoadBalancerOpts) (DeleteLoadBalancerResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DeleteLoadBalancerResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/DeleteLoadBalancer"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.DeleteLoadBalancerRequest.IsSet() {
		localVarOptionalDeleteLoadBalancerRequest, localVarOptionalDeleteLoadBalancerRequestok := localVarOptionals.DeleteLoadBalancerRequest.Value().(DeleteLoadBalancerRequest)
		if !localVarOptionalDeleteLoadBalancerRequestok {
			return localVarReturnValue, nil, reportError("deleteLoadBalancerRequest should be DeleteLoadBalancerRequest")
		}
		localVarPostBody = &localVarOptionalDeleteLoadBalancerRequest
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["Authorization"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v DeleteLoadBalancerResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// DeregisterVmsInLoadBalancerOpts Optional parameters for the method 'DeregisterVmsInLoadBalancer'
type DeregisterVmsInLoadBalancerOpts struct {
    DeregisterVmsInLoadBalancerRequest optional.Interface
}

/*
DeregisterVmsInLoadBalancer Method for DeregisterVmsInLoadBalancer
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *DeregisterVmsInLoadBalancerOpts - Optional Parameters:
 * @param "DeregisterVmsInLoadBalancerRequest" (optional.Interface of DeregisterVmsInLoadBalancerRequest) - 
@return DeregisterVmsInLoadBalancerResponse
*/
func (a *LoadBalancerApiService) DeregisterVmsInLoadBalancer(ctx _context.Context, localVarOptionals *DeregisterVmsInLoadBalancerOpts) (DeregisterVmsInLoadBalancerResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  DeregisterVmsInLoadBalancerResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/DeregisterVmsInLoadBalancer"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.DeregisterVmsInLoadBalancerRequest.IsSet() {
		localVarOptionalDeregisterVmsInLoadBalancerRequest, localVarOptionalDeregisterVmsInLoadBalancerRequestok := localVarOptionals.DeregisterVmsInLoadBalancerRequest.Value().(DeregisterVmsInLoadBalancerRequest)
		if !localVarOptionalDeregisterVmsInLoadBalancerRequestok {
			return localVarReturnValue, nil, reportError("deregisterVmsInLoadBalancerRequest should be DeregisterVmsInLoadBalancerRequest")
		}
		localVarPostBody = &localVarOptionalDeregisterVmsInLoadBalancerRequest
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["Authorization"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v DeregisterVmsInLoadBalancerResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// ReadLoadBalancersOpts Optional parameters for the method 'ReadLoadBalancers'
type ReadLoadBalancersOpts struct {
    ReadLoadBalancersRequest optional.Interface
}

/*
ReadLoadBalancers Method for ReadLoadBalancers
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ReadLoadBalancersOpts - Optional Parameters:
 * @param "ReadLoadBalancersRequest" (optional.Interface of ReadLoadBalancersRequest) - 
@return ReadLoadBalancersResponse
*/
func (a *LoadBalancerApiService) ReadLoadBalancers(ctx _context.Context, localVarOptionals *ReadLoadBalancersOpts) (ReadLoadBalancersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ReadLoadBalancersResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/ReadLoadBalancers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.ReadLoadBalancersRequest.IsSet() {
		localVarOptionalReadLoadBalancersRequest, localVarOptionalReadLoadBalancersRequestok := localVarOptionals.ReadLoadBalancersRequest.Value().(ReadLoadBalancersRequest)
		if !localVarOptionalReadLoadBalancersRequestok {
			return localVarReturnValue, nil, reportError("readLoadBalancersRequest should be ReadLoadBalancersRequest")
		}
		localVarPostBody = &localVarOptionalReadLoadBalancersRequest
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["Authorization"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v ReadLoadBalancersResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// RegisterVmsInLoadBalancerOpts Optional parameters for the method 'RegisterVmsInLoadBalancer'
type RegisterVmsInLoadBalancerOpts struct {
    RegisterVmsInLoadBalancerRequest optional.Interface
}

/*
RegisterVmsInLoadBalancer Method for RegisterVmsInLoadBalancer
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *RegisterVmsInLoadBalancerOpts - Optional Parameters:
 * @param "RegisterVmsInLoadBalancerRequest" (optional.Interface of RegisterVmsInLoadBalancerRequest) - 
@return RegisterVmsInLoadBalancerResponse
*/
func (a *LoadBalancerApiService) RegisterVmsInLoadBalancer(ctx _context.Context, localVarOptionals *RegisterVmsInLoadBalancerOpts) (RegisterVmsInLoadBalancerResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RegisterVmsInLoadBalancerResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/RegisterVmsInLoadBalancer"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.RegisterVmsInLoadBalancerRequest.IsSet() {
		localVarOptionalRegisterVmsInLoadBalancerRequest, localVarOptionalRegisterVmsInLoadBalancerRequestok := localVarOptionals.RegisterVmsInLoadBalancerRequest.Value().(RegisterVmsInLoadBalancerRequest)
		if !localVarOptionalRegisterVmsInLoadBalancerRequestok {
			return localVarReturnValue, nil, reportError("registerVmsInLoadBalancerRequest should be RegisterVmsInLoadBalancerRequest")
		}
		localVarPostBody = &localVarOptionalRegisterVmsInLoadBalancerRequest
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["Authorization"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v RegisterVmsInLoadBalancerResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// UpdateLoadBalancerOpts Optional parameters for the method 'UpdateLoadBalancer'
type UpdateLoadBalancerOpts struct {
    UpdateLoadBalancerRequest optional.Interface
}

/*
UpdateLoadBalancer Method for UpdateLoadBalancer
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *UpdateLoadBalancerOpts - Optional Parameters:
 * @param "UpdateLoadBalancerRequest" (optional.Interface of UpdateLoadBalancerRequest) - 
@return UpdateLoadBalancerResponse
*/
func (a *LoadBalancerApiService) UpdateLoadBalancer(ctx _context.Context, localVarOptionals *UpdateLoadBalancerOpts) (UpdateLoadBalancerResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UpdateLoadBalancerResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/UpdateLoadBalancer"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.UpdateLoadBalancerRequest.IsSet() {
		localVarOptionalUpdateLoadBalancerRequest, localVarOptionalUpdateLoadBalancerRequestok := localVarOptionals.UpdateLoadBalancerRequest.Value().(UpdateLoadBalancerRequest)
		if !localVarOptionalUpdateLoadBalancerRequestok {
			return localVarReturnValue, nil, reportError("updateLoadBalancerRequest should be UpdateLoadBalancerRequest")
		}
		localVarPostBody = &localVarOptionalUpdateLoadBalancerRequest
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if auth, ok := auth["Authorization"]; ok {
				var key string
				if auth.Prefix != "" {
					key = auth.Prefix + " " + auth.Key
				} else {
					key = auth.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v UpdateLoadBalancerResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
