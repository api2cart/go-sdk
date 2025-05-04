# AccountCartAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CartId** | **string** | Store’s identifier which you can get from cart_list method | 
**StoreUrl** | Pointer to **string** | A web address of a store that you would like to connect to API2Cart | [optional] 
**BridgeUrl** | Pointer to **string** | This parameter allows to set up store with custom bridge url (also you must use store_root parameter if a bridge folder is not in the root folder of the store) | [optional] 
**StoreRoot** | Pointer to **string** | Absolute path to the store root directory (used with \&quot;bridge_url\&quot; parameter) | [optional] 
**StoreKey** | Pointer to **string** | Set this parameter if bridge is already uploaded to store | [optional] 
**ValidateVersion** | Pointer to **bool** | Specify if api2cart should validate cart version | [optional] [default to false]
**Verify** | Pointer to **bool** | Enables or disables cart&#39;s verification | [optional] [default to true]
**DbTablesPrefix** | Pointer to **string** | DB tables prefix | [optional] 
**UserAgent** | Pointer to **string** | This parameter allows you to set your custom user agent, which will be used in requests to the store. Please use it cautiously, as the store&#39;s firewall may block specific values. | [optional] 
**FtpHost** | Pointer to **string** | FTP connection host | [optional] 
**FtpUser** | Pointer to **string** | FTP User | [optional] 
**FtpPassword** | Pointer to **string** | FTP Password | [optional] 
**FtpPort** | Pointer to **int32** | FTP Port | [optional] 
**FtpStoreDir** | Pointer to **string** | FTP Store dir | [optional] 
**Var3dcartPrivateKey** | Pointer to **string** | 3DCart Private Key | [optional] 
**Var3dcartAccessToken** | Pointer to **string** | 3DCart Token | [optional] 
**Var3dcartapiApiKey** | Pointer to **string** | 3DCart API Key | [optional] 
**AmazonSpClientId** | Pointer to **string** | Amazon SP API app client id | [optional] 
**AmazonSpClientSecret** | Pointer to **string** | Amazon SP API app client secret | [optional] 
**AmazonSpRefreshToken** | Pointer to **string** | Amazon SP API OAuth refresh token | [optional] 
**AmazonSpAwsRegion** | Pointer to **string** | Amazon AWS Region | [optional] 
**AmazonSpApiEnvironment** | Pointer to **string** | Amazon SP API environment | [optional] [default to "production"]
**AmazonSellerId** | Pointer to **string** | Amazon Seller ID (Merchant token) | [optional] 
**AspdotnetstorefrontApiUser** | Pointer to **string** | It&#39;s a AspDotNetStorefront account for which API is available | [optional] 
**AspdotnetstorefrontApiPass** | Pointer to **string** | AspDotNetStorefront API Password | [optional] 
**BigcommerceapiAdminAccount** | Pointer to **string** | It&#39;s a BigCommerce account for which API is enabled | [optional] 
**BigcommerceapiApiPath** | Pointer to **string** | BigCommerce API URL | [optional] 
**BigcommerceapiApiKey** | Pointer to **string** | Bigcommerce API Key | [optional] 
**BigcommerceapiClientId** | Pointer to **string** | Client ID of the requesting app | [optional] 
**BigcommerceapiAccessToken** | Pointer to **string** | Access token authorizing the app to access resources on behalf of a user | [optional] 
**BigcommerceapiContext** | Pointer to **string** | API Path section unique to the store | [optional] 
**BolApiKey** | Pointer to **string** | Bol API Key | [optional] 
**BolApiSecret** | Pointer to **string** | Bol API Secret | [optional] 
**BolRetailerId** | Pointer to **int32** | Bol Retailer ID | [optional] 
**BigcartelUserName** | **string** | Subdomain of store | 
**BigcartelPassword** | **string** | BigCartel account password | 
**DemandwareClientId** | Pointer to **string** | Demandware client id | [optional] 
**DemandwareApiPassword** | Pointer to **string** | Demandware api password | [optional] 
**DemandwareUserName** | Pointer to **string** | Demandware user name | [optional] 
**DemandwareUserPassword** | Pointer to **string** | Demandware user password | [optional] 
**EbayClientId** | Pointer to **string** | Application ID (AppID). | [optional] 
**EbayClientSecret** | Pointer to **string** | Shared Secret from eBay application | [optional] 
**EbayRuname** | Pointer to **string** | The RuName value that eBay assigns to your application. | [optional] 
**EbayAccessToken** | Pointer to **string** | Used to authenticate API requests. | [optional] 
**EbayRefreshToken** | Pointer to **string** | Used to renew the access token. | [optional] 
**EbayEnvironment** | Pointer to **string** | eBay environment | [optional] [default to "production"]
**EbaySiteId** | Pointer to **int32** | eBay global ID | [optional] [default to 0]
**WalmartClientId** | Pointer to **string** | Walmart client ID. For the region &#39;ca&#39; use Consumer ID | [optional] 
**WalmartClientSecret** | Pointer to **string** | Walmart client secret. For the region &#39;ca&#39; use Private Key | [optional] 
**WalmartEnvironment** | Pointer to **string** | Walmart environment | [optional] [default to "production"]
**WalmartChannelType** | Pointer to **string** | Walmart WM_CONSUMER.CHANNEL.TYPE header | [optional] 
**WalmartRegion** | Pointer to **string** | Walmart region | [optional] [default to "us"]
**EcwidAcessToken** | Pointer to **string** | Access token authorizing the app to access resources on behalf of a user | [optional] 
**EcwidStoreId** | Pointer to **string** | Store Id | [optional] 
**LazadaAppId** | Pointer to **string** | Lazada App ID | [optional] 
**LazadaAppSecret** | Pointer to **string** | Lazada App Secret | [optional] 
**LazadaRefreshToken** | Pointer to **string** | Lazada Refresh Token | [optional] 
**LazadaRegion** | Pointer to **string** | Lazada API endpoint Region | [optional] 
**LightspeedApiKey** | Pointer to **string** | LightSpeed api key | [optional] 
**LightspeedApiSecret** | Pointer to **string** | LightSpeed api secret | [optional] 
**EtsyKeystring** | Pointer to **string** | Etsy keystring | [optional] 
**EtsySharedSecret** | Pointer to **string** | Etsy shared secret | [optional] 
**EtsyAccessToken** | Pointer to **string** | Access token authorizing the app to access resources on behalf of a user | [optional] 
**EtsyTokenSecret** | Pointer to **string** | Secret token authorizing the app to access resources on behalf of a user | [optional] 
**EtsyClientId** | Pointer to **string** | Etsy Client Id | [optional] 
**EtsyRefreshToken** | Pointer to **string** | Etsy Refresh token | [optional] 
**FacebookAppId** | Pointer to **string** | Facebook App ID | [optional] 
**FacebookAppSecret** | Pointer to **string** | Facebook App Secret | [optional] 
**FacebookAccessToken** | Pointer to **string** | Facebook Access Token | [optional] 
**FacebookBusinessId** | Pointer to **string** | Facebook Business ID | [optional] 
**NetoApiKey** | Pointer to **string** | Neto API Key | [optional] 
**NetoApiUsername** | Pointer to **string** | Neto User Name | [optional] 
**ShoplineAccessToken** | Pointer to **string** | Shopline APP Key | [optional] 
**ShoplineAppKey** | Pointer to **string** | Shopline APP Key | [optional] 
**ShoplineAppSecret** | Pointer to **string** | Shopline App Secret | [optional] 
**ShoplineSharedSecret** | Pointer to **string** | Shopline Shared Secret | [optional] 
**ShopifyAccessToken** | Pointer to **string** | Access token authorizing the app to access resources on behalf of a user | [optional] 
**ShopifyApiKey** | Pointer to **string** | Shopify API Key | [optional] 
**ShopifyApiPassword** | Pointer to **string** | Shopify API Password | [optional] 
**ShopifySharedSecret** | Pointer to **string** | Shared secret | [optional] 
**ShoplazzaAccessToken** | Pointer to **string** | Access token authorizing the app to access resources on behalf of a user | [optional] 
**ShoplazzaSharedSecret** | Pointer to **string** | Shared secret | [optional] 
**ShopwareAccessKey** | Pointer to **string** | Shopware access key | [optional] 
**ShopwareApiKey** | Pointer to **string** | Shopware api key | [optional] 
**ShopwareApiSecret** | Pointer to **string** | Shopware client secret access key | [optional] 
**MivaAccessToken** | Pointer to **string** | Miva access token | [optional] 
**MivaSignature** | Pointer to **string** | Miva signature | [optional] 
**TiendanubeUserId** | Pointer to **int32** | Tiendanube User ID | [optional] 
**TiendanubeAccessToken** | Pointer to **string** | Tiendanube Access Token | [optional] 
**TiendanubeClientSecret** | Pointer to **string** | Tiendanube Client Secret | [optional] 
**VolusionLogin** | Pointer to **string** | It&#39;s a Volusion account for which API is enabled | [optional] 
**VolusionPassword** | Pointer to **string** | Volusion API Password | [optional] 
**HybrisClientId** | Pointer to **string** | Omni Commerce Connector Client ID | [optional] 
**HybrisClientSecret** | Pointer to **string** | Omni Commerce Connector Client Secret | [optional] 
**HybrisUsername** | Pointer to **string** | User Name | [optional] 
**HybrisPassword** | Pointer to **string** | User password | [optional] 
**HybrisWebsites** | Pointer to [**[]AccountCartAddHybrisWebsitesInner**](AccountCartAddHybrisWebsitesInner.md) | Websites to stores mapping data | [optional] 
**SquareClientId** | Pointer to **string** | Square (Weebly) Client ID | [optional] 
**SquareClientSecret** | Pointer to **string** | Square (Weebly) Client Secret | [optional] 
**SquareRefreshToken** | Pointer to **string** | Square (Weebly) Refresh Token | [optional] 
**SquarespaceApiKey** | Pointer to **string** | Squarespace API Key | [optional] 
**SquarespaceClientId** | Pointer to **string** | Squarespace Connector Client ID | [optional] 
**SquarespaceClientSecret** | Pointer to **string** | Squarespace Connector Client Secret | [optional] 
**SquarespaceAccessToken** | Pointer to **string** | Squarespace access token | [optional] 
**SquarespaceRefreshToken** | Pointer to **string** | Squarespace refresh token | [optional] 
**CommercehqApiKey** | Pointer to **string** | CommerceHQ api key | [optional] 
**CommercehqApiPassword** | Pointer to **string** | CommerceHQ api password | [optional] 
**WcConsumerKey** | Pointer to **string** | Woocommerce consumer key | [optional] 
**WcConsumerSecret** | Pointer to **string** | Woocommerce consumer secret | [optional] 
**MagentoConsumerKey** | Pointer to **string** | Magento Consumer Key | [optional] 
**MagentoConsumerSecret** | Pointer to **string** | Magento Consumer Secret | [optional] 
**MagentoAccessToken** | Pointer to **string** | Magento Access Token | [optional] 
**MagentoTokenSecret** | Pointer to **string** | Magento Token Secret | [optional] 
**PrestashopWebserviceKey** | Pointer to **string** | Prestashop webservice key | [optional] 
**WixAppId** | **string** | Wix App ID | 
**WixAppSecretKey** | **string** | Wix App Secret Key | 
**WixInstanceId** | Pointer to **string** | Wix Instance ID | [optional] 
**WixRefreshToken** | Pointer to **string** | Wix refresh token | [optional] 
**MercadoLibreAppId** | Pointer to **string** | Mercado Libre App ID | [optional] 
**MercadoLibreAppSecretKey** | Pointer to **string** | Mercado Libre App Secret Key | [optional] 
**MercadoLibreRefreshToken** | Pointer to **string** | Mercado Libre Refresh Token | [optional] 
**ZidClientId** | Pointer to **int32** | Zid Client ID | [optional] 
**ZidClientSecret** | Pointer to **string** | Zid Client Secret | [optional] 
**ZidAccessToken** | Pointer to **string** | Zid Access Token | [optional] 
**ZidAuthorization** | Pointer to **string** | Zid Authorization | [optional] 
**ZidRefreshToken** | Pointer to **string** | Zid refresh token | [optional] 
**FlipkartClientId** | Pointer to **string** | Flipkart Client ID | [optional] 
**FlipkartClientSecret** | Pointer to **string** | Flipkart Client Secret | [optional] 
**AllegroClientId** | Pointer to **string** | Allegro Client ID | [optional] 
**AllegroClientSecret** | Pointer to **string** | Allegro Client Secret | [optional] 
**AllegroAccessToken** | Pointer to **string** | Allegro Access Token | [optional] 
**AllegroRefreshToken** | Pointer to **string** | Allegro Refresh Token | [optional] 
**AllegroEnvironment** | Pointer to **string** | Allegro Environment | [optional] [default to "production"]
**ZohoClientId** | Pointer to **string** | Zoho Client ID | [optional] 
**ZohoClientSecret** | Pointer to **string** | Zoho Client Secret | [optional] 
**ZohoRefreshToken** | Pointer to **string** | Zoho Refresh Token | [optional] 
**ZohoRegion** | Pointer to **string** | Zoho API endpoint Region | [optional] 
**OttoClientId** | Pointer to **string** | Otto Client ID | [optional] 
**OttoClientSecret** | Pointer to **string** | Otto Client Secret | [optional] 
**OttoAppId** | Pointer to **string** | Otto App ID | [optional] 
**OttoRefreshToken** | Pointer to **string** | Otto Refresh Token | [optional] 
**OttoEnvironment** | Pointer to **string** | Otto Environment | [optional] [default to "production"]
**OttoAccessToken** | Pointer to **string** | Otto Access Token | [optional] 
**TiktokshopAppKey** | Pointer to **string** | TikTok Shop App Key | [optional] 
**TiktokshopAppSecret** | Pointer to **string** | TikTok Shop App Secret | [optional] 
**TiktokshopRefreshToken** | Pointer to **string** | TikTok Shop Refresh Token | [optional] 
**TiktokshopAccessToken** | Pointer to **string** | TikTok Shop Access Token | [optional] 
**SallaClientId** | Pointer to **string** | Salla Client ID | [optional] 
**SallaClientSecret** | Pointer to **string** | Salla Client Secret | [optional] 
**SallaRefreshToken** | Pointer to **string** | Salla Refresh Token | [optional] 
**SallaAccessToken** | Pointer to **string** | Salla Access Token | [optional] 

## Methods

### NewAccountCartAdd

`func NewAccountCartAdd(cartId string, bigcartelUserName string, bigcartelPassword string, wixAppId string, wixAppSecretKey string, ) *AccountCartAdd`

NewAccountCartAdd instantiates a new AccountCartAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCartAddWithDefaults

`func NewAccountCartAddWithDefaults() *AccountCartAdd`

NewAccountCartAddWithDefaults instantiates a new AccountCartAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCartId

`func (o *AccountCartAdd) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *AccountCartAdd) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *AccountCartAdd) SetCartId(v string)`

SetCartId sets CartId field to given value.


### GetStoreUrl

`func (o *AccountCartAdd) GetStoreUrl() string`

GetStoreUrl returns the StoreUrl field if non-nil, zero value otherwise.

### GetStoreUrlOk

`func (o *AccountCartAdd) GetStoreUrlOk() (*string, bool)`

GetStoreUrlOk returns a tuple with the StoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreUrl

