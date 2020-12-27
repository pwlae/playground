# protobuf scenario 1
Server backend application returns protobuf object as part of HTTP server response.
In this particualt example I'm using python based client.

## When it can be useful
File upload. For example, you want to crop and upload an image to amazon s3 in different sizes. As a result of the upload, you want to have an object with links to all copies. 

Instead of having JSON as a backend response data, like in the example below
```json
{
    "original": "link",
    "small": "link",
    "medium": "link",
    "large": "link"
}
```

You can receive an object. There are some advantages:
* no serialization needed
* strict contract of response data

## Run it locally
```sh
docker compose up --build -V
```


