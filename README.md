# github.com/bllli/logto-management-api

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/bllli/logto-management-api* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/bllli/logto-management-api&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/wtf/temp). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Logto API references: API references for Logto services.

Note: The documentation is for Logto Cloud. If you are using Logto OSS, please refer to the response of `/api/swagger.json` endpoint on your Logto instance.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/bllli/logto-management-api](#githubcomblllilogto-management-api)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/bllli/logto-management-api
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"github.com/bllli/logto-management-api/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := logtomanagementapi.New(
		logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
	)

	res, err := s.Applications.Get(ctx, operations.ListApplicationsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseBodies != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type | Scheme      | Environment Variable             |
| ------------ | ---- | ----------- | -------------------------------- |
| `BearerAuth` | http | HTTP Bearer | `LOGTOMANAGEMENTAPI_BEARER_AUTH` |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"github.com/bllli/logto-management-api/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := logtomanagementapi.New(
		logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
	)

	res, err := s.Applications.Get(ctx, operations.ListApplicationsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseBodies != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [AccountCenter](docs/sdks/accountcenter/README.md)

* [Get](docs/sdks/accountcenter/README.md#get) - Get account center settings
* [UpdateSettings](docs/sdks/accountcenter/README.md#updatesettings) - Update account center settings

### [Applications](docs/sdks/applications/README.md)

* [Get](docs/sdks/applications/README.md#get) - Get applications
* [Create](docs/sdks/applications/README.md#create) - Create an application
* [Fetch](docs/sdks/applications/README.md#fetch) - Get application
* [Update](docs/sdks/applications/README.md#update) - Update application
* [Delete](docs/sdks/applications/README.md#delete) - Delete application
* [UpdateCustomData](docs/sdks/applications/README.md#updatecustomdata) - Update application custom data
* [ListRoles](docs/sdks/applications/README.md#listroles) - Get application API resource roles
* [AssignRoles](docs/sdks/applications/README.md#assignroles) - Assign API resource roles to application
* [ReplaceRoles](docs/sdks/applications/README.md#replaceroles) - Update API resource roles for application
* [DeleteRole](docs/sdks/applications/README.md#deleterole) - Remove a API resource role from application
* [GetProtectedAppMetadataCustomDomains](docs/sdks/applications/README.md#getprotectedappmetadatacustomdomains) - Get application custom domains.
* [AddCustomDomain](docs/sdks/applications/README.md#addcustomdomain) - Add a custom domain to the application.
* [RemoveCustomDomain](docs/sdks/applications/README.md#removecustomdomain) - Remove custom domain.
* [ListOrganizations](docs/sdks/applications/README.md#listorganizations) - Get application organizations
* [DeleteLegacySecret](docs/sdks/applications/README.md#deletelegacysecret) - Delete application legacy secret
* [GetSecrets](docs/sdks/applications/README.md#getsecrets) - Get application secrets
* [AddSecret](docs/sdks/applications/README.md#addsecret) - Add application secret
* [DeleteSecret](docs/sdks/applications/README.md#deletesecret) - Delete application secret
* [UpdateSecret](docs/sdks/applications/README.md#updatesecret) - Update application secret
* [CreateUserConsentScope](docs/sdks/applications/README.md#createuserconsentscope) - Assign user consent scopes to application.
* [ListUserConsentScopes](docs/sdks/applications/README.md#listuserconsentscopes) - List all the user consent scopes of an application.
* [DeleteUserConsentScope](docs/sdks/applications/README.md#deleteuserconsentscope) - Remove user consent scope from application.
* [UpdateSignInExperience](docs/sdks/applications/README.md#updatesigninexperience) - Update application level sign-in experience
* [GetSignInExperience](docs/sdks/applications/README.md#getsigninexperience) - Get the application level sign-in experience
* [ListUserConsentOrganizations](docs/sdks/applications/README.md#listuserconsentorganizations) - List all the user consented organizations of a application.
* [ReplaceUserConsentOrganizations](docs/sdks/applications/README.md#replaceuserconsentorganizations) - Grant a list of organization access of a user for a application.
* [GrantUserConsent](docs/sdks/applications/README.md#grantuserconsent) - Grant a list of organization access of a user for a application.
* [RevokeUserConsent](docs/sdks/applications/README.md#revokeuserconsent) - Revoke a user's access to an organization for a application.

### [AuditLogs](docs/sdks/auditlogs/README.md)

* [List](docs/sdks/auditlogs/README.md#list) - Get logs
* [Get](docs/sdks/auditlogs/README.md#get) - Get log

### [Authn](docs/sdks/authn/README.md)

* [AssertSingleSignOnSaml](docs/sdks/authn/README.md#assertsinglesignonsaml) - SAML ACS endpoint (SSO)

### [AuthnHasura](docs/sdks/authnhasura/README.md)

* [Get](docs/sdks/authnhasura/README.md#get) - Hasura auth hook endpoint

### [~~AuthnSaml~~](docs/sdks/authnsaml/README.md)

* [~~Assert~~](docs/sdks/authnsaml/README.md#assert) - SAML ACS endpoint (social) :warning: **Deprecated**

### [CaptchaProvider](docs/sdks/captchaprovider/README.md)

* [Get](docs/sdks/captchaprovider/README.md#get) - Get captcha provider
* [Update](docs/sdks/captchaprovider/README.md#update) - Update captcha provider

### [CaptchaProviders](docs/sdks/captchaproviders/README.md)

* [Delete](docs/sdks/captchaproviders/README.md#delete) - Delete captcha provider

### [Configs](docs/sdks/configs/README.md)

* [GetAdminConsole](docs/sdks/configs/README.md#getadminconsole) - Get admin console config
* [UpdateAdminConsole](docs/sdks/configs/README.md#updateadminconsole) - Update admin console config
* [GetOidcKeys](docs/sdks/configs/README.md#getoidckeys) - Get OIDC keys
* [DeleteOidcKey](docs/sdks/configs/README.md#deleteoidckey) - Delete OIDC key
* [RotateOidcKeys](docs/sdks/configs/README.md#rotateoidckeys) - Rotate OIDC keys
* [GetJwtCustomizer](docs/sdks/configs/README.md#getjwtcustomizer) - Get JWT customizer
* [DeleteJwtCustomizer](docs/sdks/configs/README.md#deletejwtcustomizer) - Delete JWT customizer
* [ListJwtCustomizers](docs/sdks/configs/README.md#listjwtcustomizers) - Get all JWT customizers
* [TestJwtCustomizer](docs/sdks/configs/README.md#testjwtcustomizer) - Test JWT customizer

### [ConnectorFactories](docs/sdks/connectorfactories/README.md)

* [List](docs/sdks/connectorfactories/README.md#list) - Get connector factories
* [Get](docs/sdks/connectorfactories/README.md#get) - Get connector factory

### [Connectors](docs/sdks/connectors/README.md)

* [Create](docs/sdks/connectors/README.md#create) - Create connector
* [List](docs/sdks/connectors/README.md#list) - Get connectors
* [Get](docs/sdks/connectors/README.md#get) - Get connector
* [Update](docs/sdks/connectors/README.md#update) - Update connector
* [Delete](docs/sdks/connectors/README.md#delete) - Delete connector
* [Test](docs/sdks/connectors/README.md#test) - Test passwordless connector
* [CreateAuthorizationURI](docs/sdks/connectors/README.md#createauthorizationuri) - Get connector's authorization URI

### [CustomPhrases](docs/sdks/customphrases/README.md)

* [GetAll](docs/sdks/customphrases/README.md#getall) - Get all custom phrases
* [Get](docs/sdks/customphrases/README.md#get) - Get custom phrases
* [Upsert](docs/sdks/customphrases/README.md#upsert) - Upsert custom phrases
* [Delete](docs/sdks/customphrases/README.md#delete) - Delete custom phrase

### [Dashboard](docs/sdks/dashboard/README.md)

* [GetNewUserCounts](docs/sdks/dashboard/README.md#getnewusercounts) - Get new user count
* [GetActiveUsers](docs/sdks/dashboard/README.md#getactiveusers) - Get active user data

### [DashboardUsers](docs/sdks/dashboardusers/README.md)

* [GetTotalCount](docs/sdks/dashboardusers/README.md#gettotalcount) - Get total user count

### [Domains](docs/sdks/domains/README.md)

* [List](docs/sdks/domains/README.md#list) - Get domains
* [Create](docs/sdks/domains/README.md#create) - Create domain
* [Get](docs/sdks/domains/README.md#get) - Get domain
* [Delete](docs/sdks/domains/README.md#delete) - Delete domain

### [EmailTemplates](docs/sdks/emailtemplates/README.md)

* [Replace](docs/sdks/emailtemplates/README.md#replace) - Replace email templates
* [List](docs/sdks/emailtemplates/README.md#list) - Get email templates
* [Remove](docs/sdks/emailtemplates/README.md#remove) - Delete email templates
* [GetByID](docs/sdks/emailtemplates/README.md#getbyid) - Get email template by ID
* [Delete](docs/sdks/emailtemplates/README.md#delete) - Delete an email template
* [UpdateDetails](docs/sdks/emailtemplates/README.md#updatedetails) - Update email template details

### [EnterpriseSsoVerifications](docs/sdks/enterprisessoverifications/README.md)

* [Create](docs/sdks/enterprisessoverifications/README.md#create) - Create enterprise SSO verification

### [Experience](docs/sdks/experience/README.md)

* [InitInteraction](docs/sdks/experience/README.md#initinteraction) - Init new interaction
* [Submit](docs/sdks/experience/README.md#submit) - Submit interaction
* [CreatePasswordVerification](docs/sdks/experience/README.md#createpasswordverification) - Create password verification record
* [VerifyCode](docs/sdks/experience/README.md#verifycode) - Verify verification code
* [CreateSocialVerification](docs/sdks/experience/README.md#createsocialverification) - Create social verification
* [VerifySocial](docs/sdks/experience/README.md#verifysocial) - Verify social verification
* [VerifyEnterpriseSso](docs/sdks/experience/README.md#verifyenterprisesso) - Verify enterprise SSO verification
* [VerifyTotp](docs/sdks/experience/README.md#verifytotp) - Verify TOTP verification
* [CreateWebAuthnRegistration](docs/sdks/experience/README.md#createwebauthnregistration) - Create WebAuthn registration verification
* [VerifyWebAuthnRegistration](docs/sdks/experience/README.md#verifywebauthnregistration) - Verify WebAuthn registration verification
* [VerifyWebAuthnAuthentication](docs/sdks/experience/README.md#verifywebauthnauthentication) - Verify WebAuthn authentication verification
* [GenerateBackupCodes](docs/sdks/experience/README.md#generatebackupcodes) - Generate backup codes
* [VerifyBackupCode](docs/sdks/experience/README.md#verifybackupcode) - Verify backup code
* [VerifyOneTimeToken](docs/sdks/experience/README.md#verifyonetimetoken) - Verify one-time token
* [ResetPassword](docs/sdks/experience/README.md#resetpassword) - Reset user password
* [BindMfaVerification](docs/sdks/experience/README.md#bindmfaverification) - Bind MFA verification by verificationId

### [Hooks](docs/sdks/hooks/README.md)

* [List](docs/sdks/hooks/README.md#list) - Get hooks
* [Create](docs/sdks/hooks/README.md#create) - Create a hook
* [Get](docs/sdks/hooks/README.md#get) - Get hook
* [Update](docs/sdks/hooks/README.md#update) - Update hook
* [Delete](docs/sdks/hooks/README.md#delete) - Delete hook
* [ListRecentLogs](docs/sdks/hooks/README.md#listrecentlogs) - Get recent logs for a hook
* [Test](docs/sdks/hooks/README.md#test) - Test hook
* [UpdateSigningKey](docs/sdks/hooks/README.md#updatesigningkey) - Update signing key for a hook

### [Identifications](docs/sdks/identifications/README.md)

* [IdentifyUser](docs/sdks/identifications/README.md#identifyuser) - Identify user for the current interaction

### [Interaction](docs/sdks/interaction/README.md)

* [Update](docs/sdks/interaction/README.md#update)
* [Delete](docs/sdks/interaction/README.md#delete)
* [PutEvent](docs/sdks/interaction/README.md#putevent)
* [PatchIdentifiers](docs/sdks/interaction/README.md#patchidentifiers)
* [UpdateProfile](docs/sdks/interaction/README.md#updateprofile)
* [PatchProfile](docs/sdks/interaction/README.md#patchprofile)
* [DeleteProfile](docs/sdks/interaction/README.md#deleteprofile)
* [Submit](docs/sdks/interaction/README.md#submit)
* [CreateConsent](docs/sdks/interaction/README.md#createconsent)
* [GetConsent](docs/sdks/interaction/README.md#getconsent)
* [PostVerificationCode](docs/sdks/interaction/README.md#postverificationcode)
* [PostTotp](docs/sdks/interaction/README.md#posttotp)
* [VerifyWebAuthnRegistration](docs/sdks/interaction/README.md#verifywebauthnregistration)
* [WebauthnAuthentication](docs/sdks/interaction/README.md#webauthnauthentication)
* [BindMfa](docs/sdks/interaction/README.md#bindmfa)
* [UpdateMfa](docs/sdks/interaction/README.md#updatemfa)
* [SkipMfa](docs/sdks/interaction/README.md#skipmfa)
* [AuthenticateSso](docs/sdks/interaction/README.md#authenticatesso)
* [RegisterSingleSignOn](docs/sdks/interaction/README.md#registersinglesignon)
* [GetSingleSignOnConnectors](docs/sdks/interaction/README.md#getsinglesignonconnectors)

#### [Interaction.Sso](docs/sdks/sso/README.md)

* [PostAuthorizationURL](docs/sdks/sso/README.md#postauthorizationurl)

### [InteractionEvents](docs/sdks/interactionevents/README.md)

* [Update](docs/sdks/interactionevents/README.md#update) - Update interaction event

### [Interactions](docs/sdks/interactions/README.md)

* [VerifySocialAuthorizationURI](docs/sdks/interactions/README.md#verifysocialauthorizationuri)

### [JwtCustomizers](docs/sdks/jwtcustomizers/README.md)

* [Upsert](docs/sdks/jwtcustomizers/README.md#upsert) - Create or update JWT customizer
* [Update](docs/sdks/jwtcustomizers/README.md#update) - Update JWT customizer


### [Mfa](docs/sdks/mfa/README.md)

* [SkipBinding](docs/sdks/mfa/README.md#skipbinding) - Skip MFA binding flow

### [MyAccount](docs/sdks/myaccount/README.md)

* [GetProfile](docs/sdks/myaccount/README.md#getprofile) - Get profile
* [Modify](docs/sdks/myaccount/README.md#modify) - Update profile
* [UpdateProfile](docs/sdks/myaccount/README.md#updateprofile) - Update other profile
* [UpdatePassword](docs/sdks/myaccount/README.md#updatepassword) - Update password
* [UpdatePrimaryEmail](docs/sdks/myaccount/README.md#updateprimaryemail) - Update primary email
* [DeletePrimaryEmail](docs/sdks/myaccount/README.md#deleteprimaryemail) - Delete primary email
* [UpdatePrimaryPhone](docs/sdks/myaccount/README.md#updateprimaryphone) - Update primary phone
* [DeletePrimaryPhone](docs/sdks/myaccount/README.md#deleteprimaryphone) - Delete primary phone
* [AddUserIdentities](docs/sdks/myaccount/README.md#adduseridentities) - Add a user identity
* [DeleteIdentity](docs/sdks/myaccount/README.md#deleteidentity) - Delete a user identity

### [OneTimeTokens](docs/sdks/onetimetokens/README.md)

* [List](docs/sdks/onetimetokens/README.md#list) - Get one-time tokens
* [Add](docs/sdks/onetimetokens/README.md#add) - Create one-time token
* [Get](docs/sdks/onetimetokens/README.md#get) - Get one-time token by ID
* [Delete](docs/sdks/onetimetokens/README.md#delete) - Delete one-time token by ID
* [Verify](docs/sdks/onetimetokens/README.md#verify) - Verify one-time token
* [UpdateStatus](docs/sdks/onetimetokens/README.md#updatestatus) - Update one-time token status

### [OrganizationInvitations](docs/sdks/organizationinvitations/README.md)

* [Fetch](docs/sdks/organizationinvitations/README.md#fetch) - Get organization invitation
* [Delete](docs/sdks/organizationinvitations/README.md#delete) - Delete organization invitation
* [Get](docs/sdks/organizationinvitations/README.md#get) - Get organization invitations
* [Create](docs/sdks/organizationinvitations/README.md#create) - Create organization invitation
* [ResendMessage](docs/sdks/organizationinvitations/README.md#resendmessage) - Resend invitation message
* [ReplaceStatus](docs/sdks/organizationinvitations/README.md#replacestatus) - Update organization invitation status

### [OrganizationRoles](docs/sdks/organizationroles/README.md)

* [Get](docs/sdks/organizationroles/README.md#get) - Get organization role
* [Update](docs/sdks/organizationroles/README.md#update) - Update organization role
* [Delete](docs/sdks/organizationroles/README.md#delete) - Delete organization role
* [List](docs/sdks/organizationroles/README.md#list) - Get organization roles
* [Create](docs/sdks/organizationroles/README.md#create) - Create an organization role
* [ListScopes](docs/sdks/organizationroles/README.md#listscopes) - Get organization role scopes
* [AssignScope](docs/sdks/organizationroles/README.md#assignscope) - Assign organization scopes to organization role
* [ReplaceScopes](docs/sdks/organizationroles/README.md#replacescopes) - Replace organization scopes for organization role
* [DeleteScope](docs/sdks/organizationroles/README.md#deletescope) - Remove organization scope
* [ListResourceScopes](docs/sdks/organizationroles/README.md#listresourcescopes) - Get organization role resource scopes
* [AssignResourceScopes](docs/sdks/organizationroles/README.md#assignresourcescopes) - Assign resource scopes to organization role
* [ReplaceResourceScopes](docs/sdks/organizationroles/README.md#replaceresourcescopes) - Replace resource scopes for organization role
* [DeleteResourceScope](docs/sdks/organizationroles/README.md#deleteresourcescope) - Remove resource scope

### [Organizations](docs/sdks/organizations/README.md)

* [Create](docs/sdks/organizations/README.md#create) - Create an organization
* [List](docs/sdks/organizations/README.md#list) - Get organizations
* [Get](docs/sdks/organizations/README.md#get) - Get organization
* [Update](docs/sdks/organizations/README.md#update) - Update organization
* [Delete](docs/sdks/organizations/README.md#delete) - Delete organization
* [AddUsers](docs/sdks/organizations/README.md#addusers) - Add user members to organization
* [ReplaceUsers](docs/sdks/organizations/README.md#replaceusers) - Replace organization user members
* [ListUsers](docs/sdks/organizations/README.md#listusers) - Get organization user members
* [DeleteUser](docs/sdks/organizations/README.md#deleteuser) - Remove user member from organization
* [AssignRolesToUser](docs/sdks/organizations/README.md#assignrolestouser) - Assign roles to organization user members
* [ListUserRoles](docs/sdks/organizations/README.md#listuserroles) - Get roles for a user in an organization
* [AssignRoles](docs/sdks/organizations/README.md#assignroles) - Assign roles to a user in an organization
* [UpdateUserRoles](docs/sdks/organizations/README.md#updateuserroles) - Update roles for a user in an organization
* [RemoveUserRole](docs/sdks/organizations/README.md#removeuserrole) - Remove a role from a user in an organization
* [ListUserScopes](docs/sdks/organizations/README.md#listuserscopes) - Get scopes for a user in an organization tailored by the organization roles
* [AddApplication](docs/sdks/organizations/README.md#addapplication) - Add organization application
* [ReplaceApplications](docs/sdks/organizations/README.md#replaceapplications) - Replace organization applications
* [ListApplications](docs/sdks/organizations/README.md#listapplications) - Get organization applications
* [RemoveApplication](docs/sdks/organizations/README.md#removeapplication) - Remove organization application
* [AssignRolesToApplications](docs/sdks/organizations/README.md#assignrolestoapplications) - Assign roles to applications in an organization
* [GetApplicationRoles](docs/sdks/organizations/README.md#getapplicationroles) - Get organization application roles
* [AssignRoleToApplication](docs/sdks/organizations/README.md#assignroletoapplication) - Add organization application role
* [ReplaceApplicationRoles](docs/sdks/organizations/README.md#replaceapplicationroles) - Replace organization application roles
* [DeleteApplicationRole](docs/sdks/organizations/README.md#deleteapplicationrole) - Remove organization application role
* [ListJitEmailDomains](docs/sdks/organizations/README.md#listjitemaildomains) - Get organization JIT email domains
* [AddJitEmailDomain](docs/sdks/organizations/README.md#addjitemaildomain) - Add organization JIT email domain
* [ReplaceJitEmailDomains](docs/sdks/organizations/README.md#replacejitemaildomains) - Replace organization JIT email domains
* [DeleteJitEmailDomain](docs/sdks/organizations/README.md#deletejitemaildomain) - Remove organization JIT email domain
* [ListJitRoles](docs/sdks/organizations/README.md#listjitroles) - Get organization JIT default roles
* [CreateJitRole](docs/sdks/organizations/README.md#createjitrole) - Add organization JIT default roles
* [ReplaceJitRoles](docs/sdks/organizations/README.md#replacejitroles) - Replace organization JIT default roles
* [DeleteJitRole](docs/sdks/organizations/README.md#deletejitrole) - Remove organization JIT default role
* [GetJitSsoConnectors](docs/sdks/organizations/README.md#getjitssoconnectors) - Get organization JIT SSO connectors
* [AddJitSsoConnector](docs/sdks/organizations/README.md#addjitssoconnector) - Add organization JIT SSO connectors
* [ReplaceJitSsoConnectors](docs/sdks/organizations/README.md#replacejitssoconnectors) - Replace organization JIT SSO connectors
* [DeleteJitSsoConnector](docs/sdks/organizations/README.md#deletejitssoconnector) - Remove organization JIT SSO connector

### [OrganizationScopes](docs/sdks/organizationscopes/README.md)

* [List](docs/sdks/organizationscopes/README.md#list) - Get organization scopes
* [Create](docs/sdks/organizationscopes/README.md#create) - Create an organization scope
* [Get](docs/sdks/organizationscopes/README.md#get) - Get organization scope
* [Update](docs/sdks/organizationscopes/README.md#update) - Update organization scope
* [Delete](docs/sdks/organizationscopes/README.md#delete) - Delete organization scope

### [PasswordIdentityVerifications](docs/sdks/passwordidentityverifications/README.md)

* [Create](docs/sdks/passwordidentityverifications/README.md#create) - Create new password identity verification

### [Profiles](docs/sdks/profiles/README.md)

* [Add](docs/sdks/profiles/README.md#add) - Add user profile

### [Resources](docs/sdks/resources/README.md)

* [List](docs/sdks/resources/README.md#list) - Get API resources
* [Create](docs/sdks/resources/README.md#create) - Create an API resource
* [Get](docs/sdks/resources/README.md#get) - Get API resource
* [Update](docs/sdks/resources/README.md#update) - Update API resource
* [Delete](docs/sdks/resources/README.md#delete) - Delete API resource
* [SetAsDefault](docs/sdks/resources/README.md#setasdefault) - Set API resource as default
* [ListScopes](docs/sdks/resources/README.md#listscopes) - Get API resource scopes
* [CreateScope](docs/sdks/resources/README.md#createscope) - Create API resource scope
* [UpdateScope](docs/sdks/resources/README.md#updatescope) - Update API resource scope
* [DeleteScope](docs/sdks/resources/README.md#deletescope) - Delete API resource scope

### [Roles](docs/sdks/roles/README.md)

* [List](docs/sdks/roles/README.md#list) - Get roles
* [Create](docs/sdks/roles/README.md#create) - Create a role
* [Get](docs/sdks/roles/README.md#get) - Get role
* [Update](docs/sdks/roles/README.md#update) - Update role
* [Delete](docs/sdks/roles/README.md#delete) - Delete role
* [ListUsers](docs/sdks/roles/README.md#listusers) - Get role users
* [AssignUser](docs/sdks/roles/README.md#assignuser) - Assign role to users
* [DeleteUser](docs/sdks/roles/README.md#deleteuser) - Remove role from user
* [ListApplications](docs/sdks/roles/README.md#listapplications) - Get role applications
* [AssignApplications](docs/sdks/roles/README.md#assignapplications) - Assign role to applications
* [RemoveFromApplication](docs/sdks/roles/README.md#removefromapplication) - Remove role from application
* [GetScopes](docs/sdks/roles/README.md#getscopes) - Get role scopes
* [LinkScopes](docs/sdks/roles/README.md#linkscopes) - Link scopes to role
* [DeleteScope](docs/sdks/roles/README.md#deletescope) - Unlink scope from role

### [SamlApplications](docs/sdks/samlapplications/README.md)

* [Create](docs/sdks/samlapplications/README.md#create) - Create SAML application
* [Get](docs/sdks/samlapplications/README.md#get) - Get SAML application
* [Update](docs/sdks/samlapplications/README.md#update) - Update SAML application
* [Delete](docs/sdks/samlapplications/README.md#delete) - Delete SAML application
* [CreateSecret](docs/sdks/samlapplications/README.md#createsecret) - Create SAML application secret
* [ListSecrets](docs/sdks/samlapplications/README.md#listsecrets) - List SAML application secrets
* [DeleteSecret](docs/sdks/samlapplications/README.md#deletesecret) - Delete SAML application secret
* [UpdateSecret](docs/sdks/samlapplications/README.md#updatesecret) - Update SAML application secret
* [GetMetadata](docs/sdks/samlapplications/README.md#getmetadata) - Get SAML application metadata
* [HandleCallback](docs/sdks/samlapplications/README.md#handlecallback) - SAML application callback

### [SamlApplicationsAuthFlow](docs/sdks/samlapplicationsauthflow/README.md)

* [HandleAuthn](docs/sdks/samlapplicationsauthflow/README.md#handleauthn) - Handle SAML authentication request (Redirect binding)
* [CreateAuthn](docs/sdks/samlapplicationsauthflow/README.md#createauthn) - Handle SAML authentication request (POST binding)

### [SentinelActivities](docs/sdks/sentinelactivities/README.md)

* [BulkDelete](docs/sdks/sentinelactivities/README.md#bulkdelete) - Bulk delete sentinel activities

### [SignInExperience](docs/sdks/signinexperience/README.md)

* [Get](docs/sdks/signinexperience/README.md#get) - Get default sign-in experience settings
* [Update](docs/sdks/signinexperience/README.md#update) - Update default sign-in experience settings
* [UploadCustomUIAssets](docs/sdks/signinexperience/README.md#uploadcustomuiassets) - Upload custom UI assets

### [SignInExperiences](docs/sdks/signinexperiences/README.md)

* [CheckPassword](docs/sdks/signinexperiences/README.md#checkpassword) - Check if a password meets the password policy

### [SsoConnectorProviders](docs/sdks/ssoconnectorproviders/README.md)

* [List](docs/sdks/ssoconnectorproviders/README.md#list) - List all the supported SSO connector provider details

### [SsoConnectors](docs/sdks/ssoconnectors/README.md)

* [Create](docs/sdks/ssoconnectors/README.md#create) - Create SSO connector
* [List](docs/sdks/ssoconnectors/README.md#list) - List SSO connectors
* [Get](docs/sdks/ssoconnectors/README.md#get) - Get SSO connector
* [Delete](docs/sdks/ssoconnectors/README.md#delete) - Delete SSO connector
* [Update](docs/sdks/ssoconnectors/README.md#update) - Update SSO connector
* [GetEnabled](docs/sdks/ssoconnectors/README.md#getenabled) - Get enabled SSO connectors by the given email's domain

### [Status](docs/sdks/status/README.md)

* [Get](docs/sdks/status/README.md#get) - Health check

### [SubjectTokens](docs/sdks/subjecttokens/README.md)

* [Create](docs/sdks/subjecttokens/README.md#create) - Create a new subject token.

### [SwaggerJSON](docs/sdks/swaggerjson/README.md)

* [Get](docs/sdks/swaggerjson/README.md#get) - Get Swagger JSON

### [Systems](docs/sdks/systems/README.md)

* [GetApplicationConfig](docs/sdks/systems/README.md#getapplicationconfig) - Get the application constants.

### [TotpSecrets](docs/sdks/totpsecrets/README.md)

* [Create](docs/sdks/totpsecrets/README.md#create) - Create TOTP secret

### [UserAssets](docs/sdks/userassets/README.md)

* [GetServiceStatus](docs/sdks/userassets/README.md#getservicestatus) - Get service status
* [Upload](docs/sdks/userassets/README.md#upload) - Upload asset

### [Users](docs/sdks/users/README.md)

* [Get](docs/sdks/users/README.md#get) - Get user
* [Update](docs/sdks/users/README.md#update) - Update user
* [Delete](docs/sdks/users/README.md#delete) - Delete user
* [GetCustomData](docs/sdks/users/README.md#getcustomdata) - Get user custom data
* [UpdateCustomData](docs/sdks/users/README.md#updatecustomdata) - Update user custom data
* [UpdateProfile](docs/sdks/users/README.md#updateprofile) - Update user profile
* [Create](docs/sdks/users/README.md#create) - Create user
* [List](docs/sdks/users/README.md#list) - Get users
* [UpdatePassword](docs/sdks/users/README.md#updatepassword) - Update user password
* [VerifyPassword](docs/sdks/users/README.md#verifypassword) - Verify user password
* [HasPassword](docs/sdks/users/README.md#haspassword) - Check if user has password
* [UpdateIsSuspended](docs/sdks/users/README.md#updateissuspended) - Update user suspension status
* [GetRoles](docs/sdks/users/README.md#getroles) - Get roles for user
* [AssignRoles](docs/sdks/users/README.md#assignroles) - Assign roles to user
* [ReplaceRoles](docs/sdks/users/README.md#replaceroles) - Update roles for user
* [DeleteRole](docs/sdks/users/README.md#deleterole) - Remove role from user
* [UpdateIdentity](docs/sdks/users/README.md#updateidentity) - Update social identity of user
* [DeleteIdentity](docs/sdks/users/README.md#deleteidentity) - Delete social identity from user
* [CreateIdentity](docs/sdks/users/README.md#createidentity) - Link social identity to user
* [ListOrganizations](docs/sdks/users/README.md#listorganizations) - Get organizations for a user
* [ListMfaVerifications](docs/sdks/users/README.md#listmfaverifications) - Get user's MFA verifications
* [CreateMfaVerification](docs/sdks/users/README.md#createmfaverification) - Create an MFA verification for a user
* [DeleteMfaVerification](docs/sdks/users/README.md#deletemfaverification) - Delete an MFA verification for a user
* [ListPersonalAccessTokens](docs/sdks/users/README.md#listpersonalaccesstokens) - Get personal access tokens
* [CreatePersonalAccessToken](docs/sdks/users/README.md#createpersonalaccesstoken) - Add personal access token
* [DeletePersonalAccessToken](docs/sdks/users/README.md#deletepersonalaccesstoken) - Delete personal access token
* [UpdatePersonalAccessToken](docs/sdks/users/README.md#updatepersonalaccesstoken) - Update personal access token

### [VerificationCodes](docs/sdks/verificationcodes/README.md)

* [Request](docs/sdks/verificationcodes/README.md#request) - Request and send a verification code
* [Verify](docs/sdks/verificationcodes/README.md#verify) - Verify a verification code
* [Create](docs/sdks/verificationcodes/README.md#create) - Create and send verification code

### [Verifications](docs/sdks/verifications/README.md)

* [CreateByPassword](docs/sdks/verifications/README.md#createbypassword) - Create a record by password
* [CreateByVerificationCode](docs/sdks/verifications/README.md#createbyverificationcode) - Create a record by verification code
* [VerifyByCode](docs/sdks/verifications/README.md#verifybycode) - Verify verification code
* [CreateBySocial](docs/sdks/verifications/README.md#createbysocial) - Create a social verification record
* [VerifyBySocial](docs/sdks/verifications/README.md#verifybysocial) - Verify a social verification record

### [WebAuthnVerifications](docs/sdks/webauthnverifications/README.md)

* [CreateAuthentication](docs/sdks/webauthnverifications/README.md#createauthentication) - Create WebAuthn authentication verification

### [WellKnown](docs/sdks/wellknown/README.md)

* [~~GetSignInExperienceConfig~~](docs/sdks/wellknown/README.md#getsigninexperienceconfig) - Get full sign-in experience :warning: **Deprecated**
* [GetPhrases](docs/sdks/wellknown/README.md#getphrases) - Get localized phrases
* [GetExperience](docs/sdks/wellknown/README.md#getexperience) - Get full sign-in experience
* [GetManagementOpenapiJSON](docs/sdks/wellknown/README.md#getmanagementopenapijson) - Get Management API swagger JSON
* [GetExperienceOpenapiJSON](docs/sdks/wellknown/README.md#getexperienceopenapijson) - Get Experience API swagger JSON
* [GetUserOpenapiJSON](docs/sdks/wellknown/README.md#getuseropenapijson) - Get User API swagger JSON

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"github.com/bllli/logto-management-api/models/operations"
	"github.com/bllli/logto-management-api/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := logtomanagementapi.New(
		logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
	)

	res, err := s.Applications.Get(ctx, operations.ListApplicationsRequest{}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseBodies != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"github.com/bllli/logto-management-api/models/operations"
	"github.com/bllli/logto-management-api/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := logtomanagementapi.New(
		logtomanagementapi.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
	)

	res, err := s.Applications.Get(ctx, operations.ListApplicationsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseBodies != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `Get` function may return the following errors:

| Error Type         | Status Code | Content Type |
| ------------------ | ----------- | ------------ |
| apierrors.APIError | 4XX, 5XX    | \*/\*        |

### Example

```go
package main

import (
	"context"
	"errors"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"github.com/bllli/logto-management-api/models/apierrors"
	"github.com/bllli/logto-management-api/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := logtomanagementapi.New(
		logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
	)

	res, err := s.Applications.Get(ctx, operations.ListApplicationsRequest{})
	if err != nil {

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"github.com/bllli/logto-management-api/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := logtomanagementapi.New(
		logtomanagementapi.WithServerURL("https://[tenant_id].logto.app/"),
		logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
	)

	res, err := s.Applications.Get(ctx, operations.ListApplicationsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseBodies != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/bllli/logto-management-api&utm_campaign=go)
