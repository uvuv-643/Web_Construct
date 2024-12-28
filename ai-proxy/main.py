from __future__ import print_function

import logging


import grpc

from protogen import llmproxy_pb2_grpc, llmproxy_pb2


def run():
    print("Will try to greet world ...")
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = llmproxy_pb2_grpc.LLMProxyStub(channel)
        stub.SendRequest(llmproxy_pb2.LLMRequest(jwt='123', content='321'))


if __name__ == "__main__":
    logging.basicConfig()
    run()


