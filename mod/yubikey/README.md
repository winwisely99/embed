# yubikey

2FA independent of SMS and other stuff





## hardware
- USB 2.0 and NFC
	- yubikey-5-nfc
	- https://www.yubico.com/product/yubikey-5-nfc
	- https://support.yubico.com/support/solutions/articles/15000014174--yubikey-5-nfc
	- NFC works for IOS
	- Android with USB 2
- USB C (for android)
	- YubiKey-5C
	- https://www.yubico.com/product/yubikey-5c


## Software


https://github.com/GeertJohan/yubigo
https://en.wikipedia.org/wiki/Happy_Eyeballs
- code code makes calls to ipv4 and ipv6 endpoints. Smart approach

webauthn
https://www.yubico.com/authentication-standards/webauthn/
https://github.com/duo-labs/webauthn
- code to provide a Server to do it.
- we could do this over P2P to the administrator of a group also !! One aspect to the Ring based security approahc


Aspects: 

- Desktop. Easy using pure go.
- Mobile
	- IOS
		- Need Yubikey that works for IOS
			- https://support.yubico.com/support/solutions/articles/15000006479-getting-started-on-ios
		- NFC OR Lightening connector.
		- I can see 2 or 3 ways to do it but i askd how they did: https://github.com/GeertJohan/yubigo/issues/7
- Web
	- backed into all browsers via webauthn.
	- USB hardware channel