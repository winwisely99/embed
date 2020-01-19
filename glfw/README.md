# glfw

GUI stuff often relies on native at the opengl and glfl level.

This causes issues because lots of graphics intensive things like Webrtc relies on native for the compositing.

However we can bypass flutter and use pure glfw and openGL dart code.
It the same way we use it with go-flutter.
But we cant run go-flutter on mobiles and so end up needing native.
SO its a catch 22 !!

Use cases
- rendering 3d stuff in golang and wanting to pipe it to the GUI layer.
- vector maps for the maps module
- webrtc video

Out of Process - Called from FLutter
https://github.com/alnitak/flutter_opengl/blob/master/lib/flutteropengl.dart
- this just is a cal site from FLutter 
- The golang can be doing all the WebRTC video and audio and send back the audio and video.
- so works for Out Of Prcoess approach to pump into the Fluttr Canvas
- others
	- https://github.com/acapper/eBookReader/blob/master/example/lib/main.dart

In Process - Dart only.

- might be a viable solution:

https://pub.dev/packages/opengl
https://pub.dev/packages/glfw
https://pub.dev/packages/ffi_utils

https://gitlab.com/ext/dart-opengl
https://gitlab.com/ext/dart-glfw


How to do it for Video
I believe you can do this (if the video is decoded as RGBA or ARGB).

Just send raw address of memory (int) from Go to Flutter and use allocate() function to create a Pointer<Uint8>.
Using instantiateImageCodec and pointer.asTypedList() you create Codec and get an Image object. Don't forget to add BMP header.
Draw it using Canvas.drawImage().