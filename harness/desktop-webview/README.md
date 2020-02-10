# desktop-webview

We need to use a Webview on Desktop because WebRTC and Video does not work on FLutter Desktop yet.

We have a pathway to not use Webview but it relies on using a cusotm Flutter and we dont have time to take on that risk yet.

The only way is from the main go-flutter app to open another window, because we cant composite the Webview inside go-flutter.

SO its a hack if you want to call it that but not so bad because
- when doing video conferencing you want to have a different window
- when watching a video we can have the Screenshot of the video in go-flutter and then when the user clicks it, open a new window


WebVIEW 
https://github.com/wailsapp/wails/issues/235
- Its finally ready.

It has some other nice stuff
- We dont really want to use wails, but instead just load the FLutter Web into the WebView.
- We need an IPC / Message bux though and wails has a nice one.
- https://github.com/wailsapp/wails/blob/master/runtime/window.go
	- detects lang.
	- full scren detect
	

## Flutter web and mobile
- there is can do everything in the app like normal because Video and Flutter WebRTC works there.

