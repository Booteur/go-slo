# \BookApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBook**](BookApi.md#CreateBook) | **Post** /v1/book | 
[**DeleteBook**](BookApi.md#DeleteBook) | **Delete** /v1/book/{bookID} | 
[**GetBookById**](BookApi.md#GetBookById) | **Get** /v1/book/{bookID} | 
[**ListBooks**](BookApi.md#ListBooks) | **Get** /v1/book | 
[**UpdateBookById**](BookApi.md#UpdateBookById) | **Put** /v1/book/{bookID} | 


# **CreateBook**
> CreateBookResponse CreateBook(ctx, book)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **book** | [**CreateBookRequest**](CreateBookRequest.md)|  | 

### Return type

[**CreateBookResponse**](CreateBookResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBook**
> DeleteBookResponse DeleteBook(ctx, bookID)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bookID** | **string**|  | 

### Return type

[**DeleteBookResponse**](DeleteBookResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBookById**
> Book GetBookById(ctx, bookID)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bookID** | **string**|  | 

### Return type

[**Book**](Book.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBooks**
> ListBooksResponse ListBooks(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListBooksResponse**](ListBooksResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBookById**
> CreateBookResponse UpdateBookById(ctx, bookID, book)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bookID** | **string**|  | 
  **book** | [**CreateBookRequest**](CreateBookRequest.md)|  | 

### Return type

[**CreateBookResponse**](CreateBookResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

