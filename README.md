# Data Over Sound using Golang
runs on macos

## Setup
Install homebrew
```shell script
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```
Then install the [PortAudio](http://portaudio.com) Library and pkg-config for CGO using brew
```shell script
HOMEBREW_NO_AUTO_UPDATE=1 brew install portaudio
HOMEBREW_NO_AUTO_UPDATE=1 brew install pkg-config
```
Once you're done with that you can build and run the app.

This app is still unfinished.
Actively looking for a valid method of transcoding the audio signals to textual data and also noise cancellation