`func (o *AccountCartAdd) SetStoreUrl(v string)`

SetStoreUrl sets StoreUrl field to given value.

### HasStoreUrl

`func (o *AccountCartAdd) HasStoreUrl() bool`

HasStoreUrl returns a boolean if a field has been set.

### GetBridgeUrl

`func (o *AccountCartAdd) GetBridgeUrl() string`

GetBridgeUrl returns the BridgeUrl field if non-nil, zero value otherwise.

### GetBridgeUrlOk

`func (o *AccountCartAdd) GetBridgeUrlOk() (*string, bool)`

GetBridgeUrlOk returns a tuple with the BridgeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeUrl

`func (o *AccountCartAdd) SetBridgeUrl(v string)`

SetBridgeUrl sets BridgeUrl field to given value.

### HasBridgeUrl

`func (o *AccountCartAdd) HasBridgeUrl() bool`

HasBridgeUrl returns a boolean if a field has been set.

### GetStoreRoot

`func (o *AccountCartAdd) GetStoreRoot() string`

GetStoreRoot returns the StoreRoot field if non-nil, zero value otherwise.

### GetStoreRootOk

`func (o *AccountCartAdd) GetStoreRootOk() (*string, bool)`

GetStoreRootOk returns a tuple with the StoreRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreRoot

`func (o *AccountCartAdd) SetStoreRoot(v string)`

SetStoreRoot sets StoreRoot field to given value.

### HasStoreRoot

`func (o *AccountCartAdd) HasStoreRoot() bool`

HasStoreRoot returns a boolean if a field has been set.

### GetStoreKey

`func (o *AccountCartAdd) GetStoreKey() string`

GetStoreKey returns the StoreKey field if non-nil, zero value otherwise.

### GetStoreKeyOk

`func (o *AccountCartAdd) GetStoreKeyOk() (*string, bool)`

GetStoreKeyOk returns a tuple with the StoreKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreKey

`func (o *AccountCartAdd) SetStoreKey(v string)`

SetStoreKey sets StoreKey field to given value.

### HasStoreKey

`func (o *AccountCartAdd) HasStoreKey() bool`

HasStoreKey returns a boolean if a field has been set.

### GetValidateVersion

`func (o *AccountCartAdd) GetValidateVersion() bool`

GetValidateVersion returns the ValidateVersion field if non-nil, zero value otherwise.

### GetValidateVersionOk

`func (o *AccountCartAdd) GetValidateVersionOk() (*bool, bool)`

GetValidateVersionOk returns a tuple with the ValidateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateVersion

`func (o *AccountCartAdd) SetValidateVersion(v bool)`

SetValidateVersion sets ValidateVersion field to given value.

### HasValidateVersion

`func (o *AccountCartAdd) HasValidateVersion() bool`

HasValidateVersion returns a boolean if a field has been set.

### GetVerify

`func (o *AccountCartAdd) GetVerify() bool`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *AccountCartAdd) GetVerifyOk() (*bool, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *AccountCartAdd) SetVerify(v bool)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *AccountCartAdd) HasVerify() bool`

HasVerify returns a boolean if a field has been set.

### GetDbTablesPrefix

`func (o *AccountCartAdd) GetDbTablesPrefix() string`

GetDbTablesPrefix returns the DbTablesPrefix field if non-nil, zero value otherwise.

### GetDbTablesPrefixOk

`func (o *AccountCartAdd) GetDbTablesPrefixOk() (*string, bool)`

GetDbTablesPrefixOk returns a tuple with the DbTablesPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbTablesPrefix

`func (o *AccountCartAdd) SetDbTablesPrefix(v string)`

SetDbTablesPrefix sets DbTablesPrefix field to given value.

### HasDbTablesPrefix

`func (o *AccountCartAdd) HasDbTablesPrefix() bool`

HasDbTablesPrefix returns a boolean if a field has been set.

### GetUserAgent

`func (o *AccountCartAdd) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *AccountCartAdd) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *AccountCartAdd) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *AccountCartAdd) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetFtpHost

`func (o *AccountCartAdd) GetFtpHost() string`

GetFtpHost returns the FtpHost field if non-nil, zero value otherwise.

### GetFtpHostOk

`func (o *AccountCartAdd) GetFtpHostOk() (*string, bool)`

GetFtpHostOk returns a tuple with the FtpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpHost

`func (o *AccountCartAdd) SetFtpHost(v string)`

SetFtpHost sets FtpHost field to given value.

### HasFtpHost

`func (o *AccountCartAdd) HasFtpHost() bool`

HasFtpHost returns a boolean if a field has been set.

### GetFtpUser

`func (o *AccountCartAdd) GetFtpUser() string`

GetFtpUser returns the FtpUser field if non-nil, zero value otherwise.

### GetFtpUserOk

`func (o *AccountCartAdd) GetFtpUserOk() (*string, bool)`

GetFtpUserOk returns a tuple with the FtpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpUser

`func (o *AccountCartAdd) SetFtpUser(v string)`

SetFtpUser sets FtpUser field to given value.

### HasFtpUser

`func (o *AccountCartAdd) HasFtpUser() bool`

HasFtpUser returns a boolean if a field has been set.

### GetFtpPassword

`func (o *AccountCartAdd) GetFtpPassword() string`

GetFtpPassword returns the FtpPassword field if non-nil, zero value otherwise.

### GetFtpPasswordOk

`func (o *AccountCartAdd) GetFtpPasswordOk() (*string, bool)`

GetFtpPasswordOk returns a tuple with the FtpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpPassword

`func (o *AccountCartAdd) SetFtpPassword(v string)`

SetFtpPassword sets FtpPassword field to given value.

### HasFtpPassword

`func (o *AccountCartAdd) HasFtpPassword() bool`

HasFtpPassword returns a boolean if a field has been set.

### GetFtpPort

`func (o *AccountCartAdd) GetFtpPort() int32`

GetFtpPort returns the FtpPort field if non-nil, zero value otherwise.

### GetFtpPortOk

`func (o *AccountCartAdd) GetFtpPortOk() (*int32, bool)`

GetFtpPortOk returns a tuple with the FtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpPort

`func (o *AccountCartAdd) SetFtpPort(v int32)`

SetFtpPort sets FtpPort field to given value.

### HasFtpPort

`func (o *AccountCartAdd) HasFtpPort() bool`

HasFtpPort returns a boolean if a field has been set.

### GetFtpStoreDir

`func (o *AccountCartAdd) GetFtpStoreDir() string`

GetFtpStoreDir returns the FtpStoreDir field if non-nil, zero value otherwise.

### GetFtpStoreDirOk

`func (o *AccountCartAdd) GetFtpStoreDirOk() (*string, bool)`

GetFtpStoreDirOk returns a tuple with the FtpStoreDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpStoreDir

`func (o *AccountCartAdd) SetFtpStoreDir(v string)`

SetFtpStoreDir sets FtpStoreDir field to given value.

### HasFtpStoreDir

`func (o *AccountCartAdd) HasFtpStoreDir() bool`

HasFtpStoreDir returns a boolean if a field has been set.

### GetVar3dcartPrivateKey

`func (o *AccountCartAdd) GetVar3dcartPrivateKey() string`

GetVar3dcartPrivateKey returns the Var3dcartPrivateKey field if non-nil, zero value otherwise.

### GetVar3dcartPrivateKeyOk

`func (o *AccountCartAdd) GetVar3dcartPrivateKeyOk() (*string, bool)`

GetVar3dcartPrivateKeyOk returns a tuple with the Var3dcartPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3dcartPrivateKey

`func (o *AccountCartAdd) SetVar3dcartPrivateKey(v string)`

SetVar3dcartPrivateKey sets Var3dcartPrivateKey field to given value.

### HasVar3dcartPrivateKey

`func (o *AccountCartAdd) HasVar3dcartPrivateKey() bool`

HasVar3dcartPrivateKey returns a boolean if a field has been set.

### GetVar3dcartAccessToken

`func (o *AccountCartAdd) GetVar3dcartAccessToken() string`

GetVar3dcartAccessToken returns the Var3dcartAccessToken field if non-nil, zero value otherwise.

### GetVar3dcartAccessTokenOk

`func (o *AccountCartAdd) GetVar3dcartAccessTokenOk() (*string, bool)`

GetVar3dcartAccessTokenOk returns a tuple with the Var3dcartAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3dcartAccessToken

`func (o *AccountCartAdd) SetVar3dcartAccessToken(v string)`

SetVar3dcartAccessToken sets Var3dcartAccessToken field to given value.

### HasVar3dcartAccessToken

`func (o *AccountCartAdd) HasVar3dcartAccessToken() bool`

HasVar3dcartAccessToken returns a boolean if a field has been set.

### GetVar3dcartapiApiKey

`func (o *AccountCartAdd) GetVar3dcartapiApiKey() string`

GetVar3dcartapiApiKey returns the Var3dcartapiApiKey field if non-nil, zero value otherwise.

### GetVar3dcartapiApiKeyOk

`func (o *AccountCartAdd) GetVar3dcartapiApiKeyOk() (*string, bool)`

GetVar3dcartapiApiKeyOk returns a tuple with the Var3dcartapiApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3dcartapiApiKey

`func (o *AccountCartAdd) SetVar3dcartapiApiKey(v string)`

SetVar3dcartapiApiKey sets Var3dcartapiApiKey field to given value.

### HasVar3dcartapiApiKey

`func (o *AccountCartAdd) HasVar3dcartapiApiKey() bool`

HasVar3dcartapiApiKey returns a boolean if a field has been set.

### GetAmazonSpClientId

`func (o *AccountCartAdd) GetAmazonSpClientId() string`

GetAmazonSpClientId returns the AmazonSpClientId field if non-nil, zero value otherwise.

### GetAmazonSpClientIdOk

`func (o *AccountCartAdd) GetAmazonSpClientIdOk() (*string, bool)`

GetAmazonSpClientIdOk returns a tuple with the AmazonSpClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonSpClientId

`func (o *AccountCartAdd) SetAmazonSpClientId(v string)`

SetAmazonSpClientId sets AmazonSpClientId field to given value.

### HasAmazonSpClientId

`func (o *AccountCartAdd) HasAmazonSpClientId() bool`

HasAmazonSpClientId returns a boolean if a field has been set.

### GetAmazonSpClientSecret

`func (o *AccountCartAdd) GetAmazonSpClientSecret() string`

GetAmazonSpClientSecret returns the AmazonSpClientSecret field if non-nil, zero value otherwise.

### GetAmazonSpClientSecretOk

`func (o *AccountCartAdd) GetAmazonSpClientSecretOk() (*string, bool)`

GetAmazonSpClientSecretOk returns a tuple with the AmazonSpClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonSpClientSecret

`func (o *AccountCartAdd) SetAmazonSpClientSecret(v string)`

SetAmazonSpClientSecret sets AmazonSpClientSecret field to given value.

### HasAmazonSpClientSecret

`func (o *AccountCartAdd) HasAmazonSpClientSecret() bool`

HasAmazonSpClientSecret returns a boolean if a field has been set.

### GetAmazonSpRefreshToken

`func (o *AccountCartAdd) GetAmazonSpRefreshToken() string`

GetAmazonSpRefreshToken returns the AmazonSpRefreshToken field if non-nil, zero value otherwise.

### GetAmazonSpRefreshTokenOk

`func (o *AccountCartAdd) GetAmazonSpRefreshTokenOk() (*string, bool)`

GetAmazonSpRefreshTokenOk returns a tuple with the AmazonSpRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonSpRefreshToken

`func (o *AccountCartAdd) SetAmazonSpRefreshToken(v string)`

SetAmazonSpRefreshToken sets AmazonSpRefreshToken field to given value.

### HasAmazonSpRefreshToken

`func (o *AccountCartAdd) HasAmazonSpRefreshToken() bool`

HasAmazonSpRefreshToken returns a boolean if a field has been set.

### GetAmazonSpAwsRegion

`func (o *AccountCartAdd) GetAmazonSpAwsRegion() string`

GetAmazonSpAwsRegion returns the AmazonSpAwsRegion field if non-nil, zero value otherwise.

### GetAmazonSpAwsRegionOk

`func (o *AccountCartAdd) GetAmazonSpAwsRegionOk() (*string, bool)`

GetAmazonSpAwsRegionOk returns a tuple with the AmazonSpAwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonSpAwsRegion

`func (o *AccountCartAdd) SetAmazonSpAwsRegion(v string)`

SetAmazonSpAwsRegion sets AmazonSpAwsRegion field to given value.

### HasAmazonSpAwsRegion

`func (o *AccountCartAdd) HasAmazonSpAwsRegion() bool`

HasAmazonSpAwsRegion returns a boolean if a field has been set.

### GetAmazonSpApiEnvironment

`func (o *AccountCartAdd) GetAmazonSpApiEnvironment() string`

GetAmazonSpApiEnvironment returns the AmazonSpApiEnvironment field if non-nil, zero value otherwise.

### GetAmazonSpApiEnvironmentOk

`func (o *AccountCartAdd) GetAmazonSpApiEnvironmentOk() (*string, bool)`

GetAmazonSpApiEnvironmentOk returns a tuple with the AmazonSpApiEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonSpApiEnvironment

`func (o *AccountCartAdd) SetAmazonSpApiEnvironment(v string)`

SetAmazonSpApiEnvironment sets AmazonSpApiEnvironment field to given value.

### HasAmazonSpApiEnvironment

`func (o *AccountCartAdd) HasAmazonSpApiEnvironment() bool`

HasAmazonSpApiEnvironment returns a boolean if a field has been set.

### GetAmazonSellerId

`func (o *AccountCartAdd) GetAmazonSellerId() string`

GetAmazonSellerId returns the AmazonSellerId field if non-nil, zero value otherwise.

### GetAmazonSellerIdOk

`func (o *AccountCartAdd) GetAmazonSellerIdOk() (*string, bool)`

GetAmazonSellerIdOk returns a tuple with the AmazonSellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonSellerId

`func (o *AccountCartAdd) SetAmazonSellerId(v string)`

SetAmazonSellerId sets AmazonSellerId field to given value.

### HasAmazonSellerId

`func (o *AccountCartAdd) HasAmazonSellerId() bool`

HasAmazonSellerId returns a boolean if a field has been set.

### GetAspdotnetstorefrontApiUser

`func (o *AccountCartAdd) GetAspdotnetstorefrontApiUser() string`

GetAspdotnetstorefrontApiUser returns the AspdotnetstorefrontApiUser field if non-nil, zero value otherwise.

### GetAspdotnetstorefrontApiUserOk

`func (o *AccountCartAdd) GetAspdotnetstorefrontApiUserOk() (*string, bool)`

GetAspdotnetstorefrontApiUserOk returns a tuple with the AspdotnetstorefrontApiUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspdotnetstorefrontApiUser

`func (o *AccountCartAdd) SetAspdotnetstorefrontApiUser(v string)`

SetAspdotnetstorefrontApiUser sets AspdotnetstorefrontApiUser field to given value.

### HasAspdotnetstorefrontApiUser

`func (o *AccountCartAdd) HasAspdotnetstorefrontApiUser() bool`

HasAspdotnetstorefrontApiUser returns a boolean if a field has been set.

### GetAspdotnetstorefrontApiPass

`func (o *AccountCartAdd) GetAspdotnetstorefrontApiPass() string`

