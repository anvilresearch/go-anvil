// Package anvil provides support for validating an Anvil JWT and extracting
// the claims for authorization.
//
// An application needs to call the `/signin` and then the `/token` API calls.
// These calls authenticate the user for the application and provide the token
// required to make future calls into any webservice you are building that
// requires authentication/authorization.
//
// SignIn
//
//		You will need these values to call the signin API:
//		HOST         = 10.0.1.26:3000
//		CLIENTID     = 6b6efaae-0ab8-4152-8f92-a87c17921800
//		REDIRECT_URL = https://anvil.coralproject.net
//		EMAIL        = bill@thekennedyclan.net
//		PASSWORD     = Qfe^bJ9uD6cgnD-8
//		REFERRER     = https://anvil.coralproject.net/signin
//
// 		curl -X POST https://HOST/signin -d 'max_age=315569260&response_type=code&client_id=CLIENTID&redirect_uri=REDIRECT_URL&scope=openid%20profile%20email%20realm&provider=password&email=EMAIL&password=PASSWORD -H "referrer: REFERRER"
//
// Response
// 		Redirecting to https://anvil.coralproject.net?code=c9ce6c03ea6ad8dd3f0a%
//
// Token
//
//		You will need these values to call the token API:
//		HOST         = 10.0.1.26:3000
//		CLIENTID     = 6b6efaae-0ab8-4152-8f92-a87c17921800
//		REDIRECT_URL = https://anvil.coralproject.net
//		REFERRER     = https://anvil.coralproject.net/signin
//		CODE         = 6dafd2b59d6954849a6c  // From the response of the signin call
//
// curl -X POST https://CLIENTID:CODE@HOST/token -d 'grant_type=authorization_code&client_id=CLIENTID&code=CODE&redirect_uri=REDIRECT_URL' -H "referrer: REFERRER"
//
package anvil
