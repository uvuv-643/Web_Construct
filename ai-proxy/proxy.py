from __future__ import print_function

import os
import requests
from dotenv import load_dotenv

load_dotenv()

YA_TOKEN = os.environ.get("YA_TOKEN")
FOLDER_ID = os.environ.get("FOLDER_ID")
YA_GPT_URL = os.environ.get("YA_GPT_URL")


class YAGPTException(Exception):
    pass


def get_response_from_ya_gpt(data: str) -> str:
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
        "text": "Напиши REACT компонент, основываясь на пользовательском описании. В ответ включи только код, без форматирования, текстом. Пиши всю логику обработки пользовательских сценариев и не оставляй коментариев в тексте и нереализованных вещей. Используй typescript и первым в очереди определенном типом напиши пропсы (используй type). НЕ используй сторонние библиотеки, напиши всю внутреннюю логику собственноручно, не оставляй комментариев и пиши всю реализацию!!!"
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