GetAspdotnetstorefrontApiPass returns the AspdotnetstorefrontApiPass field if non-nil, zero value otherwise.

### GetAspdotnetstorefrontApiPassOk

`func (o *AccountCartAdd) GetAspdotnetstorefrontApiPassOk() (*string, bool)`

GetAspdotnetstorefrontApiPassOk returns a tuple with the AspdotnetstorefrontApiPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspdotnetstorefrontApiPass

`func (o *AccountCartAdd) SetAspdotnetstorefrontApiPass(v string)`

SetAspdotnetstorefrontApiPass sets AspdotnetstorefrontApiPass field to given value.

### HasAspdotnetstorefrontApiPass

`func (o *AccountCartAdd) HasAspdotnetstorefrontApiPass() bool`

HasAspdotnetstorefrontApiPass returns a boolean if a field has been set.

### GetBigcommerceapiAdminAccount

`func (o *AccountCartAdd) GetBigcommerceapiAdminAccount() string`

GetBigcommerceapiAdminAccount returns the BigcommerceapiAdminAccount field if non-nil, zero value otherwise.

### GetBigcommerceapiAdminAccountOk

`func (o *AccountCartAdd) GetBigcommerceapiAdminAccountOk() (*string, bool)`

GetBigcommerceapiAdminAccountOk returns a tuple with the BigcommerceapiAdminAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigcommerceapiAdminAccount

`func (o *AccountCartAdd) SetBigcommerceapiAdminAccount(v string)`

SetBigcommerceapiAdminAccount sets BigcommerceapiAdminAccount field to given value.

### HasBigcommerceapiAdminAccount

`func (o *AccountCartAdd) HasBigcommerceapiAdminAccount() bool`

HasBigcommerceapiAdminAccount returns a boolean if a field has been set.

### GetBigcommerceapiApiPath

`func (o *AccountCartAdd) GetBigcommerceapiApiPath() string`

GetBigcommerceapiApiPath returns the BigcommerceapiApiPath field if non-nil, zero value otherwise.

### GetBigcommerceapiApiPathOk

`func (o *AccountCartAdd) GetBigcommerceapiApiPathOk() (*string, bool)`

GetBigcommerceapiApiPathOk returns a tuple with the BigcommerceapiApiPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigcommerceapiApiPath

`func (o *AccountCartAdd) SetBigcommerceapiApiPath(v string)`

SetBigcommerceapiApiPath sets BigcommerceapiApiPath field to given value.

### HasBigcommerceapiApiPath

`func (o *AccountCartAdd) HasBigcommerceapiApiPath() bool`

HasBigcommerceapiApiPath returns a boolean if a field has been set.

### GetBigcommerceapiApiKey

`func (o *AccountCartAdd) GetBigcommerceapiApiKey() string`

GetBigcommerceapiApiKey returns the BigcommerceapiApiKey field if non-nil, zero value otherwise.

### GetBigcommerceapiApiKeyOk

`func (o *AccountCartAdd) GetBigcommerceapiApiKeyOk() (*string, bool)`

GetBigcommerceapiApiKeyOk returns a tuple with the BigcommerceapiApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigcommerceapiApiKey

`func (o *AccountCartAdd) SetBigcommerceapiApiKey(v string)`

SetBigcommerceapiApiKey sets BigcommerceapiApiKey field to given value.

### HasBigcommerceapiApiKey

`func (o *AccountCartAdd) HasBigcommerceapiApiKey() bool`

HasBigcommerceapiApiKey returns a boolean if a field has been set.

### GetBigcommerceapiClientId

`func (o *AccountCartAdd) GetBigcommerceapiClientId() string`

GetBigcommerceapiClientId returns the BigcommerceapiClientId field if non-nil, zero value otherwise.

### GetBigcommerceapiClientIdOk

`func (o *AccountCartAdd) GetBigcommerceapiClientIdOk() (*string, bool)`

GetBigcommerceapiClientIdOk returns a tuple with the BigcommerceapiClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigcommerceapiClientId

`func (o *AccountCartAdd) SetBigcommerceapiClientId(v string)`

SetBigcommerceapiClientId sets BigcommerceapiClientId field to given value.

### HasBigcommerceapiClientId

`func (o *AccountCartAdd) HasBigcommerceapiClientId() bool`

HasBigcommerceapiClientId returns a boolean if a field has been set.

### GetBigcommerceapiAccessToken

`func (o *AccountCartAdd) GetBigcommerceapiAccessToken() string`

GetBigcommerceapiAccessToken returns the BigcommerceapiAccessToken field if non-nil, zero value otherwise.

### GetBigcommerceapiAccessTokenOk

`func (o *AccountCartAdd) GetBigcommerceapiAccessTokenOk() (*string, bool)`

GetBigcommerceapiAccessTokenOk returns a tuple with the BigcommerceapiAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigcommerceapiAccessToken

`func (o *AccountCartAdd) SetBigcommerceapiAccessToken(v string)`

SetBigcommerceapiAccessToken sets BigcommerceapiAccessToken field to given value.

### HasBigcommerceapiAccessToken

`func (o *AccountCartAdd) HasBigcommerceapiAccessToken() bool`

HasBigcommerceapiAccessToken returns a boolean if a field has been set.

### GetBigcommerceapiContext

`func (o *AccountCartAdd) GetBigcommerceapiContext() string`

GetBigcommerceapiContext returns the BigcommerceapiContext field if non-nil, zero value otherwise.

### GetBigcommerceapiContextOk

`func (o *AccountCartAdd) GetBigcommerceapiContextOk() (*string, bool)`

GetBigcommerceapiContextOk returns a tuple with the BigcommerceapiContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigcommerceapiContext

`func (o *AccountCartAdd) SetBigcommerceapiContext(v string)`

SetBigcommerceapiContext sets BigcommerceapiContext field to given value.

### HasBigcommerceapiContext

`func (o *AccountCartAdd) HasBigcommerceapiContext() bool`

HasBigcommerceapiContext returns a boolean if a field has been set.

### GetBolApiKey

`func (o *AccountCartAdd) GetBolApiKey() string`

GetBolApiKey returns the BolApiKey field if non-nil, zero value otherwise.

### GetBolApiKeyOk

`func (o *AccountCartAdd) GetBolApiKeyOk() (*string, bool)`

GetBolApiKeyOk returns a tuple with the BolApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBolApiKey

`func (o *AccountCartAdd) SetBolApiKey(v string)`

SetBolApiKey sets BolApiKey field to given value.

### HasBolApiKey

`func (o *AccountCartAdd) HasBolApiKey() bool`

HasBolApiKey returns a boolean if a field has been set.

### GetBolApiSecret

`func (o *AccountCartAdd) GetBolApiSecret() string`

GetBolApiSecret returns the BolApiSecret field if non-nil, zero value otherwise.

### GetBolApiSecretOk

`func (o *AccountCartAdd) GetBolApiSecretOk() (*string, bool)`

GetBolApiSecretOk returns a tuple with the BolApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBolApiSecret

`func (o *AccountCartAdd) SetBolApiSecret(v string)`

SetBolApiSecret sets BolApiSecret field to given value.

### HasBolApiSecret

`func (o *AccountCartAdd) HasBolApiSecret() bool`

HasBolApiSecret returns a boolean if a field has been set.

### GetBolRetailerId

`func (o *AccountCartAdd) GetBolRetailerId() int32`

GetBolRetailerId returns the BolRetailerId field if non-nil, zero value otherwise.

### GetBolRetailerIdOk

`func (o *AccountCartAdd) GetBolRetailerIdOk() (*int32, bool)`

GetBolRetailerIdOk returns a tuple with the BolRetailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBolRetailerId

`func (o *AccountCartAdd) SetBolRetailerId(v int32)`

SetBolRetailerId sets BolRetailerId field to given value.

### HasBolRetailerId

`func (o *AccountCartAdd) HasBolRetailerId() bool`

HasBolRetailerId returns a boolean if a field has been set.

### GetBigcartelUserName

`func (o *AccountCartAdd) GetBigcartelUserName() string`

GetBigcartelUserName returns the BigcartelUserName field if non-nil, zero value otherwise.

### GetBigcartelUserNameOk

`func (o *AccountCartAdd) GetBigcartelUserNameOk() (*string, bool)`

GetBigcartelUserNameOk returns a tuple with the BigcartelUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigcartelUserName

`func (o *AccountCartAdd) SetBigcartelUserName(v string)`

SetBigcartelUserName sets BigcartelUserName field to given value.


### GetBigcartelPassword

`func (o *AccountCartAdd) GetBigcartelPassword() string`

GetBigcartelPassword returns the BigcartelPassword field if non-nil, zero value otherwise.

### GetBigcartelPasswordOk

`func (o *AccountCartAdd) GetBigcartelPasswordOk() (*string, bool)`

GetBigcartelPasswordOk returns a tuple with the BigcartelPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigcartelPassword

`func (o *AccountCartAdd) SetBigcartelPassword(v string)`

SetBigcartelPassword sets BigcartelPassword field to given value.


### GetDemandwareClientId

`func (o *AccountCartAdd) GetDemandwareClientId() string`

GetDemandwareClientId returns the DemandwareClientId field if non-nil, zero value otherwise.

### GetDemandwareClientIdOk

`func (o *AccountCartAdd) GetDemandwareClientIdOk() (*string, bool)`

GetDemandwareClientIdOk returns a tuple with the DemandwareClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemandwareClientId

`func (o *AccountCartAdd) SetDemandwareClientId(v string)`

SetDemandwareClientId sets DemandwareClientId field to given value.

### HasDemandwareClientId

`func (o *AccountCartAdd) HasDemandwareClientId() bool`

HasDemandwareClientId returns a boolean if a field has been set.

### GetDemandwareApiPassword

`func (o *AccountCartAdd) GetDemandwareApiPassword() string`

GetDemandwareApiPassword returns the DemandwareApiPassword field if non-nil, zero value otherwise.

### GetDemandwareApiPasswordOk

`func (o *AccountCartAdd) GetDemandwareApiPasswordOk() (*string, bool)`

GetDemandwareApiPasswordOk returns a tuple with the DemandwareApiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemandwareApiPassword

`func (o *AccountCartAdd) SetDemandwareApiPassword(v string)`

SetDemandwareApiPassword sets DemandwareApiPassword field to given value.

### HasDemandwareApiPassword

`func (o *AccountCartAdd) HasDemandwareApiPassword() bool`

HasDemandwareApiPassword returns a boolean if a field has been set.

### GetDemandwareUserName

`func (o *AccountCartAdd) GetDemandwareUserName() string`

GetDemandwareUserName returns the DemandwareUserName field if non-nil, zero value otherwise.

### GetDemandwareUserNameOk

`func (o *AccountCartAdd) GetDemandwareUserNameOk() (*string, bool)`

GetDemandwareUserNameOk returns a tuple with the DemandwareUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemandwareUserName

`func (o *AccountCartAdd) SetDemandwareUserName(v string)`

SetDemandwareUserName sets DemandwareUserName field to given value.

### HasDemandwareUserName

`func (o *AccountCartAdd) HasDemandwareUserName() bool`

HasDemandwareUserName returns a boolean if a field has been set.

### GetDemandwareUserPassword

`func (o *AccountCartAdd) GetDemandwareUserPassword() string`

GetDemandwareUserPassword returns the DemandwareUserPassword field if non-nil, zero value otherwise.

### GetDemandwareUserPasswordOk

`func (o *AccountCartAdd) GetDemandwareUserPasswordOk() (*string, bool)`

GetDemandwareUserPasswordOk returns a tuple with the DemandwareUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemandwareUserPassword

`func (o *AccountCartAdd) SetDemandwareUserPassword(v string)`

SetDemandwareUserPassword sets DemandwareUserPassword field to given value.

### HasDemandwareUserPassword

`func (o *AccountCartAdd) HasDemandwareUserPassword() bool`

HasDemandwareUserPassword returns a boolean if a field has been set.

### GetEbayClientId

`func (o *AccountCartAdd) GetEbayClientId() string`

GetEbayClientId returns the EbayClientId field if non-nil, zero value otherwise.

### GetEbayClientIdOk

`func (o *AccountCartAdd) GetEbayClientIdOk() (*string, bool)`

GetEbayClientIdOk returns a tuple with the EbayClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayClientId

`func (o *AccountCartAdd) SetEbayClientId(v string)`

SetEbayClientId sets EbayClientId field to given value.

### HasEbayClientId

`func (o *AccountCartAdd) HasEbayClientId() bool`

HasEbayClientId returns a boolean if a field has been set.

### GetEbayClientSecret

`func (o *AccountCartAdd) GetEbayClientSecret() string`

GetEbayClientSecret returns the EbayClientSecret field if non-nil, zero value otherwise.

### GetEbayClientSecretOk

`func (o *AccountCartAdd) GetEbayClientSecretOk() (*string, bool)`

GetEbayClientSecretOk returns a tuple with the EbayClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayClientSecret

`func (o *AccountCartAdd) SetEbayClientSecret(v string)`

SetEbayClientSecret sets EbayClientSecret field to given value.

### HasEbayClientSecret

`func (o *AccountCartAdd) HasEbayClientSecret() bool`

HasEbayClientSecret returns a boolean if a field has been set.

### GetEbayRuname

`func (o *AccountCartAdd) GetEbayRuname() string`

GetEbayRuname returns the EbayRuname field if non-nil, zero value otherwise.

### GetEbayRunameOk

`func (o *AccountCartAdd) GetEbayRunameOk() (*string, bool)`

GetEbayRunameOk returns a tuple with the EbayRuname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayRuname

`func (o *AccountCartAdd) SetEbayRuname(v string)`

SetEbayRuname sets EbayRuname field to given value.

### HasEbayRuname

`func (o *AccountCartAdd) HasEbayRuname() bool`

HasEbayRuname returns a boolean if a field has been set.

### GetEbayAccessToken

`func (o *AccountCartAdd) GetEbayAccessToken() string`

GetEbayAccessToken returns the EbayAccessToken field if non-nil, zero value otherwise.

### GetEbayAccessTokenOk

`func (o *AccountCartAdd) GetEbayAccessTokenOk() (*string, bool)`

GetEbayAccessTokenOk returns a tuple with the EbayAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayAccessToken

`func (o *AccountCartAdd) SetEbayAccessToken(v string)`

SetEbayAccessToken sets EbayAccessToken field to given value.

### HasEbayAccessToken

`func (o *AccountCartAdd) HasEbayAccessToken() bool`

HasEbayAccessToken returns a boolean if a field has been set.

### GetEbayRefreshToken

`func (o *AccountCartAdd) GetEbayRefreshToken() string`

GetEbayRefreshToken returns the EbayRefreshToken field if non-nil, zero value otherwise.

### GetEbayRefreshTokenOk

`func (o *AccountCartAdd) GetEbayRefreshTokenOk() (*string, bool)`

GetEbayRefreshTokenOk returns a tuple with the EbayRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayRefreshToken

`func (o *AccountCartAdd) SetEbayRefreshToken(v string)`

SetEbayRefreshToken sets EbayRefreshToken field to given value.

### HasEbayRefreshToken

`func (o *AccountCartAdd) HasEbayRefreshToken() bool`

HasEbayRefreshToken returns a boolean if a field has been set.

### GetEbayEnvironment

