#! /usr/local/bin/python3

import sys
import requests
import image_pb2

response = requests.get('http://server:8080/upload')

image = image_pb2.Image()
image.ParseFromString(response.content)

print('Original image URL : ' + image.original)
print('Small image URL    : ' + image.small)
print('Medium image URL   : ' + image.medium)
print('Large image URL    : ' + image.large)
