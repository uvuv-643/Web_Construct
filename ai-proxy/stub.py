from __future__ import print_function

import logging
import asyncio
import os
import time

import grpc
from dotenv import load_dotenv

from protogen import llmproxy_pb2_grpc, llmproxy_pb2
from google.protobuf import empty_pb2


load_dotenv()

PROXY_SOCKET = os.environ.get("AIPROXY_GRPC_SOCKET")
BACKEND_GRPC_SOCKET = os.environ.get("BACKEND_GRPC_SOCKET")

jwts = []
responses = []

class LLMProxyHandlerStub(llmproxy_pb2_grpc.LLMProxyServicer):

    async def run_llm_remote(self):
        print("Run")
        async with grpc.aio.insecure_channel(PROXY_SOCKET) as channel:
            stub = llmproxy_pb2_grpc.LLMProxyStub(channel)
            await stub.SendRequest(llmproxy_pb2.LLMRequest(jwt="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJoZWxsbzFAZ21haWwuY29tIn0.3adlNmVQfAagG7F3w5OF5sk2qXnRRXMmy433Abdb1sA", content="hello"))


    async def SendReply(
        self,
        request: llmproxy_pb2.LLMRequest,
        _,
    ) -> llmproxy_pb2.LLMReply:
        jwts.append(request.jwt)
        responses.append(request.response)
        print(request)
        return empty_pb2.Empty()


async def serve_stub() -> None:

    for i in range(5):
        asyncio.create_task(LLMProxyHandlerStub().run_llm_remote())

    server = grpc.aio.server()
    llmproxy_pb2_grpc.add_LLMProxyServicer_to_server(LLMProxyHandlerStub(), server)
    listen_addr = BACKEND_GRPC_SOCKET
    server.add_insecure_port(listen_addr)
    logging.info("Starting server on %s", listen_addr)
    await server.start()
    await server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    asyncio.run(serve_stub())