`func (o *AccountCartAdd) GetEbayEnvironment() string`

GetEbayEnvironment returns the EbayEnvironment field if non-nil, zero value otherwise.

### GetEbayEnvironmentOk

`func (o *AccountCartAdd) GetEbayEnvironmentOk() (*string, bool)`

GetEbayEnvironmentOk returns a tuple with the EbayEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayEnvironment

`func (o *AccountCartAdd) SetEbayEnvironment(v string)`

SetEbayEnvironment sets EbayEnvironment field to given value.

### HasEbayEnvironment

`func (o *AccountCartAdd) HasEbayEnvironment() bool`

HasEbayEnvironment returns a boolean if a field has been set.

### GetEbaySiteId

`func (o *AccountCartAdd) GetEbaySiteId() int32`

GetEbaySiteId returns the EbaySiteId field if non-nil, zero value otherwise.

### GetEbaySiteIdOk

`func (o *AccountCartAdd) GetEbaySiteIdOk() (*int32, bool)`

GetEbaySiteIdOk returns a tuple with the EbaySiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbaySiteId

`func (o *AccountCartAdd) SetEbaySiteId(v int32)`

SetEbaySiteId sets EbaySiteId field to given value.

### HasEbaySiteId

`func (o *AccountCartAdd) HasEbaySiteId() bool`

HasEbaySiteId returns a boolean if a field has been set.

### GetWalmartClientId

`func (o *AccountCartAdd) GetWalmartClientId() string`

GetWalmartClientId returns the WalmartClientId field if non-nil, zero value otherwise.

### GetWalmartClientIdOk

`func (o *AccountCartAdd) GetWalmartClientIdOk() (*string, bool)`

GetWalmartClientIdOk returns a tuple with the WalmartClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalmartClientId

`func (o *AccountCartAdd) SetWalmartClientId(v string)`

SetWalmartClientId sets WalmartClientId field to given value.

### HasWalmartClientId

`func (o *AccountCartAdd) HasWalmartClientId() bool`

HasWalmartClientId returns a boolean if a field has been set.

### GetWalmartClientSecret

`func (o *AccountCartAdd) GetWalmartClientSecret() string`

GetWalmartClientSecret returns the WalmartClientSecret field if non-nil, zero value otherwise.

### GetWalmartClientSecretOk

`func (o *AccountCartAdd) GetWalmartClientSecretOk() (*string, bool)`

GetWalmartClientSecretOk returns a tuple with the WalmartClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalmartClientSecret

`func (o *AccountCartAdd) SetWalmartClientSecret(v string)`

SetWalmartClientSecret sets WalmartClientSecret field to given value.

### HasWalmartClientSecret

`func (o *AccountCartAdd) HasWalmartClientSecret() bool`

HasWalmartClientSecret returns a boolean if a field has been set.

### GetWalmartEnvironment

`func (o *AccountCartAdd) GetWalmartEnvironment() string`

GetWalmartEnvironment returns the WalmartEnvironment field if non-nil, zero value otherwise.

### GetWalmartEnvironmentOk

`func (o *AccountCartAdd) GetWalmartEnvironmentOk() (*string, bool)`

GetWalmartEnvironmentOk returns a tuple with the WalmartEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalmartEnvironment

`func (o *AccountCartAdd) SetWalmartEnvironment(v string)`

SetWalmartEnvironment sets WalmartEnvironment field to given value.

### HasWalmartEnvironment

`func (o *AccountCartAdd) HasWalmartEnvironment() bool`

HasWalmartEnvironment returns a boolean if a field has been set.

### GetWalmartChannelType

`func (o *AccountCartAdd) GetWalmartChannelType() string`

GetWalmartChannelType returns the WalmartChannelType field if non-nil, zero value otherwise.

### GetWalmartChannelTypeOk

`func (o *AccountCartAdd) GetWalmartChannelTypeOk() (*string, bool)`

GetWalmartChannelTypeOk returns a tuple with the WalmartChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalmartChannelType

`func (o *AccountCartAdd) SetWalmartChannelType(v string)`

SetWalmartChannelType sets WalmartChannelType field to given value.

### HasWalmartChannelType

`func (o *AccountCartAdd) HasWalmartChannelType() bool`

HasWalmartChannelType returns a boolean if a field has been set.

### GetWalmartRegion

`func (o *AccountCartAdd) GetWalmartRegion() string`

GetWalmartRegion returns the WalmartRegion field if non-nil, zero value otherwise.

### GetWalmartRegionOk

`func (o *AccountCartAdd) GetWalmartRegionOk() (*string, bool)`

GetWalmartRegionOk returns a tuple with the WalmartRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalmartRegion

`func (o *AccountCartAdd) SetWalmartRegion(v string)`

SetWalmartRegion sets WalmartRegion field to given value.

### HasWalmartRegion

`func (o *AccountCartAdd) HasWalmartRegion() bool`

HasWalmartRegion returns a boolean if a field has been set.

### GetEcwidAcessToken

`func (o *AccountCartAdd) GetEcwidAcessToken() string`

GetEcwidAcessToken returns the EcwidAcessToken field if non-nil, zero value otherwise.

### GetEcwidAcessTokenOk

`func (o *AccountCartAdd) GetEcwidAcessTokenOk() (*string, bool)`

GetEcwidAcessTokenOk returns a tuple with the EcwidAcessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcwidAcessToken

`func (o *AccountCartAdd) SetEcwidAcessToken(v string)`

SetEcwidAcessToken sets EcwidAcessToken field to given value.

### HasEcwidAcessToken

`func (o *AccountCartAdd) HasEcwidAcessToken() bool`

HasEcwidAcessToken returns a boolean if a field has been set.

### GetEcwidStoreId

`func (o *AccountCartAdd) GetEcwidStoreId() string`

GetEcwidStoreId returns the EcwidStoreId field if non-nil, zero value otherwise.

### GetEcwidStoreIdOk

`func (o *AccountCartAdd) GetEcwidStoreIdOk() (*string, bool)`

GetEcwidStoreIdOk returns a tuple with the EcwidStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcwidStoreId

`func (o *AccountCartAdd) SetEcwidStoreId(v string)`

SetEcwidStoreId sets EcwidStoreId field to given value.

### HasEcwidStoreId

`func (o *AccountCartAdd) HasEcwidStoreId() bool`

HasEcwidStoreId returns a boolean if a field has been set.

### GetLazadaAppId

`func (o *AccountCartAdd) GetLazadaAppId() string`

GetLazadaAppId returns the LazadaAppId field if non-nil, zero value otherwise.

### GetLazadaAppIdOk

`func (o *AccountCartAdd) GetLazadaAppIdOk() (*string, bool)`

GetLazadaAppIdOk returns a tuple with the LazadaAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLazadaAppId

`func (o *AccountCartAdd) SetLazadaAppId(v string)`

SetLazadaAppId sets LazadaAppId field to given value.

### HasLazadaAppId

`func (o *AccountCartAdd) HasLazadaAppId() bool`

HasLazadaAppId returns a boolean if a field has been set.

### GetLazadaAppSecret

`func (o *AccountCartAdd) GetLazadaAppSecret() string`

GetLazadaAppSecret returns the LazadaAppSecret field if non-nil, zero value otherwise.

### GetLazadaAppSecretOk

`func (o *AccountCartAdd) GetLazadaAppSecretOk() (*string, bool)`

GetLazadaAppSecretOk returns a tuple with the LazadaAppSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLazadaAppSecret

`func (o *AccountCartAdd) SetLazadaAppSecret(v string)`

SetLazadaAppSecret sets LazadaAppSecret field to given value.

### HasLazadaAppSecret

`func (o *AccountCartAdd) HasLazadaAppSecret() bool`

HasLazadaAppSecret returns a boolean if a field has been set.

### GetLazadaRefreshToken

`func (o *AccountCartAdd) GetLazadaRefreshToken() string`

GetLazadaRefreshToken returns the LazadaRefreshToken field if non-nil, zero value otherwise.

### GetLazadaRefreshTokenOk

`func (o *AccountCartAdd) GetLazadaRefreshTokenOk() (*string, bool)`

GetLazadaRefreshTokenOk returns a tuple with the LazadaRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLazadaRefreshToken

`func (o *AccountCartAdd) SetLazadaRefreshToken(v string)`

SetLazadaRefreshToken sets LazadaRefreshToken field to given value.

### HasLazadaRefreshToken

`func (o *AccountCartAdd) HasLazadaRefreshToken() bool`

HasLazadaRefreshToken returns a boolean if a field has been set.

### GetLazadaRegion

`func (o *AccountCartAdd) GetLazadaRegion() string`

GetLazadaRegion returns the LazadaRegion field if non-nil, zero value otherwise.

### GetLazadaRegionOk

`func (o *AccountCartAdd) GetLazadaRegionOk() (*string, bool)`

GetLazadaRegionOk returns a tuple with the LazadaRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLazadaRegion

`func (o *AccountCartAdd) SetLazadaRegion(v string)`

SetLazadaRegion sets LazadaRegion field to given value.

### HasLazadaRegion

`func (o *AccountCartAdd) HasLazadaRegion() bool`

HasLazadaRegion returns a boolean if a field has been set.

### GetLightspeedApiKey

`func (o *AccountCartAdd) GetLightspeedApiKey() string`

GetLightspeedApiKey returns the LightspeedApiKey field if non-nil, zero value otherwise.

### GetLightspeedApiKeyOk

`func (o *AccountCartAdd) GetLightspeedApiKeyOk() (*string, bool)`

GetLightspeedApiKeyOk returns a tuple with the LightspeedApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLightspeedApiKey

`func (o *AccountCartAdd) SetLightspeedApiKey(v string)`

SetLightspeedApiKey sets LightspeedApiKey field to given value.

### HasLightspeedApiKey

`func (o *AccountCartAdd) HasLightspeedApiKey() bool`

HasLightspeedApiKey returns a boolean if a field has been set.

### GetLightspeedApiSecret

`func (o *AccountCartAdd) GetLightspeedApiSecret() string`

GetLightspeedApiSecret returns the LightspeedApiSecret field if non-nil, zero value otherwise.

### GetLightspeedApiSecretOk

`func (o *AccountCartAdd) GetLightspeedApiSecretOk() (*string, bool)`

GetLightspeedApiSecretOk returns a tuple with the LightspeedApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLightspeedApiSecret

`func (o *AccountCartAdd) SetLightspeedApiSecret(v string)`

SetLightspeedApiSecret sets LightspeedApiSecret field to given value.

### HasLightspeedApiSecret

`func (o *AccountCartAdd) HasLightspeedApiSecret() bool`

HasLightspeedApiSecret returns a boolean if a field has been set.

### GetEtsyKeystring

`func (o *AccountCartAdd) GetEtsyKeystring() string`

GetEtsyKeystring returns the EtsyKeystring field if non-nil, zero value otherwise.

### GetEtsyKeystringOk

`func (o *AccountCartAdd) GetEtsyKeystringOk() (*string, bool)`

GetEtsyKeystringOk returns a tuple with the EtsyKeystring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtsyKeystring

`func (o *AccountCartAdd) SetEtsyKeystring(v string)`

SetEtsyKeystring sets EtsyKeystring field to given value.

### HasEtsyKeystring

`func (o *AccountCartAdd) HasEtsyKeystring() bool`

HasEtsyKeystring returns a boolean if a field has been set.

### GetEtsySharedSecret

`func (o *AccountCartAdd) GetEtsySharedSecret() string`

GetEtsySharedSecret returns the EtsySharedSecret field if non-nil, zero value otherwise.

### GetEtsySharedSecretOk

`func (o *AccountCartAdd) GetEtsySharedSecretOk() (*string, bool)`

GetEtsySharedSecretOk returns a tuple with the EtsySharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtsySharedSecret

`func (o *AccountCartAdd) SetEtsySharedSecret(v string)`

SetEtsySharedSecret sets EtsySharedSecret field to given value.

### HasEtsySharedSecret

`func (o *AccountCartAdd) HasEtsySharedSecret() bool`

HasEtsySharedSecret returns a boolean if a field has been set.

### GetEtsyAccessToken

`func (o *AccountCartAdd) GetEtsyAccessToken() string`

GetEtsyAccessToken returns the EtsyAccessToken field if non-nil, zero value otherwise.

### GetEtsyAccessTokenOk

`func (o *AccountCartAdd) GetEtsyAccessTokenOk() (*string, bool)`

GetEtsyAccessTokenOk returns a tuple with the EtsyAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtsyAccessToken

`func (o *AccountCartAdd) SetEtsyAccessToken(v string)`

SetEtsyAccessToken sets EtsyAccessToken field to given value.

### HasEtsyAccessToken

`func (o *AccountCartAdd) HasEtsyAccessToken() bool`

HasEtsyAccessToken returns a boolean if a field has been set.

### GetEtsyTokenSecret

`func (o *AccountCartAdd) GetEtsyTokenSecret() string`

GetEtsyTokenSecret returns the EtsyTokenSecret field if non-nil, zero value otherwise.

### GetEtsyTokenSecretOk

`func (o *AccountCartAdd) GetEtsyTokenSecretOk() (*string, bool)`

GetEtsyTokenSecretOk returns a tuple with the EtsyTokenSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtsyTokenSecret

`func (o *AccountCartAdd) SetEtsyTokenSecret(v string)`

SetEtsyTokenSecret sets EtsyTokenSecret field to given value.

### HasEtsyTokenSecret

`func (o *AccountCartAdd) HasEtsyTokenSecret() bool`

HasEtsyTokenSecret returns a boolean if a field has been set.

### GetEtsyClientId

`func (o *AccountCartAdd) GetEtsyClientId() string`

GetEtsyClientId returns the EtsyClientId field if non-nil, zero value otherwise.

### GetEtsyClientIdOk

`func (o *AccountCartAdd) GetEtsyClientIdOk() (*string, bool)`

GetEtsyClientIdOk returns a tuple with the EtsyClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtsyClientId

`func (o *AccountCartAdd) SetEtsyClientId(v string)`

SetEtsyClientId sets EtsyClientId field to given value.

### HasEtsyClientId

`func (o *AccountCartAdd) HasEtsyClientId() bool`

HasEtsyClientId returns a boolean if a field has been set.

### GetEtsyRefreshToken

`func (o *AccountCartAdd) GetEtsyRefreshToken() string`

GetEtsyRefreshToken returns the EtsyRefreshToken field if non-nil, zero value otherwise.

### GetEtsyRefreshTokenOk

`func (o *AccountCartAdd) GetEtsyRefreshTokenOk() (*string, bool)`

GetEtsyRefreshTokenOk returns a tuple with the EtsyRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtsyRefreshToken

`func (o *AccountCartAdd) SetEtsyRefreshToken(v string)`

SetEtsyRefreshToken sets EtsyRefreshToken field to given value.

### HasEtsyRefreshToken

`func (o *AccountCartAdd) HasEtsyRefreshToken() bool`

HasEtsyRefreshToken returns a boolean if a field has been set.

### GetFacebookAppId

`func (o *AccountCartAdd) GetFacebookAppId() string`

GetFacebookAppId returns the FacebookAppId field if non-nil, zero value otherwise.

### GetFacebookAppIdOk

`func (o *AccountCartAdd) GetFacebookAppIdOk() (*string, bool)`

GetFacebookAppIdOk returns a tuple with the FacebookAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookAppId

