# CartCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CartId** | **string** | Store’s identifier which you can get from cart_list method | 
**StoreUrl** | Pointer to **string** | A web address of a store that you would like to connect to API2Cart | [optional] 
**BridgeUrl** | Pointer to **string** | This parameter allows to set up store with custom bridge url (also you must use store_root parameter if a bridge folder is not in the root folder of the store) | [optional] 
**StoreRoot** | Pointer to **string** | Absolute path to the store root directory (used with \&quot;bridge_url\&quot; parameter) | [optional] 
**StoreKey** | Pointer to **string** | Set this parameter if bridge is already uploaded to store | [optional] 
**SharedSecret** | Pointer to **string** | Shared secret | [optional] 
**ValidateVersion** | Pointer to **bool** | Specify if api2cart should validate cart version | [optional] [default to false]
**Verify** | Pointer to **bool** | Enables or disables cart&#39;s verification | [optional] [default to true]
**DbTablesPrefix** | Pointer to **string** | DB tables prefix | [optional] 
**UserAgent** | Pointer to **string** | This parameter allows you to set your custom user agent, which will be used in requests to the store. Please use it cautiously, as the store&#39;s firewall may block specific values. | [optional] 
**FtpHost** | Pointer to **string** | FTP connection host | [optional] 
**FtpUser** | Pointer to **string** | FTP User | [optional] 
**FtpPassword** | Pointer to **string** | FTP Password | [optional] 
**FtpPort** | Pointer to **int32** | FTP Port | [optional] 
**FtpStoreDir** | Pointer to **string** | FTP Store dir | [optional] 
**ApiKey3dcart** | Pointer to **string** | 3DCart API Key | [optional] 
**AdminAccount** | Pointer to **string** | It&#39;s a BigCommerce account for which API is enabled | [optional] 
**ApiPath** | Pointer to **string** | BigCommerce API URL | [optional] 
**ApiKeyBigcommerce** | Pointer to **string** | Bigcommerce API Key | [optional] 
**ClientId** | Pointer to **string** | Client ID of the requesting app | [optional] 
**AccessToken** | Pointer to **string** | Access token authorizing the app to access resources on behalf of a user | [optional] 
**Context** | Pointer to **string** | API Path section unique to the store | [optional] 
**AccessToken** | Pointer to **string** | Access token authorizing the app to access resources on behalf of a user | [optional] 
**ApiKeyShopify** | Pointer to **string** | Shopify API Key | [optional] 
**ApiPassword** | Pointer to **string** | Shopify API Password | [optional] 
**AccessTokenShopify** | Pointer to **string** | Access token authorizing the app to access resources on behalf of a user | [optional] 
**ApiKey** | Pointer to **string** | Neto API Key | [optional] 
**ApiUsername** | Pointer to **string** | Neto User Name | [optional] 
**EncryptedPassword** | Pointer to **string** | Volusion API Password | [optional] 
**Login** | Pointer to **string** | It&#39;s a Volusion account for which API is enabled | [optional] 
**ApiUserAdnsf** | Pointer to **string** | It&#39;s a AspDotNetStorefront account for which API is available | [optional] 
**ApiPass** | Pointer to **string** | AspDotNetStorefront API Password | [optional] 
**PrivateKey** | Pointer to **string** | 3DCart Application Private Key | [optional] 
**AppToken** | Pointer to **string** | 3DCart Token from Application | [optional] 
**EtsyKeystring** | Pointer to **string** | Etsy keystring | [optional] 
**EtsySharedSecret** | Pointer to **string** | Etsy shared secret | [optional] 
**TokenSecret** | Pointer to **string** | Secret token authorizing the app to access resources on behalf of a user | [optional] 
**EtsyClientId** | Pointer to **string** | Etsy Client Id | [optional] 
**EtsyRefreshToken** | Pointer to **string** | Etsy Refresh token | [optional] 
**EbayClientId** | Pointer to **string** | Application ID (AppID). | [optional] 
**EbayClientSecret** | Pointer to **string** | Shared Secret from eBay application | [optional] 
**EbayRuname** | Pointer to **string** | The RuName value that eBay assigns to your application. | [optional] 
**EbayAccessToken** | Pointer to **string** | Used to authenticate API requests. | [optional] 
**EbayRefreshToken** | Pointer to **string** | Used to renew the access token. | [optional] 
**EbayEnvironment** | Pointer to **string** | eBay environment | [optional] [default to "production"]
**EbaySiteId** | Pointer to **int32** | eBay global ID | [optional] [default to 0]
**DwClientId** | Pointer to **string** | Demandware client id | [optional] 
**DwApiPass** | Pointer to **string** | Demandware api password | [optional] 
**DemandwareUserName** | Pointer to **string** | Demandware user name | [optional] 
**DemandwareUserPassword** | Pointer to **string** | Demandware user password | [optional] 
**StoreId** | **string** | Store Id | 
**SellerId** | Pointer to **string** | Seller Id | [optional] 
**Environment** | Pointer to **string** |  | [optional] [default to "production"]
**HybrisClientId** | Pointer to **string** | Omni Commerce Connector Client ID | [optional] 
**HybrisClientSecret** | Pointer to **string** | Omni Commerce Connector Client Secret | [optional] 
**HybrisUsername** | Pointer to **string** | User Name | [optional] 
**HybrisPassword** | Pointer to **string** | User password | [optional] 
**HybrisWebsites** | Pointer to [**[]AccountCartAddHybrisWebsitesInner**](AccountCartAddHybrisWebsitesInner.md) | Websites to stores mapping data | [optional] 
**WalmartClientId** | Pointer to **string** | Walmart client ID. For the region &#39;ca&#39; use Consumer ID | [optional] 
**WalmartClientSecret** | Pointer to **string** | Walmart client secret. For the region &#39;ca&#39; use Private Key | [optional] 
**WalmartEnvironment** | Pointer to **string** | Walmart environment | [optional] [default to "production"]
**WalmartChannelType** | Pointer to **string** | Walmart WM_CONSUMER.CHANNEL.TYPE header | [optional] 
**WalmartRegion** | Pointer to **string** | Walmart region | [optional] [default to "us"]
**LightspeedApiKey** | Pointer to **string** | LightSpeed api key | [optional] 
**LightspeedApiSecret** | Pointer to **string** | LightSpeed api secret | [optional] 
**ShoplazzaAccessToken** | Pointer to **string** | Access token authorizing the app to access resources on behalf of a user | [optional] 
**ShoplazzaSharedSecret** | Pointer to **string** | Shared secret | [optional] 
**ShopwareAccessKey** | Pointer to **string** | Shopware access key | [optional] 
**ShopwareApiKey** | Pointer to **string** | Shopware api key | [optional] 
**ShopwareApiSecret** | Pointer to **string** | Shopware client secret access key | [optional] 
**CommercehqApiKey** | Pointer to **string** | CommerceHQ api key | [optional] 
**CommercehqApiPassword** | Pointer to **string** | CommerceHQ api password | [optional] 
**Var3dcartPrivateKey** | Pointer to **string** | 3DCart Private Key | [optional] 
**Var3dcartAccessToken** | Pointer to **string** | 3DCart Token | [optional] 
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

## Methods

### NewCartCreate

`func NewCartCreate(cartId string, storeId string, wixAppId string, wixAppSecretKey string, ) *CartCreate`

NewCartCreate instantiates a new CartCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartCreateWithDefaults

`func NewCartCreateWithDefaults() *CartCreate`

NewCartCreateWithDefaults instantiates a new CartCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCartId

`func (o *CartCreate) GetCartId() string`

GetCartId returns the CartId field if non-nil, zero value otherwise.

### GetCartIdOk

`func (o *CartCreate) GetCartIdOk() (*string, bool)`

GetCartIdOk returns a tuple with the CartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartId

`func (o *CartCreate) SetCartId(v string)`

SetCartId sets CartId field to given value.


### GetStoreUrl

`func (o *CartCreate) GetStoreUrl() string`

GetStoreUrl returns the StoreUrl field if non-nil, zero value otherwise.

### GetStoreUrlOk

`func (o *CartCreate) GetStoreUrlOk() (*string, bool)`

GetStoreUrlOk returns a tuple with the StoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreUrl

`func (o *CartCreate) SetStoreUrl(v string)`

SetStoreUrl sets StoreUrl field to given value.

### HasStoreUrl

`func (o *CartCreate) HasStoreUrl() bool`

HasStoreUrl returns a boolean if a field has been set.

### GetBridgeUrl

