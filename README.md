# Honyakusha

Translate text using a variety of translation services.


## Installation

### From Source

Ensure that you have a supported version of Go properly installed and setup. You can
find the minimum required version of Go in the go.mod file.

You can then install the latest release globally by running:

```sh
$ go install github.com/souk4711/honyakusha@latest
```


## Usage

Create a configuration in the XDG config directory or current working directory:

```sh
$ honyakusha init
Create a config file in ./honyakusha.toml
```

And then execute:

```sh
$ honyakusha trans 吟味されざる生に、生きる価値なし。
Raw:

    吟味されざる生に、生きる価値なし。

Google Translate:

    A life that is not examined is not worth living.

Bing Translator:

    A life that is not examined is not worth living.
```


## Translation Services

| Code               | Name                                 | Auth Method |
|--------------------|--------------------------------------|--------------
| bing               | [Bing Translator][bing]              | -           |
| google             | [Google Translate][google]           | -           |
| libretranslate-api | [LibreTranslate][libretranslate-api] | api-key     |


## Formatter

* html
* json
* plain


## Howto

* [Add Honyakusha to GoldenDict](./examples/goldendict/)


## License

available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).


[bing]:https://www.bing.com/translator
[google]:https://translate.google.com/
[libretranslate-api]:https://libretranslate.com/