`func (o *AccountCartAdd) SetFacebookAppId(v string)`

SetFacebookAppId sets FacebookAppId field to given value.

### HasFacebookAppId

`func (o *AccountCartAdd) HasFacebookAppId() bool`

HasFacebookAppId returns a boolean if a field has been set.

### GetFacebookAppSecret

`func (o *AccountCartAdd) GetFacebookAppSecret() string`

GetFacebookAppSecret returns the FacebookAppSecret field if non-nil, zero value otherwise.

### GetFacebookAppSecretOk

`func (o *AccountCartAdd) GetFacebookAppSecretOk() (*string, bool)`

GetFacebookAppSecretOk returns a tuple with the FacebookAppSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookAppSecret

`func (o *AccountCartAdd) SetFacebookAppSecret(v string)`

SetFacebookAppSecret sets FacebookAppSecret field to given value.

### HasFacebookAppSecret

`func (o *AccountCartAdd) HasFacebookAppSecret() bool`

HasFacebookAppSecret returns a boolean if a field has been set.

### GetFacebookAccessToken

`func (o *AccountCartAdd) GetFacebookAccessToken() string`

GetFacebookAccessToken returns the FacebookAccessToken field if non-nil, zero value otherwise.

### GetFacebookAccessTokenOk

`func (o *AccountCartAdd) GetFacebookAccessTokenOk() (*string, bool)`

GetFacebookAccessTokenOk returns a tuple with the FacebookAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookAccessToken

`func (o *AccountCartAdd) SetFacebookAccessToken(v string)`

SetFacebookAccessToken sets FacebookAccessToken field to given value.

### HasFacebookAccessToken

`func (o *AccountCartAdd) HasFacebookAccessToken() bool`

HasFacebookAccessToken returns a boolean if a field has been set.

### GetFacebookBusinessId

`func (o *AccountCartAdd) GetFacebookBusinessId() string`

GetFacebookBusinessId returns the FacebookBusinessId field if non-nil, zero value otherwise.

### GetFacebookBusinessIdOk

`func (o *AccountCartAdd) GetFacebookBusinessIdOk() (*string, bool)`

GetFacebookBusinessIdOk returns a tuple with the FacebookBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookBusinessId

`func (o *AccountCartAdd) SetFacebookBusinessId(v string)`

SetFacebookBusinessId sets FacebookBusinessId field to given value.

### HasFacebookBusinessId

`func (o *AccountCartAdd) HasFacebookBusinessId() bool`

HasFacebookBusinessId returns a boolean if a field has been set.

### GetNetoApiKey

`func (o *AccountCartAdd) GetNetoApiKey() string`

GetNetoApiKey returns the NetoApiKey field if non-nil, zero value otherwise.

### GetNetoApiKeyOk

`func (o *AccountCartAdd) GetNetoApiKeyOk() (*string, bool)`

GetNetoApiKeyOk returns a tuple with the NetoApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetoApiKey

`func (o *AccountCartAdd) SetNetoApiKey(v string)`

SetNetoApiKey sets NetoApiKey field to given value.

### HasNetoApiKey

`func (o *AccountCartAdd) HasNetoApiKey() bool`

HasNetoApiKey returns a boolean if a field has been set.

### GetNetoApiUsername

`func (o *AccountCartAdd) GetNetoApiUsername() string`

GetNetoApiUsername returns the NetoApiUsername field if non-nil, zero value otherwise.

### GetNetoApiUsernameOk

`func (o *AccountCartAdd) GetNetoApiUsernameOk() (*string, bool)`

GetNetoApiUsernameOk returns a tuple with the NetoApiUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetoApiUsername

`func (o *AccountCartAdd) SetNetoApiUsername(v string)`

SetNetoApiUsername sets NetoApiUsername field to given value.

### HasNetoApiUsername

`func (o *AccountCartAdd) HasNetoApiUsername() bool`

HasNetoApiUsername returns a boolean if a field has been set.

### GetShoplineAccessToken

`func (o *AccountCartAdd) GetShoplineAccessToken() string`

GetShoplineAccessToken returns the ShoplineAccessToken field if non-nil, zero value otherwise.

### GetShoplineAccessTokenOk

`func (o *AccountCartAdd) GetShoplineAccessTokenOk() (*string, bool)`

GetShoplineAccessTokenOk returns a tuple with the ShoplineAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShoplineAccessToken

`func (o *AccountCartAdd) SetShoplineAccessToken(v string)`

SetShoplineAccessToken sets ShoplineAccessToken field to given value.

### HasShoplineAccessToken

`func (o *AccountCartAdd) HasShoplineAccessToken() bool`

HasShoplineAccessToken returns a boolean if a field has been set.

### GetShoplineAppKey

`func (o *AccountCartAdd) GetShoplineAppKey() string`

GetShoplineAppKey returns the ShoplineAppKey field if non-nil, zero value otherwise.

### GetShoplineAppKeyOk

`func (o *AccountCartAdd) GetShoplineAppKeyOk() (*string, bool)`

GetShoplineAppKeyOk returns a tuple with the ShoplineAppKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShoplineAppKey

`func (o *AccountCartAdd) SetShoplineAppKey(v string)`

SetShoplineAppKey sets ShoplineAppKey field to given value.

### HasShoplineAppKey

`func (o *AccountCartAdd) HasShoplineAppKey() bool`

HasShoplineAppKey returns a boolean if a field has been set.

### GetShoplineAppSecret

`func (o *AccountCartAdd) GetShoplineAppSecret() string`

GetShoplineAppSecret returns the ShoplineAppSecret field if non-nil, zero value otherwise.

### GetShoplineAppSecretOk

`func (o *AccountCartAdd) GetShoplineAppSecretOk() (*string, bool)`

GetShoplineAppSecretOk returns a tuple with the ShoplineAppSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShoplineAppSecret

`func (o *AccountCartAdd) SetShoplineAppSecret(v string)`

SetShoplineAppSecret sets ShoplineAppSecret field to given value.

### HasShoplineAppSecret

`func (o *AccountCartAdd) HasShoplineAppSecret() bool`

HasShoplineAppSecret returns a boolean if a field has been set.

### GetShoplineSharedSecret

`func (o *AccountCartAdd) GetShoplineSharedSecret() string`

GetShoplineSharedSecret returns the ShoplineSharedSecret field if non-nil, zero value otherwise.

### GetShoplineSharedSecretOk

`func (o *AccountCartAdd) GetShoplineSharedSecretOk() (*string, bool)`

GetShoplineSharedSecretOk returns a tuple with the ShoplineSharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShoplineSharedSecret

`func (o *AccountCartAdd) SetShoplineSharedSecret(v string)`

SetShoplineSharedSecret sets ShoplineSharedSecret field to given value.

### HasShoplineSharedSecret

`func (o *AccountCartAdd) HasShoplineSharedSecret() bool`

HasShoplineSharedSecret returns a boolean if a field has been set.

### GetShopifyAccessToken

`func (o *AccountCartAdd) GetShopifyAccessToken() string`

GetShopifyAccessToken returns the ShopifyAccessToken field if non-nil, zero value otherwise.

### GetShopifyAccessTokenOk

`func (o *AccountCartAdd) GetShopifyAccessTokenOk() (*string, bool)`

GetShopifyAccessTokenOk returns a tuple with the ShopifyAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopifyAccessToken

`func (o *AccountCartAdd) SetShopifyAccessToken(v string)`

SetShopifyAccessToken sets ShopifyAccessToken field to given value.

### HasShopifyAccessToken

`func (o *AccountCartAdd) HasShopifyAccessToken() bool`

HasShopifyAccessToken returns a boolean if a field has been set.

### GetShopifyApiKey

`func (o *AccountCartAdd) GetShopifyApiKey() string`

GetShopifyApiKey returns the ShopifyApiKey field if non-nil, zero value otherwise.

### GetShopifyApiKeyOk

`func (o *AccountCartAdd) GetShopifyApiKeyOk() (*string, bool)`

GetShopifyApiKeyOk returns a tuple with the ShopifyApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopifyApiKey

`func (o *AccountCartAdd) SetShopifyApiKey(v string)`

SetShopifyApiKey sets ShopifyApiKey field to given value.

### HasShopifyApiKey

`func (o *AccountCartAdd) HasShopifyApiKey() bool`

HasShopifyApiKey returns a boolean if a field has been set.

### GetShopifyApiPassword

`func (o *AccountCartAdd) GetShopifyApiPassword() string`

GetShopifyApiPassword returns the ShopifyApiPassword field if non-nil, zero value otherwise.

### GetShopifyApiPasswordOk

`func (o *AccountCartAdd) GetShopifyApiPasswordOk() (*string, bool)`

GetShopifyApiPasswordOk returns a tuple with the ShopifyApiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopifyApiPassword

`func (o *AccountCartAdd) SetShopifyApiPassword(v string)`

SetShopifyApiPassword sets ShopifyApiPassword field to given value.

### HasShopifyApiPassword

`func (o *AccountCartAdd) HasShopifyApiPassword() bool`

HasShopifyApiPassword returns a boolean if a field has been set.

### GetShopifySharedSecret

`func (o *AccountCartAdd) GetShopifySharedSecret() string`

GetShopifySharedSecret returns the ShopifySharedSecret field if non-nil, zero value otherwise.

### GetShopifySharedSecretOk

`func (o *AccountCartAdd) GetShopifySharedSecretOk() (*string, bool)`

GetShopifySharedSecretOk returns a tuple with the ShopifySharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopifySharedSecret

`func (o *AccountCartAdd) SetShopifySharedSecret(v string)`

SetShopifySharedSecret sets ShopifySharedSecret field to given value.

### HasShopifySharedSecret

`func (o *AccountCartAdd) HasShopifySharedSecret() bool`

HasShopifySharedSecret returns a boolean if a field has been set.

### GetShoplazzaAccessToken

`func (o *AccountCartAdd) GetShoplazzaAccessToken() string`

GetShoplazzaAccessToken returns the ShoplazzaAccessToken field if non-nil, zero value otherwise.

### GetShoplazzaAccessTokenOk

`func (o *AccountCartAdd) GetShoplazzaAccessTokenOk() (*string, bool)`

GetShoplazzaAccessTokenOk returns a tuple with the ShoplazzaAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShoplazzaAccessToken

`func (o *AccountCartAdd) SetShoplazzaAccessToken(v string)`

SetShoplazzaAccessToken sets ShoplazzaAccessToken field to given value.

### HasShoplazzaAccessToken

`func (o *AccountCartAdd) HasShoplazzaAccessToken() bool`

HasShoplazzaAccessToken returns a boolean if a field has been set.

### GetShoplazzaSharedSecret

`func (o *AccountCartAdd) GetShoplazzaSharedSecret() string`

GetShoplazzaSharedSecret returns the ShoplazzaSharedSecret field if non-nil, zero value otherwise.

### GetShoplazzaSharedSecretOk

`func (o *AccountCartAdd) GetShoplazzaSharedSecretOk() (*string, bool)`

GetShoplazzaSharedSecretOk returns a tuple with the ShoplazzaSharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShoplazzaSharedSecret

`func (o *AccountCartAdd) SetShoplazzaSharedSecret(v string)`

SetShoplazzaSharedSecret sets ShoplazzaSharedSecret field to given value.

### HasShoplazzaSharedSecret

`func (o *AccountCartAdd) HasShoplazzaSharedSecret() bool`

HasShoplazzaSharedSecret returns a boolean if a field has been set.

### GetShopwareAccessKey

`func (o *AccountCartAdd) GetShopwareAccessKey() string`

GetShopwareAccessKey returns the ShopwareAccessKey field if non-nil, zero value otherwise.

### GetShopwareAccessKeyOk

`func (o *AccountCartAdd) GetShopwareAccessKeyOk() (*string, bool)`

GetShopwareAccessKeyOk returns a tuple with the ShopwareAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopwareAccessKey

`func (o *AccountCartAdd) SetShopwareAccessKey(v string)`

SetShopwareAccessKey sets ShopwareAccessKey field to given value.

### HasShopwareAccessKey

`func (o *AccountCartAdd) HasShopwareAccessKey() bool`

HasShopwareAccessKey returns a boolean if a field has been set.

### GetShopwareApiKey

`func (o *AccountCartAdd) GetShopwareApiKey() string`

GetShopwareApiKey returns the ShopwareApiKey field if non-nil, zero value otherwise.

### GetShopwareApiKeyOk

`func (o *AccountCartAdd) GetShopwareApiKeyOk() (*string, bool)`

GetShopwareApiKeyOk returns a tuple with the ShopwareApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopwareApiKey

`func (o *AccountCartAdd) SetShopwareApiKey(v string)`

SetShopwareApiKey sets ShopwareApiKey field to given value.

### HasShopwareApiKey

`func (o *AccountCartAdd) HasShopwareApiKey() bool`

HasShopwareApiKey returns a boolean if a field has been set.

### GetShopwareApiSecret

`func (o *AccountCartAdd) GetShopwareApiSecret() string`

GetShopwareApiSecret returns the ShopwareApiSecret field if non-nil, zero value otherwise.

### GetShopwareApiSecretOk

`func (o *AccountCartAdd) GetShopwareApiSecretOk() (*string, bool)`

GetShopwareApiSecretOk returns a tuple with the ShopwareApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopwareApiSecret

`func (o *AccountCartAdd) SetShopwareApiSecret(v string)`

SetShopwareApiSecret sets ShopwareApiSecret field to given value.

### HasShopwareApiSecret

`func (o *AccountCartAdd) HasShopwareApiSecret() bool`

HasShopwareApiSecret returns a boolean if a field has been set.

### GetMivaAccessToken

`func (o *AccountCartAdd) GetMivaAccessToken() string`

GetMivaAccessToken returns the MivaAccessToken field if non-nil, zero value otherwise.

### GetMivaAccessTokenOk

`func (o *AccountCartAdd) GetMivaAccessTokenOk() (*string, bool)`

GetMivaAccessTokenOk returns a tuple with the MivaAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMivaAccessToken

`func (o *AccountCartAdd) SetMivaAccessToken(v string)`

SetMivaAccessToken sets MivaAccessToken field to given value.

### HasMivaAccessToken

`func (o *AccountCartAdd) HasMivaAccessToken() bool`

HasMivaAccessToken returns a boolean if a field has been set.

### GetMivaSignature

`func (o *AccountCartAdd) GetMivaSignature() string`

GetMivaSignature returns the MivaSignature field if non-nil, zero value otherwise.

### GetMivaSignatureOk

`func (o *AccountCartAdd) GetMivaSignatureOk() (*string, bool)`

GetMivaSignatureOk returns a tuple with the MivaSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMivaSignature

`func (o *AccountCartAdd) SetMivaSignature(v string)`

SetMivaSignature sets MivaSignature field to given value.

### HasMivaSignature

`func (o *AccountCartAdd) HasMivaSignature() bool`

HasMivaSignature returns a boolean if a field has been set.

### GetTiendanubeUserId

`func (o *AccountCartAdd) GetTiendanubeUserId() int32`

GetTiendanubeUserId returns the TiendanubeUserId field if non-nil, zero value otherwise.

### GetTiendanubeUserIdOk

`func (o *AccountCartAdd) GetTiendanubeUserIdOk() (*int32, bool)`