`func (o *CartCreate) GetBridgeUrl() string`

GetBridgeUrl returns the BridgeUrl field if non-nil, zero value otherwise.

### GetBridgeUrlOk

`func (o *CartCreate) GetBridgeUrlOk() (*string, bool)`

GetBridgeUrlOk returns a tuple with the BridgeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeUrl

`func (o *CartCreate) SetBridgeUrl(v string)`

SetBridgeUrl sets BridgeUrl field to given value.

### HasBridgeUrl

`func (o *CartCreate) HasBridgeUrl() bool`

HasBridgeUrl returns a boolean if a field has been set.

### GetStoreRoot

`func (o *CartCreate) GetStoreRoot() string`

GetStoreRoot returns the StoreRoot field if non-nil, zero value otherwise.

### GetStoreRootOk

`func (o *CartCreate) GetStoreRootOk() (*string, bool)`

GetStoreRootOk returns a tuple with the StoreRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreRoot

`func (o *CartCreate) SetStoreRoot(v string)`

SetStoreRoot sets StoreRoot field to given value.

### HasStoreRoot

`func (o *CartCreate) HasStoreRoot() bool`

HasStoreRoot returns a boolean if a field has been set.

### GetStoreKey

`func (o *CartCreate) GetStoreKey() string`

GetStoreKey returns the StoreKey field if non-nil, zero value otherwise.

### GetStoreKeyOk

`func (o *CartCreate) GetStoreKeyOk() (*string, bool)`

GetStoreKeyOk returns a tuple with the StoreKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreKey

`func (o *CartCreate) SetStoreKey(v string)`

SetStoreKey sets StoreKey field to given value.

### HasStoreKey

`func (o *CartCreate) HasStoreKey() bool`

HasStoreKey returns a boolean if a field has been set.

### GetSharedSecret

`func (o *CartCreate) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *CartCreate) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *CartCreate) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *CartCreate) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetValidateVersion

`func (o *CartCreate) GetValidateVersion() bool`

GetValidateVersion returns the ValidateVersion field if non-nil, zero value otherwise.

### GetValidateVersionOk

`func (o *CartCreate) GetValidateVersionOk() (*bool, bool)`

GetValidateVersionOk returns a tuple with the ValidateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateVersion

`func (o *CartCreate) SetValidateVersion(v bool)`

SetValidateVersion sets ValidateVersion field to given value.

### HasValidateVersion

`func (o *CartCreate) HasValidateVersion() bool`

HasValidateVersion returns a boolean if a field has been set.

### GetVerify

`func (o *CartCreate) GetVerify() bool`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *CartCreate) GetVerifyOk() (*bool, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *CartCreate) SetVerify(v bool)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *CartCreate) HasVerify() bool`

HasVerify returns a boolean if a field has been set.

### GetDbTablesPrefix

`func (o *CartCreate) GetDbTablesPrefix() string`

GetDbTablesPrefix returns the DbTablesPrefix field if non-nil, zero value otherwise.

### GetDbTablesPrefixOk

`func (o *CartCreate) GetDbTablesPrefixOk() (*string, bool)`

GetDbTablesPrefixOk returns a tuple with the DbTablesPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbTablesPrefix

`func (o *CartCreate) SetDbTablesPrefix(v string)`

SetDbTablesPrefix sets DbTablesPrefix field to given value.

### HasDbTablesPrefix

`func (o *CartCreate) HasDbTablesPrefix() bool`

HasDbTablesPrefix returns a boolean if a field has been set.

### GetUserAgent

`func (o *CartCreate) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *CartCreate) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *CartCreate) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *CartCreate) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetFtpHost

`func (o *CartCreate) GetFtpHost() string`

GetFtpHost returns the FtpHost field if non-nil, zero value otherwise.

### GetFtpHostOk

`func (o *CartCreate) GetFtpHostOk() (*string, bool)`

GetFtpHostOk returns a tuple with the FtpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpHost

`func (o *CartCreate) SetFtpHost(v string)`

SetFtpHost sets FtpHost field to given value.

### HasFtpHost

`func (o *CartCreate) HasFtpHost() bool`

HasFtpHost returns a boolean if a field has been set.

### GetFtpUser

`func (o *CartCreate) GetFtpUser() string`

GetFtpUser returns the FtpUser field if non-nil, zero value otherwise.

### GetFtpUserOk

`func (o *CartCreate) GetFtpUserOk() (*string, bool)`

GetFtpUserOk returns a tuple with the FtpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpUser

`func (o *CartCreate) SetFtpUser(v string)`

SetFtpUser sets FtpUser field to given value.

### HasFtpUser

`func (o *CartCreate) HasFtpUser() bool`

HasFtpUser returns a boolean if a field has been set.

### GetFtpPassword

`func (o *CartCreate) GetFtpPassword() string`

GetFtpPassword returns the FtpPassword field if non-nil, zero value otherwise.

### GetFtpPasswordOk

`func (o *CartCreate) GetFtpPasswordOk() (*string, bool)`

GetFtpPasswordOk returns a tuple with the FtpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpPassword

`func (o *CartCreate) SetFtpPassword(v string)`

SetFtpPassword sets FtpPassword field to given value.

### HasFtpPassword

`func (o *CartCreate) HasFtpPassword() bool`

HasFtpPassword returns a boolean if a field has been set.

### GetFtpPort

`func (o *CartCreate) GetFtpPort() int32`

GetFtpPort returns the FtpPort field if non-nil, zero value otherwise.

### GetFtpPortOk

`func (o *CartCreate) GetFtpPortOk() (*int32, bool)`

GetFtpPortOk returns a tuple with the FtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpPort

`func (o *CartCreate) SetFtpPort(v int32)`

SetFtpPort sets FtpPort field to given value.

### HasFtpPort

`func (o *CartCreate) HasFtpPort() bool`

HasFtpPort returns a boolean if a field has been set.

### GetFtpStoreDir

`func (o *CartCreate) GetFtpStoreDir() string`

GetFtpStoreDir returns the FtpStoreDir field if non-nil, zero value otherwise.

### GetFtpStoreDirOk

`func (o *CartCreate) GetFtpStoreDirOk() (*string, bool)`

GetFtpStoreDirOk returns a tuple with the FtpStoreDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpStoreDir

`func (o *CartCreate) SetFtpStoreDir(v string)`

SetFtpStoreDir sets FtpStoreDir field to given value.

### HasFtpStoreDir

`func (o *CartCreate) HasFtpStoreDir() bool`

HasFtpStoreDir returns a boolean if a field has been set.

### GetApiKey3dcart

`func (o *CartCreate) GetApiKey3dcart() string`

GetApiKey3dcart returns the ApiKey3dcart field if non-nil, zero value otherwise.

### GetApiKey3dcartOk

`func (o *CartCreate) GetApiKey3dcartOk() (*string, bool)`

GetApiKey3dcartOk returns a tuple with the ApiKey3dcart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey3dcart

`func (o *CartCreate) SetApiKey3dcart(v string)`

SetApiKey3dcart sets ApiKey3dcart field to given value.

### HasApiKey3dcart

`func (o *CartCreate) HasApiKey3dcart() bool`

HasApiKey3dcart returns a boolean if a field has been set.

### GetAdminAccount

`func (o *CartCreate) GetAdminAccount() string`

GetAdminAccount returns the AdminAccount field if non-nil, zero value otherwise.

### GetAdminAccountOk

`func (o *CartCreate) GetAdminAccountOk() (*string, bool)`

GetAdminAccountOk returns a tuple with the AdminAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAccount

`func (o *CartCreate) SetAdminAccount(v string)`

SetAdminAccount sets AdminAccount field to given value.

### HasAdminAccount

`func (o *CartCreate) HasAdminAccount() bool`

HasAdminAccount returns a boolean if a field has been set.

### GetApiPath

`func (o *CartCreate) GetApiPath() string`

GetApiPath returns the ApiPath field if non-nil, zero value otherwise.

### GetApiPathOk

`func (o *CartCreate) GetApiPathOk() (*string, bool)`

GetApiPathOk returns a tuple with the ApiPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPath

`func (o *CartCreate) SetApiPath(v string)`

SetApiPath sets ApiPath field to given value.

### HasApiPath

`func (o *CartCreate) HasApiPath() bool`

HasApiPath returns a boolean if a field has been set.

### GetApiKeyBigcommerce

`func (o *CartCreate) GetApiKeyBigcommerce() string`

GetApiKeyBigcommerce returns the ApiKeyBigcommerce field if non-nil, zero value otherwise.

### GetApiKeyBigcommerceOk

`func (o *CartCreate) GetApiKeyBigcommerceOk() (*string, bool)`

GetApiKeyBigcommerceOk returns a tuple with the ApiKeyBigcommerce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyBigcommerce

`func (o *CartCreate) SetApiKeyBigcommerce(v string)`

SetApiKeyBigcommerce sets ApiKeyBigcommerce field to given value.

### HasApiKeyBigcommerce

`func (o *CartCreate) HasApiKeyBigcommerce() bool`

HasApiKeyBigcommerce returns a boolean if a field has been set.

### GetClientId

`func (o *CartCreate) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CartCreate) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CartCreate) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CartCreate) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetAccessToken

`func (o *CartCreate) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *CartCreate) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *CartCreate) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *CartCreate) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetContext

`func (o *CartCreate) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CartCreate) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CartCreate) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *CartCreate) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetAccessToken

`func (o *CartCreate) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *CartCreate) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *CartCreate) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *CartCreate) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetApiKeyShopify

`func (o *CartCreate) GetApiKeyShopify() string`

GetApiKeyShopify returns the ApiKeyShopify field if non-nil, zero value otherwise.

### GetApiKeyShopifyOk

`func (o *CartCreate) GetApiKeyShopifyOk() (*string, bool)`

GetApiKeyShopifyOk returns a tuple with the ApiKeyShopify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyShopify

`func (o *CartCreate) SetApiKeyShopify(v string)`

SetApiKeyShopify sets ApiKeyShopify field to given value.

### HasApiKeyShopify

`func (o *CartCreate) HasApiKeyShopify() bool`

HasApiKeyShopify returns a boolean if a field has been set.

### GetApiPassword

`func (o *CartCreate) GetApiPassword() string`

GetApiPassword returns the ApiPassword field if non-nil, zero value otherwise.

### GetApiPasswordOk

`func (o *CartCreate) GetApiPasswordOk() (*string, bool)`

GetApiPasswordOk returns a tuple with the ApiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPassword

`func (o *CartCreate) SetApiPassword(v string)`

SetApiPassword sets ApiPassword field to given value.

### HasApiPassword

`func (o *CartCreate) HasApiPassword() bool`

HasApiPassword returns a boolean if a field has been set.

### GetAccessTokenShopify

`func (o *CartCreate) GetAccessTokenShopify() string`

GetAccessTokenShopify returns the AccessTokenShopify field if non-nil, zero value otherwise.

### GetAccessTokenShopifyOk

`func (o *CartCreate) GetAccessTokenShopifyOk() (*string, bool)`

GetAccessTokenShopifyOk returns a tuple with the AccessTokenShopify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenShopify

`func (o *CartCreate) SetAccessTokenShopify(v string)`

SetAccessTokenShopify sets AccessTokenShopify field to given value.

### HasAccessTokenShopify

`func (o *CartCreate) HasAccessTokenShopify() bool`

HasAccessTokenShopify returns a boolean if a field has been set.

### GetApiKey

`func (o *CartCreate) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CartCreate) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CartCreate) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *CartCreate) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetApiUsername

`func (o *CartCreate) GetApiUsername() string`

GetApiUsername returns the ApiUsername field if non-nil, zero value otherwise.

### GetApiUsernameOk

`func (o *CartCreate) GetApiUsernameOk() (*string, bool)`

GetApiUsernameOk returns a tuple with the ApiUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUsername

`func (o *CartCreate) SetApiUsername(v string)`

SetApiUsername sets ApiUsername field to given value.

### HasApiUsername

`func (o *CartCreate) HasApiUsername() bool`

HasApiUsername returns a boolean if a field has been set.

### GetEncryptedPassword

`func (o *CartCreate) GetEncryptedPassword() string`

GetEncryptedPassword returns the EncryptedPassword field if non-nil, zero value otherwise.

### GetEncryptedPasswordOk

`func (o *CartCreate) GetEncryptedPasswordOk() (*string, bool)`

GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassword

`func (o *CartCreate) SetEncryptedPassword(v string)`

SetEncryptedPassword sets EncryptedPassword field to given value.

### HasEncryptedPassword

`func (o *CartCreate) HasEncryptedPassword() bool`

HasEncryptedPassword returns a boolean if a field has been set.

### GetLogin

`func (o *CartCreate) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *CartCreate) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *CartCreate) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *CartCreate) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetApiUserAdnsf

`func (o *CartCreate) GetApiUserAdnsf() string`

GetApiUserAdnsf returns the ApiUserAdnsf field if non-nil, zero value otherwise.

### GetApiUserAdnsfOk

`func (o *CartCreate) GetApiUserAdnsfOk() (*string, bool)`

GetApiUserAdnsfOk returns a tuple with the ApiUserAdnsf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUserAdnsf

`func (o *CartCreate) SetApiUserAdnsf(v string)`

SetApiUserAdnsf sets ApiUserAdnsf field to given value.

### HasApiUserAdnsf

`func (o *CartCreate) HasApiUserAdnsf() bool`

HasApiUserAdnsf returns a boolean if a field has been set.

### GetApiPass

`func (o *CartCreate) GetApiPass() string`

GetApiPass returns the ApiPass field if non-nil, zero value otherwise.

### GetApiPassOk

`func (o *CartCreate) GetApiPassOk() (*string, bool)`

GetApiPassOk returns a tuple with the ApiPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPass

`func (o *CartCreate) SetApiPass(v string)`

SetApiPass sets ApiPass field to given value.

### HasApiPass

`func (o *CartCreate) HasApiPass() bool`

HasApiPass returns a boolean if a field has been set.

### GetPrivateKey

`func (o *CartCreate) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CartCreate) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CartCreate) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CartCreate) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetAppToken

`func (o *CartCreate) GetAppToken() string`

GetAppToken returns the AppToken field if non-nil, zero value otherwise.

### GetAppTokenOk

`func (o *CartCreate) GetAppTokenOk() (*string, bool)`

GetAppTokenOk returns a tuple with the AppToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppToken

`func (o *CartCreate) SetAppToken(v string)`

SetAppToken sets AppToken field to given value.

### HasAppToken

`func (o *CartCreate) HasAppToken() bool`

HasAppToken returns a boolean if a field has been set.

### GetEtsyKeystring

`func (o *CartCreate) GetEtsyKeystring() string`

GetEtsyKeystring returns the EtsyKeystring field if non-nil, zero value otherwise.

### GetEtsyKeystringOk

`func (o *CartCreate) GetEtsyKeystringOk() (*string, bool)`

GetEtsyKeystringOk returns a tuple with the EtsyKeystring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtsyKeystring

`func (o *CartCreate) SetEtsyKeystring(v string)`

SetEtsyKeystring sets EtsyKeystring field to given value.

### HasEtsyKeystring

`func (o *CartCreate) HasEtsyKeystring() bool`

HasEtsyKeystring returns a boolean if a field has been set.

### GetEtsySharedSecret

`func (o *CartCreate) GetEtsySharedSecret() string`

GetEtsySharedSecret returns the EtsySharedSecret field if non-nil, zero value otherwise.

### GetEtsySharedSecretOk

`func (o *CartCreate) GetEtsySharedSecretOk() (*string, bool)`

GetEtsySharedSecretOk returns a tuple with the EtsySharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtsySharedSecret

`func (o *CartCreate) SetEtsySharedSecret(v string)`

SetEtsySharedSecret sets EtsySharedSecret field to given value.

### HasEtsySharedSecret

`func (o *CartCreate) HasEtsySharedSecret() bool`

HasEtsySharedSecret returns a boolean if a field has been set.

### GetTokenSecret

`func (o *CartCreate) GetTokenSecret() string`

GetTokenSecret returns the TokenSecret field if non-nil, zero value otherwise.

### GetTokenSecretOk

`func (o *CartCreate) GetTokenSecretOk() (*string, bool)`

GetTokenSecretOk returns a tuple with the TokenSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSecret

`func (o *CartCreate) SetTokenSecret(v string)`

SetTokenSecret sets TokenSecret field to given value.

### HasTokenSecret

`func (o *CartCreate) HasTokenSecret() bool`

HasTokenSecret returns a boolean if a field has been set.

### GetEtsyClientId

`func (o *CartCreate) GetEtsyClientId() string`

GetEtsyClientId returns the EtsyClientId field if non-nil, zero value otherwise.

### GetEtsyClientIdOk

