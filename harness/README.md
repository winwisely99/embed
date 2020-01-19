# harness

## In Process

Mobile
- use gomobile bind.

Desktop
- just use standard go-plugins.

## Out of Process

Mobile

- needs work but not too hard.

- Not sure how a App Manifest can model its dependency on a Background Service. If many apps need the same background Service then is the Service installed in all GUI Apps Or does the OS manage it at runtime ?


Desktop

- system tray works well and avoids admin rights. Google does this for Google Chrome and just sets a startup flag in the users account.

	- this means that each OS user is 100% segregated and so you can hand your laptop to someone else using it and they can see any of your data or even the binaries.


## CI

Use Bazel
 
gomobile: https://github.com/znly/rules_gomobile
dart: NONE FOUND: Watch here: https://github.com/bazelbuild/bazel/issues/10207

- bazel deps:
	- io_bazel_rules_go: https://github.com/bazelbuild/rules_go
	- build_bazel_rules_apple: https://github.com/bazelbuild/rules_apple
	- build_bazel_rules_swift: https://github.com/bazelbuild/rules_swift
	- build_bazel_apple_support: https://github.com/bazelbuild/apple_support

## Compiler approaches

Compiling golang to work on Mobiles, Desktop and Web from a single code base.

These are golang GUI systems but the compiler approaches can be used to build the golang out of process Service.

Fyne
- Web: https://fyne.io/
- Code: https://github.com/fyne-io/fyne

gioui
- Web: https://gioui.org/
- Is Faster because it uses the GPU to do heavy duty

FFI can also be used for piping / IPC

MAYN: https://pub.dev/packages?q=dependency%3Affi&sort=updated

https://github.com/dart-interop/ffi_tool
- gens the dart code from the c code. nice

https://github.com/dart-interop
- OS repo, etc

https://github.com/CanonicalLtd/unix_domain_socket.dart
- another useful channel



## Process Management approahes

https://pub.dev/packages/workmanager
- a wrapper around Android's WorkManager and iOS' performFetchWithCompletionHandler, effectively enabling headless execution of Dart code in the background.
- SO the Flutter code does not block the GUI thread when calling the Background process.
- But also useful on the Golang out of process side also. Need to check..