GetTiendanubeUserIdOk returns a tuple with the TiendanubeUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiendanubeUserId

`func (o *AccountCartAdd) SetTiendanubeUserId(v int32)`

SetTiendanubeUserId sets TiendanubeUserId field to given value.

### HasTiendanubeUserId

`func (o *AccountCartAdd) HasTiendanubeUserId() bool`

HasTiendanubeUserId returns a boolean if a field has been set.

### GetTiendanubeAccessToken

`func (o *AccountCartAdd) GetTiendanubeAccessToken() string`

GetTiendanubeAccessToken returns the TiendanubeAccessToken field if non-nil, zero value otherwise.

### GetTiendanubeAccessTokenOk

`func (o *AccountCartAdd) GetTiendanubeAccessTokenOk() (*string, bool)`

GetTiendanubeAccessTokenOk returns a tuple with the TiendanubeAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiendanubeAccessToken

`func (o *AccountCartAdd) SetTiendanubeAccessToken(v string)`

SetTiendanubeAccessToken sets TiendanubeAccessToken field to given value.

### HasTiendanubeAccessToken

`func (o *AccountCartAdd) HasTiendanubeAccessToken() bool`

HasTiendanubeAccessToken returns a boolean if a field has been set.

### GetTiendanubeClientSecret

`func (o *AccountCartAdd) GetTiendanubeClientSecret() string`

GetTiendanubeClientSecret returns the TiendanubeClientSecret field if non-nil, zero value otherwise.

### GetTiendanubeClientSecretOk

`func (o *AccountCartAdd) GetTiendanubeClientSecretOk() (*string, bool)`

GetTiendanubeClientSecretOk returns a tuple with the TiendanubeClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiendanubeClientSecret

`func (o *AccountCartAdd) SetTiendanubeClientSecret(v string)`

SetTiendanubeClientSecret sets TiendanubeClientSecret field to given value.

### HasTiendanubeClientSecret

`func (o *AccountCartAdd) HasTiendanubeClientSecret() bool`

HasTiendanubeClientSecret returns a boolean if a field has been set.

### GetVolusionLogin

`func (o *AccountCartAdd) GetVolusionLogin() string`

GetVolusionLogin returns the VolusionLogin field if non-nil, zero value otherwise.

### GetVolusionLoginOk

`func (o *AccountCartAdd) GetVolusionLoginOk() (*string, bool)`

GetVolusionLoginOk returns a tuple with the VolusionLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolusionLogin

`func (o *AccountCartAdd) SetVolusionLogin(v string)`

SetVolusionLogin sets VolusionLogin field to given value.

### HasVolusionLogin

`func (o *AccountCartAdd) HasVolusionLogin() bool`

HasVolusionLogin returns a boolean if a field has been set.

### GetVolusionPassword

`func (o *AccountCartAdd) GetVolusionPassword() string`

GetVolusionPassword returns the VolusionPassword field if non-nil, zero value otherwise.

### GetVolusionPasswordOk

`func (o *AccountCartAdd) GetVolusionPasswordOk() (*string, bool)`

GetVolusionPasswordOk returns a tuple with the VolusionPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolusionPassword

`func (o *AccountCartAdd) SetVolusionPassword(v string)`

SetVolusionPassword sets VolusionPassword field to given value.

### HasVolusionPassword

`func (o *AccountCartAdd) HasVolusionPassword() bool`

HasVolusionPassword returns a boolean if a field has been set.

### GetHybrisClientId

`func (o *AccountCartAdd) GetHybrisClientId() string`

GetHybrisClientId returns the HybrisClientId field if non-nil, zero value otherwise.

### GetHybrisClientIdOk

`func (o *AccountCartAdd) GetHybrisClientIdOk() (*string, bool)`

GetHybrisClientIdOk returns a tuple with the HybrisClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybrisClientId

`func (o *AccountCartAdd) SetHybrisClientId(v string)`

SetHybrisClientId sets HybrisClientId field to given value.

### HasHybrisClientId

`func (o *AccountCartAdd) HasHybrisClientId() bool`

HasHybrisClientId returns a boolean if a field has been set.

### GetHybrisClientSecret

`func (o *AccountCartAdd) GetHybrisClientSecret() string`

GetHybrisClientSecret returns the HybrisClientSecret field if non-nil, zero value otherwise.

### GetHybrisClientSecretOk

`func (o *AccountCartAdd) GetHybrisClientSecretOk() (*string, bool)`

GetHybrisClientSecretOk returns a tuple with the HybrisClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybrisClientSecret

`func (o *AccountCartAdd) SetHybrisClientSecret(v string)`

SetHybrisClientSecret sets HybrisClientSecret field to given value.

### HasHybrisClientSecret

`func (o *AccountCartAdd) HasHybrisClientSecret() bool`

HasHybrisClientSecret returns a boolean if a field has been set.

### GetHybrisUsername

`func (o *AccountCartAdd) GetHybrisUsername() string`

GetHybrisUsername returns the HybrisUsername field if non-nil, zero value otherwise.

### GetHybrisUsernameOk

`func (o *AccountCartAdd) GetHybrisUsernameOk() (*string, bool)`

GetHybrisUsernameOk returns a tuple with the HybrisUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybrisUsername

`func (o *AccountCartAdd) SetHybrisUsername(v string)`

SetHybrisUsername sets HybrisUsername field to given value.

### HasHybrisUsername

`func (o *AccountCartAdd) HasHybrisUsername() bool`

HasHybrisUsername returns a boolean if a field has been set.

### GetHybrisPassword

`func (o *AccountCartAdd) GetHybrisPassword() string`

GetHybrisPassword returns the HybrisPassword field if non-nil, zero value otherwise.

### GetHybrisPasswordOk

`func (o *AccountCartAdd) GetHybrisPasswordOk() (*string, bool)`

GetHybrisPasswordOk returns a tuple with the HybrisPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybrisPassword

`func (o *AccountCartAdd) SetHybrisPassword(v string)`

SetHybrisPassword sets HybrisPassword field to given value.

### HasHybrisPassword

`func (o *AccountCartAdd) HasHybrisPassword() bool`

HasHybrisPassword returns a boolean if a field has been set.

### GetHybrisWebsites

`func (o *AccountCartAdd) GetHybrisWebsites() []AccountCartAddHybrisWebsitesInner`

GetHybrisWebsites returns the HybrisWebsites field if non-nil, zero value otherwise.

### GetHybrisWebsitesOk

`func (o *AccountCartAdd) GetHybrisWebsitesOk() (*[]AccountCartAddHybrisWebsitesInner, bool)`

GetHybrisWebsitesOk returns a tuple with the HybrisWebsites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybrisWebsites

`func (o *AccountCartAdd) SetHybrisWebsites(v []AccountCartAddHybrisWebsitesInner)`

SetHybrisWebsites sets HybrisWebsites field to given value.

### HasHybrisWebsites

`func (o *AccountCartAdd) HasHybrisWebsites() bool`

HasHybrisWebsites returns a boolean if a field has been set.

### GetSquareClientId

`func (o *AccountCartAdd) GetSquareClientId() string`

GetSquareClientId returns the SquareClientId field if non-nil, zero value otherwise.

### GetSquareClientIdOk

`func (o *AccountCartAdd) GetSquareClientIdOk() (*string, bool)`

GetSquareClientIdOk returns a tuple with the SquareClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquareClientId

`func (o *AccountCartAdd) SetSquareClientId(v string)`

SetSquareClientId sets SquareClientId field to given value.

### HasSquareClientId

`func (o *AccountCartAdd) HasSquareClientId() bool`

HasSquareClientId returns a boolean if a field has been set.

### GetSquareClientSecret

`func (o *AccountCartAdd) GetSquareClientSecret() string`

GetSquareClientSecret returns the SquareClientSecret field if non-nil, zero value otherwise.

### GetSquareClientSecretOk

`func (o *AccountCartAdd) GetSquareClientSecretOk() (*string, bool)`

GetSquareClientSecretOk returns a tuple with the SquareClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquareClientSecret

`func (o *AccountCartAdd) SetSquareClientSecret(v string)`

SetSquareClientSecret sets SquareClientSecret field to given value.

### HasSquareClientSecret

`func (o *AccountCartAdd) HasSquareClientSecret() bool`

HasSquareClientSecret returns a boolean if a field has been set.

### GetSquareRefreshToken

`func (o *AccountCartAdd) GetSquareRefreshToken() string`

GetSquareRefreshToken returns the SquareRefreshToken field if non-nil, zero value otherwise.

### GetSquareRefreshTokenOk

`func (o *AccountCartAdd) GetSquareRefreshTokenOk() (*string, bool)`

GetSquareRefreshTokenOk returns a tuple with the SquareRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquareRefreshToken

`func (o *AccountCartAdd) SetSquareRefreshToken(v string)`

SetSquareRefreshToken sets SquareRefreshToken field to given value.

### HasSquareRefreshToken

`func (o *AccountCartAdd) HasSquareRefreshToken() bool`

HasSquareRefreshToken returns a boolean if a field has been set.

### GetSquarespaceApiKey

`func (o *AccountCartAdd) GetSquarespaceApiKey() string`

GetSquarespaceApiKey returns the SquarespaceApiKey field if non-nil, zero value otherwise.

### GetSquarespaceApiKeyOk

`func (o *AccountCartAdd) GetSquarespaceApiKeyOk() (*string, bool)`

GetSquarespaceApiKeyOk returns a tuple with the SquarespaceApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquarespaceApiKey

`func (o *AccountCartAdd) SetSquarespaceApiKey(v string)`

SetSquarespaceApiKey sets SquarespaceApiKey field to given value.

### HasSquarespaceApiKey

`func (o *AccountCartAdd) HasSquarespaceApiKey() bool`

HasSquarespaceApiKey returns a boolean if a field has been set.

### GetSquarespaceClientId

`func (o *AccountCartAdd) GetSquarespaceClientId() string`

GetSquarespaceClientId returns the SquarespaceClientId field if non-nil, zero value otherwise.

### GetSquarespaceClientIdOk

`func (o *AccountCartAdd) GetSquarespaceClientIdOk() (*string, bool)`

GetSquarespaceClientIdOk returns a tuple with the SquarespaceClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquarespaceClientId

`func (o *AccountCartAdd) SetSquarespaceClientId(v string)`

SetSquarespaceClientId sets SquarespaceClientId field to given value.

### HasSquarespaceClientId

`func (o *AccountCartAdd) HasSquarespaceClientId() bool`

HasSquarespaceClientId returns a boolean if a field has been set.

### GetSquarespaceClientSecret

`func (o *AccountCartAdd) GetSquarespaceClientSecret() string`

GetSquarespaceClientSecret returns the SquarespaceClientSecret field if non-nil, zero value otherwise.

### GetSquarespaceClientSecretOk

`func (o *AccountCartAdd) GetSquarespaceClientSecretOk() (*string, bool)`

GetSquarespaceClientSecretOk returns a tuple with the SquarespaceClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquarespaceClientSecret

`func (o *AccountCartAdd) SetSquarespaceClientSecret(v string)`

SetSquarespaceClientSecret sets SquarespaceClientSecret field to given value.

### HasSquarespaceClientSecret

`func (o *AccountCartAdd) HasSquarespaceClientSecret() bool`

HasSquarespaceClientSecret returns a boolean if a field has been set.

### GetSquarespaceAccessToken

`func (o *AccountCartAdd) GetSquarespaceAccessToken() string`

GetSquarespaceAccessToken returns the SquarespaceAccessToken field if non-nil, zero value otherwise.

### GetSquarespaceAccessTokenOk

`func (o *AccountCartAdd) GetSquarespaceAccessTokenOk() (*string, bool)`

GetSquarespaceAccessTokenOk returns a tuple with the SquarespaceAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquarespaceAccessToken

`func (o *AccountCartAdd) SetSquarespaceAccessToken(v string)`

SetSquarespaceAccessToken sets SquarespaceAccessToken field to given value.

### HasSquarespaceAccessToken

`func (o *AccountCartAdd) HasSquarespaceAccessToken() bool`

HasSquarespaceAccessToken returns a boolean if a field has been set.

### GetSquarespaceRefreshToken

`func (o *AccountCartAdd) GetSquarespaceRefreshToken() string`

GetSquarespaceRefreshToken returns the SquarespaceRefreshToken field if non-nil, zero value otherwise.

### GetSquarespaceRefreshTokenOk

`func (o *AccountCartAdd) GetSquarespaceRefreshTokenOk() (*string, bool)`

GetSquarespaceRefreshTokenOk returns a tuple with the SquarespaceRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquarespaceRefreshToken

`func (o *AccountCartAdd) SetSquarespaceRefreshToken(v string)`

SetSquarespaceRefreshToken sets SquarespaceRefreshToken field to given value.

### HasSquarespaceRefreshToken

`func (o *AccountCartAdd) HasSquarespaceRefreshToken() bool`

HasSquarespaceRefreshToken returns a boolean if a field has been set.

### GetCommercehqApiKey

`func (o *AccountCartAdd) GetCommercehqApiKey() string`

GetCommercehqApiKey returns the CommercehqApiKey field if non-nil, zero value otherwise.

### GetCommercehqApiKeyOk

`func (o *AccountCartAdd) GetCommercehqApiKeyOk() (*string, bool)`

GetCommercehqApiKeyOk returns a tuple with the CommercehqApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercehqApiKey

`func (o *AccountCartAdd) SetCommercehqApiKey(v string)`

SetCommercehqApiKey sets CommercehqApiKey field to given value.

### HasCommercehqApiKey

`func (o *AccountCartAdd) HasCommercehqApiKey() bool`

HasCommercehqApiKey returns a boolean if a field has been set.

### GetCommercehqApiPassword

`func (o *AccountCartAdd) GetCommercehqApiPassword() string`

GetCommercehqApiPassword returns the CommercehqApiPassword field if non-nil, zero value otherwise.

### GetCommercehqApiPasswordOk

`func (o *AccountCartAdd) GetCommercehqApiPasswordOk() (*string, bool)`

GetCommercehqApiPasswordOk returns a tuple with the CommercehqApiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercehqApiPassword

`func (o *AccountCartAdd) SetCommercehqApiPassword(v string)`

SetCommercehqApiPassword sets CommercehqApiPassword field to given value.

### HasCommercehqApiPassword

`func (o *AccountCartAdd) HasCommercehqApiPassword() bool`

HasCommercehqApiPassword returns a boolean if a field has been set.

### GetWcConsumerKey

`func (o *AccountCartAdd) GetWcConsumerKey() string`

GetWcConsumerKey returns the WcConsumerKey field if non-nil, zero value otherwise.

### GetWcConsumerKeyOk

`func (o *AccountCartAdd) GetWcConsumerKeyOk() (*string, bool)`

GetWcConsumerKeyOk returns a tuple with the WcConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWcConsumerKey

`func (o *AccountCartAdd) SetWcConsumerKey(v string)`

SetWcConsumerKey sets WcConsumerKey field to given value.

### HasWcConsumerKey

`func (o *AccountCartAdd) HasWcConsumerKey() bool`

HasWcConsumerKey returns a boolean if a field has been set.

### GetWcConsumerSecret

`func (o *AccountCartAdd) GetWcConsumerSecret() string`

GetWcConsumerSecret returns the WcConsumerSecret field if non-nil, zero value otherwise.