`func (o *CartCreate) GetEtsyClientIdOk() (*string, bool)`

GetEtsyClientIdOk returns a tuple with the EtsyClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtsyClientId

`func (o *CartCreate) SetEtsyClientId(v string)`

SetEtsyClientId sets EtsyClientId field to given value.

### HasEtsyClientId

`func (o *CartCreate) HasEtsyClientId() bool`

HasEtsyClientId returns a boolean if a field has been set.

### GetEtsyRefreshToken

`func (o *CartCreate) GetEtsyRefreshToken() string`

GetEtsyRefreshToken returns the EtsyRefreshToken field if non-nil, zero value otherwise.

### GetEtsyRefreshTokenOk

`func (o *CartCreate) GetEtsyRefreshTokenOk() (*string, bool)`

GetEtsyRefreshTokenOk returns a tuple with the EtsyRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtsyRefreshToken

`func (o *CartCreate) SetEtsyRefreshToken(v string)`

SetEtsyRefreshToken sets EtsyRefreshToken field to given value.

### HasEtsyRefreshToken

`func (o *CartCreate) HasEtsyRefreshToken() bool`

HasEtsyRefreshToken returns a boolean if a field has been set.

### GetEbayClientId

`func (o *CartCreate) GetEbayClientId() string`

GetEbayClientId returns the EbayClientId field if non-nil, zero value otherwise.

### GetEbayClientIdOk

`func (o *CartCreate) GetEbayClientIdOk() (*string, bool)`

GetEbayClientIdOk returns a tuple with the EbayClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayClientId

`func (o *CartCreate) SetEbayClientId(v string)`

SetEbayClientId sets EbayClientId field to given value.

### HasEbayClientId

`func (o *CartCreate) HasEbayClientId() bool`

HasEbayClientId returns a boolean if a field has been set.

### GetEbayClientSecret

`func (o *CartCreate) GetEbayClientSecret() string`

GetEbayClientSecret returns the EbayClientSecret field if non-nil, zero value otherwise.

### GetEbayClientSecretOk

`func (o *CartCreate) GetEbayClientSecretOk() (*string, bool)`

GetEbayClientSecretOk returns a tuple with the EbayClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayClientSecret

`func (o *CartCreate) SetEbayClientSecret(v string)`

SetEbayClientSecret sets EbayClientSecret field to given value.

### HasEbayClientSecret

`func (o *CartCreate) HasEbayClientSecret() bool`

HasEbayClientSecret returns a boolean if a field has been set.

### GetEbayRuname

`func (o *CartCreate) GetEbayRuname() string`

GetEbayRuname returns the EbayRuname field if non-nil, zero value otherwise.

### GetEbayRunameOk

`func (o *CartCreate) GetEbayRunameOk() (*string, bool)`

GetEbayRunameOk returns a tuple with the EbayRuname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayRuname

`func (o *CartCreate) SetEbayRuname(v string)`

SetEbayRuname sets EbayRuname field to given value.

### HasEbayRuname

`func (o *CartCreate) HasEbayRuname() bool`

HasEbayRuname returns a boolean if a field has been set.

### GetEbayAccessToken

`func (o *CartCreate) GetEbayAccessToken() string`

GetEbayAccessToken returns the EbayAccessToken field if non-nil, zero value otherwise.

### GetEbayAccessTokenOk

`func (o *CartCreate) GetEbayAccessTokenOk() (*string, bool)`

GetEbayAccessTokenOk returns a tuple with the EbayAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayAccessToken

`func (o *CartCreate) SetEbayAccessToken(v string)`

SetEbayAccessToken sets EbayAccessToken field to given value.

### HasEbayAccessToken

`func (o *CartCreate) HasEbayAccessToken() bool`

HasEbayAccessToken returns a boolean if a field has been set.

### GetEbayRefreshToken

`func (o *CartCreate) GetEbayRefreshToken() string`

GetEbayRefreshToken returns the EbayRefreshToken field if non-nil, zero value otherwise.

### GetEbayRefreshTokenOk

`func (o *CartCreate) GetEbayRefreshTokenOk() (*string, bool)`

GetEbayRefreshTokenOk returns a tuple with the EbayRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayRefreshToken

`func (o *CartCreate) SetEbayRefreshToken(v string)`

SetEbayRefreshToken sets EbayRefreshToken field to given value.

### HasEbayRefreshToken

`func (o *CartCreate) HasEbayRefreshToken() bool`

HasEbayRefreshToken returns a boolean if a field has been set.

### GetEbayEnvironment

`func (o *CartCreate) GetEbayEnvironment() string`

GetEbayEnvironment returns the EbayEnvironment field if non-nil, zero value otherwise.

### GetEbayEnvironmentOk

`func (o *CartCreate) GetEbayEnvironmentOk() (*string, bool)`

GetEbayEnvironmentOk returns a tuple with the EbayEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayEnvironment

`func (o *CartCreate) SetEbayEnvironment(v string)`

SetEbayEnvironment sets EbayEnvironment field to given value.

### HasEbayEnvironment

`func (o *CartCreate) HasEbayEnvironment() bool`

HasEbayEnvironment returns a boolean if a field has been set.

### GetEbaySiteId

`func (o *CartCreate) GetEbaySiteId() int32`

GetEbaySiteId returns the EbaySiteId field if non-nil, zero value otherwise.

### GetEbaySiteIdOk

`func (o *CartCreate) GetEbaySiteIdOk() (*int32, bool)`

GetEbaySiteIdOk returns a tuple with the EbaySiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbaySiteId

`func (o *CartCreate) SetEbaySiteId(v int32)`

SetEbaySiteId sets EbaySiteId field to given value.

### HasEbaySiteId

`func (o *CartCreate) HasEbaySiteId() bool`

HasEbaySiteId returns a boolean if a field has been set.

### GetDwClientId

`func (o *CartCreate) GetDwClientId() string`

GetDwClientId returns the DwClientId field if non-nil, zero value otherwise.

### GetDwClientIdOk

`func (o *CartCreate) GetDwClientIdOk() (*string, bool)`

GetDwClientIdOk returns a tuple with the DwClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDwClientId

`func (o *CartCreate) SetDwClientId(v string)`

SetDwClientId sets DwClientId field to given value.

### HasDwClientId

`func (o *CartCreate) HasDwClientId() bool`

HasDwClientId returns a boolean if a field has been set.

### GetDwApiPass

`func (o *CartCreate) GetDwApiPass() string`

GetDwApiPass returns the DwApiPass field if non-nil, zero value otherwise.

### GetDwApiPassOk

`func (o *CartCreate) GetDwApiPassOk() (*string, bool)`

GetDwApiPassOk returns a tuple with the DwApiPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDwApiPass

`func (o *CartCreate) SetDwApiPass(v string)`

SetDwApiPass sets DwApiPass field to given value.

### HasDwApiPass

`func (o *CartCreate) HasDwApiPass() bool`

HasDwApiPass returns a boolean if a field has been set.

### GetDemandwareUserName

`func (o *CartCreate) GetDemandwareUserName() string`

GetDemandwareUserName returns the DemandwareUserName field if non-nil, zero value otherwise.

### GetDemandwareUserNameOk

`func (o *CartCreate) GetDemandwareUserNameOk() (*string, bool)`

GetDemandwareUserNameOk returns a tuple with the DemandwareUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemandwareUserName

`func (o *CartCreate) SetDemandwareUserName(v string)`

SetDemandwareUserName sets DemandwareUserName field to given value.

### HasDemandwareUserName

`func (o *CartCreate) HasDemandwareUserName() bool`

HasDemandwareUserName returns a boolean if a field has been set.

### GetDemandwareUserPassword

`func (o *CartCreate) GetDemandwareUserPassword() string`

GetDemandwareUserPassword returns the DemandwareUserPassword field if non-nil, zero value otherwise.

### GetDemandwareUserPasswordOk

`func (o *CartCreate) GetDemandwareUserPasswordOk() (*string, bool)`

GetDemandwareUserPasswordOk returns a tuple with the DemandwareUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemandwareUserPassword

`func (o *CartCreate) SetDemandwareUserPassword(v string)`

SetDemandwareUserPassword sets DemandwareUserPassword field to given value.

### HasDemandwareUserPassword

`func (o *CartCreate) HasDemandwareUserPassword() bool`

HasDemandwareUserPassword returns a boolean if a field has been set.

### GetStoreId

