# embed



## Embed golang inside Flutter

Basic example for bring up that has everything needed.

https://github.com/DimitarPetrov

- golang lib: https://github.com/DimitarPetrov/stegify
- flutter plugin that uses the golang lib: https://github.com/DimitarPetrov/stegify-flutter-plugin
- flutter app that uses the flutter plugin: https://github.com/DimitarPetrov/stegify-mobile


## Running as Process

Mobile

- needs work but not too hard.

- Not sure how a App Manifest can model its dependency on a Background Service. If many apps need the same background Service then is the Service installed in all GUI Apps Or does the OS manage it at runtime ?



Desktop

- system tray works well and avoids admin rights. Google does this for Google Chrome and just sets a startup flag in the users account.

	- this means that each OS user is 100% segregated and so you can hand your laptop to someone else using it and they can see any of your data or even the binaries.