### GetWcConsumerSecretOk

`func (o *AccountCartAdd) GetWcConsumerSecretOk() (*string, bool)`

GetWcConsumerSecretOk returns a tuple with the WcConsumerSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWcConsumerSecret

`func (o *AccountCartAdd) SetWcConsumerSecret(v string)`

SetWcConsumerSecret sets WcConsumerSecret field to given value.

### HasWcConsumerSecret

`func (o *AccountCartAdd) HasWcConsumerSecret() bool`

HasWcConsumerSecret returns a boolean if a field has been set.

### GetMagentoConsumerKey

`func (o *AccountCartAdd) GetMagentoConsumerKey() string`

GetMagentoConsumerKey returns the MagentoConsumerKey field if non-nil, zero value otherwise.

### GetMagentoConsumerKeyOk

`func (o *AccountCartAdd) GetMagentoConsumerKeyOk() (*string, bool)`

GetMagentoConsumerKeyOk returns a tuple with the MagentoConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagentoConsumerKey

`func (o *AccountCartAdd) SetMagentoConsumerKey(v string)`

SetMagentoConsumerKey sets MagentoConsumerKey field to given value.

### HasMagentoConsumerKey

`func (o *AccountCartAdd) HasMagentoConsumerKey() bool`

HasMagentoConsumerKey returns a boolean if a field has been set.

### GetMagentoConsumerSecret

`func (o *AccountCartAdd) GetMagentoConsumerSecret() string`

GetMagentoConsumerSecret returns the MagentoConsumerSecret field if non-nil, zero value otherwise.

### GetMagentoConsumerSecretOk

`func (o *AccountCartAdd) GetMagentoConsumerSecretOk() (*string, bool)`

GetMagentoConsumerSecretOk returns a tuple with the MagentoConsumerSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagentoConsumerSecret

`func (o *AccountCartAdd) SetMagentoConsumerSecret(v string)`

SetMagentoConsumerSecret sets MagentoConsumerSecret field to given value.

### HasMagentoConsumerSecret

`func (o *AccountCartAdd) HasMagentoConsumerSecret() bool`

HasMagentoConsumerSecret returns a boolean if a field has been set.

### GetMagentoAccessToken

`func (o *AccountCartAdd) GetMagentoAccessToken() string`

GetMagentoAccessToken returns the MagentoAccessToken field if non-nil, zero value otherwise.

### GetMagentoAccessTokenOk

`func (o *AccountCartAdd) GetMagentoAccessTokenOk() (*string, bool)`

GetMagentoAccessTokenOk returns a tuple with the MagentoAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagentoAccessToken

`func (o *AccountCartAdd) SetMagentoAccessToken(v string)`

SetMagentoAccessToken sets MagentoAccessToken field to given value.

### HasMagentoAccessToken

`func (o *AccountCartAdd) HasMagentoAccessToken() bool`

HasMagentoAccessToken returns a boolean if a field has been set.

### GetMagentoTokenSecret

`func (o *AccountCartAdd) GetMagentoTokenSecret() string`

GetMagentoTokenSecret returns the MagentoTokenSecret field if non-nil, zero value otherwise.

### GetMagentoTokenSecretOk

`func (o *AccountCartAdd) GetMagentoTokenSecretOk() (*string, bool)`

GetMagentoTokenSecretOk returns a tuple with the MagentoTokenSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagentoTokenSecret

`func (o *AccountCartAdd) SetMagentoTokenSecret(v string)`

SetMagentoTokenSecret sets MagentoTokenSecret field to given value.

### HasMagentoTokenSecret

`func (o *AccountCartAdd) HasMagentoTokenSecret() bool`

HasMagentoTokenSecret returns a boolean if a field has been set.

### GetPrestashopWebserviceKey

`func (o *AccountCartAdd) GetPrestashopWebserviceKey() string`

GetPrestashopWebserviceKey returns the PrestashopWebserviceKey field if non-nil, zero value otherwise.

### GetPrestashopWebserviceKeyOk

`func (o *AccountCartAdd) GetPrestashopWebserviceKeyOk() (*string, bool)`

GetPrestashopWebserviceKeyOk returns a tuple with the PrestashopWebserviceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrestashopWebserviceKey

`func (o *AccountCartAdd) SetPrestashopWebserviceKey(v string)`

SetPrestashopWebserviceKey sets PrestashopWebserviceKey field to given value.

### HasPrestashopWebserviceKey

`func (o *AccountCartAdd) HasPrestashopWebserviceKey() bool`

HasPrestashopWebserviceKey returns a boolean if a field has been set.

### GetWixAppId

`func (o *AccountCartAdd) GetWixAppId() string`

GetWixAppId returns the WixAppId field if non-nil, zero value otherwise.

### GetWixAppIdOk

`func (o *AccountCartAdd) GetWixAppIdOk() (*string, bool)`

GetWixAppIdOk returns a tuple with the WixAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWixAppId

`func (o *AccountCartAdd) SetWixAppId(v string)`

SetWixAppId sets WixAppId field to given value.


### GetWixAppSecretKey

`func (o *AccountCartAdd) GetWixAppSecretKey() string`

GetWixAppSecretKey returns the WixAppSecretKey field if non-nil, zero value otherwise.

### GetWixAppSecretKeyOk

`func (o *AccountCartAdd) GetWixAppSecretKeyOk() (*string, bool)`

GetWixAppSecretKeyOk returns a tuple with the WixAppSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWixAppSecretKey

`func (o *AccountCartAdd) SetWixAppSecretKey(v string)`

SetWixAppSecretKey sets WixAppSecretKey field to given value.


### GetWixInstanceId

`func (o *AccountCartAdd) GetWixInstanceId() string`

GetWixInstanceId returns the WixInstanceId field if non-nil, zero value otherwise.

### GetWixInstanceIdOk

`func (o *AccountCartAdd) GetWixInstanceIdOk() (*string, bool)`

GetWixInstanceIdOk returns a tuple with the WixInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWixInstanceId

`func (o *AccountCartAdd) SetWixInstanceId(v string)`

SetWixInstanceId sets WixInstanceId field to given value.

### HasWixInstanceId

`func (o *AccountCartAdd) HasWixInstanceId() bool`

HasWixInstanceId returns a boolean if a field has been set.

### GetWixRefreshToken

`func (o *AccountCartAdd) GetWixRefreshToken() string`

GetWixRefreshToken returns the WixRefreshToken field if non-nil, zero value otherwise.

### GetWixRefreshTokenOk

`func (o *AccountCartAdd) GetWixRefreshTokenOk() (*string, bool)`

GetWixRefreshTokenOk returns a tuple with the WixRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWixRefreshToken

`func (o *AccountCartAdd) SetWixRefreshToken(v string)`

SetWixRefreshToken sets WixRefreshToken field to given value.

### HasWixRefreshToken

`func (o *AccountCartAdd) HasWixRefreshToken() bool`

HasWixRefreshToken returns a boolean if a field has been set.

### GetMercadoLibreAppId

`func (o *AccountCartAdd) GetMercadoLibreAppId() string`

GetMercadoLibreAppId returns the MercadoLibreAppId field if non-nil, zero value otherwise.

### GetMercadoLibreAppIdOk

`func (o *AccountCartAdd) GetMercadoLibreAppIdOk() (*string, bool)`

GetMercadoLibreAppIdOk returns a tuple with the MercadoLibreAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMercadoLibreAppId

`func (o *AccountCartAdd) SetMercadoLibreAppId(v string)`

SetMercadoLibreAppId sets MercadoLibreAppId field to given value.

### HasMercadoLibreAppId

`func (o *AccountCartAdd) HasMercadoLibreAppId() bool`

HasMercadoLibreAppId returns a boolean if a field has been set.

### GetMercadoLibreAppSecretKey

`func (o *AccountCartAdd) GetMercadoLibreAppSecretKey() string`

GetMercadoLibreAppSecretKey returns the MercadoLibreAppSecretKey field if non-nil, zero value otherwise.

### GetMercadoLibreAppSecretKeyOk

`func (o *AccountCartAdd) GetMercadoLibreAppSecretKeyOk() (*string, bool)`

GetMercadoLibreAppSecretKeyOk returns a tuple with the MercadoLibreAppSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMercadoLibreAppSecretKey

`func (o *AccountCartAdd) SetMercadoLibreAppSecretKey(v string)`

SetMercadoLibreAppSecretKey sets MercadoLibreAppSecretKey field to given value.

### HasMercadoLibreAppSecretKey

`func (o *AccountCartAdd) HasMercadoLibreAppSecretKey() bool`

HasMercadoLibreAppSecretKey returns a boolean if a field has been set.

### GetMercadoLibreRefreshToken

`func (o *AccountCartAdd) GetMercadoLibreRefreshToken() string`

GetMercadoLibreRefreshToken returns the MercadoLibreRefreshToken field if non-nil, zero value otherwise.

### GetMercadoLibreRefreshTokenOk

`func (o *AccountCartAdd) GetMercadoLibreRefreshTokenOk() (*string, bool)`

GetMercadoLibreRefreshTokenOk returns a tuple with the MercadoLibreRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMercadoLibreRefreshToken

`func (o *AccountCartAdd) SetMercadoLibreRefreshToken(v string)`

SetMercadoLibreRefreshToken sets MercadoLibreRefreshToken field to given value.

### HasMercadoLibreRefreshToken

`func (o *AccountCartAdd) HasMercadoLibreRefreshToken() bool`

HasMercadoLibreRefreshToken returns a boolean if a field has been set.

### GetZidClientId

`func (o *AccountCartAdd) GetZidClientId() int32`

GetZidClientId returns the ZidClientId field if non-nil, zero value otherwise.

### GetZidClientIdOk

`func (o *AccountCartAdd) GetZidClientIdOk() (*int32, bool)`

GetZidClientIdOk returns a tuple with the ZidClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZidClientId

`func (o *AccountCartAdd) SetZidClientId(v int32)`

SetZidClientId sets ZidClientId field to given value.

### HasZidClientId

`func (o *AccountCartAdd) HasZidClientId() bool`

HasZidClientId returns a boolean if a field has been set.

### GetZidClientSecret

`func (o *AccountCartAdd) GetZidClientSecret() string`

GetZidClientSecret returns the ZidClientSecret field if non-nil, zero value otherwise.

### GetZidClientSecretOk

`func (o *AccountCartAdd) GetZidClientSecretOk() (*string, bool)`

GetZidClientSecretOk returns a tuple with the ZidClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZidClientSecret

`func (o *AccountCartAdd) SetZidClientSecret(v string)`

SetZidClientSecret sets ZidClientSecret field to given value.

### HasZidClientSecret

`func (o *AccountCartAdd) HasZidClientSecret() bool`

HasZidClientSecret returns a boolean if a field has been set.

### GetZidAccessToken

`func (o *AccountCartAdd) GetZidAccessToken() string`

GetZidAccessToken returns the ZidAccessToken field if non-nil, zero value otherwise.

### GetZidAccessTokenOk

`func (o *AccountCartAdd) GetZidAccessTokenOk() (*string, bool)`

GetZidAccessTokenOk returns a tuple with the ZidAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZidAccessToken

`func (o *AccountCartAdd) SetZidAccessToken(v string)`

SetZidAccessToken sets ZidAccessToken field to given value.

### HasZidAccessToken

`func (o *AccountCartAdd) HasZidAccessToken() bool`

HasZidAccessToken returns a boolean if a field has been set.

### GetZidAuthorization

`func (o *AccountCartAdd) GetZidAuthorization() string`

GetZidAuthorization returns the ZidAuthorization field if non-nil, zero value otherwise.

### GetZidAuthorizationOk

`func (o *AccountCartAdd) GetZidAuthorizationOk() (*string, bool)`

GetZidAuthorizationOk returns a tuple with the ZidAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZidAuthorization

`func (o *AccountCartAdd) SetZidAuthorization(v string)`

SetZidAuthorization sets ZidAuthorization field to given value.

### HasZidAuthorization

`func (o *AccountCartAdd) HasZidAuthorization() bool`

HasZidAuthorization returns a boolean if a field has been set.

### GetZidRefreshToken

`func (o *AccountCartAdd) GetZidRefreshToken() string`

GetZidRefreshToken returns the ZidRefreshToken field if non-nil, zero value otherwise.

### GetZidRefreshTokenOk

`func (o *AccountCartAdd) GetZidRefreshTokenOk() (*string, bool)`

GetZidRefreshTokenOk returns a tuple with the ZidRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZidRefreshToken

`func (o *AccountCartAdd) SetZidRefreshToken(v string)`

SetZidRefreshToken sets ZidRefreshToken field to given value.

### HasZidRefreshToken

`func (o *AccountCartAdd) HasZidRefreshToken() bool`

HasZidRefreshToken returns a boolean if a field has been set.

### GetFlipkartClientId

`func (o *AccountCartAdd) GetFlipkartClientId() string`

GetFlipkartClientId returns the FlipkartClientId field if non-nil, zero value otherwise.

### GetFlipkartClientIdOk

`func (o *AccountCartAdd) GetFlipkartClientIdOk() (*string, bool)`

GetFlipkartClientIdOk returns a tuple with the FlipkartClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlipkartClientId

`func (o *AccountCartAdd) SetFlipkartClientId(v string)`

SetFlipkartClientId sets FlipkartClientId field to given value.

### HasFlipkartClientId

`func (o *AccountCartAdd) HasFlipkartClientId() bool`

HasFlipkartClientId returns a boolean if a field has been set.

### GetFlipkartClientSecret

`func (o *AccountCartAdd) GetFlipkartClientSecret() string`

GetFlipkartClientSecret returns the FlipkartClientSecret field if non-nil, zero value otherwise.

### GetFlipkartClientSecretOk

`func (o *AccountCartAdd) GetFlipkartClientSecretOk() (*string, bool)`

GetFlipkartClientSecretOk returns a tuple with the FlipkartClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlipkartClientSecret

`func (o *AccountCartAdd) SetFlipkartClientSecret(v string)`

SetFlipkartClientSecret sets FlipkartClientSecret field to given value.

### HasFlipkartClientSecret

`func (o *AccountCartAdd) HasFlipkartClientSecret() bool`

HasFlipkartClientSecret returns a boolean if a field has been set.

### GetAllegroClientId

`func (o *AccountCartAdd) GetAllegroClientId() string`

GetAllegroClientId returns the AllegroClientId field if non-nil, zero value otherwise.

### GetAllegroClientIdOk

`func (o *AccountCartAdd) GetAllegroClientIdOk() (*string, bool)`

GetAllegroClientIdOk returns a tuple with the AllegroClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllegroClientId

`func (o *AccountCartAdd) SetAllegroClientId(v string)`

SetAllegroClientId sets AllegroClientId field to given value.

### HasAllegroClientId

`func (o *AccountCartAdd) HasAllegroClientId() bool`

HasAllegroClientId returns a boolean if a field has been set.

### GetAllegroClientSecret

`func (o *AccountCartAdd) GetAllegroClientSecret() string`

GetAllegroClientSecret returns the AllegroClientSecret field if non-nil, zero value otherwise.

### GetAllegroClientSecretOk

`func (o *AccountCartAdd) GetAllegroClientSecretOk() (*string, bool)`

GetAllegroClientSecretOk returns a tuple with the AllegroClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllegroClientSecret

