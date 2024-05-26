# Honyakusha

Translate text using a variety of translation services.


## Installation

### From Binaries

Go to [releases page](https://github.com/souk4711/honyakusha/releases) download latest release binary file.

### From Source

Ensure that you have a supported version of Go properly installed and setup. You can
find the minimum required version of Go in the go.mod file.

You can then install the latest release globally by running:

```sh
$ go install github.com/souk4711/honyakusha/cmd/honyakusha@latest
```


## Usage

Create a configuration file in the XDG config directory:

```console
$ honyakusha init --xdg
Create a config file in $HOME/.config/honyakusha.toml
```

And then execute:

```console
$ honyakusha trans Hello, World
Raw:

    Hello, World

Google Translate:

    こんにちは世界

Bing Translator:

    ハローワールド
```

**NOTE: You can edit the config file to change the target language if it is not Japanese.**


## Translation Services

| Code               | Name                                 | Auth Method |
|--------------------|--------------------------------------|--------------
| bing               | [Bing Translator][bing]              | -           |
| google             | [Google Translate][google]           | -           |
| deepl-api          | [DeepL Translate][deepl-api]         | api-key     |
| libretranslate-api | [LibreTranslate][libretranslate-api] | api-key     |


## Howto

* [Add Honyakusha to GoldenDict](./examples/goldendict/)


## License

available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).


[bing]:https://www.bing.com/translator
[google]:https://translate.google.com/
[deepl-api]:https://www.deepl.com/
[libretranslate-api]:https://libretranslate.com/
