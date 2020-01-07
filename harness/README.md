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

## LIBS

I just tested these.

Its a way to model using GRPC but run the two endpoints in process, and so no need a port based endpoint.

SO one side is golang embedded on the Flutter side and the other is on the local Service side.

https://github.com/fullstorydev/grpchan
&
https://github.com/jhump/goprotoc