`func (o *AccountCartAdd) SetAllegroClientSecret(v string)`

SetAllegroClientSecret sets AllegroClientSecret field to given value.

### HasAllegroClientSecret

`func (o *AccountCartAdd) HasAllegroClientSecret() bool`

HasAllegroClientSecret returns a boolean if a field has been set.

### GetAllegroAccessToken

`func (o *AccountCartAdd) GetAllegroAccessToken() string`

GetAllegroAccessToken returns the AllegroAccessToken field if non-nil, zero value otherwise.

### GetAllegroAccessTokenOk

`func (o *AccountCartAdd) GetAllegroAccessTokenOk() (*string, bool)`

GetAllegroAccessTokenOk returns a tuple with the AllegroAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllegroAccessToken

`func (o *AccountCartAdd) SetAllegroAccessToken(v string)`

SetAllegroAccessToken sets AllegroAccessToken field to given value.

### HasAllegroAccessToken

`func (o *AccountCartAdd) HasAllegroAccessToken() bool`

HasAllegroAccessToken returns a boolean if a field has been set.

### GetAllegroRefreshToken

`func (o *AccountCartAdd) GetAllegroRefreshToken() string`

GetAllegroRefreshToken returns the AllegroRefreshToken field if non-nil, zero value otherwise.

### GetAllegroRefreshTokenOk

`func (o *AccountCartAdd) GetAllegroRefreshTokenOk() (*string, bool)`

GetAllegroRefreshTokenOk returns a tuple with the AllegroRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllegroRefreshToken

`func (o *AccountCartAdd) SetAllegroRefreshToken(v string)`

SetAllegroRefreshToken sets AllegroRefreshToken field to given value.

### HasAllegroRefreshToken

`func (o *AccountCartAdd) HasAllegroRefreshToken() bool`

HasAllegroRefreshToken returns a boolean if a field has been set.

### GetAllegroEnvironment

`func (o *AccountCartAdd) GetAllegroEnvironment() string`

GetAllegroEnvironment returns the AllegroEnvironment field if non-nil, zero value otherwise.

### GetAllegroEnvironmentOk

`func (o *AccountCartAdd) GetAllegroEnvironmentOk() (*string, bool)`

GetAllegroEnvironmentOk returns a tuple with the AllegroEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllegroEnvironment

`func (o *AccountCartAdd) SetAllegroEnvironment(v string)`

SetAllegroEnvironment sets AllegroEnvironment field to given value.

### HasAllegroEnvironment

`func (o *AccountCartAdd) HasAllegroEnvironment() bool`

HasAllegroEnvironment returns a boolean if a field has been set.

### GetZohoClientId

`func (o *AccountCartAdd) GetZohoClientId() string`

GetZohoClientId returns the ZohoClientId field if non-nil, zero value otherwise.

### GetZohoClientIdOk

`func (o *AccountCartAdd) GetZohoClientIdOk() (*string, bool)`

GetZohoClientIdOk returns a tuple with the ZohoClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZohoClientId

`func (o *AccountCartAdd) SetZohoClientId(v string)`

SetZohoClientId sets ZohoClientId field to given value.

### HasZohoClientId

`func (o *AccountCartAdd) HasZohoClientId() bool`

HasZohoClientId returns a boolean if a field has been set.

### GetZohoClientSecret

`func (o *AccountCartAdd) GetZohoClientSecret() string`

GetZohoClientSecret returns the ZohoClientSecret field if non-nil, zero value otherwise.

### GetZohoClientSecretOk

`func (o *AccountCartAdd) GetZohoClientSecretOk() (*string, bool)`

GetZohoClientSecretOk returns a tuple with the ZohoClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZohoClientSecret

`func (o *AccountCartAdd) SetZohoClientSecret(v string)`

SetZohoClientSecret sets ZohoClientSecret field to given value.

### HasZohoClientSecret

`func (o *AccountCartAdd) HasZohoClientSecret() bool`

HasZohoClientSecret returns a boolean if a field has been set.

### GetZohoRefreshToken

`func (o *AccountCartAdd) GetZohoRefreshToken() string`

GetZohoRefreshToken returns the ZohoRefreshToken field if non-nil, zero value otherwise.

### GetZohoRefreshTokenOk

`func (o *AccountCartAdd) GetZohoRefreshTokenOk() (*string, bool)`

GetZohoRefreshTokenOk returns a tuple with the ZohoRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZohoRefreshToken

`func (o *AccountCartAdd) SetZohoRefreshToken(v string)`

SetZohoRefreshToken sets ZohoRefreshToken field to given value.

### HasZohoRefreshToken

`func (o *AccountCartAdd) HasZohoRefreshToken() bool`

HasZohoRefreshToken returns a boolean if a field has been set.

### GetZohoRegion

`func (o *AccountCartAdd) GetZohoRegion() string`

GetZohoRegion returns the ZohoRegion field if non-nil, zero value otherwise.

### GetZohoRegionOk

`func (o *AccountCartAdd) GetZohoRegionOk() (*string, bool)`

GetZohoRegionOk returns a tuple with the ZohoRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZohoRegion

`func (o *AccountCartAdd) SetZohoRegion(v string)`

SetZohoRegion sets ZohoRegion field to given value.

### HasZohoRegion

`func (o *AccountCartAdd) HasZohoRegion() bool`

HasZohoRegion returns a boolean if a field has been set.

### GetOttoClientId

`func (o *AccountCartAdd) GetOttoClientId() string`

GetOttoClientId returns the OttoClientId field if non-nil, zero value otherwise.

### GetOttoClientIdOk

`func (o *AccountCartAdd) GetOttoClientIdOk() (*string, bool)`

GetOttoClientIdOk returns a tuple with the OttoClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOttoClientId

`func (o *AccountCartAdd) SetOttoClientId(v string)`

SetOttoClientId sets OttoClientId field to given value.

### HasOttoClientId

`func (o *AccountCartAdd) HasOttoClientId() bool`

HasOttoClientId returns a boolean if a field has been set.

### GetOttoClientSecret

`func (o *AccountCartAdd) GetOttoClientSecret() string`

GetOttoClientSecret returns the OttoClientSecret field if non-nil, zero value otherwise.

### GetOttoClientSecretOk

`func (o *AccountCartAdd) GetOttoClientSecretOk() (*string, bool)`

GetOttoClientSecretOk returns a tuple with the OttoClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOttoClientSecret

`func (o *AccountCartAdd) SetOttoClientSecret(v string)`

SetOttoClientSecret sets OttoClientSecret field to given value.

### HasOttoClientSecret

`func (o *AccountCartAdd) HasOttoClientSecret() bool`

HasOttoClientSecret returns a boolean if a field has been set.

### GetOttoAppId

`func (o *AccountCartAdd) GetOttoAppId() string`

GetOttoAppId returns the OttoAppId field if non-nil, zero value otherwise.

### GetOttoAppIdOk

`func (o *AccountCartAdd) GetOttoAppIdOk() (*string, bool)`

GetOttoAppIdOk returns a tuple with the OttoAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOttoAppId

`func (o *AccountCartAdd) SetOttoAppId(v string)`

SetOttoAppId sets OttoAppId field to given value.

### HasOttoAppId

`func (o *AccountCartAdd) HasOttoAppId() bool`

HasOttoAppId returns a boolean if a field has been set.

### GetOttoRefreshToken

`func (o *AccountCartAdd) GetOttoRefreshToken() string`

GetOttoRefreshToken returns the OttoRefreshToken field if non-nil, zero value otherwise.

### GetOttoRefreshTokenOk

`func (o *AccountCartAdd) GetOttoRefreshTokenOk() (*string, bool)`

GetOttoRefreshTokenOk returns a tuple with the OttoRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOttoRefreshToken

`func (o *AccountCartAdd) SetOttoRefreshToken(v string)`

SetOttoRefreshToken sets OttoRefreshToken field to given value.

### HasOttoRefreshToken

`func (o *AccountCartAdd) HasOttoRefreshToken() bool`

HasOttoRefreshToken returns a boolean if a field has been set.

### GetOttoEnvironment

`func (o *AccountCartAdd) GetOttoEnvironment() string`

GetOttoEnvironment returns the OttoEnvironment field if non-nil, zero value otherwise.

### GetOttoEnvironmentOk

`func (o *AccountCartAdd) GetOttoEnvironmentOk() (*string, bool)`

GetOttoEnvironmentOk returns a tuple with the OttoEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOttoEnvironment

`func (o *AccountCartAdd) SetOttoEnvironment(v string)`

SetOttoEnvironment sets OttoEnvironment field to given value.

### HasOttoEnvironment

`func (o *AccountCartAdd) HasOttoEnvironment() bool`

HasOttoEnvironment returns a boolean if a field has been set.

### GetOttoAccessToken

`func (o *AccountCartAdd) GetOttoAccessToken() string`

GetOttoAccessToken returns the OttoAccessToken field if non-nil, zero value otherwise.

### GetOttoAccessTokenOk

`func (o *AccountCartAdd) GetOttoAccessTokenOk() (*string, bool)`

GetOttoAccessTokenOk returns a tuple with the OttoAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOttoAccessToken

`func (o *AccountCartAdd) SetOttoAccessToken(v string)`

SetOttoAccessToken sets OttoAccessToken field to given value.

### HasOttoAccessToken

`func (o *AccountCartAdd) HasOttoAccessToken() bool`

HasOttoAccessToken returns a boolean if a field has been set.

### GetTiktokshopAppKey

`func (o *AccountCartAdd) GetTiktokshopAppKey() string`

GetTiktokshopAppKey returns the TiktokshopAppKey field if non-nil, zero value otherwise.

### GetTiktokshopAppKeyOk

`func (o *AccountCartAdd) GetTiktokshopAppKeyOk() (*string, bool)`

GetTiktokshopAppKeyOk returns a tuple with the TiktokshopAppKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiktokshopAppKey

`func (o *AccountCartAdd) SetTiktokshopAppKey(v string)`

SetTiktokshopAppKey sets TiktokshopAppKey field to given value.

### HasTiktokshopAppKey

`func (o *AccountCartAdd) HasTiktokshopAppKey() bool`

HasTiktokshopAppKey returns a boolean if a field has been set.

### GetTiktokshopAppSecret

`func (o *AccountCartAdd) GetTiktokshopAppSecret() string`

GetTiktokshopAppSecret returns the TiktokshopAppSecret field if non-nil, zero value otherwise.

### GetTiktokshopAppSecretOk

`func (o *AccountCartAdd) GetTiktokshopAppSecretOk() (*string, bool)`

GetTiktokshopAppSecretOk returns a tuple with the TiktokshopAppSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiktokshopAppSecret

`func (o *AccountCartAdd) SetTiktokshopAppSecret(v string)`

SetTiktokshopAppSecret sets TiktokshopAppSecret field to given value.

### HasTiktokshopAppSecret

`func (o *AccountCartAdd) HasTiktokshopAppSecret() bool`

HasTiktokshopAppSecret returns a boolean if a field has been set.

### GetTiktokshopRefreshToken

`func (o *AccountCartAdd) GetTiktokshopRefreshToken() string`

GetTiktokshopRefreshToken returns the TiktokshopRefreshToken field if non-nil, zero value otherwise.

### GetTiktokshopRefreshTokenOk

`func (o *AccountCartAdd) GetTiktokshopRefreshTokenOk() (*string, bool)`

GetTiktokshopRefreshTokenOk returns a tuple with the TiktokshopRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiktokshopRefreshToken

`func (o *AccountCartAdd) SetTiktokshopRefreshToken(v string)`

SetTiktokshopRefreshToken sets TiktokshopRefreshToken field to given value.

### HasTiktokshopRefreshToken

`func (o *AccountCartAdd) HasTiktokshopRefreshToken() bool`

HasTiktokshopRefreshToken returns a boolean if a field has been set.

### GetTiktokshopAccessToken

`func (o *AccountCartAdd) GetTiktokshopAccessToken() string`

GetTiktokshopAccessToken returns the TiktokshopAccessToken field if non-nil, zero value otherwise.

### GetTiktokshopAccessTokenOk

`func (o *AccountCartAdd) GetTiktokshopAccessTokenOk() (*string, bool)`

GetTiktokshopAccessTokenOk returns a tuple with the TiktokshopAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiktokshopAccessToken

`func (o *AccountCartAdd) SetTiktokshopAccessToken(v string)`

SetTiktokshopAccessToken sets TiktokshopAccessToken field to given value.

### HasTiktokshopAccessToken

`func (o *AccountCartAdd) HasTiktokshopAccessToken() bool`

HasTiktokshopAccessToken returns a boolean if a field has been set.

### GetSallaClientId

`func (o *AccountCartAdd) GetSallaClientId() string`

GetSallaClientId returns the SallaClientId field if non-nil, zero value otherwise.

### GetSallaClientIdOk

`func (o *AccountCartAdd) GetSallaClientIdOk() (*string, bool)`

GetSallaClientIdOk returns a tuple with the SallaClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSallaClientId

`func (o *AccountCartAdd) SetSallaClientId(v string)`

SetSallaClientId sets SallaClientId field to given value.

### HasSallaClientId

`func (o *AccountCartAdd) HasSallaClientId() bool`

HasSallaClientId returns a boolean if a field has been set.

### GetSallaClientSecret

`func (o *AccountCartAdd) GetSallaClientSecret() string`

GetSallaClientSecret returns the SallaClientSecret field if non-nil, zero value otherwise.

### GetSallaClientSecretOk

`func (o *AccountCartAdd) GetSallaClientSecretOk() (*string, bool)`

GetSallaClientSecretOk returns a tuple with the SallaClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSallaClientSecret

`func (o *AccountCartAdd) SetSallaClientSecret(v string)`

SetSallaClientSecret sets SallaClientSecret field to given value.

### HasSallaClientSecret

`func (o *AccountCartAdd) HasSallaClientSecret() bool`

HasSallaClientSecret returns a boolean if a field has been set.

### GetSallaRefreshToken

`func (o *AccountCartAdd) GetSallaRefreshToken() string`

GetSallaRefreshToken returns the SallaRefreshToken field if non-nil, zero value otherwise.

### GetSallaRefreshTokenOk

`func (o *AccountCartAdd) GetSallaRefreshTokenOk() (*string, bool)`

GetSallaRefreshTokenOk returns a tuple with the SallaRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSallaRefreshToken

`func (o *AccountCartAdd) SetSallaRefreshToken(v string)`

SetSallaRefreshToken sets SallaRefreshToken field to given value.

### HasSallaRefreshToken

`func (o *AccountCartAdd) HasSallaRefreshToken() bool`

HasSallaRefreshToken returns a boolean if a field has been set.

### GetSallaAccessToken

`func (o *AccountCartAdd) GetSallaAccessToken() string`

GetSallaAccessToken returns the SallaAccessToken field if non-nil, zero value otherwise.

### GetSallaAccessTokenOk

`func (o *AccountCartAdd) GetSallaAccessTokenOk() (*string, bool)`

GetSallaAccessTokenOk returns a tuple with the SallaAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSallaAccessToken

`func (o *AccountCartAdd) SetSallaAccessToken(v string)`

SetSallaAccessToken sets SallaAccessToken field to given value.

### HasSallaAccessToken

`func (o *AccountCartAdd) HasSallaAccessToken() bool`

HasSallaAccessToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


