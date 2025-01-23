from __future__ import print_function

import logging
import asyncio
import os
import subprocess
import time

import grpc

from dotenv import load_dotenv
from grpc.aio import AioRpcError

from protogen import llmproxy_pb2_grpc, llmproxy_pb2, sso_pb2, sso_pb2_grpc
from google.protobuf import empty_pb2

from proxy import ya_gpt

import threading
import time

load_dotenv()

PROXY_SOCKET = os.environ.get("AIPROXY_GRPC_SOCKET")
BACKEND_GRPC_SOCKET = os.environ.get("BACKEND_GRPC_SOCKET")
SSO_GRPC_SOCKET = os.environ.get("SSO_GRPC_SOCKET")
APP_UUID = os.environ.get("APP_UUID")
SERVICE_JWT = os.environ.get("SERVICE_JWT")

async def run_llm_remote(request: llmproxy_pb2.LLMRequest):
    print("Will try to greet world ...")
    async with grpc.aio.insecure_channel(BACKEND_GRPC_SOCKET) as channel:
        stub = llmproxy_pb2_grpc.LLMProxyStub(channel)
        try:
            await check_permissions(request)
            code = ya_gpt(request.content)
            print(request, request.uuid)
            await stub.SendReply(llmproxy_pb2.LLMReply(jwt=SERVICE_JWT, response=code, uuid=request.uuid))
        except PermissionError as error:
            print(error)


async def check_permissions(request: llmproxy_pb2.LLMRequest):
    print(request)
    async with grpc.aio.insecure_channel(SSO_GRPC_SOCKET) as channel:
        try:
            stub = sso_pb2_grpc.PermissionsStub(channel)
            response = await stub.GetUserPermissions(sso_pb2.GetUserPermissionsRequest(jwt=request.jwt))
            app = list(filter(lambda x: x.app_uuid == APP_UUID, response.apps))[0]
            if len(list(filter(lambda x: x == 5, app.permissions))):
                pass
            else:
                raise PermissionError('No permission found')
        except AioRpcError as error:
            raise PermissionError(error.details()) from error
        except Exception as error:
            raise PermissionError(error) from error


class LLMProxyHandler(llmproxy_pb2_grpc.LLMProxyServicer):

    async def SendRequest(
        self,
        request: llmproxy_pb2.LLMRequest,
        _,
    ) -> llmproxy_pb2.LLMReply:
        loop = asyncio.get_event_loop()
        loop.create_task(run_llm_remote(request))
        return empty_pb2.Empty()


async def serve() -> None:
    server = grpc.aio.server()
    llmproxy_pb2_grpc.add_LLMProxyServicer_to_server(LLMProxyHandler(), server)
    listen_addr = PROXY_SOCKET
    server.add_insecure_port(listen_addr)
    logging.info("Starting server on %s", listen_addr)
    await server.start()
    await server.wait_for_termination()


if __name__ == "__main__":

    subprocess.run("scripts/update_ya_token.bash")
    def delayed_task():
        subprocess.run("scripts/update_ya_token.bash")
        time.sleep(60)

    task_thread = threading.Thread(target=delayed_task)
    task_thread.start()

    logging.basicConfig(level=logging.INFO)
    asyncio.run(serve())


# INSERT INTO user_roles VALUES ('PT_SHARE', '2024-12-31 21:34:54.508775+00', '2024-12-31 21:34:54.508775+00', 'proxy', '7a9a567d-3756-4706-b1e7-d795e908c770');