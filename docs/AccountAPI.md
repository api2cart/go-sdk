# \AccountAPI

All URIs are relative to *https://api.api2cart.local.com/v1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountCartAdd**](AccountAPI.md#AccountCartAdd) | **Post** /account.cart.add.json | account.cart.add
[**AccountCartList**](AccountAPI.md#AccountCartList) | **Get** /account.cart.list.json | account.cart.list
[**AccountConfigUpdate**](AccountAPI.md#AccountConfigUpdate) | **Put** /account.config.update.json | account.config.update
[**AccountFailedWebhooks**](AccountAPI.md#AccountFailedWebhooks) | **Get** /account.failed_webhooks.json | account.failed_webhooks
[**AccountSupportedPlatforms**](AccountAPI.md#AccountSupportedPlatforms) | **Get** /account.supported_platforms.json | account.supported_platforms



## AccountCartAdd

> AccountCartAdd200Response AccountCartAdd(ctx).AccountCartAdd(accountCartAdd).Execute()

account.cart.add



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountCartAdd := *openapiclient.NewAccountCartAdd("Opencart14") // AccountCartAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AccountCartAdd(context.Background()).AccountCartAdd(accountCartAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountCartAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountCartAdd`: AccountCartAdd200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountCartAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountCartAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountCartAdd** | [**AccountCartAdd**](AccountCartAdd.md) |  | 

### Return type

[**AccountCartAdd200Response**](AccountCartAdd200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountCartList

> AccountCartList200Response AccountCartList(ctx).StoreUrl(storeUrl).StoreKey(storeKey).RequestFromDate(requestFromDate).RequestToDate(requestToDate).CustomLabel(customLabel).Params(params).Exclude(exclude).Execute()

account.cart.list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	storeUrl := "http://mystore.com" // string | A web address of a store (optional)
	storeKey := "ab37fc230bc5df63a5be1b11220949be" // string | Find store by store key (optional)
	requestFromDate := "2010-07-29" // string | Retrieve entities from their creation date (optional)
	requestToDate := "2100-08-29" // string | Retrieve entities to their creation date (optional)
	customLabel := "This is test store" // string | Defines a custom label for the store in the app (optional)
	params := "url,store_key" // string | Set this parameter in order to choose which entity fields you want to retrieve (optional) (default to "force_all")
	exclude := "url,store_key" // string | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AccountCartList(context.Background()).StoreUrl(storeUrl).StoreKey(storeKey).RequestFromDate(requestFromDate).RequestToDate(requestToDate).CustomLabel(customLabel).Params(params).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountCartList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountCartList`: AccountCartList200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountCartList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountCartListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeUrl** | **string** | A web address of a store | 
 **storeKey** | **string** | Find store by store key | 
 **requestFromDate** | **string** | Retrieve entities from their creation date | 
 **requestToDate** | **string** | Retrieve entities to their creation date | 
 **customLabel** | **string** | Defines a custom label for the store in the app | 
 **params** | **string** | Set this parameter in order to choose which entity fields you want to retrieve | [default to &quot;force_all&quot;]
 **exclude** | **string** | Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter &#x60;params&#x60; equal force_all | 

### Return type

[**AccountCartList200Response**](AccountCartList200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountConfigUpdate

> AccountConfigUpdate200Response AccountConfigUpdate(ctx).ReplaceParameters(replaceParameters).NewStoreUrl(newStoreUrl).NewStoreKey(newStoreKey).CustomLabel(customLabel).BridgeUrl(bridgeUrl).StoreRoot(storeRoot).DbTablesPrefix(dbTablesPrefix).UserAgent(userAgent).Var3dcartPrivateKey(var3dcartPrivateKey).Var3dcartAccessToken(var3dcartAccessToken).Var3dcartapiApiKey(var3dcartapiApiKey).AmazonSpClientId(amazonSpClientId).AmazonSpClientSecret(amazonSpClientSecret).AmazonSpRefreshToken(amazonSpRefreshToken).AmazonSpAwsRegion(amazonSpAwsRegion).AmazonSpApiEnvironment(amazonSpApiEnvironment).AmazonSellerId(amazonSellerId).AspdotnetstorefrontApiUser(aspdotnetstorefrontApiUser).AspdotnetstorefrontApiPass(aspdotnetstorefrontApiPass).AmericommerceAppId(americommerceAppId).AmericommerceAppSecret(americommerceAppSecret).AmericommerceAccessToken(americommerceAccessToken).AmericommerceRefreshToken(americommerceRefreshToken).BigcommerceapiAdminAccount(bigcommerceapiAdminAccount).BigcommerceapiApiPath(bigcommerceapiApiPath).BigcommerceapiApiKey(bigcommerceapiApiKey).BigcommerceapiClientId(bigcommerceapiClientId).BigcommerceapiAccessToken(bigcommerceapiAccessToken).BigcommerceapiContext(bigcommerceapiContext).BolApiKey(bolApiKey).BolApiSecret(bolApiSecret).BolRetailerId(bolRetailerId).DemandwareClientId(demandwareClientId).DemandwareApiPassword(demandwareApiPassword).DemandwareUserName(demandwareUserName).DemandwareUserPassword(demandwareUserPassword).EbayClientId(ebayClientId).EbayClientSecret(ebayClientSecret).EbayRuname(ebayRuname).EbayAccessToken(ebayAccessToken).EbayRefreshToken(ebayRefreshToken).EbayEnvironment(ebayEnvironment).EbaySiteId(ebaySiteId).EcwidAcessToken(ecwidAcessToken).EcwidStoreId(ecwidStoreId).LazadaAppId(lazadaAppId).LazadaAppSecret(lazadaAppSecret).LazadaRefreshToken(lazadaRefreshToken).LazadaRegion(lazadaRegion).EtsyKeystring(etsyKeystring).EtsySharedSecret(etsySharedSecret).EtsyAccessToken(etsyAccessToken).EtsyTokenSecret(etsyTokenSecret).EtsyClientId(etsyClientId).EtsyRefreshToken(etsyRefreshToken).FacebookAppId(facebookAppId).FacebookAppSecret(facebookAppSecret).FacebookAccessToken(facebookAccessToken).FacebookBusinessId(facebookBusinessId).NetoApiKey(netoApiKey).NetoApiUsername(netoApiUsername).ShoplineAccessToken(shoplineAccessToken).ShoplineAppKey(shoplineAppKey).ShoplineAppSecret(shoplineAppSecret).ShoplineSharedSecret(shoplineSharedSecret).ShopifyAccessToken(shopifyAccessToken).ShopifyClientId(shopifyClientId).ShopifyApiKey(shopifyApiKey).ShopifyApiPassword(shopifyApiPassword).ShopifySharedSecret(shopifySharedSecret).ShopeePartnerId(shopeePartnerId).ShopeePartnerKey(shopeePartnerKey).ShopeeShopId(shopeeShopId).ShopeeRefreshToken(shopeeRefreshToken).ShopeeRegion(shopeeRegion).ShopeeEnvironment(shopeeEnvironment).ShoplazzaAccessToken(shoplazzaAccessToken).ShoplazzaSharedSecret(shoplazzaSharedSecret).MivaAccessToken(mivaAccessToken).MivaSignature(mivaSignature).ShopwareAccessKey(shopwareAccessKey).UnasApiKey(unasApiKey).ShopwareApiKey(shopwareApiKey).ShopwareApiSecret(shopwareApiSecret).BigcartelUserName(bigcartelUserName).BigcartelPassword(bigcartelPassword).BricklinkConsumerKey(bricklinkConsumerKey).BricklinkConsumerSecret(bricklinkConsumerSecret).BricklinkToken(bricklinkToken).BricklinkTokenSecret(bricklinkTokenSecret).VolusionLogin(volusionLogin).VolusionPassword(volusionPassword).WalmartClientId(walmartClientId).WalmartClientSecret(walmartClientSecret).WalmartEnvironment(walmartEnvironment).WalmartChannelType(walmartChannelType).WalmartRegion(walmartRegion).SquareClientId(squareClientId).SquareClientSecret(squareClientSecret).SquareRefreshToken(squareRefreshToken).SquarespaceApiKey(squarespaceApiKey).SquarespaceClientId(squarespaceClientId).SquarespaceClientSecret(squarespaceClientSecret).SquarespaceAccessToken(squarespaceAccessToken).SquarespaceRefreshToken(squarespaceRefreshToken).HybrisClientId(hybrisClientId).HybrisClientSecret(hybrisClientSecret).HybrisUsername(hybrisUsername).HybrisPassword(hybrisPassword).HybrisWebsites(hybrisWebsites).LightspeedApiKey(lightspeedApiKey).LightspeedApiSecret(lightspeedApiSecret).CommercehqApiKey(commercehqApiKey).CommercehqApiPassword(commercehqApiPassword).WcConsumerKey(wcConsumerKey).WcConsumerSecret(wcConsumerSecret).MagentoConsumerKey(magentoConsumerKey).MagentoConsumerSecret(magentoConsumerSecret).MagentoAccessToken(magentoAccessToken).MagentoTokenSecret(magentoTokenSecret).PrestashopWebserviceKey(prestashopWebserviceKey).WixAppId(wixAppId).WixAppSecretKey(wixAppSecretKey).WixInstanceId(wixInstanceId).WixRefreshToken(wixRefreshToken).MercadoLibreAppId(mercadoLibreAppId).MercadoLibreAppSecretKey(mercadoLibreAppSecretKey).MercadoLibreRefreshToken(mercadoLibreRefreshToken).ZidClientId(zidClientId).ZidClientSecret(zidClientSecret).ZidAccessToken(zidAccessToken).ZidAuthorization(zidAuthorization).ZidRefreshToken(zidRefreshToken).JumpsellerClientId(jumpsellerClientId).JumpsellerClientSecret(jumpsellerClientSecret).JumpsellerRefreshToken(jumpsellerRefreshToken).JumpsellerLogin(jumpsellerLogin).JumpsellerAuthtoken(jumpsellerAuthtoken).FlipkartClientId(flipkartClientId).FlipkartClientSecret(flipkartClientSecret).AllegroClientId(allegroClientId).AllegroClientSecret(allegroClientSecret).AllegroAccessToken(allegroAccessToken).AllegroRefreshToken(allegroRefreshToken).AllegroEnvironment(allegroEnvironment).ZohoClientId(zohoClientId).ZohoClientSecret(zohoClientSecret).ZohoRefreshToken(zohoRefreshToken).ZohoRegion(zohoRegion).TiendanubeUserId(tiendanubeUserId).TiendanubeAccessToken(tiendanubeAccessToken).TiendanubeClientSecret(tiendanubeClientSecret).OttoClientId(ottoClientId).OttoClientSecret(ottoClientSecret).OttoAppId(ottoAppId).OttoRefreshToken(ottoRefreshToken).OttoEnvironment(ottoEnvironment).OttoAccessToken(ottoAccessToken).TiktokshopAppKey(tiktokshopAppKey).TiktokshopAppSecret(tiktokshopAppSecret).TiktokshopRefreshToken(tiktokshopRefreshToken).TiktokshopAccessToken(tiktokshopAccessToken).SallaClientId(sallaClientId).SallaClientSecret(sallaClientSecret).SallaRefreshToken(sallaRefreshToken).SallaAccessToken(sallaAccessToken).TemuAppKey(temuAppKey).TemuAppSecret(temuAppSecret).TemuAccessToken(temuAccessToken).TemuRegion(temuRegion).ScapiClientId(scapiClientId).ScapiClientSecret(scapiClientSecret).ScapiOrganizationId(scapiOrganizationId).ScapiShortCode(scapiShortCode).ScapiScopes(scapiScopes).IdempotencyKey(idempotencyKey).Execute()

account.config.update



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	replaceParameters := true // bool | Identifies if there is a necessity to replace parameters (optional)
	newStoreUrl := "http://mystore.com" // string | The web address of the store you want to update to connect to API2Cart (optional)
	newStoreKey := "b636495648de3086f6f57b1bd4be548f" // string | Update store key (optional)
	customLabel := "This is test store" // string | Defines a custom label for the store in the app (optional)
	bridgeUrl := "https://your-store.com/custom/bridge/path/bridge.php" // string | This parameter allows to set up store with custom bridge url (also you must use store_root parameter if a bridge folder is not in the root folder of the store) (optional)
	storeRoot := "/home/www/stores/magento1922" // string | Absolute path to the store root directory (used with \"bridge_url\" parameter) (optional)
	dbTablesPrefix := "oc_" // string | DB tables prefix (optional)
	userAgent := "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0" // string | This parameter allows you to set your custom user agent, which will be used in requests to the store. Please use it cautiously, as the store's firewall may block specific values. (optional)
	var3dcartPrivateKey := "7dba81f90bdbe25e7000e73214ca51b" // string | 3DCart Private Key (optional)
	var3dcartAccessToken := "4Grr_ZCLNNoSUuhAjesKuchxo9SL" // string | 3DCart Token (optional)
	var3dcartapiApiKey := "82cc921c6a5c67082cc921c6a5c6707e1d6e6862ba3201a" // string | 3DCart API Key (optional)
	amazonSpClientId := "amzn1.application-oa2-client.11e000e1f47d4998aca3733716d3b5a4" // string | Amazon SP API app client id (optional)
	amazonSpClientSecret := "2c987428209f235443221255bde064f4bdf8a65165a80f5d22760a83cb" // string | Amazon SP API app client secret (optional)
	amazonSpRefreshToken := "Atzr|IwEBIPUI-bwRTdDgKNQ_g56C30wGqymtx30c9MdDC7Emwmojhs20k5BBG2hHtJiGZ_7OfG7khd1RuQr6KEst4qyWbo_eXi5S_T_VOxzJUuksG1cFOGFpFK-cnhReNzAeZIpZeJT7_ROy1csEFlQfC8FJS3bsbSkkbTz2ZcTN7_7ey0HVlhyfFizgROeSeOI24Wjs9l_KKzZW0jvi_oC2cxlIcyknnHLK6KMNz2rTXqQJWRtlK9xPJDdbcUa5STA8MQru91cxNBpSkZN_cq9OOELhbsIGKD75y7nZ3yJU4uHQC_9iBQQoFm0biKgi-kEQwOhwws8" // string | Amazon SP API OAuth refresh token (optional)
	amazonSpAwsRegion := "us-east-1" // string | Amazon AWS Region (optional)
	amazonSpApiEnvironment := "sandbox" // string | Amazon SP API environment (optional) (default to "production")
	amazonSellerId := "13P636B2M1N4WR" // string | Amazon Seller ID (Merchant token) (optional)
	aspdotnetstorefrontApiUser := "admin" // string | It's a AspDotNetStorefront account for which API is available (optional)
	aspdotnetstorefrontApiPass := "f6471ef78f72b41849a8b8b67791b0b5" // string | AspDotNetStorefront API Password (optional)
	americommerceAppId := "1" // string | Americommerce App ID (optional)
	americommerceAppSecret := "9fd3d282d65a007a2b9d541b5e0e410b2cecd6199632db53503b93637b8a6000" // string | Americommerce App Secret (optional)
	americommerceAccessToken := "1e721f59b610e2666caea03094600765" // string | Americommerce Access Token (optional)
	americommerceRefreshToken := "520c011444af41d916543cdda859a5114" // string | Americommerce Refresh Token (optional)
	bigcommerceapiAdminAccount := "admin" // string | It's a BigCommerce account for which API is enabled (optional)
	bigcommerceapiApiPath := "http://mystore.bigcommerce.com/api/v1" // string | BigCommerce API URL (optional)
	bigcommerceapiApiKey := "6b89704cd75738cb0f9f6468d5462aba" // string | Bigcommerce API Key (optional)
	bigcommerceapiClientId := "p1r37bt131z86675nofv9xmhietoe4t" // string | Client ID of the requesting app (optional)
	bigcommerceapiAccessToken := "igse8e4rdmzkxdi937qe69d59en1imw" // string | Access token authorizing the app to access resources on behalf of a user (optional)
	bigcommerceapiContext := "stores/etplnf8o8v" // string | API Path section unique to the store (optional)
	bolApiKey := "51369628-feee-11ed-be56-0242ac120002" // string | Bol API Key (optional)
	bolApiSecret := "8fGzEsbEP5z2MNZubmIil87m-sWzTkj?KDQKrmzmU!fA6aAUNMdKRp7LMWHwE!G37UMfnWByHBGSXJHkAG?QcuYTO2uklv4idIHwUMLHK!OO1yfRlWh!" // string | Bol API Secret (optional)
	bolRetailerId := int32(145001) // int32 | Bol Retailer ID (optional)
	demandwareClientId := "b849eb85-v8b9-1dw8-9fe2-97e1d6ffc7b0" // string | Demandware client id (optional)
	demandwareApiPassword := "testpassword" // string | Demandware api password (optional)
	demandwareUserName := "admin" // string | Demandware user name (optional)
	demandwareUserPassword := "12345" // string | Demandware user password (optional)
	ebayClientId := "a9psel85v1wy5faeyjw03y0r" // string | Application ID (AppID). (optional)
	ebayClientSecret := "gmz3iz45x2" // string | Shared Secret from eBay application (optional)
	ebayRuname := "gmz3iz45x2" // string | The RuName value that eBay assigns to your application. (optional)
	ebayAccessToken := "v^1.1#i ... AjRV4yNjA=" // string | Used to authenticate API requests. (optional)
	ebayRefreshToken := "v^1.1#i ... rAewqVasdA=" // string | Used to renew the access token. (optional)
	ebayEnvironment := "sandbox" // string | eBay environment (optional)
	ebaySiteId := int32(101) // int32 | eBay global ID (optional) (default to 0)
	ecwidAcessToken := "igse8e4rdmzkxdi937qe69d59en1imw" // string | Access token authorizing the app to access resources on behalf of a user (optional)
	ecwidStoreId := "1" // string | Store Id (optional)
	lazadaAppId := "112577" // string | Lazada App ID (optional)
	lazadaAppSecret := "er33raICJ79Q5b0EsR9stmRnjE9XQ2WH" // string | Lazada App Secret (optional)
	lazadaRefreshToken := "EAAPP06rM2n8BO4mZBuMPnu9zS0MaMbN7ue8aUkcxw4zewU337mVVb5br" // string | Lazada Refresh Token (optional)
	lazadaRegion := "Malaysia" // string | Lazada API endpoint Region (optional)
	etsyKeystring := "a9psel85v1wy5faeyjw03y0r" // string | Etsy keystring (optional)
	etsySharedSecret := "gmz3iz45x2" // string | Etsy shared secret (optional)
	etsyAccessToken := "igse8e4rdmzkxdi937qe69d59en1imw" // string | Access token authorizing the app to access resources on behalf of a user (optional)
	etsyTokenSecret := "igse8e4rdmzkxdi937qe69d59en1imw" // string | Secret token authorizing the app to access resources on behalf of a user (optional)
	etsyClientId := "w0fi0igk2w29bjcd7ydr2s35" // string | Etsy Client Id (optional)
	etsyRefreshToken := "223577551.L07_RE-y7unmKf2dox4djsHkVxwpUfs1ikG_uQmHhF-aASEReNn_Qns1Wqn3dDa0ZMxrt9CIael3dgudeDZb31ZUdS" // string | Etsy Refresh token (optional)
	facebookAppId := "6516912365277570" // string | Facebook App ID (optional)
	facebookAppSecret := "737cf6bd2879cb6c7e5a8ff9cd63f3d46b0b5b7b" // string | Facebook App Secret (optional)
	facebookAccessToken := "EAAPP06rM2n8BO4mZBuMPnu9zS0MaMbN7ue8aUAhqbS58clzJwyp1rYRMpP31QJGziqtYbKypdVx3Cs0RpuufoUeLsbfX195XIB8VTlkcxw4zewU337mVVb5br" // string | Facebook Access Token (optional)
	facebookBusinessId := "294042786906655" // string | Facebook Business ID (optional)
	netoApiKey := "bbca57d8ff3c3677128112c15556d9e3" // string | Neto API Key (optional)
	netoApiUsername := "mylogin" // string | Neto User Name (optional)
	shoplineAccessToken := "igse8e4rdmzkxdi937qe69d59en1imw" // string | Shopline APP Key (optional)
	shoplineAppKey := "737cf6bd2879cb6c7e5a8ff9cd63f3d46b0b5b7b" // string | Shopline APP Key (optional)
	shoplineAppSecret := "1701d123bb5cc14cd2732dcaed90638316c0a09" // string | Shopline App Secret (optional)
	shoplineSharedSecret := "1701d123bb5cc14cd2732dcaed90638316c0a09" // string | Shopline Shared Secret (optional)
	shopifyAccessToken := "igse8e4rdmzkxdi937qe69d59en1imw" // string | Access token authorizing the app to access resources on behalf of a user (optional)
	shopifyClientId := "b5defe55db3f6836fb4e0e6624ff9577" // string | Shopify Client ID (optional)
	shopifyApiKey := "bbca57d8ff3c3677128112c15556d9e3" // string | Shopify API Key (optional)
	shopifyApiPassword := "860f3a6fc87632301a42cd88e4b5ab3d" // string | Shopify API Password (optional)
	shopifySharedSecret := "gmz3iz45x2" // string | Shared secret (optional)
	shopeePartnerId := "1276777" // string | Shopee Partner ID (optional)
	shopeePartnerKey := "6a46494b4d746576554646626775617a577542774850636375464d6a736d5598" // string | Shopee Partner Key (optional)
	shopeeShopId := "137968" // string | Shopee SHOP ID (optional)
	shopeeRefreshToken := "EAAPP06rM2n8BO4mZBuMPnu9zS0MaMbN7ue8aUkcxw4zewU987mVVb5br" // string | Shopee Refresh Token (optional)
	shopeeRegion := "CN" // string | Shopee API endpoint Region. Use for Chinese Mainland or Brazil. (optional)
	shopeeEnvironment := "sandbox" // string | Shopee Environment (optional)
	shoplazzaAccessToken := "igse8e4rdmzkxdi937qe69d59en1imw" // string | Access token authorizing the app to access resources on behalf of a user (optional)
	shoplazzaSharedSecret := "gmz3iz45x2" // string | Shared secret (optional)
	mivaAccessToken := "227cbe434a1e358d72db0de993x9d9fd" // string | Miva access token (optional)
	mivaSignature := "1hpkrebfdsObGTor/0Gk9XcNBUQohrxrw67Sg9AM9ps=" // string | Miva signature (optional)
	shopwareAccessKey := "SWSCS3O1RJBSRNBYQLFIYJN2ZQ" // string | Shopware access key (optional)
	unasApiKey := "c238908e29ceb6e6ad3df15f89a6234709d3f000" // string | UNAS API Key (optional)
	shopwareApiKey := "SWSCS3O1RJBSRNBYQLFIYJN2ZQ" // string | Shopware api key (optional)
	shopwareApiSecret := "V3NYNWg2b1dZdHBUWDN1cmdKdGhnenp5enVJYlJ0WlJvOFF2bnQ" // string | Shopware client secret access key (optional)
	bigcartelUserName := "subdomain" // string | Subdomain of store (optional)
	bigcartelPassword := "4GrrZCLNNoSUuhAjesKuchxo9SL" // string | BigCartel account password (optional)
	bricklinkConsumerKey := "26F6CDA087D9444EAA71AC09E7A1D39A" // string | Bricklink Consumer Key (optional)
	bricklinkConsumerSecret := "a46abc3kxyinlbggy06i9g975xqo6gjq" // string | Bricklink Consumer Secret (optional)
	bricklinkToken := "ktv4n9rgrj0evjuy2t6p2xlb1f8u5pmy" // string | Bricklink Access Token (optional)
	bricklinkTokenSecret := "a46abc3kxyinlbggy06i9g975xqo6gjq" // string | Bricklink Access Token Secret (optional)
	volusionLogin := "admin" // string | It's a Volusion account for which API is enabled (optional)
	volusionPassword := "7943CA5F3990E00D9A4CCF0BD998211F" // string | Volusion API Password (optional)
	walmartClientId := "423f6A24-123z-8654-989u-6fa96478289" // string | Walmart client ID. For the region 'ca' use Consumer ID (optional)
	walmartClientSecret := "1gf85fea-8974-2648-w12w-rt54284tdf54" // string | Walmart client secret. For the region 'ca' use Private Key (optional)
	walmartEnvironment := "production" // string | Walmart environment (optional) (default to "production")
	walmartChannelType := "0f3e4dd4-0514-4346-b39d-af0e00ea066d" // string | Walmart WM_CONSUMER.CHANNEL.TYPE header (optional)
	walmartRegion := "us" // string | Walmart region (optional) (default to "us")
	squareClientId := "sq0idp-qwer_1pvuTYe9cAf1lmxyQ" // string | Square (Weebly) Client ID (optional)
	squareClientSecret := "c8d7077fce7b2b111111111898170695a01473a2ad" // string | Square (Weebly) Client Secret (optional)
	squareRefreshToken := "EQAAlquVXMr6xIcPu7qPkIEAZ0thqChhQuowrvZIqOlwhOwhtmyh4ZRfesdRc434" // string | Square (Weebly) Refresh Token (optional)
	squarespaceApiKey := "8f7849d5-1411-47f2-9722-aa81c2a48d95" // string | Squarespace API Key (optional)
	squarespaceClientId := "9UGbUtS2V96BxRGmfOjsGAhTdsr9Vxxx" // string | Squarespace Connector Client ID (optional)
	squarespaceClientSecret := "GPZkUFkIKWg0KLE6rajsFMMYA9ma0udaaq2bYwBDXXX=" // string | Squarespace Connector Client Secret (optional)
	squarespaceAccessToken := "SWSCS3O1RJBSRNBYQLFIYJN2ZQ" // string | Squarespace access token (optional)
	squarespaceRefreshToken := "def50eyfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe657e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65de64a0c865d" // string | Squarespace refresh token (optional)
	hybrisClientId := "api_client_1" // string | Omni Commerce Connector Client ID (optional)
	hybrisClientSecret := "secret_phrase_1" // string | Omni Commerce Connector Client Secret (optional)
	hybrisUsername := "admin" // string | User Name (optional)
	hybrisPassword := "nimda" // string | User password (optional)
	hybrisWebsites := []string{"Inner_example"} // []string | Websites to stores mapping data (optional)
	lightspeedApiKey := "cf5444729c2abd6b6a5d983691767cb5" // string | LightSpeed api key (optional)
	lightspeedApiSecret := "2620ee52a8bc942f9d5d3a575f4d363e" // string | LightSpeed api secret (optional)
	commercehqApiKey := "sJrD-LM0eddhe63rfgfva0dDydXfre4" // string | CommerceHQ api key (optional)
	commercehqApiPassword := "4Grr_ZCLNNoSUuhAjesKuchxo9SL" // string | CommerceHQ api password (optional)
	wcConsumerKey := "ck_26d8e2ad604f3917e429df6961722282bdcf109d" // string | Woocommerce consumer key (optional)
	wcConsumerSecret := "cs_931ced666118a15c5f7b4a33a15gf5589cbeba55" // string | Woocommerce consumer secret (optional)
	magentoConsumerKey := "ktv4n9rgrj0evjuy2t6p2xlb1f8u5pmy" // string | Magento Consumer Key (optional)
	magentoConsumerSecret := "a46abc3kxyinlbggy06i9g975xqo6gjq" // string | Magento Consumer Secret (optional)
	magentoAccessToken := "igse8e4rdmzkxdi937qe69d59en1imw" // string | Magento Access Token (optional)
	magentoTokenSecret := "igse8e4rdmzkxdi937qe69d59en1imw" // string | Magento Token Secret (optional)
	prestashopWebserviceKey := "CKJ1ZEWRJWRLTPVBQJ9FGGRORD4AGS96" // string | Prestashop webservice key (optional)
	wixAppId := "6b0b5b7b-7d87-45b5-bf34-ac6b438e63da" // string | Wix App ID (optional)
	wixAppSecretKey := "316c0a09-f195-42be-74f6-a02cebb9cae6" // string | Wix App Secret Key (optional)
	wixInstanceId := "58b893a4-6b16-5c2f-qt78-qa3r61t32rt8" // string | Wix Instance ID (optional)
	wixRefreshToken := "
        OAUTH2.eyJraWQiOiJkZ0x3cjNRMCIsImFsZyI6IkhTMjU2In0.
        eyJkYXRhIjoie1wiaWRcIjpcImJlZjM3MmRmLTUyNGItNDI3NS05M2RkL
        Tg4NDBlOTU3ZWU2OFwifSIsImlhdCI6MTY0ODA0NTEyNiwiZXhwIjoxNzExMTE3MTI2fQ.
        VRR2lGSbcTVmaArtmyyhy6o4WRDwTn-nlDCQpZ97eYw
      " // string | Wix refresh token (optional)
	mercadoLibreAppId := "211188015100135" // string | Mercado Libre App ID (optional)
	mercadoLibreAppSecretKey := "e2qoG2zklLlfP7cEngEJ94YjhkejkjAm" // string | Mercado Libre App Secret Key (optional)
	mercadoLibreRefreshToken := "TG-63h13529vb5464110188d2x9-703754376" // string | Mercado Libre Refresh Token (optional)
	zidClientId := int32(1234) // int32 | Zid Client ID (optional)
	zidClientSecret := "nl5l1lE0vxgv6cV111fHsdlOOIfb0Ms5IR7l4Igs" // string | Zid Client Secret (optional)
	zidAccessToken := "def50eyfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe657e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65de64a0c865d" // string | Zid Access Token (optional)
	zidAuthorization := "def50eyfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe657e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65de64a0c865d" // string | Zid Authorization (optional)
	zidRefreshToken := "def50eyfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe657e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65d7e64a0cfe9fe65de64a0c865d" // string | Zid refresh token (optional)
	jumpsellerClientId := "your_client_id" // string | Jumpseller OAuth2 Client ID (optional)
	jumpsellerClientSecret := "your_client_secret" // string | Jumpseller OAuth2 Client Secret (optional)
	jumpsellerRefreshToken := "your_oauth_refresh_token" // string | Jumpseller OAuth2 refresh token (optional)
	jumpsellerLogin := "your_login" // string | Jumpseller API login (optional)
	jumpsellerAuthtoken := "your_auth_token" // string | Jumpseller API auth token (optional)
	flipkartClientId := "19414773883a13a850b6a52350b7246499a24" // string | Flipkart Client ID (optional)
	flipkartClientSecret := "nl5l1lE0vxgv6cV111fHsdlOOIfb0Ms5IR7l4Igs" // string | Flipkart Client Secret (optional)
	allegroClientId := "2915e189ce3d23d23d2327d204ae6a0bd" // string | Allegro Client ID (optional)
	allegroClientSecret := "DNHtqdL2WPIefeUhQWYgtXPS23fgbfgasdsGHHJGhg3RTFDQWFGZmVoFRT5IfkQj1E7eR5" // string | Allegro Client Secret (optional)
	allegroAccessToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IsfddfdfdeyJ1c2VyX25hbWUiOiI5NDUxMzE1MSIsInNjb3BlIjpbImFsbGVncm86YXBpOm9yZGVyczpyZWFkIiwiYWxsZWdybzphcGk6cHJvZmlsZTp3cml0ZSIsImFsbGVncm86YXBpOnNhbGU6b2ZmZXJzOndyaXRlIiwiYWxsZWdybzphcGk6YmlsbGluZzpyZWFkIiwiYWxsZWdybzphcGk6Y2FtcGFpZ25zIiwiYWxsZWdybzphcGk6ZGlzcHV0ZXMiLCJhbGxlZ3JvOmFwaTpzYWxlOm9mZmVyczpyZWFkIiwiYWxsZWdybzphcGk6YmlkcyIsImFsbGVncm86YXBpOm9yZGVyczp3cml0ZSIsImFsbGVncm86YXBpOnBheW1lbnRzOndyaXRlIiwiYWxsZWdybzphcGk6c2FsZTpzZXR0aW5nczp3cml0ZSIsImFsbGVncm86YXBpOnByb2ZpbGU6cmVhZCIsImFsbGVncm86YXBpOnJhdGluZ3MiLCJhbGxlZ3JvOmFwaTpzYWxlOnNldHRpbmdzOnJlYWQiLCJhbGxlZ3JvOmFwaTpwYXltZW50czpyZWFkIiwiYWxsZWdybzphcGk6bWVzc2FnaW5nIl0sI" // string | Allegro Access Token (optional)
	allegroRefreshToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IsfddfdfdeyJ1c2VyX25hbWUiOiI5NDUxMzE1MSIsInNjb3BlIjpbImFsbGVncm86YXBpOm9yZGVyczpyZWFkIiwiYWxsZWdybzphcGk6cHJvZmlsZTp3cml0ZSIsImFsbGVncm86YXBpOnNhbGU6b2ZmZXJzOndyaXRlIiwiYWxsZWdybzphcGk6YmlsbGluZzpyZWFkIiwiYWxsZWdybzphcGk6Y2FtcGFpZ25zIiwiYWxsZWdybzphcGk6ZGlzcHV0ZXMiLCJhbGxlZ3JvOmFwaTpzYWxlOm9mZmVyczpyZWFkIiwiYWxsZWdybzphcGk6YmlkcyIsImFsbGVncm86YXBpOm9yZGVyczp3cml0ZSIsImFsbGVncm86YXBpOnBheW1lbnRzOndyaXRlIiwiYWxsZWdybzphcGk6c2FsZTpzZXR0aW5nczp3cml0ZSIsImFsbGVncm86YXBpOnByb2ZpbGU6cmVhZCIsImFsbGVncm86YXBpOnJhdGluZ3MiLCJhbGxlZ3JvOmFwaTpzYWxlOnNldHRpbmdzOnJlYWQiLCJhbGxlZ3JvOmFwaTpwYXltZW50czpyZWFkIiwiYWxsZWdybzphcGk6bWVzc2FnaW5nIl0sI" // string | Allegro Refresh Token (optional)
	allegroEnvironment := "sandbox" // string | Allegro Environment (optional) (default to "production")
	zohoClientId := "1000.FLCHGI2LS1111111TOR4OGB697W4IX" // string | Zoho Client ID (optional)
	zohoClientSecret := "c8d7077fce7b2b111111111898170695a01473a2ad" // string | Zoho Client Secret (optional)
	zohoRefreshToken := "1000.11111111111111111111111111111111.1b3ca6f054341a111118abf928beb33b" // string | Zoho Refresh Token (optional)
	zohoRegion := "Europe" // string | Zoho API endpoint Region (optional)
	tiendanubeUserId := int32(1234) // int32 | Tiendanube User ID (optional)
	tiendanubeAccessToken := "75bde7bb0b437475423e7e87c142c06052f80199" // string | Tiendanube Access Token (optional)
	tiendanubeClientSecret := "5e3588f514a5ae0d0fa063d1b556531e25c83fa7e47472ed" // string | Tiendanube Client Secret (optional)
	ottoClientId := "911a3dbf-d261-4763-cc81-052876465b55" // string | Otto Client ID (optional)
	ottoClientSecret := "9887a82a-2879-421e-a6wc-54e986b3458c" // string | Otto Client Secret (optional)
	ottoAppId := "6eaef6a3-822e-425b-8mc9-53750063e34d" // string | Otto App ID (optional)
	ottoRefreshToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IsfddfdfdeyJ1c2VyX25hbWUiOiI5NDUxMzE1MSIsInNjb3BlIjpbImFsbGVncm86YXBpOm9yZGVyczpyZWFkIiwiYWxsZWdybzphcGk6cHJvZmlsZTp3cml0ZSIsImFsbGVncm86YXBpOnNhbGU6b2ZmZXJzOndyaXRlIiwiYWxsZWdybzphcGk6YmlsbGluZzpyZWFkIiwiYWxsZWdybzphcGk6Y2FtcGFpZ25zIiwiYWxsZWdybzphcGk6ZGlzcHV0ZXMiLCJhbGxlZ3JvOmFwaTpzYWxlOm9mZmVyczpyZWFkIiwiYWxsZWdybzphcGk6YmlkcyIsImFsbGVncm86YXBpOm9yZGVyczp3cml0ZSIsImFsbGVncm86YXBpOnBheW1lbnRzOndyaXRlIiwiYWxsZWdybzphcGk6c2FsZTpzZXR0aW5nczp3cml0ZSIsImFsbGVncm86YXBpOnByb2ZpbGU6cmVhZCIsImFsbGVncm86YXBpOnJhdGluZ3MiLCJhbGxlZ3JvOmFwaTpzYWxlOnNldHRpbmdzOnJlYWQiLCJhbGxlZ3JvOmFwaTpwYXltZW50czpyZWFkIiwiYWxsZWdybzphcGk6bWVzc2FnaW5nIl0sI" // string | Otto Refresh Token (optional)
	ottoEnvironment := "sandbox" // string | Otto Environment (optional)
	ottoAccessToken := "eyJhbGciOiJS34535f45f54f5656deyJ1c2VyX25hbWUiOiI5NDUxMzE1MSIsInNjb3BlIjpbImFsbGVncm86YXBpOm9yZGVyczpyZWFkIiwiYWxsZWdybzphcGk6cHJvZmlsZTp3cml0ZSIsImFsbGVncm86YXBpOnNhbGU6b2ZmZXJzOndyaXRlIiwiYWxsZWdybzphcGk6YmlsbGluZzpyZWFkIiwiYWxsZWdybzphcGk6Y2FtcGFpZ25zIiwiYWxsZWdybzphcGk6ZGlzcHV0ZXMiLCJhbGxlZ3JvOmFwaTpzYWxlOm9mZmVyczpyZWFkIiwiYWxsZWdybzphcGk6YmlkcyIsImFsbGVncm86YXBpOm9yZGVyczp3cml0ZSIsImFsbGVncm86YXBpOnBheW1lbnRzOndyaXRlIiwiYWxsZWdybzphcGk6c2FsZTpzZXR0aW5nczp3cml0ZSIsImFsbGVncm86YXBpOnByb2ZpbGU6cmVhZCIsImFsbGVncm86YXBpOnJhdGluZ3MiLCJhbGxlZ3JvOmFwaTpzYWxlOnNldHRpbmdzOnJlYWQiLCJhbGxlZ3JvOmFwaTpwYXltZW50czpyZWFkIiwiYWxsZWdybzphcGk6bWVzc2FnaW5nIl0sI" // string | Otto Access Token (optional)
	tiktokshopAppKey := "6arbhkzno8nbv" // string | TikTok Shop App Key (optional)
	tiktokshopAppSecret := "d95820a05a0cd54fb394fcd26fgat63999b183bc" // string | TikTok Shop App Secret (optional)
	tiktokshopRefreshToken := "TTP_NTUxZTNhYTQ2ZDk2YmRmZWNmYWY2YWY12345NGYwNjQ3YjkzYTllYjA0YmNlMw" // string | TikTok Shop Refresh Token (optional)
	tiktokshopAccessToken := "TTP_Fw8r12345kW03FYd09DG-9INtpw361hWthei12345iPJ5AUv99fLSCYD9-Uu12345TgNRzKZxi5-tfFMtdWqglEt5_iCk" // string | TikTok Shop Access Token (optional)
	sallaClientId := "1bxxxcf9-5xx4-xxx-bxxf-929b8xxxxe11" // string | Salla Client ID (optional)
	sallaClientSecret := "8x88axxxc25e1fxxxa1c06fxxx150xx5" // string | Salla Client Secret (optional)
	sallaRefreshToken := "oxy_rt_zxxxxiY2xxZWWxxxxlU-tROxxxxx2JzS2fwzxxxxxkU.p3xxxkCIyFexxxxP50WwZYfhw5_wg1xxxxV5F-8xxXc" // string | Salla Refresh Token (optional)
	sallaAccessToken := "oxy_rt_zxxxxiY2xxZWWxxxxlU-tROxxxxx2JzS2fwzxxxxxkU.p3xxxkCIyFexxxxP50WwZYfhw5_wg1xxxxV5F-8xxXc" // string | Salla Access Token (optional)
	temuAppKey := "4ebbc9190ae410443d65b4c2faca9811" // string | Temu App Key (optional)
	temuAppSecret := "4782d2d827276688bf4758bed55dbdd4bbe79a78" // string | Temu App Secret (optional)
	temuAccessToken := "uplv3hfyt5kcwoymrgnajnbl1ow5qxlz4sqhev6hl3xosz5dejrtyl2jre6" // string | Temu Access Token (optional)
	temuRegion := "US" // string | Temu API endpoint Region. (optional)
	scapiClientId := "b941ec85-v6b2-1dw8-9fe2-98e1d6eec7b1" // string | Salesforce Commerce API Client ID (optional)
	scapiClientSecret := "testpassword" // string | Salesforce Commerce API Client Secret (optional)
	scapiOrganizationId := "f_ecom_xxxx_001" // string | Salesforce Commerce Organization ID (optional)
	scapiShortCode := "zs5ksm25" // string | Salesforce Commerce Short Code (optional)
	scapiScopes := "sfcc.catalogs,sfcc.orders,sfcc.products" // string | Salesforce Commerce API Scopes (optional)
	idempotencyKey := "098f6bcd4621d373cade4e832627b4f6" // string | A unique identifier associated with a specific request. Repeated requests with the same <strong>idempotency_key</strong> return a cached response without re-executing the business logic. <strong>Please note that the cache lifetime is 15 minutes.</strong> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AccountConfigUpdate(context.Background()).ReplaceParameters(replaceParameters).NewStoreUrl(newStoreUrl).NewStoreKey(newStoreKey).CustomLabel(customLabel).BridgeUrl(bridgeUrl).StoreRoot(storeRoot).DbTablesPrefix(dbTablesPrefix).UserAgent(userAgent).Var3dcartPrivateKey(var3dcartPrivateKey).Var3dcartAccessToken(var3dcartAccessToken).Var3dcartapiApiKey(var3dcartapiApiKey).AmazonSpClientId(amazonSpClientId).AmazonSpClientSecret(amazonSpClientSecret).AmazonSpRefreshToken(amazonSpRefreshToken).AmazonSpAwsRegion(amazonSpAwsRegion).AmazonSpApiEnvironment(amazonSpApiEnvironment).AmazonSellerId(amazonSellerId).AspdotnetstorefrontApiUser(aspdotnetstorefrontApiUser).AspdotnetstorefrontApiPass(aspdotnetstorefrontApiPass).AmericommerceAppId(americommerceAppId).AmericommerceAppSecret(americommerceAppSecret).AmericommerceAccessToken(americommerceAccessToken).AmericommerceRefreshToken(americommerceRefreshToken).BigcommerceapiAdminAccount(bigcommerceapiAdminAccount).BigcommerceapiApiPath(bigcommerceapiApiPath).BigcommerceapiApiKey(bigcommerceapiApiKey).BigcommerceapiClientId(bigcommerceapiClientId).BigcommerceapiAccessToken(bigcommerceapiAccessToken).BigcommerceapiContext(bigcommerceapiContext).BolApiKey(bolApiKey).BolApiSecret(bolApiSecret).BolRetailerId(bolRetailerId).DemandwareClientId(demandwareClientId).DemandwareApiPassword(demandwareApiPassword).DemandwareUserName(demandwareUserName).DemandwareUserPassword(demandwareUserPassword).EbayClientId(ebayClientId).EbayClientSecret(ebayClientSecret).EbayRuname(ebayRuname).EbayAccessToken(ebayAccessToken).EbayRefreshToken(ebayRefreshToken).EbayEnvironment(ebayEnvironment).EbaySiteId(ebaySiteId).EcwidAcessToken(ecwidAcessToken).EcwidStoreId(ecwidStoreId).LazadaAppId(lazadaAppId).LazadaAppSecret(lazadaAppSecret).LazadaRefreshToken(lazadaRefreshToken).LazadaRegion(lazadaRegion).EtsyKeystring(etsyKeystring).EtsySharedSecret(etsySharedSecret).EtsyAccessToken(etsyAccessToken).EtsyTokenSecret(etsyTokenSecret).EtsyClientId(etsyClientId).EtsyRefreshToken(etsyRefreshToken).FacebookAppId(facebookAppId).FacebookAppSecret(facebookAppSecret).FacebookAccessToken(facebookAccessToken).FacebookBusinessId(facebookBusinessId).NetoApiKey(netoApiKey).NetoApiUsername(netoApiUsername).ShoplineAccessToken(shoplineAccessToken).ShoplineAppKey(shoplineAppKey).ShoplineAppSecret(shoplineAppSecret).ShoplineSharedSecret(shoplineSharedSecret).ShopifyAccessToken(shopifyAccessToken).ShopifyClientId(shopifyClientId).ShopifyApiKey(shopifyApiKey).ShopifyApiPassword(shopifyApiPassword).ShopifySharedSecret(shopifySharedSecret).ShopeePartnerId(shopeePartnerId).ShopeePartnerKey(shopeePartnerKey).ShopeeShopId(shopeeShopId).ShopeeRefreshToken(shopeeRefreshToken).ShopeeRegion(shopeeRegion).ShopeeEnvironment(shopeeEnvironment).ShoplazzaAccessToken(shoplazzaAccessToken).ShoplazzaSharedSecret(shoplazzaSharedSecret).MivaAccessToken(mivaAccessToken).MivaSignature(mivaSignature).ShopwareAccessKey(shopwareAccessKey).UnasApiKey(unasApiKey).ShopwareApiKey(shopwareApiKey).ShopwareApiSecret(shopwareApiSecret).BigcartelUserName(bigcartelUserName).BigcartelPassword(bigcartelPassword).BricklinkConsumerKey(bricklinkConsumerKey).BricklinkConsumerSecret(bricklinkConsumerSecret).BricklinkToken(bricklinkToken).BricklinkTokenSecret(bricklinkTokenSecret).VolusionLogin(volusionLogin).VolusionPassword(volusionPassword).WalmartClientId(walmartClientId).WalmartClientSecret(walmartClientSecret).WalmartEnvironment(walmartEnvironment).WalmartChannelType(walmartChannelType).WalmartRegion(walmartRegion).SquareClientId(squareClientId).SquareClientSecret(squareClientSecret).SquareRefreshToken(squareRefreshToken).SquarespaceApiKey(squarespaceApiKey).SquarespaceClientId(squarespaceClientId).SquarespaceClientSecret(squarespaceClientSecret).SquarespaceAccessToken(squarespaceAccessToken).SquarespaceRefreshToken(squarespaceRefreshToken).HybrisClientId(hybrisClientId).HybrisClientSecret(hybrisClientSecret).HybrisUsername(hybrisUsername).HybrisPassword(hybrisPassword).HybrisWebsites(hybrisWebsites).LightspeedApiKey(lightspeedApiKey).LightspeedApiSecret(lightspeedApiSecret).CommercehqApiKey(commercehqApiKey).CommercehqApiPassword(commercehqApiPassword).WcConsumerKey(wcConsumerKey).WcConsumerSecret(wcConsumerSecret).MagentoConsumerKey(magentoConsumerKey).MagentoConsumerSecret(magentoConsumerSecret).MagentoAccessToken(magentoAccessToken).MagentoTokenSecret(magentoTokenSecret).PrestashopWebserviceKey(prestashopWebserviceKey).WixAppId(wixAppId).WixAppSecretKey(wixAppSecretKey).WixInstanceId(wixInstanceId).WixRefreshToken(wixRefreshToken).MercadoLibreAppId(mercadoLibreAppId).MercadoLibreAppSecretKey(mercadoLibreAppSecretKey).MercadoLibreRefreshToken(mercadoLibreRefreshToken).ZidClientId(zidClientId).ZidClientSecret(zidClientSecret).ZidAccessToken(zidAccessToken).ZidAuthorization(zidAuthorization).ZidRefreshToken(zidRefreshToken).JumpsellerClientId(jumpsellerClientId).JumpsellerClientSecret(jumpsellerClientSecret).JumpsellerRefreshToken(jumpsellerRefreshToken).JumpsellerLogin(jumpsellerLogin).JumpsellerAuthtoken(jumpsellerAuthtoken).FlipkartClientId(flipkartClientId).FlipkartClientSecret(flipkartClientSecret).AllegroClientId(allegroClientId).AllegroClientSecret(allegroClientSecret).AllegroAccessToken(allegroAccessToken).AllegroRefreshToken(allegroRefreshToken).AllegroEnvironment(allegroEnvironment).ZohoClientId(zohoClientId).ZohoClientSecret(zohoClientSecret).ZohoRefreshToken(zohoRefreshToken).ZohoRegion(zohoRegion).TiendanubeUserId(tiendanubeUserId).TiendanubeAccessToken(tiendanubeAccessToken).TiendanubeClientSecret(tiendanubeClientSecret).OttoClientId(ottoClientId).OttoClientSecret(ottoClientSecret).OttoAppId(ottoAppId).OttoRefreshToken(ottoRefreshToken).OttoEnvironment(ottoEnvironment).OttoAccessToken(ottoAccessToken).TiktokshopAppKey(tiktokshopAppKey).TiktokshopAppSecret(tiktokshopAppSecret).TiktokshopRefreshToken(tiktokshopRefreshToken).TiktokshopAccessToken(tiktokshopAccessToken).SallaClientId(sallaClientId).SallaClientSecret(sallaClientSecret).SallaRefreshToken(sallaRefreshToken).SallaAccessToken(sallaAccessToken).TemuAppKey(temuAppKey).TemuAppSecret(temuAppSecret).TemuAccessToken(temuAccessToken).TemuRegion(temuRegion).ScapiClientId(scapiClientId).ScapiClientSecret(scapiClientSecret).ScapiOrganizationId(scapiOrganizationId).ScapiShortCode(scapiShortCode).ScapiScopes(scapiScopes).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountConfigUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountConfigUpdate`: AccountConfigUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountConfigUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountConfigUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **replaceParameters** | **bool** | Identifies if there is a necessity to replace parameters | 
 **newStoreUrl** | **string** | The web address of the store you want to update to connect to API2Cart | 
 **newStoreKey** | **string** | Update store key | 
 **customLabel** | **string** | Defines a custom label for the store in the app | 
 **bridgeUrl** | **string** | This parameter allows to set up store with custom bridge url (also you must use store_root parameter if a bridge folder is not in the root folder of the store) | 
 **storeRoot** | **string** | Absolute path to the store root directory (used with \&quot;bridge_url\&quot; parameter) | 
 **dbTablesPrefix** | **string** | DB tables prefix | 
 **userAgent** | **string** | This parameter allows you to set your custom user agent, which will be used in requests to the store. Please use it cautiously, as the store&#39;s firewall may block specific values. | 
 **var3dcartPrivateKey** | **string** | 3DCart Private Key | 
 **var3dcartAccessToken** | **string** | 3DCart Token | 
 **var3dcartapiApiKey** | **string** | 3DCart API Key | 
 **amazonSpClientId** | **string** | Amazon SP API app client id | 
 **amazonSpClientSecret** | **string** | Amazon SP API app client secret | 
 **amazonSpRefreshToken** | **string** | Amazon SP API OAuth refresh token | 
 **amazonSpAwsRegion** | **string** | Amazon AWS Region | 
 **amazonSpApiEnvironment** | **string** | Amazon SP API environment | [default to &quot;production&quot;]
 **amazonSellerId** | **string** | Amazon Seller ID (Merchant token) | 
 **aspdotnetstorefrontApiUser** | **string** | It&#39;s a AspDotNetStorefront account for which API is available | 
 **aspdotnetstorefrontApiPass** | **string** | AspDotNetStorefront API Password | 
 **americommerceAppId** | **string** | Americommerce App ID | 
 **americommerceAppSecret** | **string** | Americommerce App Secret | 
 **americommerceAccessToken** | **string** | Americommerce Access Token | 
 **americommerceRefreshToken** | **string** | Americommerce Refresh Token | 
 **bigcommerceapiAdminAccount** | **string** | It&#39;s a BigCommerce account for which API is enabled | 
 **bigcommerceapiApiPath** | **string** | BigCommerce API URL | 
 **bigcommerceapiApiKey** | **string** | Bigcommerce API Key | 
 **bigcommerceapiClientId** | **string** | Client ID of the requesting app | 
 **bigcommerceapiAccessToken** | **string** | Access token authorizing the app to access resources on behalf of a user | 
 **bigcommerceapiContext** | **string** | API Path section unique to the store | 
 **bolApiKey** | **string** | Bol API Key | 
 **bolApiSecret** | **string** | Bol API Secret | 
 **bolRetailerId** | **int32** | Bol Retailer ID | 
 **demandwareClientId** | **string** | Demandware client id | 
 **demandwareApiPassword** | **string** | Demandware api password | 
 **demandwareUserName** | **string** | Demandware user name | 
 **demandwareUserPassword** | **string** | Demandware user password | 
 **ebayClientId** | **string** | Application ID (AppID). | 
 **ebayClientSecret** | **string** | Shared Secret from eBay application | 
 **ebayRuname** | **string** | The RuName value that eBay assigns to your application. | 
 **ebayAccessToken** | **string** | Used to authenticate API requests. | 
 **ebayRefreshToken** | **string** | Used to renew the access token. | 
 **ebayEnvironment** | **string** | eBay environment | 
 **ebaySiteId** | **int32** | eBay global ID | [default to 0]
 **ecwidAcessToken** | **string** | Access token authorizing the app to access resources on behalf of a user | 
 **ecwidStoreId** | **string** | Store Id | 
 **lazadaAppId** | **string** | Lazada App ID | 
 **lazadaAppSecret** | **string** | Lazada App Secret | 
 **lazadaRefreshToken** | **string** | Lazada Refresh Token | 
 **lazadaRegion** | **string** | Lazada API endpoint Region | 
 **etsyKeystring** | **string** | Etsy keystring | 
 **etsySharedSecret** | **string** | Etsy shared secret | 
 **etsyAccessToken** | **string** | Access token authorizing the app to access resources on behalf of a user | 
 **etsyTokenSecret** | **string** | Secret token authorizing the app to access resources on behalf of a user | 
 **etsyClientId** | **string** | Etsy Client Id | 
 **etsyRefreshToken** | **string** | Etsy Refresh token | 
 **facebookAppId** | **string** | Facebook App ID | 
 **facebookAppSecret** | **string** | Facebook App Secret | 
 **facebookAccessToken** | **string** | Facebook Access Token | 
 **facebookBusinessId** | **string** | Facebook Business ID | 
 **netoApiKey** | **string** | Neto API Key | 
 **netoApiUsername** | **string** | Neto User Name | 
 **shoplineAccessToken** | **string** | Shopline APP Key | 
 **shoplineAppKey** | **string** | Shopline APP Key | 
 **shoplineAppSecret** | **string** | Shopline App Secret | 
 **shoplineSharedSecret** | **string** | Shopline Shared Secret | 
 **shopifyAccessToken** | **string** | Access token authorizing the app to access resources on behalf of a user | 
 **shopifyClientId** | **string** | Shopify Client ID | 
 **shopifyApiKey** | **string** | Shopify API Key | 
 **shopifyApiPassword** | **string** | Shopify API Password | 
 **shopifySharedSecret** | **string** | Shared secret | 
 **shopeePartnerId** | **string** | Shopee Partner ID | 
 **shopeePartnerKey** | **string** | Shopee Partner Key | 
 **shopeeShopId** | **string** | Shopee SHOP ID | 
 **shopeeRefreshToken** | **string** | Shopee Refresh Token | 
 **shopeeRegion** | **string** | Shopee API endpoint Region. Use for Chinese Mainland or Brazil. | 
 **shopeeEnvironment** | **string** | Shopee Environment | 
 **shoplazzaAccessToken** | **string** | Access token authorizing the app to access resources on behalf of a user | 
 **shoplazzaSharedSecret** | **string** | Shared secret | 
 **mivaAccessToken** | **string** | Miva access token | 
 **mivaSignature** | **string** | Miva signature | 
 **shopwareAccessKey** | **string** | Shopware access key | 
 **unasApiKey** | **string** | UNAS API Key | 
 **shopwareApiKey** | **string** | Shopware api key | 
 **shopwareApiSecret** | **string** | Shopware client secret access key | 
 **bigcartelUserName** | **string** | Subdomain of store | 
 **bigcartelPassword** | **string** | BigCartel account password | 
 **bricklinkConsumerKey** | **string** | Bricklink Consumer Key | 
 **bricklinkConsumerSecret** | **string** | Bricklink Consumer Secret | 
 **bricklinkToken** | **string** | Bricklink Access Token | 
 **bricklinkTokenSecret** | **string** | Bricklink Access Token Secret | 
 **volusionLogin** | **string** | It&#39;s a Volusion account for which API is enabled | 
 **volusionPassword** | **string** | Volusion API Password | 
 **walmartClientId** | **string** | Walmart client ID. For the region &#39;ca&#39; use Consumer ID | 
 **walmartClientSecret** | **string** | Walmart client secret. For the region &#39;ca&#39; use Private Key | 
 **walmartEnvironment** | **string** | Walmart environment | [default to &quot;production&quot;]
 **walmartChannelType** | **string** | Walmart WM_CONSUMER.CHANNEL.TYPE header | 
 **walmartRegion** | **string** | Walmart region | [default to &quot;us&quot;]
 **squareClientId** | **string** | Square (Weebly) Client ID | 
 **squareClientSecret** | **string** | Square (Weebly) Client Secret | 
 **squareRefreshToken** | **string** | Square (Weebly) Refresh Token | 
 **squarespaceApiKey** | **string** | Squarespace API Key | 
 **squarespaceClientId** | **string** | Squarespace Connector Client ID | 
 **squarespaceClientSecret** | **string** | Squarespace Connector Client Secret | 
 **squarespaceAccessToken** | **string** | Squarespace access token | 
 **squarespaceRefreshToken** | **string** | Squarespace refresh token | 
 **hybrisClientId** | **string** | Omni Commerce Connector Client ID | 
 **hybrisClientSecret** | **string** | Omni Commerce Connector Client Secret | 
 **hybrisUsername** | **string** | User Name | 
 **hybrisPassword** | **string** | User password | 
 **hybrisWebsites** | **[]string** | Websites to stores mapping data | 
 **lightspeedApiKey** | **string** | LightSpeed api key | 
 **lightspeedApiSecret** | **string** | LightSpeed api secret | 
 **commercehqApiKey** | **string** | CommerceHQ api key | 
 **commercehqApiPassword** | **string** | CommerceHQ api password | 
 **wcConsumerKey** | **string** | Woocommerce consumer key | 
 **wcConsumerSecret** | **string** | Woocommerce consumer secret | 
 **magentoConsumerKey** | **string** | Magento Consumer Key | 
 **magentoConsumerSecret** | **string** | Magento Consumer Secret | 
 **magentoAccessToken** | **string** | Magento Access Token | 
 **magentoTokenSecret** | **string** | Magento Token Secret | 
 **prestashopWebserviceKey** | **string** | Prestashop webservice key | 
 **wixAppId** | **string** | Wix App ID | 
 **wixAppSecretKey** | **string** | Wix App Secret Key | 
 **wixInstanceId** | **string** | Wix Instance ID | 
 **wixRefreshToken** | **string** | Wix refresh token | 
 **mercadoLibreAppId** | **string** | Mercado Libre App ID | 
 **mercadoLibreAppSecretKey** | **string** | Mercado Libre App Secret Key | 
 **mercadoLibreRefreshToken** | **string** | Mercado Libre Refresh Token | 
 **zidClientId** | **int32** | Zid Client ID | 
 **zidClientSecret** | **string** | Zid Client Secret | 
 **zidAccessToken** | **string** | Zid Access Token | 
 **zidAuthorization** | **string** | Zid Authorization | 
 **zidRefreshToken** | **string** | Zid refresh token | 
 **jumpsellerClientId** | **string** | Jumpseller OAuth2 Client ID | 
 **jumpsellerClientSecret** | **string** | Jumpseller OAuth2 Client Secret | 
 **jumpsellerRefreshToken** | **string** | Jumpseller OAuth2 refresh token | 
 **jumpsellerLogin** | **string** | Jumpseller API login | 
 **jumpsellerAuthtoken** | **string** | Jumpseller API auth token | 
 **flipkartClientId** | **string** | Flipkart Client ID | 
 **flipkartClientSecret** | **string** | Flipkart Client Secret | 
 **allegroClientId** | **string** | Allegro Client ID | 
 **allegroClientSecret** | **string** | Allegro Client Secret | 
 **allegroAccessToken** | **string** | Allegro Access Token | 
 **allegroRefreshToken** | **string** | Allegro Refresh Token | 
 **allegroEnvironment** | **string** | Allegro Environment | [default to &quot;production&quot;]
 **zohoClientId** | **string** | Zoho Client ID | 
 **zohoClientSecret** | **string** | Zoho Client Secret | 
 **zohoRefreshToken** | **string** | Zoho Refresh Token | 
 **zohoRegion** | **string** | Zoho API endpoint Region | 
 **tiendanubeUserId** | **int32** | Tiendanube User ID | 
 **tiendanubeAccessToken** | **string** | Tiendanube Access Token | 
 **tiendanubeClientSecret** | **string** | Tiendanube Client Secret | 
 **ottoClientId** | **string** | Otto Client ID | 
 **ottoClientSecret** | **string** | Otto Client Secret | 
 **ottoAppId** | **string** | Otto App ID | 
 **ottoRefreshToken** | **string** | Otto Refresh Token | 
 **ottoEnvironment** | **string** | Otto Environment | 
 **ottoAccessToken** | **string** | Otto Access Token | 
 **tiktokshopAppKey** | **string** | TikTok Shop App Key | 
 **tiktokshopAppSecret** | **string** | TikTok Shop App Secret | 
 **tiktokshopRefreshToken** | **string** | TikTok Shop Refresh Token | 
 **tiktokshopAccessToken** | **string** | TikTok Shop Access Token | 
 **sallaClientId** | **string** | Salla Client ID | 
 **sallaClientSecret** | **string** | Salla Client Secret | 
 **sallaRefreshToken** | **string** | Salla Refresh Token | 
 **sallaAccessToken** | **string** | Salla Access Token | 
 **temuAppKey** | **string** | Temu App Key | 
 **temuAppSecret** | **string** | Temu App Secret | 
 **temuAccessToken** | **string** | Temu Access Token | 
 **temuRegion** | **string** | Temu API endpoint Region. | 
 **scapiClientId** | **string** | Salesforce Commerce API Client ID | 
 **scapiClientSecret** | **string** | Salesforce Commerce API Client Secret | 
 **scapiOrganizationId** | **string** | Salesforce Commerce Organization ID | 
 **scapiShortCode** | **string** | Salesforce Commerce Short Code | 
 **scapiScopes** | **string** | Salesforce Commerce API Scopes | 
 **idempotencyKey** | **string** | A unique identifier associated with a specific request. Repeated requests with the same &lt;strong&gt;idempotency_key&lt;/strong&gt; return a cached response without re-executing the business logic. &lt;strong&gt;Please note that the cache lifetime is 15 minutes.&lt;/strong&gt; | 

### Return type

[**AccountConfigUpdate200Response**](AccountConfigUpdate200Response.md)

### Authorization

[StoreKeyAuth](../README.md#StoreKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountFailedWebhooks

> AccountFailedWebhooks200Response AccountFailedWebhooks(ctx).Start(start).Count(count).Ids(ids).Execute()

account.failed_webhooks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	start := int32(0) // int32 | This parameter sets the number from which you want to get entities (optional) (default to 0)
	count := int32(20) // int32 | This parameter sets the entity amount that has to be retrieved. Max allowed count=250 (optional) (default to 10)
	ids := "3,14,25" // string | List of сomma-separated webhook ids (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AccountFailedWebhooks(context.Background()).Start(start).Count(count).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountFailedWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountFailedWebhooks`: AccountFailedWebhooks200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountFailedWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountFailedWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | This parameter sets the number from which you want to get entities | [default to 0]
 **count** | **int32** | This parameter sets the entity amount that has to be retrieved. Max allowed count&#x3D;250 | [default to 10]
 **ids** | **string** | List of сomma-separated webhook ids | 

### Return type

[**AccountFailedWebhooks200Response**](AccountFailedWebhooks200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountSupportedPlatforms

> AccountSupportedPlatforms200Response AccountSupportedPlatforms(ctx).Execute()

account.supported_platforms



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.AccountSupportedPlatforms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AccountSupportedPlatforms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountSupportedPlatforms`: AccountSupportedPlatforms200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AccountSupportedPlatforms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountSupportedPlatformsRequest struct via the builder pattern


### Return type

[**AccountSupportedPlatforms200Response**](AccountSupportedPlatforms200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

