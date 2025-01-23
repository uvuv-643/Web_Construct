from __future__ import print_function

import os
import requests
from dotenv import load_dotenv

load_dotenv()

FOLDER_ID = os.environ.get("FOLDER_ID")
YA_GPT_URL = os.environ.get("YA_GPT_URL")
YA_TOKEN = os.environ.get("YA_TOKEN")

class YAGPTException(Exception):
    pass


def get_response_from_ya_gpt(data: str) -> str:
  global YA_TOKEN

  load_dotenv()
  YA_TOKEN = os.environ.get("YA_TOKEN")

  payload = {
    "modelUri": f"gpt://{FOLDER_ID}/yandexgpt-lite",
    "completionOptions": {
      "stream": False,
      "temperature": 0.6,
      "maxTokens": "5000"
    },
    "messages": [
      {
        "role": "system",
        "text": "Напиши HTML + inline CSS компонент, основываясь на пользовательском описании. В ответ включи только код, без форматирования, текстом. Не оставляй коментариев в тексте и нереализованных вещей. Используй HTML и стили в одном файле, без разделения на файлы. НЕ используй сторонние библиотеки и javascript, напиши всю внутреннюю логику собственноручно, не оставляй комментариев и пиши всю реализацию!!!"
      },
      {
        "role": "user",
        "text": data
      }
    ]
  }
  try:
    response = requests.post(YA_GPT_URL, json=payload, headers={'Authorization': f'Bearer {YA_TOKEN}'})
    response.raise_for_status()
    return response.json()['result']['alternatives'][0]['message']['text']
  except Exception as e:
    raise YAGPTException(e)


def filter_output(output: str) -> str:
  output = output.replace('```python', '')
  output = output.replace('```', '')
  return output


def ya_gpt(input_value: str) -> str:
  return filter_output(get_response_from_ya_gpt(input_value))