`func (o *CartCreate) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *CartCreate) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *CartCreate) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.


### GetSellerId

`func (o *CartCreate) GetSellerId() string`

GetSellerId returns the SellerId field if non-nil, zero value otherwise.

### GetSellerIdOk

`func (o *CartCreate) GetSellerIdOk() (*string, bool)`

GetSellerIdOk returns a tuple with the SellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerId

`func (o *CartCreate) SetSellerId(v string)`

SetSellerId sets SellerId field to given value.

### HasSellerId

`func (o *CartCreate) HasSellerId() bool`

HasSellerId returns a boolean if a field has been set.

### GetEnvironment

`func (o *CartCreate) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CartCreate) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CartCreate) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CartCreate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetHybrisClientId

`func (o *CartCreate) GetHybrisClientId() string`

GetHybrisClientId returns the HybrisClientId field if non-nil, zero value otherwise.

### GetHybrisClientIdOk

`func (o *CartCreate) GetHybrisClientIdOk() (*string, bool)`

GetHybrisClientIdOk returns a tuple with the HybrisClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybrisClientId

`func (o *CartCreate) SetHybrisClientId(v string)`

SetHybrisClientId sets HybrisClientId field to given value.

### HasHybrisClientId

`func (o *CartCreate) HasHybrisClientId() bool`

HasHybrisClientId returns a boolean if a field has been set.

### GetHybrisClientSecret

`func (o *CartCreate) GetHybrisClientSecret() string`

GetHybrisClientSecret returns the HybrisClientSecret field if non-nil, zero value otherwise.

### GetHybrisClientSecretOk

`func (o *CartCreate) GetHybrisClientSecretOk() (*string, bool)`

GetHybrisClientSecretOk returns a tuple with the HybrisClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybrisClientSecret

`func (o *CartCreate) SetHybrisClientSecret(v string)`

SetHybrisClientSecret sets HybrisClientSecret field to given value.

### HasHybrisClientSecret

`func (o *CartCreate) HasHybrisClientSecret() bool`

HasHybrisClientSecret returns a boolean if a field has been set.

### GetHybrisUsername

`func (o *CartCreate) GetHybrisUsername() string`

GetHybrisUsername returns the HybrisUsername field if non-nil, zero value otherwise.

### GetHybrisUsernameOk

`func (o *CartCreate) GetHybrisUsernameOk() (*string, bool)`

GetHybrisUsernameOk returns a tuple with the HybrisUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybrisUsername

`func (o *CartCreate) SetHybrisUsername(v string)`

SetHybrisUsername sets HybrisUsername field to given value.

### HasHybrisUsername

`func (o *CartCreate) HasHybrisUsername() bool`

HasHybrisUsername returns a boolean if a field has been set.

### GetHybrisPassword

`func (o *CartCreate) GetHybrisPassword() string`

GetHybrisPassword returns the HybrisPassword field if non-nil, zero value otherwise.

### GetHybrisPasswordOk

`func (o *CartCreate) GetHybrisPasswordOk() (*string, bool)`

GetHybrisPasswordOk returns a tuple with the HybrisPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybrisPassword

`func (o *CartCreate) SetHybrisPassword(v string)`

SetHybrisPassword sets HybrisPassword field to given value.

### HasHybrisPassword

`func (o *CartCreate) HasHybrisPassword() bool`

HasHybrisPassword returns a boolean if a field has been set.

### GetHybrisWebsites

`func (o *CartCreate) GetHybrisWebsites() []AccountCartAddHybrisWebsitesInner`

GetHybrisWebsites returns the HybrisWebsites field if non-nil, zero value otherwise.

### GetHybrisWebsitesOk

`func (o *CartCreate) GetHybrisWebsitesOk() (*[]AccountCartAddHybrisWebsitesInner, bool)`

GetHybrisWebsitesOk returns a tuple with the HybrisWebsites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybrisWebsites

`func (o *CartCreate) SetHybrisWebsites(v []AccountCartAddHybrisWebsitesInner)`

SetHybrisWebsites sets HybrisWebsites field to given value.

### HasHybrisWebsites

`func (o *CartCreate) HasHybrisWebsites() bool`

HasHybrisWebsites returns a boolean if a field has been set.

### GetWalmartClientId

`func (o *CartCreate) GetWalmartClientId() string`

GetWalmartClientId returns the WalmartClientId field if non-nil, zero value otherwise.

### GetWalmartClientIdOk

`func (o *CartCreate) GetWalmartClientIdOk() (*string, bool)`

GetWalmartClientIdOk returns a tuple with the WalmartClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalmartClientId

`func (o *CartCreate) SetWalmartClientId(v string)`

SetWalmartClientId sets WalmartClientId field to given value.

### HasWalmartClientId

`func (o *CartCreate) HasWalmartClientId() bool`

HasWalmartClientId returns a boolean if a field has been set.

### GetWalmartClientSecret

`func (o *CartCreate) GetWalmartClientSecret() string`

GetWalmartClientSecret returns the WalmartClientSecret field if non-nil, zero value otherwise.

### GetWalmartClientSecretOk

`func (o *CartCreate) GetWalmartClientSecretOk() (*string, bool)`

GetWalmartClientSecretOk returns a tuple with the WalmartClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalmartClientSecret

`func (o *CartCreate) SetWalmartClientSecret(v string)`

SetWalmartClientSecret sets WalmartClientSecret field to given value.

### HasWalmartClientSecret

`func (o *CartCreate) HasWalmartClientSecret() bool`

HasWalmartClientSecret returns a boolean if a field has been set.

### GetWalmartEnvironment

`func (o *CartCreate) GetWalmartEnvironment() string`

GetWalmartEnvironment returns the WalmartEnvironment field if non-nil, zero value otherwise.

### GetWalmartEnvironmentOk

`func (o *CartCreate) GetWalmartEnvironmentOk() (*string, bool)`

GetWalmartEnvironmentOk returns a tuple with the WalmartEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalmartEnvironment

`func (o *CartCreate) SetWalmartEnvironment(v string)`

SetWalmartEnvironment sets WalmartEnvironment field to given value.

### HasWalmartEnvironment

`func (o *CartCreate) HasWalmartEnvironment() bool`

HasWalmartEnvironment returns a boolean if a field has been set.

### GetWalmartChannelType

`func (o *CartCreate) GetWalmartChannelType() string`

GetWalmartChannelType returns the WalmartChannelType field if non-nil, zero value otherwise.

### GetWalmartChannelTypeOk

`func (o *CartCreate) GetWalmartChannelTypeOk() (*string, bool)`

GetWalmartChannelTypeOk returns a tuple with the WalmartChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalmartChannelType

`func (o *CartCreate) SetWalmartChannelType(v string)`

SetWalmartChannelType sets WalmartChannelType field to given value.

### HasWalmartChannelType

`func (o *CartCreate) HasWalmartChannelType() bool`

HasWalmartChannelType returns a boolean if a field has been set.

### GetWalmartRegion

`func (o *CartCreate) GetWalmartRegion() string`

GetWalmartRegion returns the WalmartRegion field if non-nil, zero value otherwise.

### GetWalmartRegionOk

`func (o *CartCreate) GetWalmartRegionOk() (*string, bool)`

GetWalmartRegionOk returns a tuple with the WalmartRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalmartRegion

`func (o *CartCreate) SetWalmartRegion(v string)`

SetWalmartRegion sets WalmartRegion field to given value.

### HasWalmartRegion

`func (o *CartCreate) HasWalmartRegion() bool`

HasWalmartRegion returns a boolean if a field has been set.

### GetLightspeedApiKey

`func (o *CartCreate) GetLightspeedApiKey() string`

GetLightspeedApiKey returns the LightspeedApiKey field if non-nil, zero value otherwise.

### GetLightspeedApiKeyOk

`func (o *CartCreate) GetLightspeedApiKeyOk() (*string, bool)`

GetLightspeedApiKeyOk returns a tuple with the LightspeedApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLightspeedApiKey

`func (o *CartCreate) SetLightspeedApiKey(v string)`

SetLightspeedApiKey sets LightspeedApiKey field to given value.

### HasLightspeedApiKey

`func (o *CartCreate) HasLightspeedApiKey() bool`

HasLightspeedApiKey returns a boolean if a field has been set.

### GetLightspeedApiSecret

`func (o *CartCreate) GetLightspeedApiSecret() string`

GetLightspeedApiSecret returns the LightspeedApiSecret field if non-nil, zero value otherwise.

### GetLightspeedApiSecretOk

`func (o *CartCreate) GetLightspeedApiSecretOk() (*string, bool)`

