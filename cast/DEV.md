## Casting

Device(s) over wifi → Video mixer ( presenter ) over wifi → google Chromecast ultra
See Embed/Cast

1. Device

- stream OS level audio and video to Presenter
- Desktop ( easy )
- Mobile ( hard )

If Flutter WebRTC works then it is done.

Flutter cast
https://github.com/terrabythia/dart_chromecast

https://pub.dev/packages/flutter_google_cast_button



2. Presenter

- Mixers all devices stream into a single stream to the Cast

Golang cast
https://github.com/vishen/go-chromecast/
- can transcode
- has CLI 
- seesm perfect

3. Remote Users

Pion watch
https://github.com/pion/rtwatch
- synchronised streaming of a video.
- can pause and sync remotly.





