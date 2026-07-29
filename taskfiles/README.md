# Quicken tasks

`quicken.yml` is the shared GoForge Task surface for applications using
Quicken. Include or copy it into a starter and provide `quicken.yaml`; no
platform compiler, packager, signing, simulator, or emulator commands belong
in the application Taskfile.

```sh
task doctor
task build:web
task build:desktop
task run:ios-simulator
task run:android-emulator
```