GetLightspeedApiSecretOk returns a tuple with the LightspeedApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLightspeedApiSecret

`func (o *CartCreate) SetLightspeedApiSecret(v string)`

SetLightspeedApiSecret sets LightspeedApiSecret field to given value.

### HasLightspeedApiSecret

`func (o *CartCreate) HasLightspeedApiSecret() bool`

HasLightspeedApiSecret returns a boolean if a field has been set.

### GetShoplazzaAccessToken

`func (o *CartCreate) GetShoplazzaAccessToken() string`

GetShoplazzaAccessToken returns the ShoplazzaAccessToken field if non-nil, zero value otherwise.

### GetShoplazzaAccessTokenOk

`func (o *CartCreate) GetShoplazzaAccessTokenOk() (*string, bool)`

GetShoplazzaAccessTokenOk returns a tuple with the ShoplazzaAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShoplazzaAccessToken

`func (o *CartCreate) SetShoplazzaAccessToken(v string)`

SetShoplazzaAccessToken sets ShoplazzaAccessToken field to given value.

### HasShoplazzaAccessToken

`func (o *CartCreate) HasShoplazzaAccessToken() bool`

HasShoplazzaAccessToken returns a boolean if a field has been set.

### GetShoplazzaSharedSecret

`func (o *CartCreate) GetShoplazzaSharedSecret() string`

GetShoplazzaSharedSecret returns the ShoplazzaSharedSecret field if non-nil, zero value otherwise.

### GetShoplazzaSharedSecretOk

`func (o *CartCreate) GetShoplazzaSharedSecretOk() (*string, bool)`

GetShoplazzaSharedSecretOk returns a tuple with the ShoplazzaSharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShoplazzaSharedSecret

`func (o *CartCreate) SetShoplazzaSharedSecret(v string)`

SetShoplazzaSharedSecret sets ShoplazzaSharedSecret field to given value.

### HasShoplazzaSharedSecret

`func (o *CartCreate) HasShoplazzaSharedSecret() bool`

HasShoplazzaSharedSecret returns a boolean if a field has been set.

### GetShopwareAccessKey

`func (o *CartCreate) GetShopwareAccessKey() string`

GetShopwareAccessKey returns the ShopwareAccessKey field if non-nil, zero value otherwise.

### GetShopwareAccessKeyOk

`func (o *CartCreate) GetShopwareAccessKeyOk() (*string, bool)`

GetShopwareAccessKeyOk returns a tuple with the ShopwareAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopwareAccessKey

`func (o *CartCreate) SetShopwareAccessKey(v string)`

SetShopwareAccessKey sets ShopwareAccessKey field to given value.

### HasShopwareAccessKey

`func (o *CartCreate) HasShopwareAccessKey() bool`

HasShopwareAccessKey returns a boolean if a field has been set.

### GetShopwareApiKey

`func (o *CartCreate) GetShopwareApiKey() string`

GetShopwareApiKey returns the ShopwareApiKey field if non-nil, zero value otherwise.

### GetShopwareApiKeyOk

`func (o *CartCreate) GetShopwareApiKeyOk() (*string, bool)`

GetShopwareApiKeyOk returns a tuple with the ShopwareApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopwareApiKey

`func (o *CartCreate) SetShopwareApiKey(v string)`

SetShopwareApiKey sets ShopwareApiKey field to given value.

### HasShopwareApiKey

`func (o *CartCreate) HasShopwareApiKey() bool`

HasShopwareApiKey returns a boolean if a field has been set.

### GetShopwareApiSecret

`func (o *CartCreate) GetShopwareApiSecret() string`

GetShopwareApiSecret returns the ShopwareApiSecret field if non-nil, zero value otherwise.

### GetShopwareApiSecretOk

`func (o *CartCreate) GetShopwareApiSecretOk() (*string, bool)`

GetShopwareApiSecretOk returns a tuple with the ShopwareApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopwareApiSecret

`func (o *CartCreate) SetShopwareApiSecret(v string)`

SetShopwareApiSecret sets ShopwareApiSecret field to given value.

### HasShopwareApiSecret

`func (o *CartCreate) HasShopwareApiSecret() bool`

HasShopwareApiSecret returns a boolean if a field has been set.

### GetCommercehqApiKey

`func (o *CartCreate) GetCommercehqApiKey() string`

GetCommercehqApiKey returns the CommercehqApiKey field if non-nil, zero value otherwise.

### GetCommercehqApiKeyOk

`func (o *CartCreate) GetCommercehqApiKeyOk() (*string, bool)`

GetCommercehqApiKeyOk returns a tuple with the CommercehqApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercehqApiKey

`func (o *CartCreate) SetCommercehqApiKey(v string)`

SetCommercehqApiKey sets CommercehqApiKey field to given value.

### HasCommercehqApiKey

`func (o *CartCreate) HasCommercehqApiKey() bool`

HasCommercehqApiKey returns a boolean if a field has been set.

### GetCommercehqApiPassword

`func (o *CartCreate) GetCommercehqApiPassword() string`

GetCommercehqApiPassword returns the CommercehqApiPassword field if non-nil, zero value otherwise.

### GetCommercehqApiPasswordOk

`func (o *CartCreate) GetCommercehqApiPasswordOk() (*string, bool)`

GetCommercehqApiPasswordOk returns a tuple with the CommercehqApiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercehqApiPassword

`func (o *CartCreate) SetCommercehqApiPassword(v string)`

SetCommercehqApiPassword sets CommercehqApiPassword field to given value.

### HasCommercehqApiPassword

`func (o *CartCreate) HasCommercehqApiPassword() bool`

HasCommercehqApiPassword returns a boolean if a field has been set.

### GetVar3dcartPrivateKey

`func (o *CartCreate) GetVar3dcartPrivateKey() string`

GetVar3dcartPrivateKey returns the Var3dcartPrivateKey field if non-nil, zero value otherwise.

### GetVar3dcartPrivateKeyOk

`func (o *CartCreate) GetVar3dcartPrivateKeyOk() (*string, bool)`

GetVar3dcartPrivateKeyOk returns a tuple with the Var3dcartPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3dcartPrivateKey

`func (o *CartCreate) SetVar3dcartPrivateKey(v string)`

SetVar3dcartPrivateKey sets Var3dcartPrivateKey field to given value.

### HasVar3dcartPrivateKey

`func (o *CartCreate) HasVar3dcartPrivateKey() bool`

HasVar3dcartPrivateKey returns a boolean if a field has been set.

### GetVar3dcartAccessToken

`func (o *CartCreate) GetVar3dcartAccessToken() string`

GetVar3dcartAccessToken returns the Var3dcartAccessToken field if non-nil, zero value otherwise.

### GetVar3dcartAccessTokenOk

`func (o *CartCreate) GetVar3dcartAccessTokenOk() (*string, bool)`

GetVar3dcartAccessTokenOk returns a tuple with the Var3dcartAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3dcartAccessToken

`func (o *CartCreate) SetVar3dcartAccessToken(v string)`

SetVar3dcartAccessToken sets Var3dcartAccessToken field to given value.

### HasVar3dcartAccessToken

`func (o *CartCreate) HasVar3dcartAccessToken() bool`

HasVar3dcartAccessToken returns a boolean if a field has been set.

### GetWcConsumerKey

`func (o *CartCreate) GetWcConsumerKey() string`

GetWcConsumerKey returns the WcConsumerKey field if non-nil, zero value otherwise.

### GetWcConsumerKeyOk

`func (o *CartCreate) GetWcConsumerKeyOk() (*string, bool)`

GetWcConsumerKeyOk returns a tuple with the WcConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWcConsumerKey

`func (o *CartCreate) SetWcConsumerKey(v string)`

SetWcConsumerKey sets WcConsumerKey field to given value.

### HasWcConsumerKey

`func (o *CartCreate) HasWcConsumerKey() bool`

HasWcConsumerKey returns a boolean if a field has been set.

### GetWcConsumerSecret

`func (o *CartCreate) GetWcConsumerSecret() string`

GetWcConsumerSecret returns the WcConsumerSecret field if non-nil, zero value otherwise.

### GetWcConsumerSecretOk

`func (o *CartCreate) GetWcConsumerSecretOk() (*string, bool)`

GetWcConsumerSecretOk returns a tuple with the WcConsumerSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWcConsumerSecret

`func (o *CartCreate) SetWcConsumerSecret(v string)`

SetWcConsumerSecret sets WcConsumerSecret field to given value.

