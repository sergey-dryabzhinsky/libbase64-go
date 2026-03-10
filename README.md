# libbase64-go
Last version: 0.0.1

Pure base64 functions nothing else. Exported from golang runtime. And some examples of code.

## Requirements
To compile library and examples you will need to install
- c99 capable compiler (for now only gcc >=4.8 supported)
- golang >=1.6

## Prepare system
Instructions for Debian/Ubuntu like systems:
```
apt install build-essential tar xz-utils m4 sed
apt install golang-go
```

## Compile

If your version of go not default:
```
GOROOT=/path/to/golang/installdir GO=go-x.y make lib
```

## Reason
Not found library with only one functionality, so I tryed make it by myself.

## Result
@tbd

## API
- **libbase64_go__version**(void): Returns doted version string.

  return: `* char`

  *since*: 0.0.1

- **libbase64_go__version_go**(void): Returns version string of go rutime used.

  return: `* char`

  *since*: 0.0.1

- **libbase64_go_nts__getLastErrorCode**(void): int

  *Not Thread Safe*

  Returns code of last error happened.

  params:
  - none

  return: `int` internal code number

  *since*: 0.0.1

- **libbase64_go_nts__getErrorDescription**(int errno):

  *Not Thread Safe*

  Returns description of error by its code number.

  params:
  - `int` errno: error internal code number.

  return: `char*` internal error description.

  *since*: 0.0.1

- **libbase64_go_nts__BASE64_encode_std**(char* inputText): Returns encoded string

  *Not Thread Safe*

  *Standard* variant of encoding.

  params:
  - inputText (`char *`): input string to encode

  return: `char *` encoded string

  *since*: 0.0.1

- **libbase64_go_nts__BASE64_decode_std**(char* inputText): Returns decoded string

  *Not Thread Safe*

  Accapts *Standard* variant of encoding.

  params:
  - inputText (`char *`): input string to decode

  return: `char *` decoded string or NULL on error.

  *since*: 0.0.1

- **libbase64_go_nts__BASE64_encode_url**(char* inputText): Returns encoded string

  *Not Thread Safe*

  *URL-safe* variant of encoding.

  params:
  - inputText (`char *`): input string to encode

  return: `char *` encoded string

  *since*: 0.0.1

- **libbase64_go_nts__BASE64_encode_raw**(char* inputText): Returns encoded string

  *Not Thread Safe*

  *RAW* variant of encoding w/o padding

  params:
  - inputText (`char *`): input string to encode

  return: `char *` encoded string

  *since*: 0.0.1

- **libbase64_go__FreeResult**(char* ptr): frees memory allocated for base64 earlier.

  return: `void`

  *since*: 0.0.1