### HasWcConsumerSecret

`func (o *CartCreate) HasWcConsumerSecret() bool`

HasWcConsumerSecret returns a boolean if a field has been set.

### GetMagentoConsumerKey

`func (o *CartCreate) GetMagentoConsumerKey() string`

GetMagentoConsumerKey returns the MagentoConsumerKey field if non-nil, zero value otherwise.

### GetMagentoConsumerKeyOk

`func (o *CartCreate) GetMagentoConsumerKeyOk() (*string, bool)`

GetMagentoConsumerKeyOk returns a tuple with the MagentoConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagentoConsumerKey

`func (o *CartCreate) SetMagentoConsumerKey(v string)`

SetMagentoConsumerKey sets MagentoConsumerKey field to given value.

### HasMagentoConsumerKey

`func (o *CartCreate) HasMagentoConsumerKey() bool`

HasMagentoConsumerKey returns a boolean if a field has been set.

### GetMagentoConsumerSecret

`func (o *CartCreate) GetMagentoConsumerSecret() string`

GetMagentoConsumerSecret returns the MagentoConsumerSecret field if non-nil, zero value otherwise.

### GetMagentoConsumerSecretOk

`func (o *CartCreate) GetMagentoConsumerSecretOk() (*string, bool)`

GetMagentoConsumerSecretOk returns a tuple with the MagentoConsumerSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagentoConsumerSecret

`func (o *CartCreate) SetMagentoConsumerSecret(v string)`

SetMagentoConsumerSecret sets MagentoConsumerSecret field to given value.

### HasMagentoConsumerSecret

`func (o *CartCreate) HasMagentoConsumerSecret() bool`

HasMagentoConsumerSecret returns a boolean if a field has been set.

### GetMagentoAccessToken

`func (o *CartCreate) GetMagentoAccessToken() string`

GetMagentoAccessToken returns the MagentoAccessToken field if non-nil, zero value otherwise.

### GetMagentoAccessTokenOk

`func (o *CartCreate) GetMagentoAccessTokenOk() (*string, bool)`

GetMagentoAccessTokenOk returns a tuple with the MagentoAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagentoAccessToken

`func (o *CartCreate) SetMagentoAccessToken(v string)`

SetMagentoAccessToken sets MagentoAccessToken field to given value.

### HasMagentoAccessToken

`func (o *CartCreate) HasMagentoAccessToken() bool`

HasMagentoAccessToken returns a boolean if a field has been set.

### GetMagentoTokenSecret

`func (o *CartCreate) GetMagentoTokenSecret() string`

GetMagentoTokenSecret returns the MagentoTokenSecret field if non-nil, zero value otherwise.

### GetMagentoTokenSecretOk

`func (o *CartCreate) GetMagentoTokenSecretOk() (*string, bool)`

GetMagentoTokenSecretOk returns a tuple with the MagentoTokenSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagentoTokenSecret

`func (o *CartCreate) SetMagentoTokenSecret(v string)`

SetMagentoTokenSecret sets MagentoTokenSecret field to given value.

### HasMagentoTokenSecret

`func (o *CartCreate) HasMagentoTokenSecret() bool`

HasMagentoTokenSecret returns a boolean if a field has been set.

### GetPrestashopWebserviceKey

`func (o *CartCreate) GetPrestashopWebserviceKey() string`

GetPrestashopWebserviceKey returns the PrestashopWebserviceKey field if non-nil, zero value otherwise.

### GetPrestashopWebserviceKeyOk

`func (o *CartCreate) GetPrestashopWebserviceKeyOk() (*string, bool)`

GetPrestashopWebserviceKeyOk returns a tuple with the PrestashopWebserviceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrestashopWebserviceKey

`func (o *CartCreate) SetPrestashopWebserviceKey(v string)`

SetPrestashopWebserviceKey sets PrestashopWebserviceKey field to given value.

### HasPrestashopWebserviceKey

`func (o *CartCreate) HasPrestashopWebserviceKey() bool`

HasPrestashopWebserviceKey returns a boolean if a field has been set.

### GetWixAppId

`func (o *CartCreate) GetWixAppId() string`

GetWixAppId returns the WixAppId field if non-nil, zero value otherwise.

### GetWixAppIdOk

`func (o *CartCreate) GetWixAppIdOk() (*string, bool)`

GetWixAppIdOk returns a tuple with the WixAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWixAppId

`func (o *CartCreate) SetWixAppId(v string)`

SetWixAppId sets WixAppId field to given value.


### GetWixAppSecretKey

`func (o *CartCreate) GetWixAppSecretKey() string`

GetWixAppSecretKey returns the WixAppSecretKey field if non-nil, zero value otherwise.

### GetWixAppSecretKeyOk

`func (o *CartCreate) GetWixAppSecretKeyOk() (*string, bool)`

GetWixAppSecretKeyOk returns a tuple with the WixAppSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWixAppSecretKey

`func (o *CartCreate) SetWixAppSecretKey(v string)`

SetWixAppSecretKey sets WixAppSecretKey field to given value.


### GetWixInstanceId

`func (o *CartCreate) GetWixInstanceId() string`

GetWixInstanceId returns the WixInstanceId field if non-nil, zero value otherwise.

### GetWixInstanceIdOk

`func (o *CartCreate) GetWixInstanceIdOk() (*string, bool)`

GetWixInstanceIdOk returns a tuple with the WixInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWixInstanceId

`func (o *CartCreate) SetWixInstanceId(v string)`

SetWixInstanceId sets WixInstanceId field to given value.

### HasWixInstanceId

`func (o *CartCreate) HasWixInstanceId() bool`

HasWixInstanceId returns a boolean if a field has been set.

### GetWixRefreshToken

`func (o *CartCreate) GetWixRefreshToken() string`

GetWixRefreshToken returns the WixRefreshToken field if non-nil, zero value otherwise.

### GetWixRefreshTokenOk

`func (o *CartCreate) GetWixRefreshTokenOk() (*string, bool)`

GetWixRefreshTokenOk returns a tuple with the WixRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWixRefreshToken

`func (o *CartCreate) SetWixRefreshToken(v string)`

SetWixRefreshToken sets WixRefreshToken field to given value.

### HasWixRefreshToken

`func (o *CartCreate) HasWixRefreshToken() bool`

HasWixRefreshToken returns a boolean if a field has been set.

### GetMercadoLibreAppId

`func (o *CartCreate) GetMercadoLibreAppId() string`

GetMercadoLibreAppId returns the MercadoLibreAppId field if non-nil, zero value otherwise.

### GetMercadoLibreAppIdOk

`func (o *CartCreate) GetMercadoLibreAppIdOk() (*string, bool)`

GetMercadoLibreAppIdOk returns a tuple with the MercadoLibreAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMercadoLibreAppId

`func (o *CartCreate) SetMercadoLibreAppId(v string)`

SetMercadoLibreAppId sets MercadoLibreAppId field to given value.

### HasMercadoLibreAppId

`func (o *CartCreate) HasMercadoLibreAppId() bool`

HasMercadoLibreAppId returns a boolean if a field has been set.

### GetMercadoLibreAppSecretKey

`func (o *CartCreate) GetMercadoLibreAppSecretKey() string`

GetMercadoLibreAppSecretKey returns the MercadoLibreAppSecretKey field if non-nil, zero value otherwise.

### GetMercadoLibreAppSecretKeyOk

`func (o *CartCreate) GetMercadoLibreAppSecretKeyOk() (*string, bool)`

GetMercadoLibreAppSecretKeyOk returns a tuple with the MercadoLibreAppSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMercadoLibreAppSecretKey

`func (o *CartCreate) SetMercadoLibreAppSecretKey(v string)`

SetMercadoLibreAppSecretKey sets MercadoLibreAppSecretKey field to given value.

### HasMercadoLibreAppSecretKey

`func (o *CartCreate) HasMercadoLibreAppSecretKey() bool`

HasMercadoLibreAppSecretKey returns a boolean if a field has been set.

### GetMercadoLibreRefreshToken

`func (o *CartCreate) GetMercadoLibreRefreshToken() string`

GetMercadoLibreRefreshToken returns the MercadoLibreRefreshToken field if non-nil, zero value otherwise.

### GetMercadoLibreRefreshTokenOk

`func (o *CartCreate) GetMercadoLibreRefreshTokenOk() (*string, bool)`

GetMercadoLibreRefreshTokenOk returns a tuple with the MercadoLibreRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMercadoLibreRefreshToken

`func (o *CartCreate) SetMercadoLibreRefreshToken(v string)`

SetMercadoLibreRefreshToken sets MercadoLibreRefreshToken field to given value.

### HasMercadoLibreRefreshToken

`func (o *CartCreate) HasMercadoLibreRefreshToken() bool`

HasMercadoLibreRefreshToken returns a boolean if a field has been set.

### GetZidClientId

`func (o *CartCreate) GetZidClientId() int32`

GetZidClientId returns the ZidClientId field if non-nil, zero value otherwise.

### GetZidClientIdOk

`func (o *CartCreate) GetZidClientIdOk() (*int32, bool)`

GetZidClientIdOk returns a tuple with the ZidClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZidClientId

`func (o *CartCreate) SetZidClientId(v int32)`

SetZidClientId sets ZidClientId field to given value.

### HasZidClientId

`func (o *CartCreate) HasZidClientId() bool`

HasZidClientId returns a boolean if a field has been set.

### GetZidClientSecret

`func (o *CartCreate) GetZidClientSecret() string`

GetZidClientSecret returns the ZidClientSecret field if non-nil, zero value otherwise.

### GetZidClientSecretOk

`func (o *CartCreate) GetZidClientSecretOk() (*string, bool)`

GetZidClientSecretOk returns a tuple with the ZidClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZidClientSecret

`func (o *CartCreate) SetZidClientSecret(v string)`

SetZidClientSecret sets ZidClientSecret field to given value.

### HasZidClientSecret

`func (o *CartCreate) HasZidClientSecret() bool`

HasZidClientSecret returns a boolean if a field has been set.

### GetZidAccessToken

`func (o *CartCreate) GetZidAccessToken() string`

GetZidAccessToken returns the ZidAccessToken field if non-nil, zero value otherwise.

### GetZidAccessTokenOk

`func (o *CartCreate) GetZidAccessTokenOk() (*string, bool)`

GetZidAccessTokenOk returns a tuple with the ZidAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZidAccessToken

`func (o *CartCreate) SetZidAccessToken(v string)`

SetZidAccessToken sets ZidAccessToken field to given value.

### HasZidAccessToken

`func (o *CartCreate) HasZidAccessToken() bool`

HasZidAccessToken returns a boolean if a field has been set.

### GetZidAuthorization

`func (o *CartCreate) GetZidAuthorization() string`

GetZidAuthorization returns the ZidAuthorization field if non-nil, zero value otherwise.

### GetZidAuthorizationOk

`func (o *CartCreate) GetZidAuthorizationOk() (*string, bool)`

GetZidAuthorizationOk returns a tuple with the ZidAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZidAuthorization

`func (o *CartCreate) SetZidAuthorization(v string)`

SetZidAuthorization sets ZidAuthorization field to given value.

### HasZidAuthorization

`func (o *CartCreate) HasZidAuthorization() bool`

HasZidAuthorization returns a boolean if a field has been set.

### GetZidRefreshToken

`func (o *CartCreate) GetZidRefreshToken() string`

GetZidRefreshToken returns the ZidRefreshToken field if non-nil, zero value otherwise.

### GetZidRefreshTokenOk

`func (o *CartCreate) GetZidRefreshTokenOk() (*string, bool)`

GetZidRefreshTokenOk returns a tuple with the ZidRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZidRefreshToken

`func (o *CartCreate) SetZidRefreshToken(v string)`

SetZidRefreshToken sets ZidRefreshToken field to given value.

### HasZidRefreshToken

`func (o *CartCreate) HasZidRefreshToken() bool`

HasZidRefreshToken returns a boolean if a field has been set.

### GetFlipkartClientId

`func (o *CartCreate) GetFlipkartClientId() string`

GetFlipkartClientId returns the FlipkartClientId field if non-nil, zero value otherwise.

### GetFlipkartClientIdOk

`func (o *CartCreate) GetFlipkartClientIdOk() (*string, bool)`

GetFlipkartClientIdOk returns a tuple with the FlipkartClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlipkartClientId

`func (o *CartCreate) SetFlipkartClientId(v string)`

SetFlipkartClientId sets FlipkartClientId field to given value.

### HasFlipkartClientId

`func (o *CartCreate) HasFlipkartClientId() bool`

HasFlipkartClientId returns a boolean if a field has been set.

### GetFlipkartClientSecret

`func (o *CartCreate) GetFlipkartClientSecret() string`

GetFlipkartClientSecret returns the FlipkartClientSecret field if non-nil, zero value otherwise.

### GetFlipkartClientSecretOk

`func (o *CartCreate) GetFlipkartClientSecretOk() (*string, bool)`

GetFlipkartClientSecretOk returns a tuple with the FlipkartClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlipkartClientSecret

`func (o *CartCreate) SetFlipkartClientSecret(v string)`

SetFlipkartClientSecret sets FlipkartClientSecret field to given value.

### HasFlipkartClientSecret

`func (o *CartCreate) HasFlipkartClientSecret() bool`

HasFlipkartClientSecret returns a boolean if a field has been set.

### GetAllegroClientId

`func (o *CartCreate) GetAllegroClientId() string`

GetAllegroClientId returns the AllegroClientId field if non-nil, zero value otherwise.

### GetAllegroClientIdOk

`func (o *CartCreate) GetAllegroClientIdOk() (*string, bool)`

GetAllegroClientIdOk returns a tuple with the AllegroClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllegroClientId

`func (o *CartCreate) SetAllegroClientId(v string)`

SetAllegroClientId sets AllegroClientId field to given value.

### HasAllegroClientId

`func (o *CartCreate) HasAllegroClientId() bool`

HasAllegroClientId returns a boolean if a field has been set.

### GetAllegroClientSecret

`func (o *CartCreate) GetAllegroClientSecret() string`

GetAllegroClientSecret returns the AllegroClientSecret field if non-nil, zero value otherwise.

### GetAllegroClientSecretOk

`func (o *CartCreate) GetAllegroClientSecretOk() (*string, bool)`

GetAllegroClientSecretOk returns a tuple with the AllegroClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllegroClientSecret

`func (o *CartCreate) SetAllegroClientSecret(v string)`

SetAllegroClientSecret sets AllegroClientSecret field to given value.

### HasAllegroClientSecret

`func (o *CartCreate) HasAllegroClientSecret() bool`

HasAllegroClientSecret returns a boolean if a field has been set.

### GetAllegroAccessToken

`func (o *CartCreate) GetAllegroAccessToken() string`

GetAllegroAccessToken returns the AllegroAccessToken field if non-nil, zero value otherwise.

### GetAllegroAccessTokenOk

`func (o *CartCreate) GetAllegroAccessTokenOk() (*string, bool)`

GetAllegroAccessTokenOk returns a tuple with the AllegroAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllegroAccessToken

`func (o *CartCreate) SetAllegroAccessToken(v string)`

SetAllegroAccessToken sets AllegroAccessToken field to given value.

### HasAllegroAccessToken

`func (o *CartCreate) HasAllegroAccessToken() bool`

HasAllegroAccessToken returns a boolean if a field has been set.

### GetAllegroRefreshToken

`func (o *CartCreate) GetAllegroRefreshToken() string`

GetAllegroRefreshToken returns the AllegroRefreshToken field if non-nil, zero value otherwise.

### GetAllegroRefreshTokenOk

`func (o *CartCreate) GetAllegroRefreshTokenOk() (*string, bool)`

GetAllegroRefreshTokenOk returns a tuple with the AllegroRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllegroRefreshToken

`func (o *CartCreate) SetAllegroRefreshToken(v string)`

SetAllegroRefreshToken sets AllegroRefreshToken field to given value.

### HasAllegroRefreshToken

`func (o *CartCreate) HasAllegroRefreshToken() bool`

HasAllegroRefreshToken returns a boolean if a field has been set.

### GetAllegroEnvironment

`func (o *CartCreate) GetAllegroEnvironment() string`

GetAllegroEnvironment returns the AllegroEnvironment field if non-nil, zero value otherwise.

### GetAllegroEnvironmentOk

`func (o *CartCreate) GetAllegroEnvironmentOk() (*string, bool)`

GetAllegroEnvironmentOk returns a tuple with the AllegroEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllegroEnvironment

`func (o *CartCreate) SetAllegroEnvironment(v string)`

SetAllegroEnvironment sets AllegroEnvironment field to given value.

### HasAllegroEnvironment

`func (o *CartCreate) HasAllegroEnvironment() bool`

HasAllegroEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


